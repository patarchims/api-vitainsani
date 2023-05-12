package usecase

import (
	"vincentcoreapi/modules/billing/entity"
)

type billingUseCase struct {
	billingRepository entity.BillingRepository
	billingMapper     entity.BillingMapper
}

func NewBillingUseCase(br entity.BillingRepository, nm entity.BillingMapper) entity.BillingUsecase {
	return &billingUseCase{
		billingRepository: br,
		billingMapper:     nm,
	}
}
