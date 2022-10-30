package config

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/sqs"
	"os"
	"testing"
)

func TestNewSQSQueue(t *testing.T) {
	session, err := GetSessionConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	SQSQueueImpl := NewSQSQueue(session)
	result, err := SQSQueueImpl.CreateQueue("test2")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
}

func TestSQSQueueImpl_ListQueues(t *testing.T) {
	session, err := GetSessionConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	client := sqs.New(session)
	result, err := ListQueues(client)
	if err != nil {
		fmt.Println(err.Error())
	}
	for i, url := range result.QueueUrls {
		fmt.Printf("%d: %s\n", i, *url)
	}
}
func TestGetQueueName(t *testing.T) {
	session, err := GetSessionConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	client := sqs.New(session)
	result, err := GetQueueName(client, "test2")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(*result.QueueUrl)
}

func TestSendMessage(t *testing.T) {
	session, err := GetSessionConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	client := sqs.New(session)
	result, err := GetQueueName(client, "test2")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(*result.QueueUrl)
	message := "00eisjdwewrsdfsks844"
	err = SendMessage(client, &message, result.QueueUrl)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestReceiveMessage(t *testing.T) {
	session, err := GetSessionConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	client := sqs.New(session)
	result, err := GetQueueName(client, "test2")
	if err != nil {
		fmt.Println(err.Error())
	}
	results, err := ReceiveMessage(client, result.QueueUrl)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(results.Messages[0])
	// removing message from queue

	//err = RemoveMessageFromQueue(client, result.QueueUrl, results.Messages[0].ReceiptHandle)
	//fmt.Println(err)
}
