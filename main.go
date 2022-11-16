package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
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

	blogRepository := repository.NewBlogRepository()
	blogService := service.NewBlogService(blogRepository, db, validate)
	blogController := controller.NewBlogController(blogService)

	router := httprouter.New()

	// FEATURE USERS
	router.GET("/api/v1/users", userController.FindAll)
	router.POST("/api/v1/users", userController.Create)
	router.GET("/api/v1/users/:userId", userController.FindById)
	router.PUT("/api/v1/users/:userId", userController.Update)
	router.DELETE("/api/v1/users/:userId", userController.Delete)

	// FEATURE BLOG POSTS
	router.GET("/api/v1/posts", blogController.FindAll)
	router.POST("/api/v1/posts", blogController.Create)
	router.GET("/api/v1/posts/:postId", blogController.FindById)
	router.PUT("/api/v1/posts/:postId", blogController.Update)
	router.DELETE("/api/v1/posts/:postId", blogController.Delete)

	address := fmt.Sprintf("localhost:%s", app.EnvVariable("APP_PORT"))
	server := http.Server{
		Addr:    address,
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
