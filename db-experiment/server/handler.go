package server

import handler "db-experiment/handlers"

type handlers struct {
	todoHandler handler.TodoHandler
	todoHandlerV2 handler.TodoHandlerV2
}

func setupHandlers(uscs *usecases) *handlers {
	todoHandler := handler.InitializeTodoHandler(uscs.todoUsecase)
	todoHandlerV2 := handler.InitializeTodoHandlerV2(uscs.todoUsecaseV2)

	return &handlers{
		todoHandler: todoHandler,
		todoHandlerV2: todoHandlerV2,
	}
}
