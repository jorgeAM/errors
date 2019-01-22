package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jorgeAM/error/lib"
	"github.com/jorgeAM/error/models"
)

// GetError -> get error by code
func GetError(w http.ResponseWriter, r *http.Request) {
	myError := models.Error{}
	params := mux.Vars(r)
	db := lib.GetConnection()
	defer db.Close()
	db.Where("code=?", params["code"]).First(&myError)
	j, err := json.Marshal(&myError)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("hubo un error."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// NewError -> create new error
func NewError(w http.ResponseWriter, r *http.Request) {
	myError := models.Error{}
	err := json.NewDecoder(r.Body).Decode(&myError)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error al leer error."))
		return
	}

	db := lib.GetConnection()
	defer db.Close()
	err = db.Create(&myError).Error
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error al registrar error."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Todo salio bien."))
}

// SaveErrorsFromFile -> save errors from json file
func SaveErrorsFromFile(w http.ResponseWriter, r *http.Request) {
	file, handle, err := r.FormFile("file")
	if err != nil {
		log.Fatal("Hubo un error al leer el archivo: ", err)
		return
	}

	defer file.Close()
	mineType := handle.Header.Get("Content-type")
	if mineType == "application/json" {
		fmt.Println("Es un archivo json")
		return
	}
	fmt.Println("archivo invalido")
}
