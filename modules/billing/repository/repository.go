package repository

import (
	"context"
	"log"
	"vincentcoreapi/modules/antrian"
	"vincentcoreapi/modules/billing/entity"

	"gorm.io/gorm"
)

type billingRepository struct {
	DB *gorm.DB
}

func NewAntrianRepository(db *gorm.DB) entity.BillingRepository {
	return &billingRepository{
		DB: db,
	}
}

func (ar *billingRepository) CekPoli(ctx context.Context, value string) (isTrue bool, err error) {
	kp := antrian.Kpoli{}

	query := "SELECT bpjs " +
		"FROM his.kpoli " +
		"WHERE bpjs = ? " +
		"LIMIT 1"
	rs := ar.DB.WithContext(ctx).Raw(query, value).Scan(&kp)
	if rs.Error != nil {
		log.Println(ctx, "[Error Query]", rs.Error)
	}
	if rs.RowsAffected > 0 {
		return true, err
	}

	return false, err
}
