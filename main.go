package main

import (
	"bytes"
	"fmt"
	"text/template"

	"gopkg.in/gomail.v2"
)

const (
	host     = "smtp.gmail.com"
	port     = 587
	username = "bremen.mcn@gmail.com"
	password = "tetryvqveqozheas"
)

func main() {

	// dialer para conexão

	dialer := gomail.NewDialer(host, port, username, password)

	// criar uma mensagem

	msg := gomail.NewMessage()
	msg.SetHeader("From", username)
	msg.SetHeader("To", username)
	msg.SetHeader("Subject", "Relatório das Ob's pendentes") // Assunto
	msg.SetBody("text/html", getBody())

	// Imagem no corpo do email em anexo tmb.

	msg.Embed("DelRio.png")

	if err := dialer.DialAndSend(msg); err != nil {
		panic(err)
	}

	fmt.Println("Mensagem enviada.")

}

// função para atribuir o html

func getBody() string {
	t := template.Must(template.ParseFiles("mail.html"))
	var buff bytes.Buffer
	t.Execute(&buff, nil)
	return buff.String()
}
