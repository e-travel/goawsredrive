package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

type StackRequest struct {
	Name                     string `json:"name"`
	HttpVerb                 string `json:"http_verb"`
	Endpoint                 string `json:"endpoint"`
	MessageVisibilityTimeout string `json:"message_visibility_timeout"`
	AlarmSubscriber          string `json:"alarm_subscriber"`
	TemplateURL              string `json:"template_url"`
	StackId                  string `json:"new_stack_id"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	log.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	log.Printf("Body size = %d.\n", len(request.Body))

	// If no name is provided in the HTTP request body, throw an error
	if len(request.Body) < 1 {
		return Response{StatusCode: 400}, nil
	}

	stackRequest := &StackRequest{}

	err := json.Unmarshal([]byte(request.Body), stackRequest)

	if err != nil {
		return Response{StatusCode: 400}, err
	}

	if len(stackRequest.Endpoint) < 1 || len(stackRequest.HttpVerb) < 1 || len(stackRequest.Name) < 1 {
		return Response{StatusCode: 400}, nil
	}

	createStackInput := &cloudformation.CreateStackInput{
		StackName:   aws.String(stackRequest.Name),
		TemplateURL: aws.String(stackRequest.TemplateURL),
		Parameters: []*cloudformation.Parameter{
			&cloudformation.Parameter{
				ParameterKey:   aws.String("ResourcesName"),
				ParameterValue: aws.String(stackRequest.Name),
			},
			&cloudformation.Parameter{
				ParameterKey:   aws.String("SQSQueueVisibilityTimeout"),
				ParameterValue: aws.String(stackRequest.MessageVisibilityTimeout),
			},
			&cloudformation.Parameter{
				ParameterKey:   aws.String("SingleStack"),
				ParameterValue: aws.String("false"),
			},
			&cloudformation.Parameter{
				ParameterKey:   aws.String("AlarmSubscriber"),
				ParameterValue: aws.String(stackRequest.AlarmSubscriber),
			},
		},
	}

	// create aws session
	sess := session.Must(session.NewSession())

	client := cloudformation.New(sess)

	createStackOutput, err := client.CreateStack(createStackInput)

	if err != nil {
		log.Fatal(err)
		return Response{StatusCode: 500}, nil
	}

	stackRequest.StackId = *createStackOutput.StackId

	describeStacksInput := &cloudformation.DescribeStacksInput{
		StackName: aws.String(stackRequest.Name),
	}

	err = client.WaitUntilStackExists(describeStacksInput)

	if err != nil {
		log.Fatal(err)
		return Response{StatusCode: 500}, nil
	}

	log.Println("Stack created: " + stackRequest.StackId)

	describeStackResourceInput := &cloudformation.DescribeStackResourceInput{
		LogicalResourceId: aws.String("Queue"),
		StackName:         aws.String(stackRequest.Name),
	}

	describeStackResourceOutput, err := client.DescribeStackResource(describeStackResourceInput)

	if err != nil {
		log.Fatal(err)
		return Response{StatusCode: 500}, nil
	}

	log.Println(describeStackResourceOutput.StackResourceDetail.PhysicalResourceId)

	echo, err := json.Marshal(stackRequest)

	resp := Response{
		StatusCode:      201,
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
