package entity

import (
	"vincentcoreapi/modules/user"
)

// UserUseCase
type UserUseCase interface {
	GetByUserUsecase(userName string) (res user.ApiUser, exist bool)
}

// UserRepository
type UserRepository interface {
	GetByIDRepository(userID string) (res user.ApiUser, exist bool)
	GetByUserRepository(userName string) (res user.ApiUser, exist bool)
}

type UserMapper interface{}
