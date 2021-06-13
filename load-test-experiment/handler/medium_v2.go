package handler

import (
	"github.com/gin-gonic/gin"
	"load-test-experiment/model"
	"load-test-experiment/usecase"
	util "load-test-experiment/utils"
	"log"
	"net/http"
)

type mediumV2Handler struct {
	mediumv1Usecase usecase.MediumV2Usecase
}

type MediumV2Handler interface{
	Create() gin.HandlerFunc
	Get() gin.HandlerFunc
}

func NewMediumV2Handler(mediumv1Usecase usecase.MediumV2Usecase) MediumV2Handler {
	return &mediumV2Handler{mediumv1Usecase}
}

func (h *mediumV2Handler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var r model.MediumV2Request
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

func (h *mediumV2Handler) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var r model.MediumV2Request
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

func (h *mediumV2Handler) create(ctx *gin.Context, r model.MediumV2Request) {
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

func (h *mediumV2Handler) get(ctx *gin.Context, r model.MediumV2Request) {
	filterQuery, err := util.CreateQueryFilter(&r.Query.Filters, nil)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "Bad request bruh!",
			Data: make([]model.MediumV2Model, 0),
		})
	}

	result, err := h.mediumv1Usecase.Get(filterQuery, r.Query.Skip, r.Query.Take)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "Bad request bruh!",
			Data: make([]model.MediumV2Model, 0),
		})
		return
	}

	if len(*result) == 0 {
		log.Println("len(*result) == 0")
		ctx.JSON(http.StatusOK, model.Response{
			Message: "Success get data",
			Data: make([]model.MediumV2Model, 0),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Message: "Success get data",
		Data: *result,
	})
}