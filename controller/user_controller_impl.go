package controller

import (
	"github.com/julienschmidt/httprouter"
	"golang-freelance/helper"
	"golang-freelance/model/web"
	"golang-freelance/service"
	"net/http"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserControllerImpl(userService service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRequest := web.UserCreateRequest{}
	helper.ReadFromRequest(request, &userRequest)

	userResponse := controller.UserService.Save(request.Context(), userRequest)

	jsonResponse := web.JSONResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponse(writer, jsonResponse)
}

func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRequest := web.UserUpdateRequest{}
	helper.ReadFromRequest(request, &userRequest)

	userRequest.Id = params.ByName("userId")
	userResponse := controller.UserService.Update(request.Context(), userRequest)

	jsonResponse := web.JSONResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponse(writer, jsonResponse)
}

func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	controller.UserService.Delete(request.Context(), userId)

	jsonResponse := web.JSONResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponse(writer, jsonResponse)
}

func (controller *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	userResponse := controller.UserService.FindById(request.Context(), userId)

	jsonResponse := web.JSONResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponse(writer, jsonResponse)
}

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userResponse := controller.UserService.FindAll(request.Context())

	jsonResponse := web.JSONResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponse(writer, jsonResponse)
}
