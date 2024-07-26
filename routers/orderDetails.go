package routers

import (
	"encoding/json"
	"strconv"

	"github.com/berpeda/comercialbermudez/database"
	"github.com/berpeda/comercialbermudez/models"
)

func GetOrderDetail(idOrder int) (int, string) {
	result, err := database.SelectOrderDetail(idOrder)
	if err != nil {
		return 400, "Error trying to SELECT the Order with ID > " + strconv.Itoa(idOrder) +
			"\nError > " + err.Error()
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return 400, "Error trying to resialize result to JSON format."
	}

	return 200, string(jsonData)
}

func GetAllOrderDetails() (int, string) {
	result, err := database.SelectAllOrdersDetails()
	if err != nil {
		return 400, "Error trying to SELECT all the Orders details. \nError > " + err.Error()
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return 400, "Error trying to serialize result to JSON format."
	}

	return 200, string(jsonData)
}

func PostOrderDetail(user, body string) (int, string) {

	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return 400, issue
	}

	var oDetail models.OrderDetails

	err := json.Unmarshal([]byte(body), &oDetail)
	if err != nil {
		return 400, "There is an error with received data " + err.Error()
	}

	if oDetail.IdOrder == 0 || oDetail.IdProduct == 0 ||
		oDetail.QuantityOrderDetail <= 0 || oDetail.PriceOrderDetail <= 0.0 {
		return 400, "Any of the fields are empty or less than 0."
	}

	result, err := database.InsertOrderDetail(oDetail)

	if err != nil {
		return 400, "Error trying to INSERT the orderDetail ID > " + strconv.Itoa(oDetail.IdOrder) + "\nError > " + err.Error()
	}

	return 200, "{ IdOrderDetail: " + strconv.Itoa(int(result)) + "}"
}

func PutOrderDetail(user, body string, idOrderDetail int) (int, string) {
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return 400, issue
	}

	var updateOrderDetail models.OrderDetails

	err := json.Unmarshal([]byte(body), &updateOrderDetail)
	if err != nil {
		return 400, "There is an issue with received data " + err.Error()
	}

	if updateOrderDetail.IdOrder == 0 || updateOrderDetail.IdProduct == 0 ||
		updateOrderDetail.QuantityOrderDetail <= 0 || updateOrderDetail.PriceOrderDetail <= 0.0 {
		return 400, "Any of the order's attributes needs to be filled (The creation is not necessary to be filled)"
	}

	result, err := database.UpdateOrderDetail(updateOrderDetail, idOrderDetail)
	if err != nil {
		return 400, "Error trying to UPDATE the order detail > " + strconv.Itoa(idOrderDetail)
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return 400, "Error trying to serialize result to JSON format."
	}

	return 200, string(jsonData)
}

func DeleteOrderDetail(user string, idOrderDetail int) (int, string) {
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return 400, issue
	}

	result, err := database.DeleteOrderDetail(idOrderDetail)
	if err != nil {
		return 400, "Error trying to DELETE the order detail with ID > " + strconv.Itoa(int(result))
	}

	return 200, "{ IdOrderDetail: " + strconv.Itoa(int(result)) + "}"
}
