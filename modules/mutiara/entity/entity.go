package entity

import (
	"context"
	"vincentcoreapi/modules/mutiara"
	"vincentcoreapi/modules/mutiara/dto"
)

type MutiaraRepository interface {
	GetGaji(ctx context.Context, userID string) (res []mutiara.DGaji, err error)
	GetKaryawan(ctx context.Context, userID string) (res mutiara.DKaryawan, err error)
}

type MutiaraUseCase interface {
	GetDataKaryawan(ctx context.Context, userID string) (data interface{}, err error)
}

type MutiaraMapper interface {
	ToDataGaji(gaji []mutiara.DGaji) (res []dto.ResDataGaji)
}
