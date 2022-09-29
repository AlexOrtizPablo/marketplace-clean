package interfaces

import "marketplace-clean/domain"

type ShippingRepository interface {
	GetById(id string) domain.Shipping
}
