package models

import (
	backend "APIEmail/backend/config"

	_ "github.com/denisenkom/go-mssqldb" // Driver do banco SQL Server
)

type CamposAlerta struct {
	NReq           int
	DataFim        string
	DataRequisicao string
	Solicitante    *string
	NomeArtigo     *string
	Composicao     *string
	Cor            *string
	TamanhoAmostra *string
}

// ConsultarDadosAlerta consulta os dados dos alertas no banco
func ConsultarDadosAlerta() ([]CamposAlerta, error) {
	var camposAlerta []CamposAlerta

	// Exemplo de consulta SQL para pegar os dados
	query := `SELECT ID, CONVERT(VARCHAR, DATEADD(DAY, 45, DATA_REQUISICAO), 103) AS DATA_FIM, 
			  CONVERT(VARCHAR, DATA_REQUISICAO, 103) AS DATA_REQUISICAO, SOLICITANTE, NOME_ARTIGO, 
			  COMPOSICAO, COR, TAMANHO_AMOSTRA FROM DRI_REQUISICAO_DESENVOLVIMENTO drd
				WHERE STATUS NOT IN ('Concluído')`

	// A variável db representa a conexão com o banco (definida anteriormente no seu código)
	rows, err := backend.GetDB().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Faz a leitura dos resultados
	for rows.Next() {
		var ca CamposAlerta
		var dataReq string
		var dataFim string

		// Lê os dados da linha
		if err := rows.Scan(&ca.NReq, &dataFim, &dataReq, &ca.Solicitante, &ca.NomeArtigo,
			&ca.Composicao, &ca.Cor, &ca.TamanhoAmostra); err != nil {
			return nil, err
		}

		// Atribui as datas formatadas
		ca.DataRequisicao = dataReq
		ca.DataFim = dataFim

		// Adiciona ao slice de alertas
		camposAlerta = append(camposAlerta, ca)
	}

	return camposAlerta, nil
}
