package repository

import (
	"errors"
	"fmt"
	"vincentcoreapi/modules/mutiara"
	"vincentcoreapi/modules/mutiara/entity"

	"golang.org/x/net/context"
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

func (mr *mutiaraRepository) GetKaryawan(ctx context.Context, userID string) (res mutiara.DKaryawan, err error) {

	var karyawan mutiara.DKaryawan
	errs := mr.DB.Where("id = ?", userID).Find(&karyawan).Error

	if errs != nil {
		return karyawan, err
	}

	return karyawan, nil

}

func (mr *mutiaraRepository) GetGaji(ctx context.Context, userID string) (res []mutiara.DGaji, err error) {

	query := "SELECT * FROM mutiara.dgaji  WHERE id=?;"
	result := mr.DB.WithContext(ctx).Raw(query, userID).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", err.Error())
		return res, errors.New(message)
	}

	return res, nil

}

func (mr *mutiaraRepository) GetPengajar(ctx context.Context) (res []mutiara.DKaryawan, err error) {
	query := "SELECT * FROM mutiara.pengajar;"
	result := mr.DB.WithContext(ctx).Raw(query).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", err.Error())
		return res, errors.New(message)
	}

	return res, nil
}
