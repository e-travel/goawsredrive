package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/e-travel/goawsredrive/models"
)

var s3bucket = "TODO://"

var client = &http.Client{
	// TODO: Revise/parameterize timeout settings
	Timeout: 10 * time.Second,
}

var sess = session.Must(session.NewSession())
var s3client = s3.New(sess)

func Handler(ctx context.Context, event events.SQSEvent) error {
	var err error
	for _, msg := range event.Records {
		err = ProcessMessage(msg)
		if err != nil {
			fmt.Printf(err.Error())
			return err
		}
	}
	return err
}

func ProcessMessage(message events.SQSMessage) error {
	// load config from s3
	input := &s3.GetObjectInput{
		Bucket: aws.String(s3bucket),
		Key:    aws.String(message.EventSourceARN),
	}
	result, err := s3client.GetObject(input)
	if err != nil {
		return err
	}
	// deserialize config
	config := models.StackConfiguration{}
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &config)
	if err != nil {
		return err
	}

	// good! now, send the message to the application
	request, _ := http.NewRequest(config.HttpVerb, config.Endpoint, strings.NewReader(message.Body))
	response, err := client.Do(request)
	if response.StatusCode >= 300 {
		return fmt.Errorf("%s %s code:%d error:%s",
			config.HttpVerb, config.Endpoint, response.StatusCode, err.Error())
	}
	return nil
}

func main() {
	lambda.Start(Handler)
}
