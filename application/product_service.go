package application

type ProductService struct {
	ProductPersistence ProductPersistenceInterface
}

fun (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.ProductPersistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}