package mapper

import (
	"vincentcoreapi/modules/farmasi"
	"vincentcoreapi/modules/farmasi/dto"
)

type IFarmasiMapper interface {
	ToFarmasiAntreanResep(data farmasi.AntreanResep) (ambilAntreanResponse dto.AmbilAntreanFarmasiResponse)
	ToStatusAntranFarmasiResponse(data farmasi.AntreanResep, statusAntrea farmasi.StatusAntrean) (statusAntrean dto.StatusAntreanFarmasiResponse)
}
