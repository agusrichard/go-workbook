package server

import handler "db-experiment/handlers"

type handlers struct {
	todoHandler handler.TodoHandler
}

func setupHandlers(uscs *usecases) *handlers {
	todoHandler := handler.InitializeTodoHandler(uscs.todoUsecase)

	return &handlers{
		todoHandler: todoHandler,
	}
}
