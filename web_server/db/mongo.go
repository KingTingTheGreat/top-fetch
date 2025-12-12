package db

import (
	"context"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/kingtingthegreat/top-fetch/env"
	"github.com/kingtingthegreat/top-fetch/providers"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DBUser struct {
	ProviderId   string             `bson:"providerId"`
	Provider     providers.Provider `bson:"provider"`
	Id           string             `bson:"id"`
	AccessToken  string             `bson:"accessToken"`
	RefreshToken string             `bson:"refreshToken"`
}

func generateId() string {
	randStr := func() string {
		chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		b := make([]byte, 32)
		for i := range b {
			b[i] = chars[rand.Intn(len(chars))]
		}
		return string(b)
	}

	for {
		id := randStr()
		_, err := GetUserById(id)
		if err != nil {
			return id
		}
	}
}

func ConnectDB() *mongo.Client {
	env.LoadEnv()

	log.Println("Environment Variables:")
	for i, env := range os.Environ() {
		log.Println(i, env)
	}

	bg := context.Background()
	wT, cancel := context.WithTimeout(bg, 10000*time.Millisecond)
	defer func() { cancel() }()
	client, err := mongo.Connect(options.Client().ApplyURI(env.EnvVal("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(wT, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("connected to mongodb")
	return client
}

var db *mongo.Client = ConnectDB()
var userColletion *mongo.Collection = getCollection(env.EnvVal("COLLECTION_NAME"))

func getCollection(collectionName string) *mongo.Collection {
	environment := env.EnvVal("ENVIRONMENT")
	if environment == "" {
		environment = "dev"
	}

	return db.Database(env.EnvVal("DB_NAME") + environment).Collection(collectionName)
}

func GetUserById(id string) (*DBUser, error) {
	var user DBUser
	err := userColletion.FindOne(context.Background(), bson.M{"id": id}).Decode(&user)
	return &user, err
}

func GetUserByProviderId(providerId string) (*DBUser, error) {
	var user DBUser
	err := userColletion.FindOne(context.Background(), bson.M{"providerId": providerId}).Decode(&user)
	return &user, err
}

func InsertUser(user *DBUser) (string, error) {
	user.Id = generateId()

	_, err := userColletion.InsertOne(context.Background(), *user)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return user.Id, nil
}

func UpdateUser(user *DBUser) error {
	_, err := userColletion.UpdateOne(context.Background(), bson.M{"id": user.Id}, bson.M{"$set": *user})
	return err
}
