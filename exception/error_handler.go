package exception

import (
	"github.com/go-playground/validator/v10"
	"golang-freelance/helper"
	"golang-freelance/model/web"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if validationErrors(writer, request, err) {
		return
	}

	if notFoundError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	_, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		var errString []string
		for _, err := range err.(validator.ValidationErrors) {
			errString = append(errString, err.Field()+" is "+err.Tag()+" "+err.Param())
		}

		jsonResponse := web.JSONResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   errString,
		}

		helper.WriteToResponse(writer, jsonResponse)
		return true
	}

	return false
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	_, ok := err.(NotFoundError)
	if !ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		jsonResponse := web.JSONResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
		}

		helper.WriteToResponse(writer, jsonResponse)
		return true
	}

	return false
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	jsonResponse := web.JSONResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponse(writer, jsonResponse)
}
