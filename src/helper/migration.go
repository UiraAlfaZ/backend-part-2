package helper

import (
	"backend/src/config"
	"backend/src/models"
)

func Migration() {
	config.DB.AutoMigrate(&models.Product{})
}
