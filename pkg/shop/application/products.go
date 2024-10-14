package application

import "errors"

type productReadModel interface {
    AllProducts() ([]products.Product, error)
}

type ProductsService struct {
    repo products.Repository
    readModel productReadModel
}

func NewProductsService() ProductsService {

}

func (s ProductsService) AllProducts() () {
    
}

type AddProductCommand struct {
    ID              string
    Name            string
    Description     string
    PriceCents      uint
    PriceCurrency   string
}

func (s ProductsService) AddProduct(cmd AddProductCommand) error {
    price, err := products.NewPrice(cmd.PriceCents, cmd.PriceCurrency)
    if err != nil {
        return errors.Wrap(err, "invalid product price")
    }

    p, err := products.NewProduct(products.ID(cmd.ID), cmd.Name, cmd.Description, price)
    if err != nil {
        return errors.Wrap(err, "cannot create product")
    }

    if err := s.repo.Save(p); err != nil {
        return errors.Wrap(err, "cannot save product")
    }

    return nil
}
