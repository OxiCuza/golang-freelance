package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"golang-freelance/app"
	"golang-freelance/controller"
	"golang-freelance/helper"
	"golang-freelance/repository"
	"golang-freelance/service"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserControllerImpl(userService)

	rdb := helper.InitRedis()
	blogRepository := repository.NewBlogRepository(rdb)
	blogService := service.NewBlogService(blogRepository, userRepository, db, validate)
	blogController := controller.NewBlogController(blogService)

	router := app.NewRouter(userController, blogController)

	address := fmt.Sprintf("localhost:%s", app.EnvVariable("APP_PORT"))
	server := http.Server{
		Addr:    address,
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
