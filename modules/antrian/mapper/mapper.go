package mapper

import (
	"vincentcoreapi/modules/antrian"
	"vincentcoreapi/modules/antrian/dto"
)

type IAntrianMapper interface {
	ToAntrianOlModel(data dto.RegisterPasienBaruRequest) (antrianOL antrian.AntrianOl)
	ToJadwalOperasiDTO(jadopOls []antrian.JadopOl, isForPasien bool) (jadopOlsDTO []dto.JadwalOperasiDTO)
	ToSisaAntranDTO(res map[string]any) (data dto.SisaANtreanDTO)
}
