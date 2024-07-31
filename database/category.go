package database

import (
	"fmt"

	"github.com/berpeda/comercialbermudez/models"
)

// InsertCategory adds a new category to the database
func InsertCategory(category models.Category) (int64, error) {
	fmt.Println("Insert Category function starts...")

	// Connect to the database
	err := DatabaseConnect()
	if err != nil {
		return 0, err
	}
	defer Database.Close()

	// Query to insert a new category
	query := "INSERT INTO Categorias (Nombre, Descripcion) VALUES (?,?)"
	result, err := Database.Exec(query, category.NameCategory, category.DescriptionCategory)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return 0, err
	}

	// Get the last inserted ID and rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error retrieving the number of rows affected > ", err.Error())
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error retrieving the last inserted ID > ", err.Error())
	}

	fmt.Printf("Category inserted successfully.\nIndex inserted > %d\n The row(s) affected > %d", lastInsertId, rowsAffected)

	return lastInsertId, nil
}

// SelectCategory retrieves products associated with a specific category
func SelectCategory(idCategory int) (models.ProductDetails, error) {
	fmt.Println("Select all products with the same category function starts...")

	var productsCat models.ProductDetails

	// Connect to the database
	err := DatabaseConnect()
	if err != nil {
		return productsCat, err
	}
	defer Database.Close()

	// Query to select products with the given category ID
	query := "SELECT * FROM Productos WHERE Id_categoria = ?"
	result, err := Database.Query(query, idCategory)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return productsCat, err
	}
	defer result.Close()

	// Iterate over the results and scan into Product struct
	for result.Next() {
		var p models.Product
		err = result.Scan(
			&p.IdProduct,
			&p.IdProvider,
			&p.IdCategory,
			&p.CodeProduct,
			&p.NameProduct,
			&p.DescriptionProduct,
			&p.PriceProduct,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.Stock,
			&p.PathProduct)
		if err != nil {
			fmt.Println("result.Scan is having issues...")
			return productsCat, err
		}
		productsCat.TotalProducts = append(productsCat.TotalProducts, p)
	}

	fmt.Printf("Products with a specific category (%d) selected successfully.", idCategory)
	return productsCat, nil
}

// SelectAllCategories retrieves all categories from the database
func SelectAllCategories() ([]models.Category, error) {
	fmt.Println("Select All Categories is starting...")

	var categories []models.Category

	// Connect to the database
	err := DatabaseConnect()
	if err != nil {
		return categories, err
	}
	defer Database.Close()

	// Query to select all categories
	query := "SELECT * FROM Categorias"
	result, err := Database.Query(query)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return categories, err
	}
	defer result.Close()

	// Iterate over the results and scan into Category struct
	for result.Next() {
		var category models.Category
		err = result.Scan(&category.IdCategory, &category.NameCategory, &category.DescriptionCategory)
		if err != nil {
			fmt.Println("Unable to Scan all the categories > " + err.Error())
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

// DeleteCategory removes a category from the database
func DeleteCategory(idCategory int) (int64, error) {
	fmt.Println("Delete a category and its products starting...")

	// Connect to the database
	err := DatabaseConnect()
	if err != nil {
		return 0, err
	}
	defer Database.Close()

	// Query to delete a category
	query := "DELETE FROM Categorias WHERE Id_categoria = ?"
	result, err := Database.Exec(query, idCategory)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return 0, err
	}

	// Get the number of rows affected and the last inserted ID (not applicable here)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error retrieving the number of rows affected > ", err.Error())
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error retrieving the last inserted ID > ", err.Error())
	}

	fmt.Printf("Category deleted successfully.\nIndex deleted > %d\n The row(s) affected > %d", lastInsertId, rowsAffected)

	return int64(idCategory), nil
}

// UpdateCategory updates an existing category in the database
func UpdateCategory(nCategory models.Category, idCategory int) (int64, error) {
	fmt.Println("Update a category function starts...")

	// Connect to the database
	err := DatabaseConnect()
	if err != nil {
		return 0, err
	}
	defer Database.Close()

	// Query to update a category
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

	// Get the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error retrieving the number of rows affected > ", err.Error())
		return 0, err
	}

	fmt.Printf("Category updated successfully.\nIndex updated > %d\n The row(s) affected > %d", idCategory, rowsAffected)

	return int64(idCategory), nil
}
