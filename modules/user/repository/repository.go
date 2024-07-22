package repository

import (
	"errors"
	"fmt"
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

func (ur *userRepository) GetAllUserKlinikRepository() (user []user.KlinikUsers, err error) {
	query := "SELECT * FROM klinik.users;"
	result := ur.DB.Raw(query).Scan(&user)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error)
		return user, errors.New(message)
	}

	return user, nil
}

// UPDATE DATA
func (ur *userRepository) OnUpdateUserKlinikRepository(userID string, app string, userDat user.KlinikUsers) (err error) {
	err1 := ur.DB.Model(user.KlinikUsers{}).Where("id=? AND app=?", userID, app).Updates(userDat).Error

	if err1 != nil {
		return err1
	}

	return nil

}
