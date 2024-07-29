package database

import (
	"fmt"

	"github.com/berpeda/comercialbermudez/models"
)

func InsertCategory(category models.Category) (int64, error) {
	fmt.Println("Insert Category function starts...")

	err := DatabaseConnect()
	if err != nil {
		return 0, err
	}

	defer Database.Close()

	query := "INSERT INTO Categorias (Nombre, Descripcion) VALUES (?,?)"

	result, err := Database.Exec(query, category.NameCategory, category.DescriptionCategory)
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

	fmt.Printf("Category inserted successfully.\nIndex inserted > %d\n The row(s) affected > %d", lastInsertId, rowsAffected)

	return lastInsertId, nil
}

func SelectCategory(idCategory int) (models.ProductDetails, error) {
	fmt.Println("Select all products with the same category function starts...")

	var productsCat models.ProductDetails

	err := DatabaseConnect()
	if err != nil {
		return productsCat, err
	}

	defer Database.Close()

	query := "SELECT Id_producto, Id_proveedor, Id_categoria, Codigo, Nombre, Precio FROM Productos WHERE Id_categoria = ?"

	result, err := Database.Query(query, idCategory)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return productsCat, err
	}

	for result.Next() {
		var p models.Product
		err = result.Scan(
			&p.IdProduct,
			&p.IdProvider,
			&p.IdCategory,
			&p.CodeProduct,
			&p.NameProduct,
			&p.PriceProduct)

		if err != nil {
			fmt.Println("result.Scan is having issues...")
			return productsCat, err
		}
		productsCat.TotalProducts = append(productsCat.TotalProducts, p)
	}

	fmt.Printf("Products with a specific category (%d) selected successfully.", idCategory)
	return productsCat, nil
}

func SelectAllCategories() ([]models.Category, error) {
	fmt.Println("Select All Categories is starting...")

	var categories []models.Category

	err := DatabaseConnect()
	if err != nil {
		return categories, err
	}

	defer Database.Close()

	query := "SELECT * FROM Categorias"
	result, err := Database.Query(query)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return categories, err
	}

	for result.Next() {
		var category models.Category
		err = result.Scan(&category.IdCategory, &category.NameCategory, &category.DescriptionCategory)
		if err != nil {
			fmt.Println("Unable to Scan all the products > " + err.Error())
			panic(err)
		}
		categories = append(categories, category)
	}

	if err = result.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Categories Selected successfully!")

	return categories, nil
}

func DeleteCategory(idCategory int) (int64, error) {
	fmt.Println("Delete a category and its products starting...")

	err := DatabaseConnect()
	if err != nil {
		return 0, err
	}

	defer Database.Close()

	query := "DELETE FROM Categorias WHERE Id_categoria = ?"
	result, err := Database.Exec(query, idCategory)
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

	fmt.Printf("Category deleted successfully.\nIndex deleted > %d\n The row(s) affected > %d", lastInsertId, rowsAffected)

	return int64(idCategory), nil

}

func UpdateCategory(nCategory models.Category, idCategory int) (int64, error) {
	fmt.Println("Delete a category and its products starting...")

	err := DatabaseConnect()
	if err != nil {
		return 0, err
	}

	defer Database.Close()

	query := "UPDATE Categorias SET Nombre = ?, Descripcion = ? WHERE Id_categoria = ?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		fmt.Println("Error preparing the query > ", err.Error())
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(nCategory.NameCategory, nCategory.DescriptionCategory, idCategory)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error retrieving the number of rows affected > ", err.Error())
		return 0, err
	}

	fmt.Printf("Category updated successfully.\nIndex updated > %d\n The row(s) affected > %d", idCategory, rowsAffected)

	return int64(idCategory), nil

}
