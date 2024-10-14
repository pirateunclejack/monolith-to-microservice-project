package orders

import "errors"

type ID string

var ErrEmptyID = errors.New("empty order id")

type Order struct {
    id      ID
    product Product
    address Address
    paid    bool
}

func (o *Order) ID() ID {
    return o.id
}

func (o *Order) Product() Product {
    return o.product
}

func (o *Order) Address() Address {
    return o.address
}

func (o *Order) Paid() bool {
    return o.paid
}

func (o *Order) MarkAsPaid()  {
    o.paid = true
}

func NewOrder(id ID, product Product, address Address) (*Order, error){
    if len(id) == 0 {
        return nil, ErrEmptyID
    }

    return &Order{
        id:      id,
        product: product,
        address: address,
        paid:    false,
    }, nil
}
