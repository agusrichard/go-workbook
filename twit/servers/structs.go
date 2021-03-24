package servers

import (
	"twit/handlers"
	"twit/repositories"
	"twit/usecases"
)

type Repositories struct {
	UserRepository repositories.UserRepository
}

type Usecases struct {
	UserUsecase usecases.UserUsecase
}

type Handlers struct {
	UserHandler handlers.UserHandler
}
