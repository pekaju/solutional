package handlers

import (
	"net/http"
	"github.com/pekaju/solutional/structs"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	"github.com/pekaju/solutional/helpers"
)
/// GET api/products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(structs.Products)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

// GET api/orders/{order_id}
func GetSingleOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    orderID := vars["order_id"]
	orderIndex := helpers.FindOrderByID(orderID)
	if orderIndex != -1 {
		jsonResp, err := json.Marshal(structs.Orders[orderIndex])
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)
	} else {
		customErrorString(w, "Not Found");
	}
}

// GET api/orders/{order_id}/products
func GetOrderProductsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["order_id"]
	orderIndex := helpers.FindOrderByID(orderID)
	if orderIndex != -1 {
		jsonResp, err := json.Marshal(structs.Orders[orderIndex].Products)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)
	} else {
		customErrorString(w, "Not Found");
	}
}