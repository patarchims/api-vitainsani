package usecase

import (
	"vincentcoreapi/modules/user"
	"vincentcoreapi/modules/user/entity"

	"github.com/sirupsen/logrus"
)

type userUseCase struct {
	userRepository entity.UserRepository
	logging        *logrus.Logger
}

func NewUserUseCase(ur entity.UserRepository, log *logrus.Logger) entity.UserUseCase {
	return &userUseCase{
		userRepository: ur,
		logging:        log,
	}
}

func (uu *userUseCase) GetByUserUsecase(userName string) (user user.ApiUser, exist bool) {
	user, exist = uu.userRepository.GetByUserRepository(userName)
	return
}

func (uu *userUseCase) OnChangeUserUseCase() (message string, err error) {
	// ====s
	users, er := uu.userRepository.GetAllUserKlinikRepository()

	if er != nil {
		uu.logging.Info(er)
	}

	// PARSING DATA
	for _, V := range users {
		uu.logging.Info(V.Id)
		// OnUpdateUserKlinikRepository(userID string, app string, userDat user.KlinikUsers) (err error)
		// UPDATE DATA
		dataUser := user.KlinikUsers{
			App:      V.App,
			Id:       V.Id + "-00",
			User:     V.User,
			Password: V.Password,
		}
		update := uu.userRepository.OnUpdateUserKlinikRepository(V.Id, V.App, dataUser)

		if update != nil {
			uu.logging.Info(update)
		}
	}

	return "Data berhasil diubah", nil
}
