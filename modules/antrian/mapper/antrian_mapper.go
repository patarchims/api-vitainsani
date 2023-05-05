package mapper

import (
	"strconv"
	"time"
	"vincentcoreapi/modules/antrian"
	"vincentcoreapi/modules/antrian/dto"
)

type AntrianMapperImpl struct {
}

func NewAntrianMapperImpl() IAntrianMapper {
	return &AntrianMapperImpl{}
}

func (m *AntrianMapperImpl) ToAntrianOlModel(data dto.RegisterPasienBaruRequest) (antrianOL antrian.AntrianOl) {
	return antrian.AntrianOl{
		Nik:             data.Nik,
		Noka:            data.Nomorkartu,
		NoHp:            data.Nohp,
		Nama:            data.Nama,
		Gender:          data.Jeniskelamin,
		Dob:             data.Tanggallahir,
		Alamat:          data.Alamat,
		KodeDebitur:     "BPJS",
		Debitur:         "JKN-BPJS",
		BookDate:        time.Now().Format("2006-01-02 15:04:05"),
		JknNomorkk:      data.Nomorkk,
		JknTanggallahir: data.Tanggallahir,
		JknKodeprop:     data.Kodeprop,
		JknNamaprop:     data.Namaprop,
		JknKodedati2:    data.Kodedati2,
		JknNamadati2:    data.Namadati2,
		JknKodekec:      data.Kodekec,
		JknNamakec:      data.Namakec,
		JknKodekel:      data.Kodekel,
		JknNamakel:      data.Namakel,
		JknRw:           data.Rw,
		JknRt:           data.Rt,
	}
}

func (m *AntrianMapperImpl) ToSisaAntranDTO(res map[string]any) (data dto.SisaANtreanDTO) {
	return data
}

func (m *AntrianMapperImpl) ToJadwalOperasiDTO(jadopOls []antrian.JadopOl, isForPasien bool) (jadopOlsDTO []dto.JadwalOperasiDTO) {
	for _, V := range jadopOls {

		terlaksana, _ := strconv.Atoi(V.Status)
		if isForPasien {
			jadopOlsDTO = append(jadopOlsDTO, dto.JadwalOperasiDTO{
				Kodebooking:    V.NoBook,
				Tanggaloperasi: V.TglOperasi.Format("2006-01-02"),
				Jenistindakan:  V.JenisTindakan,
				Kodepoli:       V.KdTujuan,
				Namapoli:       V.Tujuan,
				Terlaksana:     terlaksana,
			})
		} else {
			jadopOlsDTO = append(jadopOlsDTO, dto.JadwalOperasiDTO{
				Kodebooking:    V.NoBook,
				Tanggaloperasi: V.TglOperasi.Format("2006-01-02"),
				Jenistindakan:  V.JenisTindakan,
				Kodepoli:       V.KdTujuan,
				Namapoli:       V.Tujuan,
				Terlaksana:     terlaksana,
				Nopeserta:      V.Noka,
				Lastupdate:     V.UpdDttm.UnixNano() / int64(time.Millisecond),
			})
		}
	}

	return jadopOlsDTO
}
