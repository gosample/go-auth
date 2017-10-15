package models

import "gopkg.in/mgo.v2/bson"

const (
	// CollectionBook holds the name of the articles collection
	CollectionUser = "users"
)

// User model
type User struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreatedOn int64         `json:"created_on" bson:"created_od"`
	UpdatedOn int64         `json:"updated_on" bson:"updated_on"`
}