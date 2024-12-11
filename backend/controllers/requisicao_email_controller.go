package controllers

import (
	"APIEmail/backend/models"
	"log"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

// requisicaoHandler consulta a requisicao os retorna como HTML

func RequisicaoHandler(c *gin.Context) {
	requisicoes, err := models.ConsultarRequisicao()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar as requisições"})
		return
	}

	tmpl, err := template.ParseFiles("./frontend/mail.html")
	if err != nil {
		log.Print("Erro ao carregar template:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao carregar o template"})
		return
	}

	c.Writer.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(c.Writer, requisicoes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao renderizar o template"})
		return
	}

	// Envia o e-mail passando a lista de requisições
	if err := EnviarEmail(requisicoes); err != nil {
		log.Fatal("Erro ao enviar e-mail:", err)
	}
}
