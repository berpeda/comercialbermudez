package database

import (
	"fmt"
	"math"

	"github.com/berpeda/comercialbermudez/models"
)

// SelectOrderDetail retrieves a single order detail from the database based on the provided ID.
func SelectOrderDetail(idOrderDetail int) (models.OrderDetails, error) {
	fmt.Println("Select a single Order detail function starts...")

	var oDetail models.OrderDetails

	// Connect to the database
	err := DatabaseConnect()
	if err != nil {
		return oDetail, err
	}

	defer Database.Close()

	// Query to select the order detail by its ID
	query := "SELECT * FROM Detalle_pedido WHERE Id_detalle_pedido = ?"

	result, err := Database.Query(query, idOrderDetail)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return oDetail, err
	}

	defer result.Close()

	// Scan the result into the order detail struct
	result.Next()
	err2 := result.Scan(&oDetail.IdOrderDetail,
		&oDetail.IdOrder,
		&oDetail.IdProduct,
		&oDetail.QuantityOrderDetail,
		&oDetail.PriceOrderDetail)

	if err2 != nil {
		fmt.Println("result.Scan is having issues...")
		return oDetail, err2
	}

	fmt.Printf("Order detail selected successfully.")

	return oDetail, nil
}

// SelectAllOrdersDetails retrieves all order details from the database.
func SelectAllOrdersDetails() ([]models.OrderDetails, error) {
	fmt.Println("Select all Order detail function starts...")

	var oDetails []models.OrderDetails

	// Connect to the database
	err := DatabaseConnect()
	if err != nil {
		return oDetails, err
	}

	defer Database.Close()

	// Query to select all order details
	query := "SELECT * FROM Detalle_pedido"

	result, err := Database.Query(query)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return oDetails, err
	}

	defer result.Close()

	// Scan each result into the order details slice
	for result.Next() {
		var oDetail models.OrderDetails
		err2 := result.Scan(&oDetail.IdOrderDetail,
			&oDetail.IdOrder,
			&oDetail.IdProduct,
			&oDetail.QuantityOrderDetail,
			&oDetail.PriceOrderDetail)
		if err2 != nil {
			fmt.Println("Unable to Scan all the order details > " + err2.Error())
			return oDetails, err2
		}
		oDetails = append(oDetails, oDetail)
	}

	if err = result.Err(); err != nil {
		panic(err)
	}

	fmt.Println("All Order details Selected successfully!")

	return oDetails, nil
}

// InsertOrderDetail adds a new order detail to the database and returns its ID.
func InsertOrderDetail(oDetail models.OrderDetails) (int64, error) {
	fmt.Println("Insert Order details function starts...")

	err := DatabaseConnect()
	if err != nil {
		return 0, err
	}

	defer Database.Close()

	// Round the price to 2 decimal places
	roundTotal := math.Round(oDetail.PriceOrderDetail*100) / 100

	// Query to insert a new order detail
	query := "INSERT INTO Detalle_pedido (Id_pedido, Id_producto, Cantidad, Precio) VALUES (?, ?, ?, ?)"
	result, err := Database.Exec(query, oDetail.IdOrder, oDetail.IdProduct, oDetail.QuantityOrderDetail, roundTotal)
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

	fmt.Printf("Order details inserted successfully.\nIndex inserted > %d\n The row(s) affected > %d", lastInsertId, rowsAffected)

	return lastInsertId, nil
}

// UpdateOrderDetail updates an existing order detail based on its ID.
func UpdateOrderDetail(updateOrderDetail models.OrderDetails, idOrderDetail int) (models.OrderDetails, error) {
	fmt.Println("Update order is starting...")

	err := DatabaseConnect()
	if err != nil {
		return updateOrderDetail, err
	}

	defer Database.Close()

	// Query to update the order detail
	query := "UPDATE Detalle_pedido SET Id_pedido = ?, Id_producto = ?, Cantidad = ?, Precio = ? WHERE Id_detalle_pedido = ?"
	_, err = Database.Exec(query, updateOrderDetail.IdOrder, updateOrderDetail.IdProduct, updateOrderDetail.QuantityOrderDetail, updateOrderDetail.PriceOrderDetail, idOrderDetail)
	if err != nil {
		return updateOrderDetail, err
	}

	// Query to retrieve the updated order detail
	query = "SELECT * FROM Detalle_pedido WHERE Id_detalle_pedido = ?"
	result, err2 := Database.Query(query, idOrderDetail)
	if err2 != nil {
		return updateOrderDetail, err
	}

	defer result.Close()

	// Scan the updated result into the order detail struct
	result.Next()
	err = result.Scan(&updateOrderDetail.IdOrderDetail,
		&updateOrderDetail.IdOrder,
		&updateOrderDetail.IdProduct,
		&updateOrderDetail.QuantityOrderDetail,
		&updateOrderDetail.PriceOrderDetail)

	if err != nil {
		return updateOrderDetail, err
	}

	fmt.Println("The order detail has been updated successfully!")
	return updateOrderDetail, nil
}

// DeleteOrderDetail removes an order detail from the database by its ID.
func DeleteOrderDetail(idOrderDetail int) (int64, error) {
	fmt.Println("Delete order detail is starting...")

	err := DatabaseConnect()
	if err != nil {
		return 0, err
	}

	defer Database.Close()

	// Query to delete the order detail by its ID
	query := "DELETE FROM Detalle_pedido WHERE Id_detalle_pedido = ?"
	result, err := Database.Exec(query, idOrderDetail)
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

	fmt.Printf("Order detail deleted successfully.\nIndex deleted > %d\n The row(s) affected > %d", lastInsertId, rowsAffected)
	return int64(idOrderDetail), nil
}
