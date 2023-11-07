package repository

import (
	"errors"
	"fmt"
	"vincentcoreapi/modules/mutiara"
	"vincentcoreapi/modules/mutiara/entity"

	"gorm.io/gorm"
)

type mutiaraRepository struct {
	DB *gorm.DB
}

func NewMutiaraRepository(db *gorm.DB) entity.MutiaraRepository {
	return &mutiaraRepository{
		DB: db,
	}
}

func (mr *mutiaraRepository) GetKaryawanRepository(userID string) (res mutiara.DKaryawan, err error) {

	var karyawan mutiara.DKaryawan
	errs := mr.DB.Where("id = ?", userID).Find(&karyawan).Error

	if errs != nil {
		return karyawan, err
	}

	return karyawan, nil

}

func (mr *mutiaraRepository) GetGajiRepository(userID string) (res []mutiara.DGaji, err error) {

	query := "SELECT * FROM mutiara.dgaji  WHERE id=?;"
	result := mr.DB.Raw(query, userID).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", err.Error())
		return res, errors.New(message)
	}

	return res, nil

}

func (mr *mutiaraRepository) GetPengajarRepository() (res []mutiara.DKaryawan, err error) {
	query := "SELECT * FROM mutiara.pengajar;"
	result := mr.DB.Raw(query).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", err.Error())
		return res, errors.New(message)
	}

	return res, nil
}
