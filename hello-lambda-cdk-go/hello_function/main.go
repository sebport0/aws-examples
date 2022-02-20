package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/google/uuid"
)

type MyEvent struct {
	Message string `json:"message"`
}

type MessageRecord struct {
	ID      string
	MESSAGE string
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	// New dynamoDB client.
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("Unable to load SDK config, " + err.Error())
	}
	dynamoClient := dynamodb.NewFromConfig(cfg)

	tableName := os.Getenv("TABLE")
	messageRecord := MessageRecord{
		ID:      uuid.NewString(),
		MESSAGE: event.Message,
	}
	item, err := attributevalue.MarshalMap(messageRecord)
	if err != nil {
		panic(fmt.Sprintf("Something went wrong with DynamoDB marshal %v", err))
	}

	_, err = dynamoClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	})
	if err != nil {
		panic(fmt.Sprintf("Something went wrong with DynamoDB putItem operation, %v", err))
	}

	return "Ok", nil
}

func main() {
	lambda.Start(HandleRequest)
}
