package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"strconv"
)

// DynamoMapper describes a type that can be converted into a
// DynamoDB Item
type DynamoMapper interface {
	DynamoMap() map[string]*dynamodb.AttributeValue
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
	UUID string `json:"uuid"`

	// Datetime registration timestamp
	Datetime string `json:"datetime"`
}

// DynamoMap map registration values to a DynamoDB item
// Change tags here
func (r *Registration) DynamoMap() map[string]*dynamodb.AttributeValue {

	return map[string]*dynamodb.AttributeValue{
		"Zipcode":       {N: aws.String(strconv.Itoa(r.Zipcode))},
		"Id":            {S: aws.String(r.UUID)},
		"Householdsize": {N: aws.String(strconv.Itoa(r.HouseholdSize))},
	}
}

// Progress data message to update the user's
// star count for a trackable category
type Progress struct {
	// Number of stars
	Stars int `json:"stars"`

	// Associated user
	UUID string `json:"uuid"`

	// Datetime event timestamp
	Datetime string `json:"datetime"`
}
