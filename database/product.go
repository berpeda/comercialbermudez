package database

import (
	"fmt"
	"math"
	"strings"

	"github.com/berpeda/comercialbermudez/models"
	"github.com/berpeda/comercialbermudez/tools"
)

// InsertProduct adds a new product to the database and returns its ID.
func InsertProduct(product models.Product) (int64, error) {
	fmt.Println("Insert Product function starts...")

	// Connect to the database
	err := DatabaseConnect()
	if err != nil {
		return 0, err
	}

	defer Database.Close()

	// Round the price to 2 decimal places
	roundPrice := math.Round(product.PriceProduct*100) / 100

	// Query to insert a new product
	query := "INSERT INTO Productos (Id_proveedor, Id_categoria, Codigo, Nombre, Descripcion, Precio, Creado, Stock, Ruta) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := Database.Exec(query, product.IdProvider, product.IdCategory, product.CodeProduct, product.NameProduct, product.DescriptionProduct, roundPrice, tools.DateMySQL(), product.Stock, product.PathProduct)

	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error retrieving the number of rows affected > ", err.Error())
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error retrieving the number of rows affected > ", err.Error())
	}

	fmt.Printf("Product inserted successfully.\nIndex inserted > %d\n The row(s) affected > %d", lastInsertId, rowsAffected)

	return lastInsertId, nil
}

// SelectProduct retrieves product(s) based on the specified criteria.
func SelectProduct(product models.Product, action string, page, pageSize int, order, orderField string) (models.ProductDetails, error) {
	fmt.Println("Select a single Product function starts...")

	var nProduct models.Product
	var resultSelect models.ProductDetails
	var productsT []models.Product

	// Connect to the database
	err := DatabaseConnect()
	if err != nil {
		return resultSelect, err
	}

	defer Database.Close()

	// Base query to select products
	query := "SELECT * FROM Productos"
	where := ""
	params := []interface{}{}

	// Build the WHERE clause based on the action
	if action == "P" {
		where = " WHERE Id_producto = ?"
		params = append(params, product.IdProduct)
	} else if action == "S" {
		where = " WHERE UCASE(CONCAT(Nombre, Descripcion)) LIKE ?"
		params = append(params, "%"+strings.ToUpper(product.SearchProduct)+"%")
	}

	if where != "" {
		query += where
	}

	// Add ordering to the query
	if len(orderField) != 0 {
		if orderField == "P" {
			query += " ORDER BY Precio"
		} else if orderField == "N" {
			query += " ORDER BY Nombre"
		}
		if order == "Desc" {
			query += " DESC"
		}
	}

	// Add pagination to the query
	if page > 0 && pageSize > 0 {
		offset := pageSize * (page - 1)
		query += fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, offset)
	}

	fmt.Println("The sentence is > ", query)
	fmt.Println("The params for the sentence are > ", params)

	result, err := Database.Query(query, params...)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return resultSelect, err
	}

	defer result.Close()

	// Scan the results into the products slice
	for result.Next() {
		err2 := result.Scan(&nProduct.IdProduct,
			&nProduct.IdProvider,
			&nProduct.IdCategory,
			&nProduct.CodeProduct,
			&nProduct.NameProduct,
			&nProduct.DescriptionProduct,
			&nProduct.PriceProduct,
			&nProduct.CreatedAt,
			&nProduct.UpdatedAt,
			&nProduct.Stock,
			&nProduct.PathProduct)

		if err2 != nil {
			fmt.Println("result.Scan is having issues...")
			return resultSelect, err2
		}

		productsT = append(productsT, nProduct)
	}

	resultSelect.TotalProducts = productsT

	fmt.Printf("Products selected successfully.")

	return resultSelect, nil
}

// UpdateProduct updates the product details based on the provided ID.
func UpdateProduct(p models.Product, idProduct int) (models.Product, error) {
	fmt.Println("Update Product is starting...")

	// Connect to the database
	err := DatabaseConnect()
	if err != nil {
		return p, err
	}
	defer Database.Close()

	// Build the UPDATE query dynamically based on provided fields
	query := "UPDATE Productos SET"
	params := []interface{}{}
	if p.IdProvider != 0 {
		query += " Id_proveedor = ?,"
		params = append(params, p.IdProvider)
	}

	if p.IdCategory != 0 {
		query += " Id_categoria = ?,"
		params = append(params, p.IdCategory)
	}

	if len(p.CodeProduct) != 0 {
		query += " Codigo = ?,"
		params = append(params, p.CodeProduct)
	}

	if len(p.NameProduct) != 0 {
		query += " Nombre = ?,"
		params = append(params, p.NameProduct)
	}

	if len(p.DescriptionProduct) != 0 {
		query += " Descripcion = ?,"
		params = append(params, p.DescriptionProduct)
	}

	if p.PriceProduct != 0 {
		query += " Precio = ?,"
		params = append(params, p.PriceProduct)
	}

	if p.Stock > 0 {
		query += " Stock = ?,"
		params = append(params, p.Stock)
	}

	if len(p.PathProduct) != 0 {
		query += " Ruta = ?,"
		params = append(params, p.PathProduct)
	}

	// Remove the trailing comma
	query = query[:len(query)-1]
	query += " WHERE Id_producto = ?"
	params = append(params, idProduct)

	_, err = Database.Exec(query, params...)
	if err != nil {
		return p, err
	}

	// Query to retrieve the updated product
	query = "SELECT Id_producto, Id_proveedor, Id_categoria, Codigo, Nombre, Descripcion, Precio, Creado, Actualizado, Stock, Ruta FROM Productos WHERE Id_producto = ?"
	result, err2 := Database.Query(query, idProduct)
	if err2 != nil {
		return p, err2
	}

	defer result.Close()

	// Scan the updated result into the product struct
	result.Next()
	err = result.Scan(&p.IdProduct, &p.IdProvider, &p.IdCategory, &p.CodeProduct, &p.NameProduct, &p.DescriptionProduct, &p.PriceProduct, &p.CreatedAt, &p.UpdatedAt, &p.Stock, &p.PathProduct)
	if err != nil {
		return p, err
	}

	fmt.Println("The product has been updated successfully!")
	return p, nil
}

// DeleteProduct removes a product from the database based on its ID.
func DeleteProduct(idProduct int) (int64, error) {
	fmt.Println("Delete Product is starting...")

	// Connect to the database
	err := DatabaseConnect()
	if err != nil {
		return 0, err
	}

	defer Database.Close()

	// Query to delete the product by its ID
	query := "DELETE FROM Productos WHERE Id_producto = ?"
	result, err := Database.Exec(query, idProduct)
	if err != nil {
		fmt.Println("There is an error in query > " + err.Error())
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error retrieving the number of rows affected > ", err.Error())
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error retrieving the number of rows affected > ", err.Error())
	}

	fmt.Printf("Product deleted successfully.\nIndex deleted > %d\n The row(s) affected > %d", lastInsertId, rowsAffected)

	return int64(idProduct), nil
}
