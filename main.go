package main

import (
	"jin-example/config"
	"jin-example/controller"
	"jin-example/helper"
	"jin-example/model"
	tagsRepository "jin-example/repository/tags"
	userRepository "jin-example/repository/user"
	"jin-example/router"
	tagsService "jin-example/service/tags"
	userService "jin-example/service/user"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api
func main() {

	log.Info().Msg("Started Server!")
	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})
	db.Table("users").AutoMigrate(&model.User{})

	tagsRepository := tagsRepository.NewTagsRepositoryImpl(db)
	tagsService := tagsService.NewTagsServiceImpl(tagsRepository, validate)
	tagsController := controller.NewTagsController(tagsService)

	userRepository := userRepository.NewUserRepositoryImpl(db)
	userService := userService.NewUserService(userRepository, validate)
	userController := controller.NewUserController(userService)

	// Router
	routes := router.NewRouter(tagsController, userController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
