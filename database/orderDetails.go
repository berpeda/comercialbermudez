package database

import (
	"fmt"
	"math"

	"github.com/berpeda/comercialbermudez/models"
)

func SelectOrderDetail(idOrderDetail int) (models.OrderDetails, error) {
	fmt.Println("Select a single Order detail function starts...")

	var oDetail models.OrderDetails

	err := DatabaseConnect()
	if err != nil {
		return oDetail, err
	}

	defer Database.Close()

	query := "SELECT * FROM Detalle_pedido WHERE Id_detalle_pedido = ?"

	result, err := Database.Query(query, idOrderDetail)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return oDetail, err
	}

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

func SelectAllOrdersDetails() ([]models.OrderDetails, error) {
	fmt.Println("Select all Order detail function starts...")

	var oDetails []models.OrderDetails

	err := DatabaseConnect()
	if err != nil {
		return oDetails, err
	}

	defer Database.Close()

	query := "SELECT * FROM Detalle_pedido"

	result, err := Database.Query(query)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return oDetails, err
	}

	for result.Next() {
		var oDetail models.OrderDetails
		err2 := result.Scan(&oDetail.IdOrderDetail,
			&oDetail.IdOrder,
			&oDetail.IdProduct,
			&oDetail.QuantityOrderDetail,
			&oDetail.PriceOrderDetail)
		if err2 != nil {
			fmt.Println("Unable to Scan all the order details > " + err.Error())
			panic(err)
		}
		oDetails = append(oDetails, oDetail)
	}

	if err = result.Err(); err != nil {
		panic(err)
	}

	fmt.Println("All Order details Selected successfully!")

	return oDetails, nil
}

func InsertOrderDetail(oDetail models.OrderDetails) (int64, error) {
	fmt.Println("Insert Order details function starts...")

	err := DatabaseConnect()
	if err != nil {
		return 0, err
	}

	defer Database.Close()

	// This round the price to 2 decimals in case the order price have more than 2
	roundTotal := math.Round(oDetail.PriceOrderDetail*100) / 100

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

func UpdateOrderDetail(updateOrderDetail models.OrderDetails, idOrderDetail int) (models.OrderDetails, error) {
	fmt.Println("Update order is starting...")

	err := DatabaseConnect()
	if err != nil {
		return updateOrderDetail, err
	}

	defer Database.Close()

	query := "UPDATE Detalle_pedido SET Id_pedido = ?, Id_producto = ?, Cantidad = ?, Precio = ? WHERE Id_detalle_pedido = ?"
	_, err = Database.Exec(query, updateOrderDetail.IdOrder, updateOrderDetail.IdProduct, updateOrderDetail.QuantityOrderDetail, updateOrderDetail.PriceOrderDetail, idOrderDetail)
	if err != nil {
		return updateOrderDetail, err
	}

	query = "SELECT * FROM Detalle_pedido WHERE Id_detalle_pedido = ?"
	result, err2 := Database.Query(query, idOrderDetail)
	if err2 != nil {
		return updateOrderDetail, err
	}

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

func DeleteOrderDetail(idOrderDetail int) (int64, error) {
	fmt.Println("Delete order detail is starting...")

	err := DatabaseConnect()
	if err != nil {
		return 0, err
	}

	defer Database.Close()

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
