package main

import (
	backend "APIEmail/backend/config"
	"APIEmail/backend/controllers"
	"APIEmail/backend/routes"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env")
	}

	controllers.LoadEmailConfig()
}

func main() {
	// Inicializar a conexão com o banco
	backend.InitDB()

	// Inicia o cron job
	iniciarCron()

	// Configura o roteador Gin
	r := gin.Default()

	routes.ConfigurarRotas(r)

	// Log para confirmar que o servidor está iniciado
	fmt.Println("Servidor iniciado em http://localhost:8080")
	log.Fatal(r.Run(":8080")) // Inicializa o servidor Gin na porta 8080
}

// Função para iniciar o cron (verificação a cada 24 horas) e o servidor HTTP
func iniciarCron() {
	// Configura o cron para iniciar a verificação de alertas a cada 24 horas
	go verificarAlertasPeriodicamente()

	// Log para saber que o cron foi iniciado
	log.Println("Verificação de alertas iniciada a cada 24 horas")
}

// Função que será chamada a cada 24 horas
func verificarAlertasPeriodicamente() {
	ticker := time.NewTicker(24 * time.Hour) // Ticker a cada 24 horas
	defer ticker.Stop()                      // Certifique-se de parar o ticker quando terminar

	// Use for range para iterar sobre os valores recebidos no canal
	for range ticker.C {
		// Faz a requisição GET para a rota /alerta
		resp, err := http.Get("http://localhost:8080/alerta")
		if err != nil {
			fmt.Println("Erro ao fazer a requisição:", err)
			continue
		}

		// Exibe a resposta do servidor
		fmt.Printf("Resposta do servidor: %s\n", resp.Status)

		// Certifique-se de fechar o corpo da resposta após o uso
		resp.Body.Close()
	}
}
