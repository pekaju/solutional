## Solutional 

This is my solution to homework application found [here](https://homework.solutional.ee/)

## Bugs

This api is missing the functionality to add a replacement product.

I was not able to get a positive response from the original api with this call:
```
curl -H "Content-Type: application/json" \
  -X PATCH \
  --data '{"replaced_with": {"product_id": 123, "quantity": 6}}' \
  https://homework.solutional.ee/api/orders/:order_id/products/:product_id
  ``` 
## Guide

This application runs a local server on default port 8080. Port can be configured when running the application with ```go run```, by adding an argument.
The argument must be a number from 1 to 65535.

## How to run

Clone the project:
````
git clone https://github.com/pekaju/solutional
````
Download the missing modules if needed
````
go mod download
````
And run the project
````
go run main.go
````
