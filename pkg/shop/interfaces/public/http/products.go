package http

import "net/http"

type productsReadModel interface {
    AllProducts() ([]products.Product, error)
}

type productsResource struct {
    readModel productsReadModel
}

type productView struct {
    ID          uint      `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Price       priceView `json:"price"`
}

type priceView struct {
    Cents    uint   `json:"cents"`
    Currency string `json:"currency"`
}

func priceViewFromPrice(p price.Price) priceView {
    return priceView{
        p.Cents(),
        p.Currency(),
    }
}

func (p productsResource) GetAll(w http.ResponseWriter, r *http.Request){
    products, err := p.readModel.AllProducts()
    if err != nil {
        _ = render.Render(w, r, common_http.ErrInternal(err))
        return
    }

    view := []productView{}
    for _, product := range products {
        view = append(view, productView{
            string(product.ID()),
            product.Name(),
            product.Description(),
            priceViewFromPrice(product.Price()),
        })
    }

    render.Response(w, r, view)
}

func AddRoutes(router *chi.Mux, productsReadModel productsReadModel) {
    resource := productsResource{productsReadModel}
    router.Get("/products", resource.GetAll)
}
