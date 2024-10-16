package orders

import (
	"errors"

	"github.com/pirateunclejack/monolith-to-microservice-project/pkg/common/price"
)

type ProductID string

var ErrEmptyProductID = errors.New("empty product ID")

type Product struct {
    id ProductID
    name string
    price price.Price
}

func NewProduct(id ProductID, name string, price price.Price) (Product, error){
    if len(id) == 0 {
        return Product{}, ErrEmptyProductID
    }

    return Product {
        id: id,
        name: name,
        price: price,
    }, nil
}

func (p Product) ID() ProductID {
    return p.id
}

func (p Product) Name() string {
    return p.name
}

func (p Product) Price() price.Price {
    return p.price
}
