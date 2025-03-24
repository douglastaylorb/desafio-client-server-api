package models

import (
	"time"

	"gorm.io/gorm"
)

type CotacaoResponse struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

type Cotacao struct {
	gorm.Model
	Bid       string
	Timestamp time.Time
}
