package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	handler "github.com/kingtingthegreat/top-fetch/api"
)

//	func h(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
//		handler.Handler()
//		return events.APIGatewayProxyResponse{
//			StatusCode: 200,
//			Body:       "Hello AWS Lambda and Netlify",
//		}, nil
//		package main
//
// Lambda handler wrapper
func h(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Build an *http.Request from the API Gateway event
	r, err := http.NewRequest(
		request.HTTPMethod,
		request.Path,
		bytes.NewBufferString(request.Body),
	)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	// Copy headers
	for k, v := range request.Headers {
		r.Header.Set(k, v)
	}

	// For query string parameters
	q := r.URL.Query()
	for k, v := range request.QueryStringParameters {
		q.Set(k, v)
	}
	r.URL.RawQuery = q.Encode()

	// Use a ResponseRecorder to capture the output
	rec := httptest.NewRecorder()

	// Call the existing handler
	handler.Handler(rec, r)

	// Convert ResponseRecorder → APIGatewayProxyResponse
	resp := events.APIGatewayProxyResponse{
		StatusCode: rec.Code,
		Headers:    map[string]string{},
		Body:       rec.Body.String(),
	}

	// Copy headers back to API Gateway response
	for k, v := range rec.Header() {
		// If multiple values, join using comma per HTTP conventions
		resp.Headers[k] = joinHeaderValues(v)
	}

	return resp, nil
}

// Helper to combine multiple header values
func joinHeaderValues(values []string) string {
	if len(values) == 1 {
		return values[0]
	}
	var buf bytes.Buffer
	for i, v := range values {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(v)
	}
	return buf.String()
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(h)
}
