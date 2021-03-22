package repositories

type userRepository struct {
}

type UserRepository interface {
}

func InitUserRepository() UserRepository {
	return &userRepository{}
}
