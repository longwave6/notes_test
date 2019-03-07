package routes

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	. "notes_test/dao"
	. "notes_test/models"
)

var notesDao = NotesDAO{}

// GET list of notes
func AllNotesEndPoint(w http.ResponseWriter, r *http.Request) {
	notes, err := notesDao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, notes)
}

// GET a note by its ID
func FindNoteEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	notes, err := notesDao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Note ID")
		return
	}
	respondWithJson(w, http.StatusOK, notes)
}

// POST a new notes
func CreateNoteEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var note Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	note.ID = bson.NewObjectId()
	if err := notesDao.Insert(note); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, note)
}

// PUT update an existing note
func UpdateNoteEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var note Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := notesDao.Update(note); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing note
func DeleteNoteEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var note Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := notesDao.Delete(note); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
