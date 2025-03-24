package database

import (
	"github.com/douglastaylorb/desafio-client-server-api/tree/main/server/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("cotacoes.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Cotacao{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
