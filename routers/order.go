package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/berpeda/comercialbermudez/database"
	"github.com/berpeda/comercialbermudez/models"
)

func GetOrder(idOrder int) (int, string) {
	result, err := database.SelectOrder(idOrder)
	if err != nil {
		return http.StatusBadRequest, "Error trying to SELECT the Order with ID > " + strconv.Itoa(idOrder) +
			"\nError > " + err.Error()
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return http.StatusBadRequest, "Error trying to resialize result to JSON format."
	}

	return http.StatusOK, string(jsonData)
}

func GetAllOrders() (int, string) {

	result, err := database.SelectAllOrders()
	if err != nil {
		return http.StatusBadRequest, "Error trying to SELECT all the Orders. \nError > " + err.Error()
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return http.StatusBadRequest, "Error trying to serialize result to JSON format."
	}

	return http.StatusOK, string(jsonData)
}

func PostOrder(user, body string) (int, string) {

	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return http.StatusBadRequest, issue
	}

	var order models.Order

	err := json.Unmarshal([]byte(body), &order)
	if err != nil {
		return http.StatusBadRequest, "There is an error with received data " + err.Error()
	}

	if len(order.UUIDUser) == 0 || order.IdAddress == 0 || order.Total <= 0.0 {
		return http.StatusBadRequest, "The UUID User or the address ID is not filled"
	}

	result, err := database.InsertOrder(order)

	if err != nil {
		return http.StatusBadRequest, "Error trying to INSERT the order ID > " + strconv.Itoa(order.IdOrder) + "\nError > " + err.Error()
	}

	return http.StatusOK, "{ IdOrder: " + strconv.Itoa(int(result)) + "}"
}

func PutOrder(user, body string, idOrder int) (int, string) {
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return http.StatusBadRequest, issue
	}

	var updateOrder models.Order
	err := json.Unmarshal([]byte(body), &updateOrder)
	if err != nil {
		return http.StatusBadRequest, "There is an issue with received data " + err.Error()
	}

	if len(updateOrder.UUIDUser) == 0 || updateOrder.IdAddress == 0 || updateOrder.Total == 0 {

		return http.StatusBadRequest, "Any of the order's attributes needs to be filled (The creation is not necessary to be filled)"
	}

	result, err := database.UpdateOrder(updateOrder, idOrder)
	if err != nil {
		return http.StatusBadRequest, "Error trying to UPDATE the order > " + strconv.Itoa(idOrder)
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return http.StatusBadRequest, "Error trying to serialize result to JSON format."
	}

	return http.StatusOK, string(jsonData)
}

func DeleteOrder(user string, idOrder int) (int, string) {
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return http.StatusBadRequest, issue
	}

	result, err := database.DeleteOrder(idOrder)
	if err != nil {
		return http.StatusBadRequest, "Error trying to DELETE the order with ID > " + strconv.Itoa(int(result))
	}

	return http.StatusOK, "{ IdOrder: " + strconv.Itoa(int(result)) + "}"
}
