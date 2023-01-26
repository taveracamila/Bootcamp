package store	


type storage struct{}




func (s *storage) cargarJSON() (list []domain.Product, err error){

	fname := os.Getenv("FILE_NAME")

	obj, err := os.ReadFile(path)

	if err != nil {
		return
	}

	json.Unmarshal(obj, &list)
 	return
}

//save products



func NewStorage(path string) Storage {
	return &storage{
		json: path,
	}
}




func (s *storage) GetAll() ([]domain.Product, error) {

	return cargarJSON()
}

func (s *storage) GetProductById(id int) (domain.Product, err error) {

	list, err := s.cargarJSON()
	if err != nil {
		return 
	}

	for _, item := range list {
		if item.Id == id {
			return item, nil
		}
	}

	return domain.Product{}, errors.New("No se encontro el producto")
}


func (s *storage) Create(p domain.Product) error {
	list, err := s.cargarJSON()

	if err != nil {
		return err
	}

	p.Id = len(products) + 1
	list = append(list, p)
	return s.saveProducts(p)
}

func (s *storage) Update(product domain.Product) error {
	products, err := s.loadProducts()
	if err != nil {
		return err
	}
	for i, p := range products {
		if p.Id == product.Id {
			products[i] = product
			return s.saveProducts(products)
		}
	}
	return errors.New("product not found")
}

func (s *storage) UpdatePrice(product domain.Product, price float64) (domain.Product, error) {
	products, err := s.loadProducts()
	if err != nil {
		return domain.Product{}, err
	}
	for i, p := range products {
		if p.Id == product.Id {
			products[i].Price = price
			return products[i], s.saveProducts(products)
		}
	}
	return domain.Product{}, errors.New("product not found")
}

func (s *storage) Delete(id int) error {
	products, err := s.loadProducts()
	if err != nil {
		return err
	}
	for i, p := range products {
		if p.Id == id {
			products = append(products[:i], products[i+1:]...)
			return s.saveProducts(products)
		}
	}
	return errors.New("product not found")
}