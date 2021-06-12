package handler

import (
	"github.com/gin-gonic/gin"
	"load-test-experiment/model"
	"load-test-experiment/utils/actions"
	"net/http"
)

type heavyV1Handler struct {}

type HeavyV1Handler interface{
	Handle() gin.HandlerFunc
}

func NewHeavyV1Handler() HeavyV1Handler {
	return &heavyV1Handler{}
}

func (h *heavyV1Handler) Handle() gin.HandlerFunc {
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

func (h *heavyV1Handler) registerRoutes() []model.Route {
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

func (h *heavyV1Handler) create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.Response{
		Message: "Create",
		Data: struct{}{},
	})
}

func (h *heavyV1Handler) get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.Response{
		Message: "Get",
		Data: struct{}{},
	})
}