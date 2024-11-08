package backend

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/denisenkom/go-mssqldb" // Driver para SQL Server
)

var db *sql.DB // Variável global para a conexão com o banco

// initDB inicializa a conexão com o banco de dados
func InitDB() error {
	// Monta a string de conexão com base nas variáveis de ambiente
	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_HOST_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err error
	db, err = sql.Open("sqlserver", connectionString)
	if err != nil {
		return fmt.Errorf("erro ao conectar ao banco: %v", err)
	}

	// Verifica a conexão
	if err := db.Ping(); err != nil {
		return fmt.Errorf("erro ao verificar a conexão com o banco: %v", err)
	}

	return nil
}

// Função para obter a conexão com o banco de dados
func GetDB() *sql.DB {
	return db
}
