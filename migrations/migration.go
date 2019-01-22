package migrations

import (
	"github.com/jorgeAM/error/lib"
	"github.com/jorgeAM/error/models"
)

// Migrate -> migrate db
func Migrate() {
	db := lib.GetConnection()
	defer db.Close()
	db.CreateTable(&models.Error{})
}
