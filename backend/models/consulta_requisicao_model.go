package models

import (
	backend "APIEmail/backend/config"
	"time"
)

type Requisicao struct {
	NReq             int
	DataRequisicao   string
	Solicitante      string
	NivelNecessidade string
	NomeArtigo       *string
	Composicao       *string
	Cor              *string
	TamanhoAmostra   *string
}

// consultarRequisicao faz uma consulta ao banco e retorna a requisicão
func ConsultarRequisicao() ([]Requisicao, error) {
	var requisicao []Requisicao

	query := `
		SELECT TOP 1 ID AS 'NREQ', DATA_REQUISICAO, SOLICITANTE, NIVEL_NECESSIDADE, NOME_ARTIGO, COMPOSICAO, COR,TAMANHO_AMOSTRA 
		FROM DRI_REQUISICAO_DESENVOLVIMENTO drd 
		ORDER BY ID DESC
	`
	rows, err := backend.GetDB().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var r Requisicao
		var data time.Time
		if err := rows.Scan(&r.NReq, &data, &r.Solicitante, &r.NivelNecessidade, &r.NomeArtigo, &r.Composicao, &r.Cor, &r.TamanhoAmostra); err != nil {
			return nil, err
		}
		r.DataRequisicao = data.Format("02-01-2006")
		requisicao = append(requisicao, r)
	}

	return requisicao, nil
}
