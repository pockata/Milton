package main

import (
	"errors"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type CreateEntryResponse struct {
	Entry *interface{} `json:"entry"`
}

func createEntry(
	rw http.ResponseWriter,
	r *http.Request,
	db gorm.DB,
	entry interface{},
) {
	res := db.Create(entry)

	if res.Error != nil {
		errorResponse(rw, res.Error)
		return
	}

	successResponse(rw, CreateEntryResponse{
		Entry: &entry,
	})
}

func deleteEntry(
	rw http.ResponseWriter,
	r *http.Request,
	db gorm.DB,
	entry interface{},
) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Error parsing form data", err)
		return
	}

	ID := r.Form.Get("ID")

	if !checkParams(ID) {
		errorResponse(rw, errors.New("Invalid request. Missing parameters"))
		return
	}

	res := db.Unscoped().Where("ID = ?", ID).Delete(entry)

	if res.Error != nil {
		errorResponse(rw, res.Error)
		return
	}

	successResponse(rw, nil)
}
