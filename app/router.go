package app

import (
	"github.com/julienschmidt/httprouter"
	"golang-freelance/controller"
	"golang-freelance/exception"
)

func NewRouter(userController controller.UserController, blogController controller.BlogController) *httprouter.Router {
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

	// ERROR HANDLING
	router.PanicHandler = exception.ErrorHandler

	return router
}
