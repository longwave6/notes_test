package dao

import (
	"gopkg.in/mgo.v2/bson"
	. "notes_test/models"
)

type UsersDAO struct {
	Server   string
	Database string
}

// Find list of users
func (m *UsersDAO) FindAll() ([]User, error) {
	var users []User
	err := db.C(USERS).Find(bson.M{}).All(&users)
	return users, err
}

// Find a user by its id
func (m *UsersDAO) FindById(id string) (User, error) {
	var user User
	err := db.C(USERS).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// Insert a user into database
func (m *UsersDAO) Insert(user User) error {
	err := db.C(USERS).Insert(&user)
	return err
}

// Delete an existing user
func (m *UsersDAO) Delete(user User) error {
	err := db.C(USERS).Remove(&user)
	return err
}

// Update an existing user
func (m *UsersDAO) Update(user User) error {
	err := db.C(USERS).UpdateId(user.ID, &user)
	return err
}
