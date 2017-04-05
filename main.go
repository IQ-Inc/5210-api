package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"net/http"
)

const (
	Port = ":9000"
)

func main() {

	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-west-2"),
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

	log.Println("Found tables:")
	for _, table := range result.TableNames {
		log.Println(*table)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		NotImplemented(w)
	})

	http.HandleFunc("/registration", HandleRegistration)

	log.Printf("Starting server on localhost%s", Port)
	http.ListenAndServe(Port, nil)
}
