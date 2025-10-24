package database

import (
	"fmt"
	"log"
	"os"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectaComBancoDeDados() {
	stringDeConexao := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("DBNAME"),
		os.Getenv("PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
	if err != nil {
		log.Panicf("Erro ao conectar com banco de dados: %v", err)
	}

	if err := DB.AutoMigrate(&models.Aluno{}); err != nil {
		log.Panicf("Erro ao migrar tabela: %v", err)
	}
}
