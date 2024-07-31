package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/berpeda/comercialbermudez/database"
	"github.com/berpeda/comercialbermudez/models"
)

// GetAddress retrieves the address of a specific user.
func GetAddress(user string) (int, string) {
	// Fetch the address from the database
	result, err := database.SelectAddress(user)
	if err != nil {
		// Return a bad request status and error message if the database query fails
		return http.StatusBadRequest, "Error trying to SELECT the Address of the user > " + user +
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

// GetAllAddress retrieves all addresses, but only if the user is an admin.
func GetAllAddress(user string) (int, string) {
	// Check if the user is an admin
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		// Return a bad request status and issue message if the user is not an admin
		return http.StatusBadRequest, issue
	}

	// Fetch all addresses from the database
	result, err := database.SelectAllAddress()
	if err != nil {
		// Return a bad request status and error message if the database query fails
		return http.StatusBadRequest, "Error trying to SELECT all the addresses. \nError > " + err.Error()
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

// PostAddress inserts a new address, but only if the user is an admin.
func PostAddress(user, body string) (int, string) {
	// Check if the user is an admin
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		// Return a bad request status and issue message if the user is not an admin
		return http.StatusBadRequest, issue
	}

	var address models.Address

	// Deserialize the JSON body into an Address struct
	err := json.Unmarshal([]byte(body), &address)
	if err != nil {
		// Return a bad request status and error message if deserialization fails
		return http.StatusBadRequest, "There is an error with received data " + err.Error()
	}

	// Check if all required fields are filled
	if len(address.NameAddress) == 0 ||
		len(address.StateAddress) == 0 || len(address.CityAddress) == 0 ||
		len(address.PhoneAddress) == 0 || len(address.PostalCodeAddress) == 0 {
		// Return a bad request status if any required field is missing
		return http.StatusBadRequest, "The UUID of user, Name, State, City, Phone and Postal Code must be filled."
	}

	// Insert the address into the database
	result, err := database.InsertAddress(address, user)
	if err != nil {
		// Return a bad request status and error message if the database insert fails
		return http.StatusBadRequest, "Error trying to INSERT the address ID > " + strconv.Itoa(address.IdAddress) + "\nError > " + err.Error()
	}

	// Return a success status and the ID of the inserted address
	return http.StatusOK, "{ IdAddress: " + strconv.Itoa(int(result)) + "}"
}

// PutAddress updates an existing address, but only if the user is an admin.
func PutAddress(user, body string, id int) (int, string) {
	// Check if the user is an admin
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		// Return a bad request status and issue message if the user is not an admin
		return http.StatusBadRequest, issue
	}

	var updateAddress models.Address

	// Deserialize the JSON body into an Address struct
	err := json.Unmarshal([]byte(body), &updateAddress)
	if err != nil {
		// Return a bad request status and error message if deserialization fails
		return http.StatusBadRequest, "There is an issue with received data " + err.Error()
	}

	// Check if all required fields are filled
	if len(updateAddress.NameAddress) == 0 ||
		len(updateAddress.StateAddress) == 0 || len(updateAddress.CityAddress) == 0 ||
		len(updateAddress.PhoneAddress) == 0 || len(updateAddress.PostalCodeAddress) == 0 {
		// Return a bad request status if any required field is missing
		return http.StatusBadRequest, "Any of the Address's attributes needs to be filled."
	}

	// Update the address in the database
	result, err := database.UpdateAddress(updateAddress, id)
	if err != nil {
		// Return a bad request status and error message if the database update fails
		return http.StatusBadRequest, "Error trying to UPDATE the address > " + strconv.Itoa(id)
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

// DeleteAddress deletes an address, but only if the user is an admin.
func DeleteAddress(user string, id int) (int, string) {
	// Check if the user is an admin
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		// Return a bad request status and issue message if the user is not an admin
		return http.StatusBadRequest, issue
	}

	// Delete the address from the database
	result, err := database.DeleteAddress(id)
	if err != nil {
		// Return a bad request status and error message if the database delete fails
		return http.StatusBadRequest, "Error trying to DELETE the address with ID > " + strconv.Itoa(int(result))
	}

	// Return a success status and the ID of the deleted address
	return http.StatusOK, "{ IdAddress: " + strconv.Itoa(int(result)) + "}"
}
