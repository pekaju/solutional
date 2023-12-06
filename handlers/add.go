package handlers

import (
	"net/http"
	"github.com/pekaju/solutional/helpers"
	"github.com/pekaju/solutional/structs"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	"fmt"
	"strings"
)
// POST api/orders
func AddOrder(w http.ResponseWriter, r *http.Request) {
	order := helpers.CreateNewOrder()
	jsonResp, err := json.Marshal(order)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

// POST api/orders/{order_id}/products
func AddProductsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    orderID := vars["order_id"]
	orderIndex := helpers.FindOrderByID(orderID)
	if orderIndex == -1 {
		customErrorString(w, "Not found");
		return;
	}
	if helpers.CheckPostPatchHeaders(r.Header) == false {
		customErrorString(w, "Invalid parameters");
		return;
	}
	var productIDs []int = decodeProduct(w, r);
	if len(productIDs) == 0 {
		return;
	}
	if structs.Orders[orderIndex].Status == "PAID" {
		customErrorString(w, "Invalid parameters");
		return;
	}
	helpers.AddProductsToOrder(orderIndex, productIDs)
	helpers.CalculateOrderTotal(orderIndex)
	okResponse(w)
}

func decodeProduct(w http.ResponseWriter, r *http.Request) []int{
	var productIDs []int
	err := json.NewDecoder(r.Body).Decode(&productIDs)
	if err != nil {
		fmt.Println(err)
		if strings.Contains(err.Error(), "EOF") {
			BadRequest(w, r)
			return []int{}
		}
		if strings.Contains(err.Error(), "invalid character") {
			BadRequest(w, r)
			return []int{}
		}
		customErrorString(w, "Invalid parameters");
		return []int{}
	}
	if helpers.CheckProductIDs(productIDs) == false {
		customErrorString(w, "Invalid parameters")
		return []int{}
	}
	return productIDs
}
