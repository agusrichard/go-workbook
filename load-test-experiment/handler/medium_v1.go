package handler

import (
	"github.com/gin-gonic/gin"
	"load-test-experiment/model"
	"load-test-experiment/usecase"
	util "load-test-experiment/utils"
	"log"
	"net/http"
)

type mediumV1Handler struct {
	mediumv1Usecase usecase.MediumV1Usecase
}

type MediumV1Handler interface{
	Create() gin.HandlerFunc
	Get() gin.HandlerFunc
}

func NewMediumV1Handler(mediumv1Usecase usecase.MediumV1Usecase) MediumV1Handler {
	return &mediumV1Handler{mediumv1Usecase}
}

func (h *mediumV1Handler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var r model.MediumV1Request
		if err := ctx.ShouldBindJSON(&r); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, model.Response{
				Message: "Bad Request",
				Data: struct{}{},
			})
			return
		}

		h.create(ctx, r)
	}
}

func (h *mediumV1Handler) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var r model.MediumV1Request
		if err := ctx.ShouldBindJSON(&r); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, model.Response{
				Message: "Bad Request",
				Data: struct{}{},
			})
			return
		}

		h.get(ctx, r)
	}
}

func (h *mediumV1Handler) create(ctx *gin.Context, r model.MediumV1Request) {
	err := h.mediumv1Usecase.Create(&r.Data)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: err.Error(),
			Data: struct{}{},
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Message: "Success to create",
		Data: struct{}{},
	})
}

func (h *mediumV1Handler) get(ctx *gin.Context, r model.MediumV1Request) {
	filterQuery, err := util.CreateQueryFilter(&r.Query.Filters, nil)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "Bad request bruh!",
			Data: make([]model.MediumV1Model, 0),
		})
	}

	result, err := h.mediumv1Usecase.Get(filterQuery, r.Query.Skip, r.Query.Take)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "Bad request bruh!",
			Data: make([]model.MediumV1Model, 0),
		})
		return
	}

	if len(*result) == 0 {
		log.Println("len(*result) == 0")
		ctx.JSON(http.StatusOK, model.Response{
			Message: "Success get data",
			Data: make([]model.MediumV1Model, 0),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Message: "Success get data",
		Data: *result,
	})
}