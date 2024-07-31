package routers

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/berpeda/comercialbermudez/database"
	"github.com/berpeda/comercialbermudez/models"
)

// PutUser updates user information in the database
func PutUser(body, user string) (int, string) {
	var us models.User

	// Deserialize the JSON body into a User struct
	err := json.Unmarshal([]byte(body), &us)
	if err != nil {
		return http.StatusBadRequest, "There is an error with received data " + err.Error()
	}

	// Check if at least Name or Surname is filled
	if len(us.NameUser) == 0 && len(us.SurnameUser) == 0 {
		return http.StatusBadRequest, "Name or Surname must be filled to update the user correctly."
	}

	// Update the user in the database
	result, err := database.UpdateUser(us, user)
	if err != nil {
		return http.StatusBadRequest, "Error trying to UPDATE the User > " + us.NameUser + "\n" + err.Error()
	}

	// Serialize the result to JSON format
	jsonData, err := json.Marshal(result)
	if err != nil {
		return http.StatusBadRequest, "Error trying to serialize result to JSON format."
	}
	return http.StatusOK, string(jsonData)
}

// GetAllUsers retrieves all users from the database, but only if the requester is an admin
func GetAllUsers(user string, request events.APIGatewayV2HTTPRequest) (int, string) {
	// Verify if the user is an admin
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return http.StatusBadRequest, issue
	}

	// Fetch all users from the database
	result, err := database.SelectAllUsers()
	if err != nil {
		return http.StatusInternalServerError, "Error trying to SELECT All Users > " + err.Error()
	}

	// Serialize the result to JSON format
	jsonData, err := json.Marshal(result)
	if err != nil {
		return http.StatusInternalServerError, "Error trying to serialize user"
	}

	return http.StatusOK, string(jsonData)
}

// GetMyUser retrieves the information of the user making the request
func GetMyUser(user string) (int, string) {
	// Fetch the user's information from the database
	result, err := database.SelectMyUser(user)
	if err != nil {
		return http.StatusBadRequest, err.Error()
	}

	// Serialize the result to JSON format
	jsonData, err := json.Marshal(result)
	if err != nil {
		return http.StatusInternalServerError, "Error trying to serialize user"
	}

	return http.StatusOK, string(jsonData)
}
