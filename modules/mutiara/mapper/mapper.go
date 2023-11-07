package mapper

import (
	"vincentcoreapi/helper"
	"vincentcoreapi/modules/mutiara"
	"vincentcoreapi/modules/mutiara/dto"
	"vincentcoreapi/modules/mutiara/entity"
)

type MutiaraMapperImpl struct {
}

func NewMutiaranMapperImpl() entity.MutiaraMapper {
	return &MutiaraMapperImpl{}
}

func (m *MutiaraMapperImpl) ToDataGajiMapper(gaji []mutiara.DGaji) (res []dto.ResDataGaji) {
	for _, V := range gaji {
		res = append(res, dto.ResDataGaji{
			Jumlah:    helper.FormatRupiah(V.Jumlah),
			Gajipokok: helper.FormatRupiah(V.Gajipokok),
			Makan:     helper.FormatRupiah(V.Makan),
			Tansport:  helper.FormatRupiah(V.Tansport),
			Tglbayar:  V.Tglbayar.Format("2006-01-02"),
			Jamlem:    helper.FormatRupiah(V.Jamlem),
			Bulan:     V.Bulan,
			Bruto:     V.Bulan,
			Potot:     helper.FormatRupiah(V.Potot),
		})
	}

	return res
}
