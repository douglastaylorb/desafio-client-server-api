package services

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/douglastaylorb/desafio-client-server-api/tree/main/server/config"
	"github.com/douglastaylorb/desafio-client-server-api/tree/main/server/models"
	"gorm.io/gorm"
)

type CotacaoService struct {
	db *gorm.DB
}

func NewCotacaoService(db *gorm.DB) *CotacaoService {
	return &CotacaoService{db: db}
}

func (s *CotacaoService) GetAndSaveCotacao(ctx context.Context) (string, error) {
	cotacaoCtx, cancel := context.WithTimeout(ctx, config.APITimeout)
	defer cancel()

	cotacao, err := s.getCotacao(cotacaoCtx)
	if err != nil {
		return "", err
	}

	dbCtx, cancelDB := context.WithTimeout(context.Background(), config.DBTimeout)
	defer cancelDB()

	err = s.saveCotacao(dbCtx, cotacao.USDBRL.Bid)
	if err != nil {
		return "", err
	}

	return cotacao.USDBRL.Bid, nil
}

func (s *CotacaoService) getCotacao(ctx context.Context) (*models.CotacaoResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", config.CotacaoAPIURL, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cotacao models.CotacaoResponse
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		return nil, err
	}

	return &cotacao, nil
}

func (s *CotacaoService) saveCotacao(ctx context.Context, bid string) error {
	cotacao := models.Cotacao{
		Bid:       bid,
		Timestamp: time.Now(),
	}

	result := s.db.WithContext(ctx).Create(&cotacao)
	return result.Error
}
