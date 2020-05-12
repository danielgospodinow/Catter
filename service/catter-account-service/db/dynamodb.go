package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/danielgospodinow/Catter/service/catter-account-service/constants"
)

var dynamodbClient *dynamodb.DynamoDB

// InitDynamoClient initializes the DynamoDB client.
func InitDynamoClient() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	dynamodbClient = dynamodb.New(sess, aws.NewConfig().WithRegion(constants.Region))
}

// GetDynamoClient returns the DynamoDB client.
func GetDynamoClient() *dynamodb.DynamoDB {
	return dynamodbClient
}

// CloseDynamoClient closes the DynamoDB client.
func CloseDynamoClient() {
	// TODO: No such functionality ? Ensure that there's no need to close this client.
}
