package mapper

import (
	"vincentcoreapi/modules/user/dto"
	"vincentcoreapi/modules/user/entity"
)

type UserMapperImpl struct {
}

func NewAntrianMapperImpl() entity.UserMapper {
	return &UserMapperImpl{}
}

func (um *UserMapperImpl) ToUserMapperImple() (res dto.PasienResponse) {
	return dto.PasienResponse{}
}
