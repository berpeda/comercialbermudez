package routers

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/berpeda/comercialbermudez/database"
	"github.com/berpeda/comercialbermudez/models"
)

func PutUser(body, user string) (int, string) {
	var us models.User

	err := json.Unmarshal([]byte(body), &us)
	if err != nil {
		return http.StatusBadRequest, "There is an error with received data " + err.Error()
	}

	if len(us.NameUser) == 0 && len(us.SurnameUser) == 0 {
		return http.StatusBadRequest, "Name or Surname must be filled to update correctly the user."
	}

	result, err := database.UpdateUser(us, user)
	if err != nil {
		return http.StatusBadRequest, "Error trying to UPDATE the User > " + us.NameUser + "\n" + err.Error()
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return http.StatusBadRequest, "Error trying to serialize result to JSON format."
	}
	return http.StatusOK, string(jsonData)
}

func GetAllUsers(user string, request events.APIGatewayV2HTTPRequest) (int, string) {

	// Verify if user is Admin
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return http.StatusBadRequest, issue
	}

	// Select All users
	result, err := database.SelectAllUsers()
	if err != nil {
		return http.StatusInternalServerError, "Error trying to SELECT All Users > " + err.Error()
	}

	// Serialize to JSON the data
	jsonData, err := json.Marshal(result)
	if err != nil {
		return http.StatusInternalServerError, "Error trying to serialize user"
	}

	return http.StatusOK, string(jsonData)
}

func GetMyUser(user string) (int, string) {

	result, err := database.SelectMyUser(user)
	if err != nil {
		return http.StatusBadRequest, err.Error()
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return http.StatusInternalServerError, "Error trying to serialize user"
	}

	return http.StatusOK, string(jsonData)
}
