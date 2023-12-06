package handlers

import (
	"fmt"
	"net/http"
	"github.com/pekaju/solutional/helpers"
	"github.com/pekaju/solutional/structs"
	"strconv"
	"strings"
	"github.com/gorilla/mux"
)

// PATCH /api/orders/{order_id}
func UpdateOrderStatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["order_id"]
	orderIndex := helpers.FindOrderByID(orderID)
	if orderIndex == -1 {
		customErrorString(w, "Not found")
		return
	}
	if helpers.CheckPostPatchHeaders(r.Header) == false {
		customErrorString(w, "Invalid parameters")
		return
	}
	strBody, success := helpers.ParseRequestBody(r)
	if success == false {
		BadRequest(w, r)
		return
	}
	if !helpers.CheckIfStruct(strBody) {
		customErrorString(w, "Invalid order status")
		return
	}
	//find first 4 double quotes to get first 2 strings.
	var indices []int = helpers.FindQuotesIndices(strBody, 4)
	if len(indices) != 4 {
		customErrorString(w, "Invalid order status")
		return
	}
	//value of first string which has to be the first key
	keyValue := strBody[indices[0]+1 : indices[1]]
	//value of second string
	statusValue := strBody[indices[2]+1 : indices[3]]
	//check if only whitespaces and ":" are between first key and second string
	//that means second string is value for first key.
	checker := strBody[indices[1]+1 : indices[2]]
	for _, char := range checker {
		if char != ' ' && char != ':' && char != '\t' {
			customErrorString(w, "Invalid order status")
			return
		}
	}
	if keyValue != "status" || statusValue != "PAID" {
		customErrorString(w, "Invalid order status")
		return
	}
	if helpers.UpdateOrderStatus(orderIndex) == false {
		customErrorString(w, "Invalid order status")
		return
	}
	okResponse(w)
}

// PATCH /api/orders/{order_id}/products/{product_id}
func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["order_id"]
	orderIndex := helpers.FindOrderByID(orderID)
	if orderIndex == -1 {
		customErrorString(w, "Not found")
		return
	}
	productID := vars["product_id"]
	productIndex := helpers.FindProductByOrder(orderIndex, productID)
	if productIndex == -1 {
		customErrorString(w, "Not found")
		return
	}
	if helpers.CheckPostPatchHeaders(r.Header) == false {
		customErrorString(w, "Invalid parameters")
		return
	}
	strBody, check := helpers.ParseRequestBody(r)
	if check == false {
		BadRequest(w, r)
		return
	}
	if !helpers.CheckIfStruct(strBody) {
		customErrorString(w, "Invalid order status")
		return
	}
	// Finding the first 2 quotes to determine the first string, 
	// which has to be the first key value
	var indices []int = helpers.FindQuotesIndices(strBody, 2)
	if len(indices) != 2 {
		customErrorString(w, "Invalid parameters")
		return
	}
	//checking if the first key value is "quantity"
	keyValue := strBody[indices[0]+1 : indices[1]]
	if keyValue != "quantity" {
		customErrorString(w, "Invalid parameters")
		return
	}
	// clear whitespaces in the request body
	noSpaces := strings.ReplaceAll(strBody, " ", "")
	newIndices := helpers.FindQuotesIndices(noSpaces, 2)
	// find the first index of first value
	firstIntIndex := newIndices[1] + 2
	// make full integer recursively
	var quantityValue = helpers.GetNumber(firstIntIndex, "", noSpaces)
	// float is invalid
	if strings.Contains(quantityValue, ".") {
		customErrorString(w, "Invalid parameters")
		return
	}
	quantityValueInt, err := strconv.Atoi(quantityValue)
	if err != nil {
		fmt.Println("err converting from string to int")
		return 
	}
	if structs.Orders[orderIndex].Status == "PAID" {
		customErrorString(w, "Invalid parameters")
		return
	}
	helpers.UpdateProductQuantity(orderIndex, productIndex, quantityValueInt)
	helpers.CalculateOrderTotal(orderIndex)
	okResponse(w)
}