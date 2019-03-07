package dao

import (
	. "longwave_api/models"
	"gopkg.in/mgo.v2/bson"
)

type NotesDAO struct {
	Server   string
	Database string
}

// Find list of notes
func (m *NotesDAO) FindAll() ([]Note, error) {
	var notes []Note
	err := db.C(NOTES).Find(bson.M{}).All(&notes)
	return notes, err
}

// Find a note by its id
func (m *NotesDAO) FindById(id string) (Note, error) {
	var note Note
	err := db.C(NOTES).FindId(bson.ObjectIdHex(id)).One(&note)
	return note, err
}

// Insert a note into database
func (m *NotesDAO) Insert(note Note) error {
	err := db.C(NOTES).Insert(&note)
	return err
}

// Delete an existing note
func (m *NotesDAO) Delete(note Note) error {
	err := db.C(NOTES).Remove(&note)
	return err
}

// Update an existing note
func (m *NotesDAO) Update(note Note) error {
	err := db.C(NOTES).UpdateId(note.ID, &note)
	return err
}
