package interfaces

import "marketplace-clean/domain"

type ProductRepository interface {
	CreateProduct(product domain.Product)
	DeleteProduct(product domain.Product)
	GetAllRatingsFromProduct(product domain.Product)
}
