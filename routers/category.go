package routers

import (
	"encoding/json"
	"strconv"

	"github.com/berpeda/comercialbermudez/database"
	"github.com/berpeda/comercialbermudez/models"
)

// ----------------------  POST --------------------------
func PostCategory(user, body string) (int, string) {

	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return 400, issue
	}

	var cat models.Category

	err := json.Unmarshal([]byte(body), &cat)
	if err != nil {
		return 400, "There is an issue with received data " + err.Error()
	}

	if len(cat.NameCategory) == 0 || len(cat.DescriptionCategory) == 0 {
		return 400, "Category Name and Description needs to be filled"
	}

	result, err := database.InsertCategory(cat)
	if err != nil {
		return 400, "Error trying to INSERT the Category > " + cat.NameCategory + "\nError > " + err.Error()
	}

	return 200, "{ IdCategory: " + strconv.Itoa(int(result)) + "}"
}

// ----------------------  GETS --------------------------
func GetCategory(IdCategory int) (int, string) {

	result, err := database.SelectCategory(IdCategory)
	if err != nil {
		return 400, "Error trying to SELECT the category with Id > " + strconv.Itoa(IdCategory) +
			"\nError > " + err.Error()
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return 400, "Error trying to resialize result to JSON format."
	}

	return 200, string(jsonData)
}

func GetAllCategories() (int, string) {
	result, err := database.SelectAllCategories()
	if err != nil {
		return 400, "Error trying to SELECT all the Products. \nError > " + err.Error()
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return 400, "Error trying to serialize result to JSON format."
	}

	return 200, string(jsonData)
}

// ----------------------  PUT --------------------------
func PutCategory(user, body string, id int) (int, string) {

	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return 400, issue
	}

	var cat models.Category

	err := json.Unmarshal([]byte(body), &cat)
	if err != nil {
		return 400, "There is an issue with received data " + err.Error()
	}

	if len(cat.NameCategory) == 0 && len(cat.DescriptionCategory) == 0 {
		return 400, "Category Name and Description needs to be filled"
	}

	result, err := database.UpdateCategory(cat, id)
	if err != nil {
		return 400, "Error trying to UPDATE the Category > " + cat.NameCategory + "\nError > " + err.Error()
	}

	return 200, "{ IdCategory: " + strconv.Itoa(int(result)) + "}"
}

// ----------------------  DELETE --------------------------
func DeleteCategory(user string, IdCategory int) (int, string) {

	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		return 400, issue
	}

	result, err := database.DeleteCategory(IdCategory)
	if err != nil {
		return 400, "Error trying to DELETE a Category with ID > " + strconv.Itoa(int(result)) + " . \nError > " + err.Error()
	}

	return 200, "{ IdCategory: " + strconv.Itoa(int(result)) + " }"
}
