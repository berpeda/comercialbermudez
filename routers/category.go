package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/berpeda/comercialbermudez/database"
	"github.com/berpeda/comercialbermudez/models"
)

// ----------------------  POST --------------------------
func PostCategory(user, body string) (int, string) {

	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return http.StatusBadRequest, issue
	}

	var cat models.Category

	err := json.Unmarshal([]byte(body), &cat)
	if err != nil {
		return http.StatusBadRequest, "There is an issue with received data " + err.Error()
	}

	if len(cat.NameCategory) == 0 || len(cat.DescriptionCategory) == 0 {
		return http.StatusBadRequest, "Category Name and Description needs to be filled"
	}

	result, err := database.InsertCategory(cat)
	if err != nil {
		return http.StatusBadRequest, "Error trying to INSERT the Category > " + cat.NameCategory + "\nError > " + err.Error()
	}

	return http.StatusOK, "{ IdCategory: " + strconv.Itoa(int(result)) + "}"
}

// ----------------------  GETS --------------------------
func GetCategory(IdCategory int) (int, string) {

	result, err := database.SelectCategory(IdCategory)
	if err != nil {
		return http.StatusBadRequest, "Error trying to SELECT the category with Id > " + strconv.Itoa(IdCategory) +
			"\nError > " + err.Error()
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return http.StatusBadRequest, "Error trying to resialize result to JSON format."
	}

	return http.StatusOK, string(jsonData)
}

func GetAllCategories() (int, string) {
	result, err := database.SelectAllCategories()
	if err != nil {
		return http.StatusBadRequest, "Error trying to SELECT all the Products. \nError > " + err.Error()
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return http.StatusBadRequest, "Error trying to serialize result to JSON format."
	}

	return http.StatusOK, string(jsonData)
}

// ----------------------  PUT --------------------------
func PutCategory(user, body string, id int) (int, string) {

	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return http.StatusBadRequest, issue
	}

	var cat models.Category

	err := json.Unmarshal([]byte(body), &cat)
	if err != nil {
		return http.StatusBadRequest, "There is an issue with received data " + err.Error()
	}

	if len(cat.NameCategory) == 0 && len(cat.DescriptionCategory) == 0 {
		return http.StatusBadRequest, "Category Name and Description needs to be filled"
	}

	result, err := database.UpdateCategory(cat, id)
	if err != nil {
		return http.StatusBadRequest, "Error trying to UPDATE the Category > " + cat.NameCategory + "\nError > " + err.Error()
	}

	return http.StatusOK, "{ IdCategory: " + strconv.Itoa(int(result)) + "}"
}

// ----------------------  DELETE --------------------------
func DeleteCategory(user string, IdCategory int) (int, string) {

	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return http.StatusBadRequest, issue
	}

	result, err := database.DeleteCategory(IdCategory)
	if err != nil {
		return http.StatusBadRequest, "Error trying to DELETE a Category with ID > " + strconv.Itoa(int(result)) + " . \nError > " + err.Error()
	}

	return http.StatusOK, "{ IdCategory: " + strconv.Itoa(int(result)) + " }"
}
