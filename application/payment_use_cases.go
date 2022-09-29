package application

import (
	"errors"
	"marketplace-clean/application/interfaces"
	"marketplace-clean/domain"
)

type PaymentUseCases struct {
	PaymentRepository interfaces.PaymentRepository
	ShippingUseCase   *ShippingUseCase
}

func (p *PaymentUseCases) GetPaymentById(shippingId string) (domain.Payment, error) {
	payment := p.PaymentRepository.GetById(shippingId)
	if !payment.IsValid() {
		return domain.Payment{}, errors.New("payment invalid")
	}

	return payment, nil
}

func (p *PaymentUseCases) CreatePayment(payment domain.Payment) error {
	paymentCreated := domain.CreatePayment(payment)

	if !paymentCreated.IsValid() {
		return errors.New("payment invalid")
	}

	p.PaymentRepository.Save(paymentCreated)
	return nil
}

func (p *PaymentUseCases) ApprovePayment(paymentId string) error {
	payment, err := p.GetPaymentById(paymentId)
	if err != nil {
		return err
	}

	err = payment.Approve()
	if err != nil {
		return err
	}

	p.PaymentRepository.Save(payment)
	err = p.ShippingUseCase.CreateShippingByUser(payment.Payer)
	if err != nil {
		return err
	}

	return nil
}

func (p *PaymentUseCases) CancelPayment(paymentId string) error {
	payment, err := p.GetPaymentById(paymentId)
	if err != nil {
		return err
	}

	err = payment.Cancel()
	if err != nil {
		return err
	}

	p.PaymentRepository.Cancel(payment)

	return nil
}
