package models

import backend "APIEmail/backend/config"

// Importa o pacote db

// consultarEmails faz uma consulta ao banco e retorna os e-mails
func ConsultarEmails() ([]string, error) {
	var emails []string
	// Usa a conexão com o banco de dados obtida do pacote db
	query := `SELECT EMAIL FROM DRI_EMAIL_NOTIFICACAO`
	rows, err := backend.GetDB().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Processa os resultados da consulta
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			return nil, err
		}
		emails = append(emails, email)
	}

	return emails, nil
}
