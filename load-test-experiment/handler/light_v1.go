package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"load-test-experiment/model"
	"load-test-experiment/usecase"
	util "load-test-experiment/utils"
	"load-test-experiment/utils/actions"
	"net/http"
	"strings"
)

type lightV1Handler struct {
	lightv1Usecase usecase.LightV1Usecase
}

type LightV1Handler interface{
	Handle() gin.HandlerFunc
}

func NewLightV1Handler(lightv1Usecase usecase.LightV1Usecase) LightV1Handler {
	return &lightV1Handler{lightv1Usecase}
}

func (h *lightV1Handler) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var r model.LightV1Request
		if err := ctx.ShouldBindJSON(&r); err != nil {
			fmt.Println("err", err)
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
	fmt.Printf("%+v\n", r.Query)
	filterQuery, err := util.CreateQueryFilter(&r.Query.Filters, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "Bad request bruh!",
			Data:make([]interface{}, 0),
		})
	}

	fmt.Println("filterQuery", filterQuery)

	result, err := h.lightv1Usecase.Get(filterQuery, r.Query.Skip, r.Query.Take)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "Bad request bruh!",
			Data: make([]interface{}, 0),
		})
		return
	}

	if len(*result) == 0 {
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