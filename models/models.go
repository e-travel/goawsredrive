package models

type StackConfiguration struct {
	SqsArn   string `json:"sqs_arn"`
	HttpVerb string `json:"http_verb"`
	Endpoint string `json:"endpoint"`
}

type RedriveBasicTemplate struct {
	AWSTemplateFormatVersion string `yaml:"AWSTemplateFormatVersion"`
	Description              string `yaml:"Description"`
	Parameters               struct {
		ResourcesName struct {
			Type                  string `yaml:"Type"`
			Description           string `yaml:"Description"`
			MinLength             string `yaml:"MinLength"`
			MaxLength             string `yaml:"MaxLength"`
			AllowedPattern        string `yaml:"AllowedPattern"`
			ConstraintDescription string `yaml:"ConstraintDescription"`
		} `yaml:"ResourcesName"`
		AlarmSubscriber struct {
			Type        string `yaml:"Type"`
			Description string `yaml:"Description"`
		} `yaml:"AlarmSubscriber"`
		Environment struct {
			Type          string   `yaml:"Type"`
			Description   string   `yaml:"Description"`
			Default       string   `yaml:"Default"`
			AllowedValues []string `yaml:"AllowedValues"`
		} `yaml:"Environment"`
		SNSSubscriptionRawMessageDelivery struct {
			Type          string   `yaml:"Type"`
			Description   string   `yaml:"Description"`
			Default       string   `yaml:"Default"`
			AllowedValues []string `yaml:"AllowedValues"`
		} `yaml:"SNSSubscriptionRawMessageDelivery"`
		SQSQueueMessageRetentionPeriod struct {
			Type                  string `yaml:"Type"`
			Description           string `yaml:"Description"`
			Default               int    `yaml:"Default"`
			MinValue              int    `yaml:"MinValue"`
			MaxValue              int    `yaml:"MaxValue"`
			ConstraintDescription string `yaml:"ConstraintDescription"`
		} `yaml:"SQSQueueMessageRetentionPeriod"`
		SQSQueueVisibilityTimeout struct {
			Type                  string `yaml:"Type"`
			Description           string `yaml:"Description"`
			Default               int    `yaml:"Default"`
			MinValue              int    `yaml:"MinValue"`
			MaxValue              int    `yaml:"MaxValue"`
			ConstraintDescription string `yaml:"ConstraintDescription"`
		} `yaml:"SQSQueueVisibilityTimeout"`
		SQSQueueReceiveMessageWaitTimeSeconds struct {
			Type                  string `yaml:"Type"`
			Description           string `yaml:"Description"`
			Default               int    `yaml:"Default"`
			MinValue              int    `yaml:"MinValue"`
			MaxValue              int    `yaml:"MaxValue"`
			ConstraintDescription string `yaml:"ConstraintDescription"`
		} `yaml:"SQSQueueReceiveMessageWaitTimeSeconds"`
		SQSQueueMaximumMessageSize struct {
			Type                  string `yaml:"Type"`
			Description           string `yaml:"Description"`
			Default               int    `yaml:"Default"`
			MinValue              int    `yaml:"MinValue"`
			MaxValue              int    `yaml:"MaxValue"`
			ConstraintDescription string `yaml:"ConstraintDescription"`
		} `yaml:"SQSQueueMaximumMessageSize"`
		SQSQueueDelaySeconds struct {
			Type                  string `yaml:"Type"`
			Description           string `yaml:"Description"`
			Default               int    `yaml:"Default"`
			MinValue              int    `yaml:"MinValue"`
			MaxValue              int    `yaml:"MaxValue"`
			ConstraintDescription string `yaml:"ConstraintDescription"`
		} `yaml:"SQSQueueDelaySeconds"`
		SQSQueueRedrivePolicyMaxReceiveCount struct {
			Type                  string `yaml:"Type"`
			Description           string `yaml:"Description"`
			Default               int    `yaml:"Default"`
			MinValue              int    `yaml:"MinValue"`
			MaxValue              int    `yaml:"MaxValue"`
			ConstraintDescription string `yaml:"ConstraintDescription"`
		} `yaml:"SQSQueueRedrivePolicyMaxReceiveCount"`
		SingleStack struct {
			Type          string   `yaml:"Type"`
			Description   string   `yaml:"Description"`
			Default       string   `yaml:"Default"`
			AllowedValues []string `yaml:"AllowedValues"`
		} `yaml:"SingleStack"`
	} `yaml:"Parameters"`
	Conditions struct {
		Production  interface{} `yaml:"Production"`
		SingleStack interface{} `yaml:"SingleStack"`
	} `yaml:"Conditions"`
	Resources struct {
		Topic struct {
			Type       string `yaml:"Type"`
			Properties struct {
				TopicName interface{} `yaml:"TopicName"`
			} `yaml:"Properties"`
		} `yaml:"Topic"`
		QueueSubscription struct {
			Type       string `yaml:"Type"`
			Properties struct {
				Endpoint           interface{} `yaml:"Endpoint"`
				Protocol           string      `yaml:"Protocol"`
				RawMessageDelivery interface{} `yaml:"RawMessageDelivery"`
				TopicArn           interface{} `yaml:"TopicArn"`
			} `yaml:"Properties"`
		} `yaml:"QueueSubscription"`
		Queue struct {
			Type       string `yaml:"Type"`
			Properties struct {
				QueueName                     interface{} `yaml:"QueueName"`
				MessageRetentionPeriod        interface{} `yaml:"MessageRetentionPeriod"`
				VisibilityTimeout             interface{} `yaml:"VisibilityTimeout"`
				ReceiveMessageWaitTimeSeconds interface{} `yaml:"ReceiveMessageWaitTimeSeconds"`
				MaximumMessageSize            interface{} `yaml:"MaximumMessageSize"`
				DelaySeconds                  interface{} `yaml:"DelaySeconds"`
				RedrivePolicy                 struct {
					DeadLetterTargetArn interface{} `yaml:"deadLetterTargetArn"`
					MaxReceiveCount     interface{} `yaml:"maxReceiveCount"`
				} `yaml:"RedrivePolicy"`
			} `yaml:"Properties"`
		} `yaml:"Queue"`
		DeadLetterQueue struct {
			Type       string `yaml:"Type"`
			Properties struct {
				QueueName interface{} `yaml:"QueueName"`
			} `yaml:"Properties"`
		} `yaml:"DeadLetterQueue"`
		QueuePolicy struct {
			Type       string `yaml:"Type"`
			Properties struct {
				PolicyDocument struct {
					Statement []struct {
						Sid       string `yaml:"Sid"`
						Effect    string `yaml:"Effect"`
						Principal struct {
							AWS string `yaml:"AWS"`
						} `yaml:"Principal"`
						Action    string `yaml:"Action"`
						Resource  string `yaml:"Resource"`
						Condition struct {
							ArnEquals struct {
								AwsSourceArn interface{} `yaml:"aws:SourceArn"`
							} `yaml:"ArnEquals"`
						} `yaml:"Condition"`
					} `yaml:"Statement"`
				} `yaml:"PolicyDocument"`
				Queues []interface{} `yaml:"Queues"`
			} `yaml:"Properties"`
		} `yaml:"QueuePolicy"`
		AccessSQS struct {
			Condition  string `yaml:"Condition"`
			Type       string `yaml:"Type"`
			Properties struct {
				Path           string `yaml:"Path"`
				PolicyDocument struct {
					Version struct {
					} `yaml:"Version"`
					Statement struct {
						Effect    string        `yaml:"Effect"`
						Action    []string      `yaml:"Action"`
						Resource  []interface{} `yaml:"Resource"`
						Condition struct {
							IPAddress struct {
								AwsSourceIP []string `yaml:"aws:SourceIp"`
							} `yaml:"IpAddress"`
						} `yaml:"Condition"`
					} `yaml:"Statement"`
				} `yaml:"PolicyDocument"`
			} `yaml:"Properties"`
		} `yaml:"AccessSQS"`
		AccessSNS struct {
			Condition  string `yaml:"Condition"`
			Type       string `yaml:"Type"`
			Properties struct {
				Path           string `yaml:"Path"`
				PolicyDocument struct {
					Version struct {
					} `yaml:"Version"`
					Statement struct {
						Effect    string      `yaml:"Effect"`
						Action    []string    `yaml:"Action"`
						Resource  interface{} `yaml:"Resource"`
						Condition struct {
							IPAddress struct {
								AwsSourceIP []string `yaml:"aws:SourceIp"`
							} `yaml:"IpAddress"`
						} `yaml:"Condition"`
					} `yaml:"Statement"`
				} `yaml:"PolicyDocument"`
			} `yaml:"Properties"`
		} `yaml:"AccessSNS"`
		DeadLetterQueueAlarm struct {
			Type       string `yaml:"Type"`
			Condition  string `yaml:"Condition"`
			Properties struct {
				AlarmDescription   string        `yaml:"AlarmDescription"`
				AlarmActions       []interface{} `yaml:"AlarmActions"`
				OKActions          []interface{} `yaml:"OKActions"`
				MetricName         string        `yaml:"MetricName"`
				Namespace          string        `yaml:"Namespace"`
				Statistic          string        `yaml:"Statistic"`
				Period             string        `yaml:"Period"`
				EvaluationPeriods  string        `yaml:"EvaluationPeriods"`
				Threshold          string        `yaml:"Threshold"`
				ComparisonOperator string        `yaml:"ComparisonOperator"`
				Dimensions         []struct {
					Name  string      `yaml:"Name"`
					Value interface{} `yaml:"Value"`
				} `yaml:"Dimensions"`
				TreatMissingData string `yaml:"TreatMissingData"`
			} `yaml:"Properties"`
		} `yaml:"DeadLetterQueueAlarm"`
	} `yaml:"Resources"`
	Outputs struct {
		TopicName struct {
			Value       interface{} `yaml:"Value"`
			Description string      `yaml:"Description"`
		} `yaml:"TopicName"`
		QueueName struct {
			Value       interface{} `yaml:"Value"`
			Description string      `yaml:"Description"`
		} `yaml:"QueueName"`
		QueueURL struct {
			Value       interface{} `yaml:"Value"`
			Description string      `yaml:"Description"`
		} `yaml:"QueueUrl"`
		DeadLetterQueueName struct {
			Value       interface{} `yaml:"Value"`
			Description string      `yaml:"Description"`
		} `yaml:"DeadLetterQueueName"`
		TopicARN struct {
			Value       interface{} `yaml:"Value"`
			Description string      `yaml:"Description"`
		} `yaml:"TopicARN"`
		QueueARN struct {
			Value       interface{} `yaml:"Value"`
			Description string      `yaml:"Description"`
		} `yaml:"QueueARN"`
		DeadLetterQueueARN struct {
			Value       interface{} `yaml:"Value"`
			Description string      `yaml:"Description"`
		} `yaml:"DeadLetterQueueARN"`
		AccessSQSPolicyARN struct {
			Condition   string      `yaml:"Condition"`
			Value       interface{} `yaml:"Value"`
			Description string      `yaml:"Description"`
		} `yaml:"AccessSQSPolicyARN"`
		AccessSNSPolicyARN struct {
			Condition   string      `yaml:"Condition"`
			Value       interface{} `yaml:"Value"`
			Description string      `yaml:"Description"`
		} `yaml:"AccessSNSPolicyARN"`
	} `yaml:"Outputs"`
}
