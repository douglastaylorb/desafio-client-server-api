package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/douglastaylorb/desafio-client-server-api/tree/main/server/config"
	"github.com/douglastaylorb/desafio-client-server-api/tree/main/server/controllers"
	database "github.com/douglastaylorb/desafio-client-server-api/tree/main/server/db"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Erro ao inicializar o banco de dados: %v", err)
	}

	cotacaoController := controllers.NewCotacaoController(db)

	http.HandleFunc("/cotacao", cotacaoController.GetCotacao)

	fmt.Printf("Servidor rodando na porta %d\n", config.ServerPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ServerPort), nil))
}
