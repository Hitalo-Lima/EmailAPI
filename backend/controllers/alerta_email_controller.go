package controllers

import (
	"APIEmail/backend/models"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

// AlertaHandler trata a requisição para exibir os alertas e enviar e-mails
func AlertaHandler(c *gin.Context) {
	// Chama o modelo para consultar os dados de alerta
	alertas, err := models.ConsultarDadosAlerta()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar os alertas"})
		return
	}

	// Renderiza o template usando os dados recebidos
	tmpl, err := template.ParseFiles("./frontend/alerta.html")
	if err != nil {
		log.Print("Erro ao carregar template:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao carregar template"})
		return
	}

	// Gera a resposta no contexto HTTP de Gin
	c.Writer.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(c.Writer, alertas); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao renderizar template"})
		return
	}

	// Filtra os alertas para aqueles que devem ser enviados por e-mail
	var alertasParaEnvio []models.CamposAlerta
	for _, alerta := range alertas {
		dataRequisicao, err := time.Parse("02/01/2006", alerta.DataRequisicao)
		if err != nil {
			log.Println("Erro ao analisar DataRequisicao:", err)
			continue
		}

		// Calcula a diferença em dias entre a data de requisição e a data atual
		diasDesdeRequisicao := time.Since(dataRequisicao).Hours() / 24
		if diasDesdeRequisicao >= 40 {
			alertasParaEnvio = append(alertasParaEnvio, alerta)
		}
	}

	// Se houver alertas para envio, chama a função para enviar o e-mail
	if len(alertasParaEnvio) > 0 {
		if err := EnviarEmail(alertasParaEnvio); err != nil {
			log.Fatal("Erro ao enviar e-mail:", err)
		}
	}
}
