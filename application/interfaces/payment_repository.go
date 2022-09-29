package interfaces

import "marketplace-clean/domain"

type PaymentRepository interface {
	GetById(id string) domain.Payment
	Save(payment domain.Payment)
	Cancel(payment domain.Payment)
}
