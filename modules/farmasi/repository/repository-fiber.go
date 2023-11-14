package repository

import (
	"fmt"
	"vincentcoreapi/modules/farmasi"
	"vincentcoreapi/modules/farmasi/dto"
)

func (fr *farmasiRepository) CekKodeBookingRepositoryV2(req dto.GetAntreanFarmasiRequestV2) (res farmasi.AntreanOL, err error) {

	query := "SELECT * FROM rekam.antrian_ol WHERE no_book=? LIMIT 1"
	rs := fr.DB.Raw(query, req.Kodebooking).Scan(&res)

	if rs.Error != nil {
		return res, rs.Error
	}

	if rs.RowsAffected > 0 {
		return res, nil
	}

	return res, nil
}

func (fr *farmasiRepository) CekKodeBookingAntreanResepRepositoryV2(req dto.GetAntreanFarmasiRequestV2) (res farmasi.AntreanResep, err error) {

	value := farmasi.AntreanResep{}

	query := "SELECT * FROM posfar.antrean_resep WHERE  kode_booking_ref = ? LIMIT 1"

	rs := fr.DB.Raw(query, req.Kodebooking).Scan(&value)

	fmt.Println(rs)

	if rs.Error != nil {
		return res, rs.Error
	}

	if rs.RowsAffected > 0 {
		return value, nil
	}

	return res, rs.Error
}
