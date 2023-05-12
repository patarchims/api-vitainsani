package handler

import "vincentcoreapi/modules/billing/entity"

type BillingHadler struct {
	BillingUsecase    entity.BillingUsecase
	BillingRepository entity.BillingRepository
	BillingMapper     entity.BillingMapper
}
