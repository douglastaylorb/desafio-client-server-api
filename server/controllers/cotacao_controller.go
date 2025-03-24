package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/douglastaylorb/desafio-client-server-api/tree/main/server/services"
	"gorm.io/gorm"
)

type CotacaoController struct {
	db *gorm.DB
}

func NewCotacaoController(db *gorm.DB) *CotacaoController {
	return &CotacaoController{db: db}
}

func (c *CotacaoController) GetCotacao(w http.ResponseWriter, r *http.Request) {
	cotacaoService := services.NewCotacaoService(c.db)

	bid, err := cotacaoService.GetAndSaveCotacao(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"bid": bid})
}
