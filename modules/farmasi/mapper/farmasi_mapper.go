package mapper

import (
	"vincentcoreapi/modules/farmasi"
	"vincentcoreapi/modules/farmasi/dto"
)

type FarmasiMapperImpl struct {
}

func NewAntrianMapperImpl() IFarmasiMapper {
	return &FarmasiMapperImpl{}
}

func (m *FarmasiMapperImpl) ToFarmasiAntreanResep(data farmasi.AntreanResep) (ambilAntreanResponse dto.AmbilAntreanFarmasiResponse) {
	return dto.AmbilAntreanFarmasiResponse{
		JenisResep:   "Non Racikan",
		NomorAntrean: data.NoAntreanAngka,
		Keterangan:   "-",
	}
}

func (m *FarmasiMapperImpl) ToStatusAntranFarmasiResponse(data farmasi.AntreanResep, statusAntrea farmasi.StatusAntrean) (statusAntrean dto.StatusAntreanFarmasiResponse) {
	return dto.StatusAntreanFarmasiResponse{
		JenisResep:     "Non Racikan",
		AntreanPanggil: statusAntrea.Antreanpanggil,
		SisaAntrean:    statusAntrea.Sisaantrean,
		TotalAntrean:   statusAntrea.Totalantrean,
		Keterangan:     "-",
	}
}
