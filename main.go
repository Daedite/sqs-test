package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"os"
	"sqs-test/config"
)

func main() {
	session, err := config.GetSessionConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	client := sqs.New(session)
	resp, err := config.GetQueueName(client, "test2")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for {
		output, err := client.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            resp.QueueUrl,
			MaxNumberOfMessages: aws.Int64(1),
			WaitTimeSeconds:     aws.Int64(3000),
		})

		if err != nil {
			fmt.Printf("failed to fetch sqs message %v\n", err)
		}
		fmt.Println(output.Messages[0])
		if output.Messages[0].String() != "" {
			err = config.RemoveMessageFromQueue(client, resp.QueueUrl, output.Messages[0].ReceiptHandle)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
