package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type CotacaoResponse struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var cotacao CotacaoResponse
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		log.Fatal(err)
	}

	err = saveCotacao(cotacao.Bid)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Cotação atual do dólar: %s\n", cotacao.Bid)
}

func saveCotacao(bid string) error {
	content := fmt.Sprintf("Cotação atual do dólar: %s", bid)
	return os.WriteFile("cotacao.txt", []byte(content), 0644)
}
