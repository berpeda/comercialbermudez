package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/berpeda/comercialbermudez/database"
	"github.com/berpeda/comercialbermudez/models"
)

// GetProvider retrieves a provider by its ID from the database
func GetProvider(idProvider int) (int, string) {
	// Fetch the provider from the database
	result, err := database.SelectProvider(idProvider)
	if err != nil {
		return http.StatusBadRequest, "Error trying to SELECT the Provider with ID > " + strconv.Itoa(idProvider) +
			"\nError > " + err.Error()
	}

	// Serialize the result to JSON format
	jsonData, err := json.Marshal(result)
	if err != nil {
		return http.StatusBadRequest, "Error trying to serialize result to JSON format."
	}

	// Return a success status and the JSON result
	return http.StatusOK, string(jsonData)
}

// GetAllProviders retrieves all providers from the database
func GetAllProviders() (int, string) {
	// Fetch all providers from the database
	result, err := database.SelectAllProviders()
	if err != nil {
		return http.StatusBadRequest, "Error trying to SELECT all the Providers. \nError > " + err.Error()
	}

	// Serialize the result to JSON format
	jsonData, err := json.Marshal(result)
	if err != nil {
		return http.StatusBadRequest, "Error trying to serialize result to JSON format."
	}

	// Return a success status and the JSON result
	return http.StatusOK, string(jsonData)
}

// PostProvider adds a new provider to the database if the user is an admin
func PostProvider(user, body string) (int, string) {
	// Check if the user is an admin
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return http.StatusBadRequest, issue
	}

	var provider models.Provider

	// Deserialize the JSON body into a Provider struct
	err := json.Unmarshal([]byte(body), &provider)
	if err != nil {
		return http.StatusBadRequest, "There is an error with received data " + err.Error()
	}

	// Check if all required fields are filled
	if len(provider.NameProvider) == 0 || len(provider.EmailProvider) == 0 || len(provider.PhoneNumberProvider) == 0 {
		return http.StatusBadRequest, "The Name, Email or Phone of the provider is not filled"
	}

	// Insert the provider into the database
	result, err := database.InsertProvider(provider)
	if err != nil {
		return http.StatusBadRequest, "Error trying to INSERT the provider ID > " + strconv.Itoa(provider.IdProvider) + "\nError > " + err.Error()
	}

	// Return a success status and the ID of the inserted provider
	return http.StatusOK, "{ IdProvider: " + strconv.Itoa(int(result)) + "}"
}

// PutProvider updates an existing provider if the user is an admin
func PutProvider(user, body string, idProvider int) (int, string) {
	// Check if the user is an admin
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return http.StatusBadRequest, issue
	}

	var provider models.Provider

	// Deserialize the JSON body into a Provider struct
	err := json.Unmarshal([]byte(body), &provider)
	if err != nil {
		return http.StatusBadRequest, "There is an issue with received data " + err.Error()
	}

	// Check if all required fields are filled
	if len(provider.NameProvider) == 0 || len(provider.EmailProvider) == 0 || len(provider.PhoneNumberProvider) == 0 {
		return http.StatusBadRequest, "Any of the provider's attributes need to be filled."
	}

	// Update the provider in the database
	result, err := database.UpdateProvider(provider, idProvider)
	if err != nil {
		return http.StatusBadRequest, "Error trying to UPDATE the provider > " + strconv.Itoa(idProvider)
	}

	// Serialize the result to JSON format
	jsonData, err := json.Marshal(result)
	if err != nil {
		return http.StatusBadRequest, "Error trying to serialize result to JSON format."
	}

	// Return a success status and the JSON result
	return http.StatusOK, string(jsonData)
}

// DeleteProvider deletes a provider if the user is an admin
func DeleteProvider(user string, idProvider int) (int, string) {
	// Check if the user is an admin
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return http.StatusBadRequest, issue
	}

	// Delete the provider from the database
	result, err := database.DeleteProvider(idProvider)
	if err != nil {
		return http.StatusBadRequest, "Error trying to DELETE the provider with ID > " + strconv.Itoa(int(result))
	}

	// Return a success status and the ID of the deleted provider
	return http.StatusOK, "{ IdProvider: " + strconv.Itoa(int(result)) + "}"
}
