package structs

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

var Products []Product;

func init() {
	productsData := []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Price string `json:"price"`
	}{
		{ID: 123, Name: "Ketchup", Price: "0.45"},
		{ID: 456, Name: "Beer", Price: "2.33"},
		{ID: 879, Name: "Õllesnäkk", Price: "0.42"},
		{ID: 999, Name: "75\" OLED TV", Price: "1333.37"},
	}

	for _, data := range productsData {
		product := Product{
			ID:    data.ID,
			Name:  data.Name,
			Price: data.Price,
		}
		Products = append(Products, product)
	}

}