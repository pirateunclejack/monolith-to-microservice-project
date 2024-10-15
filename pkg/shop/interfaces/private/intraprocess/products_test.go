package intraprocess

import (
	"testing"

	"github.com/pirateunclejack/monolith-to-microservice-project/pkg/common/price"
	"github.com/pirateunclejack/monolith-to-microservice-project/pkg/shop/domain/products"
	"github.com/stretchr/testify/assert"
)

func TestProductFromDomainProduct(t *testing.T) {
	productPrice := price.NewPricePanic(42, "USD")
	domainProduct, err := products.NewProduct("123", "name", "desc", productPrice)
	assert.NoError(t, err)

	p := ProductFromDomainProduct(*domainProduct)

	assert.EqualValues(t, Product{
		"123",
		"name",
		"desc",
		productPrice,
	}, p)
}
