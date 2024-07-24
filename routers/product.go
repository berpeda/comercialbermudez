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

func PutProduct(user, body string, idProduct int) (int, string) {

	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return 400, issue
	}

	var updateProduct models.Product
	err := json.Unmarshal([]byte(body), &updateProduct)
	if err != nil {
		return 400, "There is an issue with received data " + err.Error()
	}

	if updateProduct.IdCategory == 0 || updateProduct.IdProvider == 0 ||
		len(updateProduct.NameProduct) == 0 || len(updateProduct.DescriptionProduct) == 0 ||
		len(updateProduct.CodeProduct) == 0 || updateProduct.PriceProduct == 0 ||
		updateProduct.Stock < 0 {

		return 400, "Any of the product's attributes needs to be filled (The creation and the update isn't necessary to be filled)"
	}

	result, err := database.UpdateProduct(updateProduct, idProduct)
	if err != nil {
		return 400, "Error trying to UPDATE the product > " + strconv.Itoa(idProduct)
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return 400, "Error trying to serialize result to JSON format."
	}

	return 200, string(jsonData)
}

func DeleteProduct(user string, idProduct int) (int, string) {

	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return 400, issue
	}

	result, err := database.DeleteProduct(idProduct)
	if err != nil {
		return 400, "Error trying to DELETE the product with ID > " + strconv.Itoa(int(result))
	}

	return 200, "{ IdProduct: " + strconv.Itoa(int(result)) + "}"
}
