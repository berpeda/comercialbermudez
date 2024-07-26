package routers

import (
	"encoding/json"
	"strconv"

	"github.com/berpeda/comercialbermudez/database"
	"github.com/berpeda/comercialbermudez/models"
)

func GetAddress(id int) (int, string) {
	result, err := database.SelectAddress(id)
	if err != nil {
		return 400, "Error trying to SELECT the Address with ID > " + strconv.Itoa(id) +
			"\nError > " + err.Error()
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return 400, "Error trying to resialize result to JSON format."
	}

	return 200, string(jsonData)
}

func GetAllAddress() (int, string) {

	result, err := database.SelectAllAddress()
	if err != nil {
		return 400, "Error trying to SELECT all the addresses. \nError > " + err.Error()
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return 400, "Error trying to serialize result to JSON format."
	}

	return 200, string(jsonData)
}

func PostAddress(user, body string) (int, string) {

	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return 400, issue
	}

	var address models.Address

	err := json.Unmarshal([]byte(body), &address)
	if err != nil {
		return 400, "There is an error with received data " + err.Error()
	}

	if len(address.UUIDUser) == 0 || len(address.NameAddress) == 0 ||
		len(address.StateAddress) == 0 || len(address.CityAddress) == 0 ||
		len(address.PhoneAddress) == 0 || len(address.PostalCodeAddress) == 0 {
		return 400, "The UUUID of user, Name, State, City, Phone and Postal Code must be filled."
	}

	result, err := database.InsertAddress(address)
	if err != nil {
		return 400, "Error trying to INSERT the address ID > " + strconv.Itoa(address.IdAddress) + "\nError > " + err.Error()
	}

	return 200, "{ IdAddress: " + strconv.Itoa(int(result)) + "}"
}

func PutAddress(user, body string, id int) (int, string) {
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return 400, issue
	}

	var updateAddress models.Address
	err := json.Unmarshal([]byte(body), &updateAddress)
	if err != nil {
		return 400, "There is an issue with received data " + err.Error()
	}

	if len(updateAddress.UUIDUser) == 0 || len(updateAddress.NameAddress) == 0 ||
		len(updateAddress.StateAddress) == 0 || len(updateAddress.CityAddress) == 0 ||
		len(updateAddress.PhoneAddress) == 0 || len(updateAddress.PostalCodeAddress) == 0 {

		return 400, "Any of the Address's attributes needs to be filled."
	}

	result, err := database.UpdateAddress(updateAddress, id)
	if err != nil {
		return 400, "Error trying to UPDATE the address > " + strconv.Itoa(id)
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return 400, "Error trying to serialize result to JSON format."
	}

	return 200, string(jsonData)
}

func DeleteAddress(user string, id int) (int, string) {
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return 400, issue
	}

	result, err := database.DeleteAddress(id)
	if err != nil {
		return 400, "Error trying to DELETE the address with ID > " + strconv.Itoa(int(result))
	}

	return 200, "{ IdAddress: " + strconv.Itoa(int(result)) + "}"
}
