package payments

import (
	"github.com/pirateunclejack/monolith-to-microservice-project/pkg/common/price"
	"github.com/pirateunclejack/monolith-to-microservice-project/pkg/orders/domain/orders"
	"github.com/pirateunclejack/monolith-to-microservice-project/pkg/payments/interfaces/intraprocess"
)

type IntraprocessService struct {
	orders chan<- intraprocess.OrderToProcess
}

func NewIntraprocessService(ordersChannel chan<- intraprocess.OrderToProcess) IntraprocessService {
	return IntraprocessService{ordersChannel}
}

func (i IntraprocessService) InitializeOrderPayment(id orders.ID, price price.Price) error {
	i.orders <- intraprocess.OrderToProcess{ID: string(id), Price: price}
	return nil
}
