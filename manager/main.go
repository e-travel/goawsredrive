package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

type StackRequest struct {
	Name     string `json:"name"`
	HttpVerb string `json:"http_verb"`
	Endpoint string `json:"endpoint"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	log.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	log.Printf("Body size = %d.\n", len(request.Body))
	log.Println("Headers:")
	for key, value := range request.Headers {
		log.Printf("    %s: %s\n", key, value)
	}

	// If no name is provided in the HTTP request body, throw an error
	if len(request.Body) < 1 {
		return Response{StatusCode: 400}, nil
	}

	stackRequest := &StackRequest{}

	err := json.Unmarshal([]byte(request.Body), stackRequest)

	if err != nil {
		return Response{StatusCode: 400}, err
	}

	echo, err := json.Marshal(stackRequest)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(echo),
		Headers: map[string]string{
			"Content-Type": "application/json",
			"X-eTraveli-goawsredrive-createstack-reply": "sta",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
