package handler

import (
	"github.com/gin-gonic/gin"
	"load-test-experiment/model"
	"load-test-experiment/usecase"
	util "load-test-experiment/utils"
	"log"
	"net/http"
)

type lightV2Handler struct {
	lightv1Usecase usecase.LightV2Usecase
}

type LightV2Handler interface{
	Create() gin.HandlerFunc
	Get() gin.HandlerFunc
}

func NewLightV2Handler(lightv1Usecase usecase.LightV2Usecase) LightV2Handler {
	return &lightV2Handler{lightv1Usecase}
}

func (h *lightV2Handler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var r model.LightV2Request
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

func (h *lightV2Handler) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var r model.LightV2Request
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

func (h *lightV2Handler) create(ctx *gin.Context, r model.LightV2Request) {
	err := h.lightv1Usecase.Create(&r.Data)
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

func (h *lightV2Handler) get(ctx *gin.Context, r model.LightV2Request) {
	filterQuery, err := util.CreateQueryFilter(&r.Query.Filters, nil)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "Bad request bruh!",
			Data:make([]interface{}, 0),
		})
	}

	result, err := h.lightv1Usecase.Get(filterQuery, r.Query.Skip, r.Query.Take)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "Bad request bruh!",
			Data: make([]interface{}, 0),
		})
		return
	}

	if len(*result) == 0 {
		log.Println("len(*result) == 0")
		ctx.JSON(http.StatusOK, model.Response{
			Message: "Success get data",
			Data: make([]interface{}, 0),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Message: "Success get data",
		Data: result,
	})
}