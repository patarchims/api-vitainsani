package mapper

import (
	"vincentcoreapi/modules/antrian"
	"vincentcoreapi/modules/antrian/dto"
)

type IAntrianMapper interface {
	ToAntrianOlModelMapper(data dto.RegisterPasienBaruRequest) (antrianOL antrian.AntrianOl)
	ToJadwalOperasiDTOMapper(jadopOls []antrian.JadopOl, isForPasien bool) (jadopOlsDTO []dto.JadwalOperasiDTO)
	ToSisaAntranDTOMapepr(res map[string]any) (data dto.SisaANtreanDTO)
}
