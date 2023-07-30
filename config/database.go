package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func DatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Error().Err(err).Msg("error loading .env file")
	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Error().Err(err).Msg("DATABASE_URL environment variable is not set")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("There was an error connecting to the database")
	}
	log.Info().Msg("Database connection established")
	return db
}
