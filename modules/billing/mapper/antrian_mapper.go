package mapper

import (
	"vincentcoreapi/modules/billing/entity"
)

type BillingMapperImpl struct {
}

func NewBillingMapperImpl() entity.BillingMapper {
	return &BillingMapperImpl{}
}
