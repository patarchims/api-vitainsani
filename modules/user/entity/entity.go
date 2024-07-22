package entity

import (
	"vincentcoreapi/modules/user"
)

// UserUseCase
type UserUseCase interface {
	GetByUserUsecase(userName string) (res user.ApiUser, exist bool)
	OnChangeUserUseCase() (message string, err error)
}

// UserRepository
type UserRepository interface {
	GetByIDRepository(userID string) (res user.ApiUser, exist bool)
	GetAllUserKlinikRepository() (user []user.KlinikUsers, err error)
	GetByUserRepository(userName string) (res user.ApiUser, exist bool)
	OnUpdateUserKlinikRepository(userID string, app string, userDat user.KlinikUsers) (err error)
}

type UserMapper interface{}
