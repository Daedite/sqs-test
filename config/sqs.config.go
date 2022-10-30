package config

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQSQueueImpl struct {
	Session   *session.Session
	SqsClient *sqs.SQS
	QueueName string
	QueueUrl  *string
}

func NewSQSQueue(Session *session.Session) *SQSQueueImpl {
	return &SQSQueueImpl{
		Session: Session,
	}
}

func GetSessionConfig() (*session.Session, error) {
	result, err := session.NewSession(&aws.Config{
		Region:                        aws.String("us-west-2"),
		Endpoint:                      aws.String("http://localhost:4566"),
		CredentialsChainVerboseErrors: aws.Bool(true),
	})
	if err != nil {
		fmt.Println(err.Error())
		return result, err
	}
	client := sqs.New(result)
	sqsImpl := SQSQueueImpl{
		Session:   result,
		SqsClient: client,
	}
	return sqsImpl.Session, err
}

func (s *SQSQueueImpl) CreateQueue(name string) (*sqs.CreateQueueOutput, error) {
	//creating a new session.
	s.SqsClient = sqs.New(s.Session)
	s.QueueName = name
	////Now we proceed with queue creation.
	//queue := flag.String("q", "", name)
	//flag.Parse()

	svc := sqs.New(s.Session)

	result, err := svc.CreateQueue(&sqs.CreateQueueInput{
		QueueName: &name,
		Attributes: map[string]*string{
			"DelaySeconds":           aws.String("0"),
			"MessageRetentionPeriod": aws.String("86400"),
			"VisibilityTimeout":      aws.String("10"),
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return result, err
	}
	s.QueueUrl = result.QueueUrl
	return result, nil
}

func RemoveMessageFromQueue(s *sqs.SQS, queueUrl, receiverHandle *string) error {
	_, err := s.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      queueUrl,
		ReceiptHandle: receiverHandle,
	})
	if err != nil {
		return err
	}
	return nil
}

func ListQueues(s *sqs.SQS) (*sqs.ListQueuesOutput, error) {
	result, err := s.ListQueues(&sqs.ListQueuesInput{
		MaxResults:      aws.Int64(10),
		NextToken:       aws.String(""),
		QueueNamePrefix: aws.String(""),
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return result, nil
	//to view the ListQueueOutput do:
	//for i, url := range result.QueueUrls {
	//	fmt.Printf("%d: %s\n", i, *url)
	//}
}
func SendMessage(s *sqs.SQS, payload, queueUrl *string) error {
	_, err := s.SendMessage(&sqs.SendMessageInput{
		MessageBody: payload,
		QueueUrl:    queueUrl,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *SQSQueueImpl) DeleteQueue(queueURL string) error {
	_, err := s.SqsClient.DeleteQueue(&sqs.DeleteQueueInput{
		QueueUrl: &queueURL,
	})
	return err
}

func GetQueueName(s *sqs.SQS, queueName string) (*sqs.GetQueueUrlOutput, error) {
	result, err := s.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})
	if err != nil {
		fmt.Println(err.Error())
		return result, err
	}
	return result, nil
}

func PurgeQueue(s *sqs.SQS, url *string) error {
	params := &sqs.PurgeQueueInput{
		QueueUrl: url,
	}
	_, err := s.PurgeQueue(params)

	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func ReceiveMessage(s *sqs.SQS, url *string) (*sqs.ReceiveMessageOutput, error) {
	params := &sqs.ReceiveMessageInput{
		QueueUrl: url,
	}
	resp, err := s.ReceiveMessage(params)

	if err != nil {
		fmt.Println(err.Error())
		return resp, err
	}
	return resp, nil
}

func (s *SQSQueueImpl) GetUrl() *string {
	return s.QueueUrl
}

func (s *SQSQueueImpl) GetName() string {
	return s.QueueName
}

func (s *SQSQueueImpl) GetSession() *session.Session {
	return s.Session
}

func (s *SQSQueueImpl) GetSQSClient() *sqs.SQS {
	return s.SqsClient
}
