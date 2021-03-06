package handler

import (
	"github.com/gin-gonic/gin"
	"load-test-experiment/model"
	"load-test-experiment/usecase"
	util "load-test-experiment/utils"
	"load-test-experiment/utils/actions"
	"log"
	"net/http"
	"strings"
)

type lightV1Handler struct {
	lightv1Usecase usecase.LightV1Usecase
}

type LightV1Handler interface{
	Handle() gin.HandlerFunc
	Create() gin.HandlerFunc
	Get() gin.HandlerFunc
}

func NewLightV1Handler(lightv1Usecase usecase.LightV1Usecase) LightV1Handler {
	return &lightV1Handler{lightv1Usecase}
}

func (h *lightV1Handler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var r model.LightV1Request
		if err := ctx.ShouldBindJSON(&r); err != nil {
			ctx.JSON(http.StatusBadRequest, model.Response{
				Message: "Bad Request",
				Data: struct{}{},
			})
			return
		}

		h.create(ctx, r)
	}
}

func (h *lightV1Handler) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var r model.LightV1Request
		if err := ctx.ShouldBindJSON(&r); err != nil {
			ctx.JSON(http.StatusBadRequest, model.Response{
				Message: "Bad Request",
				Data: struct{}{},
			})
			return
		}

		h.get(ctx, r)
	}
}

func (h *lightV1Handler) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var r model.LightV1Request
		if err := ctx.ShouldBindJSON(&r); err != nil {
			ctx.JSON(http.StatusBadRequest, model.Response{
				Message: "Bad Request",
				Data: struct{}{},
			})
			return
		}

		switch strings.ToUpper(r.Action) {
		case actions.CREATE:
			h.create(ctx, r)
		case actions.GET:
			h.get(ctx, r)
		default:
			ctx.JSON(http.StatusNotFound, model.Response{
				Message: "Action Not Found",
				Data: struct{}{},
			})
		}
	}
}

func (h *lightV1Handler) create(ctx *gin.Context, r model.LightV1Request) {
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

func (h *lightV1Handler) get(ctx *gin.Context, r model.LightV1Request) {
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