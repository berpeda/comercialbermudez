package database

import (
	"fmt"

	"github.com/berpeda/comercialbermudez/models"
	"github.com/berpeda/comercialbermudez/tools"
)

func InsertProduct(product models.Product) (int64, error) {
	fmt.Println("Insert Product function starts...")

	err := DatabaseConnect()
	if err != nil {
		return 0, err
	}

	defer Database.Close()

	query := "INSERT INTO Productos (Id_proveedor, Id_categoria, Codigo, Nombre, Descripcion, Precio, Creado, Stock) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := Database.Exec(query, product.IdProvider, product.IdCategory, product.CodeProduct, product.NameProduct, product.DescriptionProduct, product.PriceProduct, tools.DateMySQL(), product.Stock)
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

func SelectProduct(idProduct int) (models.Product, error) {
	fmt.Println("Select a single Product function starts...")

	var nProduct models.Product

	err := DatabaseConnect()
	if err != nil {
		return nProduct, err
	}

	defer Database.Close()

	query := "SELECT * FROM Productos WHERE Id_producto = ?"

	result, err := Database.Query(query, idProduct)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return nProduct, err
	}

	result.Next()
	err2 := result.Scan(&nProduct.IdProduct,
		&nProduct.IdProvider,
		&nProduct.IdCategory,
		&nProduct.CodeProduct,
		&nProduct.NameProduct,
		&nProduct.DescriptionProduct,
		&nProduct.PriceProduct,
		&nProduct.CreatedAt,
		&nProduct.UpdatedAt,
		&nProduct.Stock)

	if err2 != nil {
		fmt.Println("result.Scan is having issues...")
		return nProduct, err2
	}

	// rowsAffected, err := result.RowsAffected()
	// if err != nil {
	// 	fmt.Println("Error retrieving the number of rows affected > ", err.Error())
	// 	return nProduct, err
	// }

	// lastInsertId, err := result.LastInsertId()
	// if err != nil {
	// 	fmt.Println("Error retrieving the number of rows affected > ", err.Error())
	// }

	// fmt.Printf("Product selected successfully.\n Index product selected > %d\n The row(s) affected > %d", lastInsertId, rowsAffected)
	fmt.Printf("Product selected successfully.")

	return nProduct, nil
}

func SelectAllProducts() ([]models.Product, error) {
	fmt.Println("Select All Products is starting...")

	var products []models.Product

	err := DatabaseConnect()
	if err != nil {
		return products, err
	}

	defer Database.Close()

	query := "SELECT * FROM Productos"
	result, err := Database.Query(query)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return products, err
	}

	for result.Next() {
		var product models.Product
		err = result.Scan(&product.IdProduct,
			&product.IdProvider,
			&product.IdCategory,
			&product.CodeProduct,
			&product.NameProduct,
			&product.DescriptionProduct,
			&product.PriceProduct,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.Stock)

		if err != nil {
			fmt.Println("Unable to Scan all the products > " + err.Error())
			panic(err)
		}
		products = append(products, product)
	}

	if err = result.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Products Selected successfully!")

	return products, nil
}
