package handler

import (
	"github.com/gin-gonic/gin"
	"load-test-experiment/model"
	"load-test-experiment/utils/actions"
	"net/http"
)

type mediumV1Handler struct {}

type MediumV1Handler interface{
	Handle() gin.HandlerFunc
}

func NewMediumV1Handler() MediumV1Handler {
	return &mediumV1Handler{}
}

func (h *mediumV1Handler) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		r := model.Request{}
		if err := ctx.ShouldBindJSON(&r); err != nil {
			ctx.JSON(http.StatusBadRequest, model.Response{
				Message: "Bad Request",
				Data: struct{}{},
			})
		}

		for _, val := range h.registerRoutes() {
			if val.Action == r.Action {
				val.Handler(ctx)
				return
			}
		}

		ctx.JSON(http.StatusNotFound, model.Response{
			Message: "Action Not Found",
			Data: struct{}{},
		})
	}
}

func (h *mediumV1Handler) registerRoutes() []model.Route {
	return []model.Route{
		{
			Action: actions.CREATE,
			Handler: h.create,
		},
		{
			Action: actions.GET,
			Handler: h.get,
		},
	}
}

func (h *mediumV1Handler) create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.Response{
		Message: "Create",
		Data: struct{}{},
	})
}

func (h *mediumV1Handler) get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.Response{
		Message: "Get",
		Data: struct{}{},
	})
}