package models

type Collection struct {
	Name     string
	Price    string
	Image    string
	Features map[string]string
	Products []Product
}

func NewCollection(name string, price string, image string, features map[string]string, products []Product) Collection {
	return Collection{
		Name:     name,
		Price:    price,
		Image:    image,
		Features: features,
		Products: products,
	}
}
