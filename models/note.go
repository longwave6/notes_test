package models

import "gopkg.in/mgo.v2/bson"

// Represents a note, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Note struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Title       string        `bson:"title" json:"title"`
	Content     string        `bson:"content" json:"content"`
	Users       []string      `bson:"users" json:"users"`
}

