package main

import (
	"log"
	"os"
	"path/filepath"

	dynamodb "github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	iam "github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	lambda "github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	logs "github.com/aws/aws-cdk-go/awscdk/v2/awslogs"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"

	awscdk "github.com/aws/aws-cdk-go/awscdk/v2"
)

type HelloLambdaCdkGoStackProps struct {
	awscdk.StackProps
}

func NewHelloLambdaCdkGoStack(
	scope constructs.Construct, id string, props *HelloLambdaCdkGoStackProps,
) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// ===== DynamoDB table =====
	tableName := "Messages"
	tableId := "DynamoDBTable"
	partitionKeyName := "ID"

	messagesTable := dynamodb.NewTable(
		stack,
		aws.String(tableId),
		&dynamodb.TableProps{
			TableName: aws.String(tableName),
			PartitionKey: &dynamodb.Attribute{
				Name: aws.String(partitionKeyName),
				Type: dynamodb.AttributeType_STRING,
			},
			BillingMode: dynamodb.BillingMode_PAY_PER_REQUEST,
			TableClass:  dynamodb.TableClass_STANDARD,
		},
	)

	// ===== IAM =====
	dynamodbActions := aws.StringSlice([]string{
		"dynamodb:PutItem",
		"dynamodb:UpdateItem",
		"dynamodb:GetItem",
	})
	resources := []*string{messagesTable.TableArn()}
	lambdaMessagesTableStatement := iam.NewPolicyStatement(&iam.PolicyStatementProps{
		Effect:    iam.Effect_ALLOW,
		Actions:   &dynamodbActions,
		Resources: &resources,
	})

	// ===== Lambda =====
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	lambdaPath := filepath.Join(path, "../dist/main.zip")

	lambdaName := "hello-go-lambda"
	helloLambda := lambda.NewFunction(
		stack,
		aws.String(lambdaName),
		&lambda.FunctionProps{
			Description:  aws.String("Lambda example, it's alive!!"),
			FunctionName: aws.String(lambdaName),
			LogRetention: logs.RetentionDays_ONE_DAY,
			MemorySize:   aws.Float64(128),
			Timeout:      awscdk.Duration_Seconds(aws.Float64(10)),
			Code:         lambda.Code_FromAsset(&lambdaPath, &awss3assets.AssetOptions{}),
			Handler:      aws.String("main"),
			Runtime:      lambda.Runtime_GO_1_X(),
			Environment:  &map[string]*string{"TABLE": aws.String(tableName)},
		},
	)
	helloLambda.AddToRolePolicy(lambdaMessagesTableStatement)

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewHelloLambdaCdkGoStack(app, "HelloLambdaCdkGoStack", &HelloLambdaCdkGoStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return nil
}
