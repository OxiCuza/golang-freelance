package controller

import (
	"github.com/julienschmidt/httprouter"
	"golang-freelance/helper"
	"golang-freelance/model/web"
	"golang-freelance/service"
	"net/http"
)

type BlogControllerImpl struct {
	BlogService service.BlogService
}

func NewBlogController(blogService service.BlogService) BlogController {
	return &BlogControllerImpl{BlogService: blogService}
}

func (controller *BlogControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	blogRequest := web.BlogCreateRequest{}
	helper.ReadFromRequest(request, &blogRequest)

	blogResponse := controller.BlogService.Save(request.Context(), blogRequest)

	jsonResponse := web.JSONResponse{
		Code:   200,
		Status: "OK",
		Data:   blogResponse,
	}

	helper.WriteToResponse(writer, jsonResponse)
}

func (controller *BlogControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	blogRequest := web.BlogUpdateRequest{}
	helper.ReadFromRequest(request, &blogRequest)

	blogRequest.Id = params.ByName("postId")
	blogResponse := controller.BlogService.Update(request.Context(), blogRequest)

	jsonResponse := web.JSONResponse{
		Code:   200,
		Status: "OK",
		Data:   blogResponse,
	}

	helper.WriteToResponse(writer, jsonResponse)
}

func (controller *BlogControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	blogId := params.ByName("postId")
	controller.BlogService.Delete(request.Context(), blogId)

	jsonResponse := web.JSONResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponse(writer, jsonResponse)
}

func (controller *BlogControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	blogId := params.ByName("postId")
	blogResponse := controller.BlogService.FindById(request.Context(), blogId)

	jsonResponse := web.JSONResponse{
		Code:   200,
		Status: "OK",
		Data:   blogResponse,
	}

	helper.WriteToResponse(writer, jsonResponse)
}

func (controller *BlogControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	blogResponse := controller.BlogService.FindAll(request.Context())

	jsonResponse := web.JSONResponse{
		Code:   200,
		Status: "OK",
		Data:   blogResponse,
	}

	helper.WriteToResponse(writer, jsonResponse)
}
