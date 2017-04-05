package main

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"net/http"
)

const (
	// Server port
	Port = ":9000"

	// Context key of the database
	ContextKeyDB = "db"
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

	db := dynamodb.New(sess)

	ctx := context.WithValue(context.Background(), ContextKeyDB, db)

	// Root endpoint not implemented
	http.HandleFunc("/", NotImplemented)

	// Add endpoints here
	http.Handle("/registration", &ContextAdapter{
		ctx,
		ContextHandlerFunc(HandleRegistration),
	})

	log.Printf("Starting server on localhost%s", Port)
	http.ListenAndServe(Port, nil)
}
