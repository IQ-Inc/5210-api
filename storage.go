package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Store uses a database context to store a DynamoMapper data item
// Panics if the database was not found in the context.
func Store(ctx context.Context, data DynamoMapper) error {

	log.Printf("Storing %v", data)

	cmd := &dynamodb.PutItemInput{
		TableName: aws.String(data.GetTable()),
		Item:      data.DynamoMap(),
	}

	db, ok := ctx.Value(ContextKeyDB).(*dynamodb.DynamoDB)
	if !ok {
		log.Fatal("Could not obtain DB from context!")
	}

	_, err := db.PutItem(cmd)
	return err
}
