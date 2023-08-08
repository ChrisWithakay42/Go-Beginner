package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"go_rest/config"
	"go_rest/controller"
	"go_rest/model"
	"go_rest/repository"
	"go_rest/router"
	"go_rest/service"
	"go_rest/utils"
	"net/http"
)

func main() {
	log.Info().Msg("Server is up and running!!")
	db := config.DatabaseConnection()

	err := db.Table("users").AutoMigrate(&model.User{})

	validate := validator.New()

	// Repository
	userRepository := repository.NewUserRepoImpl(db)

	// Service
	userService := service.NewUserServiceImpl(userRepository, validate)

	// Controller
	userController := controller.NewUserController(userService)

	// Router
	routes := router.NewRouter(userController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err = server.ListenAndServe()
	utils.ErrorPanic(err)
}
