package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"github.com/pekaju/solutional/handlers"
	"strconv"
	"strings"
)

//golang seems to have a problem with multiple slashes IN THE MIDDLE OF A REQUEST path
//this middleware calls the error handler if multple slashes are detected.
func RemoveTrailingSlashMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimRight(r.URL.Path, "/")
		if !strings.Contains(path, "//") {
			r.URL.Path = path
			next.ServeHTTP(w, r)
			return
		}
		handlers.NotFound(w, r)
		return
	})
}

func main() {
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(handlers.NotFound)
	router.MethodNotAllowedHandler = http.HandlerFunc(handlers.NotFound)
	router.HandleFunc("/api/orders/{order_id}/products/{product_id}", handlers.UpdateProductHandler).Methods("PATCH")
	router.HandleFunc("/api/orders/{order_id}/products", handlers.AddProductsHandler).Methods("POST")
	router.HandleFunc("/api/orders/{order_id}/products", handlers.GetOrderProductsHandler).Methods("GET")
	router.HandleFunc("/api/orders/{order_id}", handlers.GetSingleOrder).Methods("GET")
	router.HandleFunc("/api/orders/{order_id}", handlers.UpdateOrderStatusHandler).Methods("PATCH")
	router.HandleFunc("/api/products", handlers.GetProducts).Methods("GET")
	router.HandleFunc("/api/orders", handlers.AddOrder).Methods("POST")

	port := 8080
	if len(os.Args) > 1 {
		argPort := os.Args[1]
		numPort, err := strconv.Atoi(argPort)
		if err == nil {
			if numPort >= 1 && numPort <= 65535 {
				port = numPort
			} else {
				fmt.Println("Port number out of range. Please provide a valid port number between 1 and 65535.")
				return
			}
		} else {
			fmt.Println("Invalid port number. Please provide a valid integer.")
			return
		}
	}

	fmt.Printf("Server is running on port %d\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), RemoveTrailingSlashMiddleware(router))
	if err != nil {
		fmt.Println(err)
	}
}