package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/berpeda/comercialbermudez/database"
	"github.com/berpeda/comercialbermudez/models"
)

// GetOrderDetail retrieves the details of a specific order by its ID.
func GetOrderDetail(idOrder int) (int, string) {
	// Fetch the order details from the database
	result, err := database.SelectOrderDetail(idOrder)
	if err != nil {
		// Return a bad request status and error message if the database query fails
		return http.StatusBadRequest, "Error trying to SELECT the Order with ID > " + strconv.Itoa(idOrder) +
			"\nError > " + err.Error()
	}

	// Serialize the result to JSON format
	jsonData, err := json.Marshal(result)
	if err != nil {
		// Return a bad request status if JSON serialization fails
		return http.StatusBadRequest, "Error trying to serialize result to JSON format."
	}

	// Return a success status and the JSON result
	return http.StatusOK, string(jsonData)
}

// GetAllOrderDetails retrieves the details of all orders.
func GetAllOrderDetails() (int, string) {
	// Fetch all order details from the database
	result, err := database.SelectAllOrdersDetails()
	if err != nil {
		// Return a bad request status and error message if the database query fails
		return http.StatusBadRequest, "Error trying to SELECT all the Orders details. \nError > " + err.Error()
	}

	// Serialize the result to JSON format
	jsonData, err := json.Marshal(result)
	if err != nil {
		// Return a bad request status if JSON serialization fails
		return http.StatusBadRequest, "Error trying to serialize result to JSON format."
	}

	// Return a success status and the JSON result
	return http.StatusOK, string(jsonData)
}

// PostOrderDetail inserts a new order detail, but only if the user is an admin.
func PostOrderDetail(user, body string) (int, string) {
	// Check if the user is an admin
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		// Return a bad request status and issue message if the user is not an admin
		return http.StatusBadRequest, issue
	}

	var oDetail models.OrderDetails

	// Deserialize the JSON body into an OrderDetails struct
	err := json.Unmarshal([]byte(body), &oDetail)
	if err != nil {
		// Return a bad request status and error message if deserialization fails
		return http.StatusBadRequest, "There is an error with received data " + err.Error()
	}

	// Check if all required fields are filled
	if oDetail.IdOrder == 0 || oDetail.IdProduct == 0 ||
		oDetail.QuantityOrderDetail <= 0 || oDetail.PriceOrderDetail <= 0.0 {
		// Return a bad request status if any required field is missing or invalid
		return http.StatusBadRequest, "Any of the fields are empty or less than 0."
	}

	// Insert the order detail into the database
	result, err := database.InsertOrderDetail(oDetail)
	if err != nil {
		// Return a bad request status and error message if the database insert fails
		return http.StatusBadRequest, "Error trying to INSERT the orderDetail ID > " + strconv.Itoa(oDetail.IdOrder) + "\nError > " + err.Error()
	}

	// Return a success status and the ID of the inserted order detail
	return http.StatusOK, "{ IdOrderDetail: " + strconv.Itoa(int(result)) + "}"
}

// PutOrderDetail updates an existing order detail, but only if the user is an admin.
func PutOrderDetail(user, body string, idOrderDetail int) (int, string) {
	// Check if the user is an admin
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		// Return a bad request status and issue message if the user is not an admin
		return http.StatusBadRequest, issue
	}

	var updateOrderDetail models.OrderDetails

	// Deserialize the JSON body into an OrderDetails struct
	err := json.Unmarshal([]byte(body), &updateOrderDetail)
	if err != nil {
		// Return a bad request status and error message if deserialization fails
		return http.StatusBadRequest, "There is an issue with received data " + err.Error()
	}

	// Check if all required fields are filled
	if updateOrderDetail.IdOrder == 0 || updateOrderDetail.IdProduct == 0 ||
		updateOrderDetail.QuantityOrderDetail <= 0 || updateOrderDetail.PriceOrderDetail <= 0.0 {
		// Return a bad request status if any required field is missing or invalid
		return http.StatusBadRequest, "Any of the order's attributes need to be filled (The creation is not necessary to be filled)"
	}

	// Update the order detail in the database
	result, err := database.UpdateOrderDetail(updateOrderDetail, idOrderDetail)
	if err != nil {
		// Return a bad request status and error message if the database update fails
		return http.StatusBadRequest, "Error trying to UPDATE the order detail > " + strconv.Itoa(idOrderDetail)
	}

	// Serialize the result to JSON format
	jsonData, err := json.Marshal(result)
	if err != nil {
		// Return a bad request status if JSON serialization fails
		return http.StatusBadRequest, "Error trying to serialize result to JSON format."
	}

	// Return a success status and the JSON result
	return http.StatusOK, string(jsonData)
}

// DeleteOrderDetail deletes an order detail, but only if the user is an admin.
func DeleteOrderDetail(user string, idOrderDetail int) (int, string) {
	// Check if the user is an admin
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		// Return a bad request status and issue message if the user is not an admin
		return http.StatusBadRequest, issue
	}

	// Delete the order detail from the database
	result, err := database.DeleteOrderDetail(idOrderDetail)
	if err != nil {
		// Return a bad request status and error message if the database delete fails
		return http.StatusBadRequest, "Error trying to DELETE the order detail with ID > " + strconv.Itoa(int(result))
	}

	// Return a success status and the ID of the deleted order detail
	return http.StatusOK, "{ IdOrderDetail: " + strconv.Itoa(int(result)) + "}"
}
