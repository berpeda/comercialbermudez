package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/berpeda/comercialbermudez/database"
	"github.com/berpeda/comercialbermudez/models"
)

// PostOrder handles the insertion of a new order.
func PostOrder(user, body string) (int, string) {
	var order models.Order

	// Deserialize the JSON body into an Order struct
	err := json.Unmarshal([]byte(body), &order)
	if err != nil {
		// Return a bad request status and error message if deserialization fails
		return http.StatusBadRequest, "There is an error with received data " + err.Error()
	}

	// Check if the total amount or order items are zero
	if order.Total == 0 || len(order.OrderItems) == 0 {
		// Return a bad request status if any required field is missing
		return http.StatusBadRequest, "Error, the order total bill or the slice of items shouldn't be 0"
	}

	// Insert the order into the database
	result, err := database.InsertOrder(order, user)
	if err != nil {
		// Return a bad request status and error message if the database insert fails
		return http.StatusBadRequest, "Error trying to INSERT the order ID > " + strconv.Itoa(order.IdOrder) + "\nError > " + err.Error()
	}

	// Return a success status and the ID of the inserted order
	return http.StatusOK, "{ IdOrder: " + strconv.Itoa(int(result)) + "}"
}

// DeleteOrder handles the deletion of an order, but only if the user is an admin.
func DeleteOrder(user string, idOrder int) (int, string) {
	// Check if the user is an admin
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		// Return a bad request status and issue message if the user is not an admin
		return http.StatusBadRequest, issue
	}

	// Delete the order from the database
	result, err := database.DeleteOrder(idOrder)
	if err != nil {
		// Return a bad request status and error message if the database delete fails
		return http.StatusBadRequest, "Error trying to DELETE the order with ID > " + strconv.Itoa(int(result))
	}

	// Return a success status and the ID of the deleted order
	return http.StatusOK, "{ IdOrder: " + strconv.Itoa(int(result)) + "}"
}

// GetOrders retrieves orders based on user and query parameters.
func GetOrders(user string, request events.APIGatewayV2HTTPRequest) (int, string) {
	params := request.QueryStringParameters
	idOrder, page := 0, 0

	// Parse the "page" query parameter
	if v, _ := strconv.Atoi(params["page"]); v > 0 {
		page = v
	}

	// Parse the "idOrder" query parameter
	if v2, _ := strconv.Atoi(params["idOrder"]); v2 > 0 {
		idOrder = v2
	}

	// Fetch orders from the database
	result, err := database.SelectOrders(user, idOrder, page)
	if err != nil {
		// Return a bad request status and error message if the database query fails
		return http.StatusBadRequest, "Error trying to Select the order > " + strconv.Itoa(idOrder)
	}

	// Serialize the result to JSON format
	jsonData, err := json.Marshal(result)
	if err != nil {
		// Return a bad request status if JSON serialization fails
		return http.StatusBadRequest, "Error trying to serialize result to JSON"
	}

	// Return a success status and the JSON result
	return http.StatusOK, string(jsonData)
}
