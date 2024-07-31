package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/berpeda/comercialbermudez/database"
	"github.com/berpeda/comercialbermudez/models"
)

// This function tries to INSERT a new Product into the Database
func PostProduct(user, body string) (int, string) {
	var product models.Product

	// Deserialize the JSON body into a Product struct
	err := json.Unmarshal([]byte(body), &product)
	if err != nil {
		return http.StatusBadRequest, "There is an error with received data " + err.Error()
	}

	// Check if all required fields are filled
	if product.IdCategory == 0 || product.IdProvider == 0 ||
		len(product.NameProduct) == 0 || len(product.DescriptionProduct) == 0 ||
		len(product.CodeProduct) == 0 || product.PriceProduct == 0.0 ||
		product.Stock == 0 || len(product.PathProduct) == 0 {
		return http.StatusBadRequest, "Product Name, Description, Code, Price, Stock, Path and the IDs " +
			"from the Provider and the Category need to be filled in the Product fields"
	}

	// Check if the user is an admin
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return http.StatusBadRequest, issue
	}

	// Insert the product into the database
	result, err := database.InsertProduct(product)
	if err != nil {
		return http.StatusBadRequest, "Error trying to INSERT the Product > " + product.NameProduct + "\nError > " + err.Error()
	}

	// Return a success status and the ID of the inserted product
	return http.StatusOK, "{ IdProduct: " + strconv.Itoa(int(result)) + "}"
}

// GetProduct retrieves a product or a list of products based on query parameters
func GetProduct(request events.APIGatewayV2HTTPRequest) (int, string) {
	var product models.Product

	// Extract query parameters from the request
	params := request.QueryStringParameters

	order := params["Order"]           // Desc or Asc/nil
	orderField := params["OrderField"] // 'N' Name 'P' Price
	page, _ := strconv.Atoi(params["Page"])
	pageSize, _ := strconv.Atoi(params["Size"])
	action := ""

	// Check if the IdProduct parameter is present
	if len(params["IdProduct"]) > 0 {
		action = "P"
		product.IdProduct, _ = strconv.Atoi(params["IdProduct"])
	}

	// Check if the Search parameter is present
	if len(params["Search"]) > 0 {
		action = "S"
		product.SearchProduct = params["Search"]
	}

	// Fetch the product(s) from the database based on the provided parameters
	result, err := database.SelectProduct(product, action, page, pageSize, order, orderField)
	if err != nil {
		return http.StatusBadRequest, "Error trying to SELECT the product with Id > " + strconv.Itoa(product.IdProduct) +
			"\nError > " + err.Error()
	}

	// Serialize the result to JSON format
	jsonData, err := json.Marshal(result)
	if err != nil {
		return http.StatusBadRequest, "Error trying to serialize result to JSON format."
	}

	// Return a success status and the JSON result
	return http.StatusOK, string(jsonData)
}

// PutProduct updates an existing product, but only if the user is an admin
func PutProduct(user, body string, idProduct int) (int, string) {
	// Check if the user is an admin
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return http.StatusBadRequest, issue
	}

	var updateProduct models.Product

	// Deserialize the JSON body into a Product struct
	err := json.Unmarshal([]byte(body), &updateProduct)
	if err != nil {
		return http.StatusBadRequest, "There is an issue with received data " + err.Error()
	}

	// Update the product in the database
	result, err := database.UpdateProduct(updateProduct, idProduct)
	if err != nil {
		return http.StatusBadRequest, "Error trying to UPDATE the product > " + strconv.Itoa(idProduct)
	}

	// Serialize the result to JSON format
	jsonData, err := json.Marshal(result)
	if err != nil {
		return http.StatusBadRequest, "Error trying to serialize result to JSON format."
	}

	// Return a success status and the JSON result
	return http.StatusOK, string(jsonData)
}

// DeleteProduct deletes a product, but only if the user is an admin
func DeleteProduct(user string, idProduct int) (int, string) {
	// Check if the user is an admin
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return http.StatusBadRequest, issue
	}

	// Delete the product from the database
	result, err := database.DeleteProduct(idProduct)
	if err != nil {
		return http.StatusBadRequest, "Error trying to DELETE the product with ID > " + strconv.Itoa(int(result))
	}

	// Return a success status and the ID of the deleted product
	return http.StatusOK, "{ IdProduct: " + strconv.Itoa(int(result)) + "}"
}
