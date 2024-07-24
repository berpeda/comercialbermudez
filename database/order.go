package database

import (
	"fmt"
	"math"

	"github.com/berpeda/comercialbermudez/models"
	"github.com/berpeda/comercialbermudez/tools"
)

func SelectOrder(idOrder int) (models.Order, error) {
	fmt.Println("Select a single Order function starts...")

	var order models.Order

	err := DatabaseConnect()
	if err != nil {
		return order, err
	}

	defer Database.Close()

	query := "SELECT * FROM Pedidos WHERE Id_pedido = ?"

	result, err := Database.Query(query, idOrder)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return order, err
	}

	result.Next()
	err2 := result.Scan(&order.IdOrder,
		&order.UUIDUser,
		&order.IdAddress,
		&order.Total,
		&order.CreatedAt)

	if err2 != nil {
		fmt.Println("result.Scan is having issues...")
		return order, err2
	}

	fmt.Printf("Order selected successfully.")

	return order, nil
}

func SelectAllOrders() ([]models.Order, error) {
	fmt.Println("Select all Orders function starts...")

	var orders []models.Order

	err := DatabaseConnect()
	if err != nil {
		return orders, err
	}

	defer Database.Close()

	query := "SELECT * FROM Pedidos"
	result, err := Database.Query(query)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return orders, err
	}

	for result.Next() {
		var order models.Order
		err = result.Scan(&order.IdOrder,
			&order.UUIDUser,
			&order.IdAddress,
			&order.Total,
			&order.CreatedAt)

		if err != nil {
			fmt.Println("Unable to Scan all the orders > " + err.Error())
			panic(err)
		}
		orders = append(orders, order)
	}

	if err = result.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Orders Selected successfully!")

	return orders, nil
}

func InsertOrder(order models.Order) (int64, error) {
	fmt.Println("Insert Order function starts...")

	err := DatabaseConnect()
	if err != nil {
		return 0, err
	}

	defer Database.Close()

	// This round the price to 2 decimals in case the order price have more than 2
	roundTotal := math.Round(order.Total*100) / 100

	query := "INSERT INTO Pedidos (UUID_usuario, Id_direccion, Total, Creado) VALUES (?, ?, ?, ?)"
	result, err := Database.Exec(query, order.UUIDUser, order.IdAddress, roundTotal, tools.DateMySQL())
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

	fmt.Printf("Order inserted successfully.\nIndex inserted > %d\n The row(s) affected > %d", lastInsertId, rowsAffected)

	return lastInsertId, nil
}

func UpdateOrder(order models.Order, idOrder int) (models.Order, error) {
	fmt.Println("Update order is starting...")

	err := DatabaseConnect()
	if err != nil {
		return order, err
	}

	defer Database.Close()

	query := "UPDATE Pedidos SET UUID_usuario = ?, Id_direccion = ?, Total = ?, Creado = ? WHERE Id_pedido = ?"
	_, err = Database.Exec(query, order.UUIDUser, order.IdAddress, order.Total, order.CreatedAt, idOrder)
	if err != nil {
		return order, err
	}

	query = "SELECT * FROM Pedidos WHERE Id_pedido = ?"
	result, err2 := Database.Query(query, idOrder)
	if err2 != nil {
		return order, err
	}

	result.Next()
	err = result.Scan(&order.IdOrder, &order.UUIDUser, &order.IdAddress, &order.Total, &order.CreatedAt)
	if err != nil {
		return order, err
	}

	fmt.Println("The order has been updated successfully!")
	return order, nil
}

func DeleteOrder(idOrder int) (int64, error) {
	fmt.Println("Delete order is starting...")

	err := DatabaseConnect()
	if err != nil {
		return 0, err
	}

	defer Database.Close()

	query := "DELETE FROM Pedidos WHERE Id_pedido = ?"
	result, err := Database.Exec(query, idOrder)
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

	fmt.Printf("Order deleted successfully.\nIndex deleted > %d\n The row(s) affected > %d", lastInsertId, rowsAffected)

	return int64(idOrder), nil

}
