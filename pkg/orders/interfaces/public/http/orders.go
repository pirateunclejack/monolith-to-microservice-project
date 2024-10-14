package http

import "net/http"


type ordersResource struct {
    service    application.OrdersService
    repository orders.Repository
}

type PostOrderAddress struct {
    Name        string `json:"name"`
    Street      string `json:"street"`
    City        string `json:"city"`
    PostCode    string `json:"post_code"`
    Country     string `json:"country"`
}

type PostOrderRequest struct {
    ProductID orders.ProductID `json:"product_id"`
    Address   PostOrderAddress `json:"address"`
}

type PostOrderResponse struct {
    OrderID string
}

type OrderPaidView struct {
    ID string `json:"id"`
    IsPaid bool `json:"is_paid"`
}

func (o ordersResource) GetPaid(w http.ResponseWriter, r *http.Request) {
    order, err := o.repository.ByID(orders.ID(chi.URLParam(r, "id")))
    if err != nil {
        _ = render.Render(w, r, common_http.ErrBadRequest(err))
        return
    }

    render.Response(w, r, OrderPaidView{
        ID: string(order.ID()),
        IsPaid: order.Paid(),
    })
}

func (o ordersResource) Post(w http.ResponseWriter, r *http.Request){
    req := PostOrderRequest{}
    if err := render.Decode(r, &req); nil!= err {
        _ = render.Render(w, r, common_http.ErrBadRequest(err))
        return
    }

    cmd := application.PlaceOrderCommand{
        OrderID:     orders.OrderID(uuid.NewV1().string()),
        ProductID:   req.ProductID(),
        Address:     application.PlaceOrderCommandAddress(req.Address),
    }

    if err := o.service.PlaceOrder(cmd); nil!= err {
        _ = render.Render(w, r, common_http.ErrInternal(err))
         return
    }

    w.WriteHeader(http.StatusOK)
    render.JSON(w, r, PostOrderResponse{
        OrderID: string(cmd.OrderID),
    })
}

func AddRoutes(
    router     *chi.Mus,
    service    application.OrdersService,
    repository orders.Repository,
) {
    resource := orders.Resource{
        service, repository,
    }
    router.Post("/orders", resource.Post)
    router.Get("/orders/{id}/paid", resource.GetPaid)
}



