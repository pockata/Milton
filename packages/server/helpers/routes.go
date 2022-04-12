package helpers

import (
	"errors"
	"log"
	"net/http"

	"gorm.io/gorm"

	"milton/models"
)

type APIWrapperFunc func(http.ResponseWriter, *http.Request, models.DB)

func CreateAPIWrapHandler(db models.DB) func(APIWrapperFunc) http.HandlerFunc {
	return func(handler APIWrapperFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			handler(rw, r, db)
		}
	}
}

type CreateEntryResponse struct {
	Entry *interface{} `json:"entry"`
}

func CreateEntry(
	rw http.ResponseWriter,
	r *http.Request,
	db gorm.DB,
	entry interface{},
) {
	res := db.Create(entry)

	if res.Error != nil {
		ErrorResponse(rw, r, res.Error)
		return
	}

	SuccessResponse(rw, r, CreateEntryResponse{
		Entry: &entry,
	})
}

func DeleteEntry(
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

	if !CheckParams(ID) {
		ErrorResponse(rw, r, errors.New("Invalid request. Missing parameters"))
		return
	}

	res := db.Unscoped().Where("ID = ?", ID).Delete(entry)

	if res.Error != nil {
		ErrorResponse(rw, r, res.Error)
		return
	}

	SuccessResponse(rw, r, nil)
}
