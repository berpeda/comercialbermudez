package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/berpeda/comercialbermudez/database"
	"github.com/berpeda/comercialbermudez/models"
)

// This function tryes to INSERT a new Product at the Database
func PostProduct(user, body string) (int, string) {
	var product models.Product

	err := json.Unmarshal([]byte(body), &product)
	if err != nil {
		return http.StatusBadRequest, "There is an error with received data " + err.Error()
	}

	if product.IdCategory == 0 || product.IdProvider == 0 ||
		len(product.NameProduct) == 0 || len(product.DescriptionProduct) == 0 ||
		len(product.CodeProduct) == 0 || product.PriceProduct == 0.0 ||
		product.Stock == 0 || len(product.PathProduct) == 0 {
		return http.StatusBadRequest, "Product Name, Description, Code, Price, Stock, Path and the IDs " +
			"from the Provider and the Category needs to be filled at the Product fileds"
	}

	isAdmin, issue := database.IsAdmin(user)

	if !isAdmin {
		return http.StatusBadRequest, issue
	}

	result, err := database.InsertProduct(product)

	if err != nil {
		return http.StatusBadRequest, "Error trying to INSERT the Product > " + product.NameProduct + "\nError > " + err.Error()
	}

	return http.StatusOK, "{ IdProduct: " + strconv.Itoa(int(result)) + "}"
}

func GetProduct(request events.APIGatewayV2HTTPRequest) (int, string) {
	var product models.Product

	params := request.QueryStringParameters

	order := params["Order"]           // Desc or Asc/nil
	orderField := params["OrderField"] // 'N' Name 'P' Price
	page, _ := strconv.Atoi(params["Page"])
	pageSize, _ := strconv.Atoi(params["Size"])
	action := ""

	if len(params["IdProduct"]) > 0 {
		action = "P"
		product.IdProduct, _ = strconv.Atoi(params["IdProduct"])
	}

	if len(params["Search"]) > 0 {
		action = "S"
		product.SearchProduct, _ = params["Search"]
	}

	result, err := database.SelectProduct(product, action, page, pageSize, order, orderField)
	if err != nil {
		return http.StatusBadRequest, "Error trying to SELECT the product with Id > " + strconv.Itoa(product.IdProduct) +
			"\nError > " + err.Error()
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return http.StatusBadRequest, "Error trying to resialize result to JSON format."
	}

	return http.StatusOK, string(jsonData)
}

// func GetAllProducts() (int, string) {

// 	result, err := database.SelectAllProducts()
// 	if err != nil {
// 		return http.StatusBadRequest, "Error trying to SELECT all the Products. \nError > " + err.Error()
// 	}

// 	jsonData, err := json.Marshal(result)
// 	if err != nil {
// 		return http.StatusBadRequest, "Error trying to serialize result to JSON format."
// 	}

// 	return http.StatusOK, string(jsonData)
// }

func PutProduct(user, body string, idProduct int) (int, string) {

	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return http.StatusBadRequest, issue
	}

	var updateProduct models.Product
	err := json.Unmarshal([]byte(body), &updateProduct)
	if err != nil {
		return http.StatusBadRequest, "There is an issue with received data " + err.Error()
	}

	result, err := database.UpdateProduct(updateProduct, idProduct)
	if err != nil {
		return http.StatusBadRequest, "Error trying to UPDATE the product > " + strconv.Itoa(idProduct)
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return http.StatusBadRequest, "Error trying to serialize result to JSON format."
	}

	return http.StatusOK, string(jsonData)
}

func DeleteProduct(user string, idProduct int) (int, string) {

	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return http.StatusBadRequest, issue
	}

	result, err := database.DeleteProduct(idProduct)
	if err != nil {
		return http.StatusBadRequest, "Error trying to DELETE the product with ID > " + strconv.Itoa(int(result))
	}

	return http.StatusOK, "{ IdProduct: " + strconv.Itoa(int(result)) + "}"
}
