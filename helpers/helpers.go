package helpers

import (
	"github.com/google/uuid"
	"github.com/pekaju/solutional/structs"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"io"
	"encoding/json"
)

func FindOrderByID(id string) int {
	for index, order := range structs.Orders {
		if order.ID == id {
			return index
		}
	}
	return -1
}
func findProductByID(id int) structs.Product{
	for _, product := range structs.Products {
		if product.ID == id {
			return product
		}
	}
	return structs.Product{}
}

func FindProductByOrder(orderIndex int, productID string) int{

	for index, orderProduct := range structs.Orders[orderIndex].Products {
		if orderProduct.ID == productID {
			return index;
		}
	}
	return -1;
}

func CreateNewOrder() structs.Order {
	newOrderID := uuid.New().String()

	newOrder := structs.Order{
		Amount: struct {
			Discount string `json:"discount"`
			Paid     string `json:"paid"`
			Returns  string `json:"returns"`
			Total    string `json:"total"`
		}{
			Discount: "0.00",
			Paid:     "0.00",
			Returns:  "0.00",
			Total:    "0.00",
		},
		ID:       newOrderID,
		Products: []structs.OrderProduct{},
		Status:   "NEW",
	}
	structs.Orders = append(structs.Orders, newOrder)
	return newOrder
}

func UpdateOrderStatus(orderIndex int) bool {
	if structs.Orders[orderIndex].Status != "NEW" {
		return false
	}
	structs.Orders[orderIndex].Status = "PAID"
	structs.Orders[orderIndex].Amount.Paid = structs.Orders[orderIndex].Amount.Total
	return true
}
func GetNumber(index int, integer string, s string ) string{
	if (s[index] >= '0' && s[index] <= '9') || s[index] == '.' {
		integer = integer + string(s[index])
		return GetNumber(index+1, integer, s)
	}
	return integer
}

func FindQuotesIndices(input string, nr int) []int {
	var indices []int

	for i, char := range input {
		if char == '"' {
			indices = append(indices, i)
			if len(indices) == nr {
				break
			}
		}
	}

	return indices
}

func CheckProductIDs(productIDs []int) bool{
	validProductIDs := make(map[int]bool)
	for _, product := range structs.Products {
		validProductIDs[product.ID] = true
	}

	for _, productID := range productIDs {
		if _, exists := validProductIDs[productID]; !exists {
			return false
		}
	}
	return true;
}

func AddProductsToOrder(orderIndex int, productIDs []int) {
	for _, productID := range productIDs {
		found := false
		for i, orderProduct := range structs.Orders[orderIndex].Products {
			if orderProduct.ProductID == productID {
				structs.Orders[orderIndex].Products[i].Quantity++
				found = true
				break
			}
		}

		if !found {
			product := findProductByID(productID)
			newProductID := uuid.New().String()
			if product.Name != "" {
				orderProduct := structs.OrderProduct{
					ID:        newProductID,
					Name:      product.Name,
					Price:     product.Price,
					ProductID: product.ID,
					Quantity:  1,
					ReplacedWith: nil,
				}
				structs.Orders[orderIndex].Products = append(structs.Orders[orderIndex].Products, orderProduct)
			}
		}
	}
}

func CheckPostPatchHeaders(header http.Header) bool {
	contentLength := header.Get("Content-Length")
	if contentLength == "" {
		return false;
	}
	if contentLength == "0" {
		return false
	}
	contentType := header.Get("Content-Type")
	if contentType == "" {
		return false;
	}
	if contentType != "application/json" && contentType != "application/x-www-form-urlencoded" {
		return false
	}

	return true
}

func CheckKeyValuePair(data map[string]string) bool {
	if len(data) == 0 {
		return false
	}
	//Only check first key/value pair. The rest will be ignored.
	for key, value := range data {
		if key != "status" || value != "PAID" {
			return false
		}else {
			return true
		}
	}
	return true
}

func CheckQuantityMap(data map[string]int) bool {
	if len(data) == 0 {
		return false
	}
	for key, value := range data {
		if key != "quantity" || value < 0 {
			return false
		}else {
			return true;
		}
	}
	return true
}

func CalculateOrderTotal(orderIndex int) {
	var total float64
	for _, orderProduct := range structs.Orders[orderIndex].Products {
		val, _ := strconv.ParseFloat(orderProduct.Price, 64)
		total += val * float64(orderProduct.Quantity)
	}
	structs.Orders[orderIndex].Amount.Total = fmt.Sprintf("%.2f", total)
}

func UpdateProductQuantity(orderIndex int, productIndex int, quantity int){
	structs.Orders[orderIndex].Products[productIndex].Quantity = quantity
}

func CheckIfStruct(strBody string) bool {
	return strBody[0] != '{' && strBody[len(strBody)-1] != '}'
}


func ParseRequestBody(r *http.Request) (string, bool) {
	var bodyCopy strings.Builder
	teeReader := io.TeeReader(r.Body, &bodyCopy)

	// Error check for invalid body by json decoding into a temporary empty struct
	if err := json.NewDecoder(teeReader).Decode(&struct{}{}); err != nil {
		return "", false
	}

	return bodyCopy.String(), true
}