package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/berpeda/comercialbermudez/database"
	"github.com/berpeda/comercialbermudez/models"
)

func GetAddress(user string) (int, string) {
	result, err := database.SelectAddress(user)
	if err != nil {
		return http.StatusBadRequest, "Error trying to SELECT the Address of the user > " + user +
			"\nError > " + err.Error()
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return http.StatusBadRequest, "Error trying to resialize result to JSON format."
	}

	return http.StatusOK, string(jsonData)
}

func GetAllAddress(user string) (int, string) {

	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return http.StatusBadRequest, issue
	}

	result, err := database.SelectAllAddress()
	if err != nil {
		return http.StatusBadRequest, "Error trying to SELECT all the addresses. \nError > " + err.Error()
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return http.StatusBadRequest, "Error trying to serialize result to JSON format."
	}

	return http.StatusOK, string(jsonData)
}

func PostAddress(user, body string) (int, string) {

	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return http.StatusBadRequest, issue
	}

	var address models.Address

	err := json.Unmarshal([]byte(body), &address)
	if err != nil {
		return http.StatusBadRequest, "There is an error with received data " + err.Error()
	}

	if len(address.NameAddress) == 0 ||
		len(address.StateAddress) == 0 || len(address.CityAddress) == 0 ||
		len(address.PhoneAddress) == 0 || len(address.PostalCodeAddress) == 0 {
		return http.StatusBadRequest, "The UUUID of user, Name, State, City, Phone and Postal Code must be filled."
	}

	result, err := database.InsertAddress(address, user)
	if err != nil {
		return http.StatusBadRequest, "Error trying to INSERT the address ID > " + strconv.Itoa(address.IdAddress) + "\nError > " + err.Error()
	}

	return http.StatusOK, "{ IdAddress: " + strconv.Itoa(int(result)) + "}"
}

func PutAddress(user, body string, id int) (int, string) {
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return http.StatusBadRequest, issue
	}

	var updateAddress models.Address
	err := json.Unmarshal([]byte(body), &updateAddress)
	if err != nil {
		return http.StatusBadRequest, "There is an issue with received data " + err.Error()
	}

	if len(updateAddress.NameAddress) == 0 ||
		len(updateAddress.StateAddress) == 0 || len(updateAddress.CityAddress) == 0 ||
		len(updateAddress.PhoneAddress) == 0 || len(updateAddress.PostalCodeAddress) == 0 {

		return http.StatusBadRequest, "Any of the Address's attributes needs to be filled."
	}

	result, err := database.UpdateAddress(updateAddress, id)
	if err != nil {
		return http.StatusBadRequest, "Error trying to UPDATE the address > " + strconv.Itoa(id)
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return http.StatusBadRequest, "Error trying to serialize result to JSON format."
	}

	return http.StatusOK, string(jsonData)
}

func DeleteAddress(user string, id int) (int, string) {
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return http.StatusBadRequest, issue
	}

	result, err := database.DeleteAddress(id)
	if err != nil {
		return http.StatusBadRequest, "Error trying to DELETE the address with ID > " + strconv.Itoa(int(result))
	}

	return http.StatusOK, "{ IdAddress: " + strconv.Itoa(int(result)) + "}"
}
