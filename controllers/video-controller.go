package controllers

import (
	"net/http"
	"practice/entities"
	"practice/services"
	"practice/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entities.Video
	Save(ctx *gin.Context) interface{}
	ShowAll(ctx *gin.Context)
}

type VideoControllerImpl struct {
	service services.VideoService
}

var validate *validator.Validate

func New(service services.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateIsTitleCool)
	return &VideoControllerImpl{
		service: service,
	}
}

func (controller *VideoControllerImpl) FindAll() []entities.Video {
	return controller.service.FindAll()
}

func (controller *VideoControllerImpl) Save(ctx *gin.Context) interface{} {

	var video entities.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	err = validate.Struct(video)
	if err != nil {
		return err
	}

	return controller.service.Save(&video)
}

func (controller *VideoControllerImpl) ShowAll(ctx *gin.Context) {
	videos := controller.service.FindAll()
	data := gin.H{
		"title":  "Videos Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
