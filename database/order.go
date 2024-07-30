package database

import (
	"fmt"

	"github.com/berpeda/comercialbermudez/models"
	"github.com/berpeda/comercialbermudez/tools"
)

func SelectOrders(user string, idOrder, page int) ([]models.Order, error) {
	fmt.Println("Select Order function starts...")

	var orders []models.Order

	query := "SELECT * FROM Pedidos"
	params := []interface{}{}

	if idOrder > 0 {
		query += " WHERE Id_pedido = ?"
		params = append(params, idOrder)
	} else {
		offset := 0
		if page == 0 {
			page = 1
		}
		if page > 1 {
			offset = (10 * (page - 1))
		}

		where := ""
		whereUser := " UUID_usuario = ?"

		if where != "" {
			where += " AND " + whereUser
			params = append(params, user)
		} else {
			where = " WHERE " + whereUser
			params = append(params, user)
		}

		limit := " LIMIT 10"
		if offset > 0 {
			limit += " OFFSET ?"
			params = append(params, offset)
		}

		query += where + limit
	}

	fmt.Println(query)
	fmt.Println(params)

	err := DatabaseConnect()

	if err != nil {
		return orders, err
	}
	defer Database.Close()

	result, err := Database.Query(query, params...)
	if err != nil {
		fmt.Println("Error trying to do the query.")
		return orders, err
	}

	defer result.Close()

	for result.Next() {
		var o models.Order
		err2 := result.Scan(&o.IdOrder,
			&o.UUIDUser,
			&o.IdAddress,
			&o.Total,
			&o.CreatedAt)

		if err2 != nil {
			fmt.Println("Error trying to scan the order.")
			return orders, err2
		}

		query = "SELECT * FROM Detalle_pedido WHERE Id_pedido = ?"
		r2, err := Database.Query(query, o.IdOrder)
		if err != nil {
			return orders, err
		}

		defer r2.Close()

		for r2.Next() {
			var od models.OrderDetails
			err3 := r2.Scan(
				&od.IdOrderDetail,
				&od.IdOrder,
				&od.IdProduct,
				&od.QuantityOrderDetail,
				&od.PriceOrderDetail)

			if err3 != nil {
				fmt.Println("Error trying to scan the order detail.")
				return orders, err3
			}
			o.OrderItems = append(o.OrderItems, od)
		}

		orders = append(orders, o)
	}

	fmt.Println("Order selected succesfully!")

	return orders, nil
}

func InsertOrder(order models.Order, idUser string) (int64, error) {
	fmt.Println("Insert Order function starts...")

	err := DatabaseConnect()
	if err != nil {
		return 0, err
	}

	defer Database.Close()

	query := "INSERT INTO Pedidos (UUID_usuario, Id_direccion, Total, Creado) VALUES (?, ?, ?, ?)"
	result, err := Database.Exec(query, idUser, order.IdAddress, order.Total, tools.DateMySQL())
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

	for _, orderD := range order.OrderItems {
		query = "INSERT INTO Detalle_pedido (Id_pedido, Id_producto, Cantidad, Precio) VALUES (?, ?, ?, ?)"
		_, err = Database.Exec(query, lastInsertId, orderD.IdProduct, orderD.QuantityOrderDetail, orderD.PriceOrderDetail)
		if err != nil {
			fmt.Println("Error trying to INSERT the Order Detail > ", err.Error())
			return 0, err
		}
	}

	fmt.Printf("Order inserted successfully.\nIndex inserted > %d\n The row(s) affected > %d", lastInsertId, rowsAffected)

	return lastInsertId, nil
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
