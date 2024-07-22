package routers

import (
	"encoding/json"
	"strconv"

	"github.com/berpeda/comercialbermudez/database"
	"github.com/berpeda/comercialbermudez/models"
)

// This function tryes to INSERT a new Product at the Database
func PostProduct(user, body string) (int, string) {
	var product models.Product

	err := json.Unmarshal([]byte(body), &product)
	if err != nil {
		return 400, "There is an error with received data " + err.Error()
	}

	if product.IdCategory == 0 || product.IdProvider == 0 ||
		len(product.NameProduct) == 0 || len(product.DescriptionProduct) == 0 ||
		len(product.CodeProduct) == 0 || product.PriceProduct == 0.0 ||
		product.Stock == 0 {
		return 400, "Product Name, Description, Code, Price, Stock and the IDs " +
			"from the Provider and the Category needs to be filled at the Product fileds"
	}

	isAdmin, issue := database.IsAdmin(user)

	if !isAdmin {
		return 400, issue
	}

	result, err := database.InsertProduct(product)

	if err != nil {
		return 400, "Error trying to INSERT the Product > " + product.NameProduct + "\nError > " + err.Error()
	}

	return 200, "{IdProduct: " + strconv.Itoa(int(result)) + "}"
}

func GetProduct(idProduct int) (int, string) {

	result, err := database.SelectProduct(idProduct)
	if err != nil {
		return 400, "Error trying to SELECT the product with Id > " + strconv.Itoa(idProduct) +
			"\nError > " + err.Error()
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return 400, "Error trying to resialize result to JSON format."
	}

	return 200, string(jsonData)
}

func GetAllProducts() (int, string) {

	result, err := database.SelectAllProducts()
	if err != nil {
		return 400, "Error trying to SELECT all the Products. \nError > " + err.Error()
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return 400, "Error trying to serialize result to JSON format."
	}

	return 200, string(jsonData)
}

// func UpdateProduct(user, body string, idProduct int) (int, string) {
// 	var updateProduct models.Product
// 	r, productToUpdate := GetProduct(idProduct)

// 	if r == 400 {
// 		return 400, "No products found with id > " + strconv.Itoa(idProduct)
// 	}

// 	err := json.Marshal(productToUpdate)
// 	// CONTINUAR AQUI
// }
