package models

import (
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/danielgospodinow/Catter/service/catter-account-service/db"
)

const (
	tableName          = "accounts"
	tableReadCapacity  = 10
	tableWriteCapacity = 10
)

// Account is the model of an account object.
type Account struct {
	Name     string
	Email    string
	Password string
	Picture  string
}

// CreateAccount creates an Account instance.
func CreateAccount(acc Account) (Account, error) {
	accJSON, merr := dynamodbattribute.MarshalMap(acc)
	if merr != nil {
		fmt.Printf("Error while marshaling object %s\n", acc)
		return Account{}, merr
	}

	input := dynamodb.PutItemInput{
		Item:      accJSON,
		TableName: aws.String(tableName),
	}

	_, perr := db.GetDynamoClient().PutItem(&input)
	if perr != nil {
		fmt.Printf("Failed to create item '%s' in the database\n", acc)
		return Account{}, perr
	}

	return acc, nil
}

// GetAccount retrieves an Account instance by a given email.
func GetAccount(email string) (Account, error) {
	input := dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Email": {
				S: aws.String(email),
			},
		},
	}

	output, gerr := db.GetDynamoClient().GetItem(&input)
	if gerr != nil {
		fmt.Printf("There was an error retrieving item with Email '%s'\n", email)
		return Account{}, gerr
	}

	acc := Account{}

	merr := dynamodbattribute.UnmarshalMap(output.Item, &acc)
	if merr != nil {
		fmt.Printf("Failed to unmarshal map for Account with Email '%s'!\n", email)
		return acc, merr
	}

	return acc, nil
}

// UpdateAccount updates an Account instance by a given email.
func UpdateAccount(email string, acc Account) (Account, error) {
	return Account{}, errors.New("Method not implemented")
}

// DeleteAccount deletes an Account instance by a given email.
func DeleteAccount(email string) (Account, error) {
	acc, gerr := GetAccount(email)
	if gerr != nil {
		fmt.Printf("Failed to delete Account with Email '%s'!\n", email)
		return Account{}, gerr
	}

	input := dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Email": {
				S: aws.String(email),
			},
		},
	}

	_, derr := db.GetDynamoClient().DeleteItem(&input)
	if derr != nil {
		fmt.Printf("Failed to delete Account with Email '%s'\n", email)
		return Account{}, derr
	}

	return acc, nil
}

// InitAccountsTable initializes the Accounts table.
func InitAccountsTable() {
	_, err := db.GetDynamoClient().DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	})

	if err != nil {
		if _, ok := err.(*dynamodb.ResourceNotFoundException); ok {
			fmt.Printf("No such table '%s' exists!\n", tableName)
			fmt.Println("Creating one ...")

			db.GetDynamoClient().CreateTable(&dynamodb.CreateTableInput{
				AttributeDefinitions: []*dynamodb.AttributeDefinition{
					{
						AttributeName: aws.String("Email"),
						AttributeType: aws.String("S"),
					},
				},
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("Email"),
						KeyType:       aws.String("HASH"),
					},
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(tableReadCapacity),
					WriteCapacityUnits: aws.Int64(tableWriteCapacity),
				},
				TableName: aws.String(tableName),
			})
		}
	} else if err != nil {
		fmt.Println("Error while describing table!")
		log.Fatal(err)
	} else {
		fmt.Printf("Table '%s' exists and is ready to be used!\n", tableName)
	}
}
