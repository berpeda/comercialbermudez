package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/berpeda/comercialbermudez/database"
	"github.com/berpeda/comercialbermudez/models"
)

func PostOrder(user, body string) (int, string) {

	var order models.Order

	err := json.Unmarshal([]byte(body), &order)
	if err != nil {
		return http.StatusBadRequest, "There is an error with received data " + err.Error()
	}

	if order.Total == 0 || len(order.OrderItems) == 0 {
		return http.StatusBadRequest, "Error, the order total bill or the slice of items shouldn't be 0"
	}

	result, err := database.InsertOrder(order, user)

	if err != nil {
		return http.StatusBadRequest, "Error trying to INSERT the order ID > " + strconv.Itoa(order.IdOrder) + "\nError > " + err.Error()
	}

	return http.StatusOK, "{ IdOrder: " + strconv.Itoa(int(result)) + "}"
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

func GetOrders(user string, request events.APIGatewayV2HTTPRequest) (int, string) {
	params := request.QueryStringParameters
	idOrder, page := 0, 0

	if v, _ := strconv.Atoi(params["page"]); v > 0 {
		page = v
	}

	if v2, _ := strconv.Atoi(params["idOrder"]); v2 > 0 {
		idOrder = v2
	}

	result, err := database.SelectOrders(user, idOrder, page)
	if err != nil {
		return http.StatusBadRequest, "Error trying to Select the order > " + strconv.Itoa(idOrder)
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return http.StatusBadRequest, "Error trying to serialize result to JSON"
	}

	return http.StatusOK, string(jsonData)
}
