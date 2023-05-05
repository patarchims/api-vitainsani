package usecase

import (
	"context"
	"vincentcoreapi/modules/mutiara/entity"
)

type mutiaraUseCase struct {
	mutiaraRepository entity.MutiaraRepository
	mutiaraMapper     entity.MutiaraMapper
}

func MutiaraUseCase(fr entity.MutiaraRepository, mm entity.MutiaraMapper) entity.MutiaraUseCase {
	return &mutiaraUseCase{
		mutiaraRepository: fr,
		mutiaraMapper:     mm,
	}
}

func (mu *mutiaraUseCase) GetDataKaryawan(ctx context.Context, userID string) (data interface{}, err error) {
	m := map[string]any{}
	gaji, _ := mu.mutiaraRepository.GetGaji(ctx, userID)
	karyawan, _ := mu.mutiaraRepository.GetKaryawan(ctx, userID)

	gajiMapper := mu.mutiaraMapper.ToDataGaji(gaji)

	if err != nil {
		return m, err
	}
	m["karyawan"] = karyawan
	if len(gaji) > 0 {
		m["gaji"] = gajiMapper
	} else {
		m["gaji"] = []any{}
	}

	return m, nil
}
