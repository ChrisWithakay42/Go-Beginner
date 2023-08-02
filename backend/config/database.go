package config

import (
	"fmt"
	"go_rest/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//	func DatabaseConnection() *gorm.DB {
//		err := godotenv.Load()
//		if err != nil {
//			log.Error().Err(err).Msg("error loading .env file")
//		}
//		dsn := os.Getenv("DATABASE_URL")
//		if dsn == "" {
//			log.Error().Err(err).Msg("DATABASE_URL environment variable is not set")
//		}
//		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//		if err != nil {
//			log.Error().Err(err).Msg("There was an error connecting to the database")
//		}
//		log.Info().Msg("Database connection established")
//		return db
//	}
const (
	host     = "localhost"
	port     = 5432
	user     = "kris"
	password = "post123"
	dbName   = "go_test"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	utils.ErrorPanic(err)

	return db
}
