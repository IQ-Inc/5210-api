package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		// local DB is us-west-2
		Region: aws.String("us-west-2"),

		// CHANGE ME
		// if you are running against a different port
		Endpoint: aws.String("http://localhost:8000")})
	if err != nil {
		log.Println(err)
		return
	}
	dbSvc := dynamodb.New(sess)

	result, err := dbSvc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Tables:")
	for _, table := range result.TableNames {
		log.Println(">>", *table)
	}
}
