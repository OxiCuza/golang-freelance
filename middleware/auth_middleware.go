package middleware

import (
	"golang-freelance/app"
	"golang-freelance/helper"
	"golang-freelance/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	apiKey := app.EnvVariable("X-API-Key")
	headerKey := request.Header.Get("X-API-Key")

	if apiKey == headerKey {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		//	ERROR
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		jsonResponse := web.JSONResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponse(writer, jsonResponse)
	}
}
