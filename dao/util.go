package dao

import (
	"log"
	mgo "gopkg.in/mgo.v2"
)

//pointer for db client DBO
var db *mgo.Database

//struct for consuming config.toml later
type GeneralDAO struct {
	Server   string
	Database string
}

//const values for accessing mongo db collections
const (
	USERS = "users"
	NOTES = "notes"
)

// Establish a connection to database
func (m *GeneralDAO) Connect() {
  session, err := mgo.Dial(m.Server)
  if err != nil {
    log.Fatal(err)
  }
  db = session.DB(m.Database)
}

