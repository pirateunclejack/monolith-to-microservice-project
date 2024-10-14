package application

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

func (s ProductsService) AddProduct() error {}
