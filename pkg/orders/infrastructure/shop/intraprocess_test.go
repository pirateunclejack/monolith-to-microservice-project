package shop_test

import (
	"testing"

	"github.com/pirateunclejack/monolith-to-microservice-project/pkg/common/price"
	"github.com/pirateunclejack/monolith-to-microservice-project/pkg/orders/domain/orders"
	"github.com/pirateunclejack/monolith-to-microservice-project/pkg/orders/infrastructure/shop"
	"github.com/pirateunclejack/monolith-to-microservice-project/pkg/shop/interfaces/private/intraprocess"
	"github.com/stretchr/testify/assert"
)

func TestOrderProductFromShopProduct(t *testing.T) {
	shopProduct := intraprocess.Product{
		ID: "123",
		Name: "name",
		Description: "desc",
		Price: price.NewPricePanic(42, "EUR"),
	}
	orderProduct, err := shop.OrderProductFromIntraprocess(shopProduct)
	assert.NoError(t, err)

	expectedOrderProduct, err := orders.NewProduct("123", "name", price.NewPricePanic(42, "EUR"))
	assert.NoError(t, err)

	assert.EqualValues(t, expectedOrderProduct, orderProduct)
}
