package main

import (
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DynamoMapper describes a type that can be converted into a
// DynamoDB Item
type DynamoMapper interface {
	DynamoMap() map[string]*dynamodb.AttributeValue
	GetTable() string
}

// Registration metrics captured during signup
type Registration struct {
	// Zipcode of user
	Zipcode int `json:"zipcode"`

	// HouseholdSize size of family
	HouseholdSize int `json:"householdsize"`

	// Ages of child / children
	ChildAges []int `json:"childages"`

	// Unique identifier of the user
	ID string `json:"ID"`

	// timestamp not exposed
	timestamp time.Time

	// table DynamoDB table not exposed
	table string
}

// DynamoMap map registration values to a DynamoDB item
// Change tags here
func (r *Registration) DynamoMap() map[string]*dynamodb.AttributeValue {

	return map[string]*dynamodb.AttributeValue{
		"Zipcode":       {N: aws.String(strconv.Itoa(r.Zipcode))},
		"Id":            {S: aws.String(r.ID)},
		"Householdsize": {N: aws.String(strconv.Itoa(r.HouseholdSize))},
		"TimestampUTC":  {S: aws.String(r.timestamp.UTC().String())},
		// TODO stringify slice of ints
	}
}

// GetTable get the registration table
func (r *Registration) GetTable() string {
	return r.table
}

// Progress data message to update the user's
// star count for a trackable category
type Progress struct {
	// Stars5 number of stars for fruits & veggies
	Stars5 int `json:"stars5"`

	// Stars2 number if stars for screen time
	Stars2 int `json:"stars2"`

	// Stars1 number of stars for physical activity
	Stars1 int `json:"stars1"`

	// Stars0 number of stars for screen time
	Stars0 int `json:"start0"`

	// Associated user
	ID string `json:"ID"`

	// timestamp not exposed
	timestamp time.Time

	// table DynamoDB table not exposed
	table string
}

// DynamoMap map progress values to DynamoDB item
func (p *Progress) DynamoMap() map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		"Stars5":       {N: aws.String(strconv.Itoa(p.Stars5))},
		"Stars2":       {N: aws.String(strconv.Itoa(p.Stars2))},
		"Stars1":       {N: aws.String(strconv.Itoa(p.Stars1))},
		"Stars0":       {N: aws.String(strconv.Itoa(p.Stars0))},
		"Id":           {S: aws.String(p.ID)},
		"TimestampUTC": {S: aws.String(p.timestamp.UTC().String())},
	}
}

// GetTable get the table with the associated
func (p *Progress) GetTable() string {
	return p.table
}
