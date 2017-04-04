package main

type TypeUUID string

// Registration metrics captured during signup
type Registration struct {
	// Zipcode of user
	Zipcode int `json:"zipcode"`

	// HouseholdSize size of family
	HouseholdSize int `json:"householdsize"`

	// Ages of child / children
	ChildAges []int `json:"childages"`

	// Unique identifier of the user
	UUID TypeUUID `json:"uuid"`

	// Datetime registration timestamp
	Datetime string `json:"datetime"`
}

// Progress data message to update the user's
// star count for a trackable category
type Progress struct {
	// Number of stars
	Stars int `json:"stars"`

	// Associated user
	UUID TypeUUID `json:"uuid"`

	// Datetime event timestamp
	Datetime string `json:"datetime"`
}
