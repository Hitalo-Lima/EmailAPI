package controllers

import (
	"APIEmail/backend/models"
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"

	"gopkg.in/gomail.v2"
)

// Constantes de configuração do e-mail
const (
	host = "smtp.gmail.com"
	port = 587
)

var (
	username string
	password string
)

func LoadEmailConfig() {
	// Carregar as variáveis de configuração do e-mail
	username = os.Getenv("EMAIL_API_USERNAME")
	password = os.Getenv("EMAIL_API_PASSWORD")

	if username == "" || password == "" {
		log.Fatal("As variáveis de ambiente EMAIL_API_USERNAME ou EMAIL_API_PASSWORD não estão configuradas corretamente.")
	}
}

// Função genérica de envio de e-mail, que aceita diferentes tipos de dados
func EnviarEmail(dados interface{}) error {

	// Consulta os e-mails no banco de dados
	emails, err := models.ConsultarEmails()
	if err != nil {
		return fmt.Errorf("erro ao consultar e-mails: %w", err)
	}

	// Configura o dialer para conexão com o servidor SMTP
	dialer := gomail.NewDialer(host, port, username, password)

	// Cria a mensagem de e-mail
	msg := gomail.NewMessage()
	msg.SetHeader("From", username)
	msg.SetHeader("To", emails...)

	// Variáveis para determinar o template e assunto
	var subject string
	var templateFile string
	var bodyContent string

	// Switch para definir o template e assunto baseado no tipo de dados
	switch v := dados.(type) {
	case []models.Requisicao:
		subject = fmt.Sprintf("Nova requisição solicitada, Nº %d", v[0].NReq)
		templateFile = "./frontend/mail.html"
		bodyContent = getBody(templateFile, v)
	case []models.CamposAlerta:
		subject = "Alerta: Solicitações Pendentes"
		templateFile = "./frontend/alerta.html"
		bodyContent = getBody(templateFile, v)
	default:
		return fmt.Errorf("tipo de dados não suportado para envio de e-mail")
	}

	// Define o assunto e o corpo do e-mail
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", bodyContent)

	// Anexa uma imagem ao corpo do e-mail
	msg.Embed("DelRio.png")

	// Envia a mensagem
	if err := dialer.DialAndSend(msg); err != nil {
		return err
	}

	log.Println("Mensagem enviada.")
	return nil
}

// Função para renderizar o template com base nos dados
func getBody(templateFile string, data interface{}) string {
	t := template.Must(template.ParseFiles(templateFile))
	var buff bytes.Buffer
	t.Execute(&buff, data)
	return buff.String()
}
