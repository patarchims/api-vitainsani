package usecase

import (
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

func (mu *mutiaraUseCase) GetDataKaryawanUsecase(userID string) (data interface{}, err error) {
	m := map[string]any{}
	gaji, _ := mu.mutiaraRepository.GetGajiRepository(userID)
	karyawan, _ := mu.mutiaraRepository.GetKaryawanRepository(userID)

	gajiMapper := mu.mutiaraMapper.ToDataGajiMapper(gaji)

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
