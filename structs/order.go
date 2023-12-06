package structs

type Order struct {
	Amount   struct {
		Discount string `json:"discount"`
		Paid     string `json:"paid"`
		Returns  string `json:"returns"`
		Total    string `json:"total"`
	} `json:"amount"`
	ID       string   `json:"id"`
	Products []OrderProduct `json:"products"`
	Status   string   `json:"status"`
}

type OrderProduct struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Price        string  `json:"price"`
	ProductID    int     `json:"product_id"`
	Quantity     int     `json:"quantity"`
	ReplacedWith *string `json:"replaced_with"`
}

var Orders []Order;