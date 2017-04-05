// go run utils/listandscan
//
// Performs a list of all tables, then scans all tables for all content.
// This will be really, really slow for big tables! Use during preliminary
// development.

package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

const (
	// Location of your server
	// For local DynamoDB, it's probably http://localhost:8000,
	// but you may need to change the port.
	DynamoDBEndpoint = "http://localhost:8000"

	// Region of your server instance
	// NOTE: local dev server of DynamoDB is "us-west-2"
	DynamoDBRegion = "us-west-2"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		// local DB is us-west-2
		Region: aws.String(DynamoDBRegion),

		// CHANGE ME
		// if you are running against a different port
		Endpoint: aws.String(DynamoDBEndpoint)})
	if err != nil {
		log.Println(err)
		return
	}
	db := dynamodb.New(sess)

	result, err := db.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Tables and scans...")
	for _, table := range result.TableNames {
		log.Println("Table", *table)

		scanner := dynamodb.ScanInput{
			TableName: table,
		}

		scan, err := db.Scan(&scanner)

		if nil != err {
			log.Printf(">> Error scanning table %s... continue", *table)
			continue
		}

		log.Printf(">> Count: %d", *scan.Count)
		log.Printf(">> Items: %v", scan.Items)
	}
}
