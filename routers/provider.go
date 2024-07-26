package routers

import (
	"encoding/json"
	"strconv"

	"github.com/berpeda/comercialbermudez/database"
	"github.com/berpeda/comercialbermudez/models"
)

func GetProvider(idProvider int) (int, string) {
	result, err := database.SelectProvider(idProvider)
	if err != nil {
		return 400, "Error trying to SELECT the Provider with ID > " + strconv.Itoa(idProvider) +
			"\nError > " + err.Error()
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return 400, "Error trying to resialize result to JSON format."
	}

	return 200, string(jsonData)
}

func GetAllProviders() (int, string) {
	result, err := database.SelectAllProviders()
	if err != nil {
		return 400, "Error trying to SELECT all the Providers. \nError > " + err.Error()
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return 400, "Error trying to serialize result to JSON format."
	}

	return 200, string(jsonData)
}

func PostProvider(user, body string) (int, string) {
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return 400, issue
	}

	var provider models.Provider

	err := json.Unmarshal([]byte(body), &provider)
	if err != nil {
		return 400, "There is an error with received data " + err.Error()
	}

	if len(provider.NameProvider) == 0 || len(provider.EmailProvider) == 0 || len(provider.PhoneNumberProvider) == 0 {
		return 400, "The Name, Email or Phone of the provider is not filled"
	}

	result, err := database.InsertProvider(provider)

	if err != nil {
		return 400, "Error trying to INSERT the provider ID > " + strconv.Itoa(provider.IdProvider) + "\nError > " + err.Error()
	}

	return 200, "{ IdProvider: " + strconv.Itoa(int(result)) + "}"
}

func PutProvider(user, body string, idProvider int) (int, string) {
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return 400, issue
	}

	var provider models.Provider
	err := json.Unmarshal([]byte(body), &provider)
	if err != nil {
		return 400, "There is an issue with received data " + err.Error()
	}

	if len(provider.NameProvider) == 0 || len(provider.EmailProvider) == 0 || len(provider.PhoneNumberProvider) == 0 {

		return 400, "Any of the provider's attributes needs to be filled."
	}

	result, err := database.UpdateProvider(provider, idProvider)
	if err != nil {
		return 400, "Error trying to UPDATE the provider > " + strconv.Itoa(idProvider)
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return 400, "Error trying to serialize result to JSON format."
	}

	return 200, string(jsonData)
}

func DeleteProvider(user string, idProvider int) (int, string) {
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return 400, issue
	}

	result, err := database.DeleteProvider(idProvider)
	if err != nil {
		return 400, "Error trying to DELETE the provider with ID > " + strconv.Itoa(int(result))
	}

	return 200, "{ IdProvider: " + strconv.Itoa(int(result)) + "}"
}
