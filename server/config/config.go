package config

import "time"

const (
	ServerPort    = 8080
	CotacaoAPIURL = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	APITimeout    = 200 * time.Millisecond
	DBTimeout     = 10 * time.Millisecond
)
