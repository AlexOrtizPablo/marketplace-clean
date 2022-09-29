package application

import (
	"errors"
	"github.com/google/uuid"
	"marketplace-clean/application/interfaces"
	"marketplace-clean/domain"
)

type ShippingUseCase struct {
	ShippingRepository interfaces.ShippingRepository
}

func (p *ShippingUseCase) CreateShippingByUser(user domain.User) error {
	err := p.CreateShipping(domain.Shipping{
		Id:   uuid.New().String(),
		User: user,
	})
	if err != nil {
		return err
	}

	return nil
}

func (p *ShippingUseCase) CreateShipping(shipping domain.Shipping) error {
	shippingCreated := domain.CreateShipping(shipping)
	if !shippingCreated.IsValid() {
		return errors.New("shipping invalid")
	}

	return nil
}

func (p *ShippingUseCase) GetShippingById(shippingId string) (domain.Shipping, error) {
	shipping := p.ShippingRepository.GetById(shippingId)
	if !shipping.IsValid() {
		return domain.Shipping{}, errors.New("shipping invalid")
	}

	return shipping, nil
}

func (p *ShippingUseCase) ChangeToInProgress(shippingId string) error {
	shipping, err := p.GetShippingById(shippingId)
	if err != nil {
		return err
	}

	err = shipping.ChangeToInProgress()
	if err != nil {
		return err
	}
	if !shipping.IsValid() {
		return errors.New("shipping invalid")
	}

	return nil
}

func (p *ShippingUseCase) ChangeToDelivered(shippingId string) error {
	shipping, err := p.GetShippingById(shippingId)
	if err != nil {
		return err
	}

	err = shipping.ChangeToInDelivered()
	if err != nil {
		return err
	}
	if !shipping.IsValid() {
		return errors.New("shipping invalid")
	}

	return nil
}
