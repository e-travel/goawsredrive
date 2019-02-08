package models

type StackConfiguration struct {
	SqsArn   string `json:"sqs_arn"`
	HttpVerb string `json:"http_verb"`
	Endpoint string `json:"endpoint"`
}
