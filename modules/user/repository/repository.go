package repository

import (
	"vincentcoreapi/modules/user"
	"vincentcoreapi/modules/user/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) entity.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) GetByIDRepository(userID string) (user user.ApiUser, exist bool) {
	rs := ur.DB.First(&user, userID).RowsAffected
	if rs > 0 {
		return user, true
	} else {
		return user, false
	}
}

func (ur *userRepository) GetByUserRepository(userName string) (user user.ApiUser, exist bool) {
	rs := ur.DB.Where("username = ? ", userName).First(&user).RowsAffected
	if rs > 0 {
		return user, true
	} else {
		return user, false
	}
}
