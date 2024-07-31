package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/berpeda/comercialbermudez/database"
	"github.com/berpeda/comercialbermudez/models"
)

// PostCategory inserts a new category, but only if the user is an admin.
func PostCategory(user, body string) (int, string) {
	// Check if the user is an admin
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		// Return a bad request status and issue message if the user is not an admin
		return http.StatusBadRequest, issue
	}

	var cat models.Category

	// Deserialize the JSON body into a Category struct
	err := json.Unmarshal([]byte(body), &cat)
	if err != nil {
		// Return a bad request status and error message if deserialization fails
		return http.StatusBadRequest, "There is an issue with received data " + err.Error()
	}

	// Check if all required fields are filled
	if len(cat.NameCategory) == 0 || len(cat.DescriptionCategory) == 0 {
		// Return a bad request status if any required field is missing
		return http.StatusBadRequest, "Category Name and Description need to be filled"
	}

	// Insert the category into the database
	result, err := database.InsertCategory(cat)
	if err != nil {
		// Return a bad request status and error message if the database insert fails
		return http.StatusBadRequest, "Error trying to INSERT the Category > " + cat.NameCategory + "\nError > " + err.Error()
	}

	// Return a success status and the ID of the inserted category
	return http.StatusOK, "{ IdCategory: " + strconv.Itoa(int(result)) + "}"
}

// GetCategory retrieves a specific category by its ID.
func GetCategory(IdCategory int) (int, string) {
	// Fetch the category from the database
	result, err := database.SelectCategory(IdCategory)
	if err != nil {
		// Return a bad request status and error message if the database query fails
		return http.StatusBadRequest, "Error trying to SELECT the category with Id > " + strconv.Itoa(IdCategory) +
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

// GetAllCategories retrieves all categories.
func GetAllCategories() (int, string) {
	// Fetch all categories from the database
	result, err := database.SelectAllCategories()
	if err != nil {
		// Return a bad request status and error message if the database query fails
		return http.StatusBadRequest, "Error trying to SELECT all the Products. \nError > " + err.Error()
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

// PutCategory updates an existing category, but only if the user is an admin.
func PutCategory(user, body string, id int) (int, string) {
	// Check if the user is an admin
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		// Return a bad request status and issue message if the user is not an admin
		return http.StatusBadRequest, issue
	}

	var cat models.Category

	// Deserialize the JSON body into a Category struct
	err := json.Unmarshal([]byte(body), &cat)
	if err != nil {
		// Return a bad request status and error message if deserialization fails
		return http.StatusBadRequest, "There is an issue with received data " + err.Error()
	}

	// Check if all required fields are filled
	if len(cat.NameCategory) == 0 || len(cat.DescriptionCategory) == 0 {
		// Return a bad request status if any required field is missing
		return http.StatusBadRequest, "Category Name and Description need to be filled"
	}

	// Update the category in the database
	result, err := database.UpdateCategory(cat, id)
	if err != nil {
		// Return a bad request status and error message if the database update fails
		return http.StatusBadRequest, "Error trying to UPDATE the Category > " + cat.NameCategory + "\nError > " + err.Error()
	}

	// Return a success status and the ID of the updated category
	return http.StatusOK, "{ IdCategory: " + strconv.Itoa(int(result)) + "}"
}

// DeleteCategory deletes a category, but only if the user is an admin.
func DeleteCategory(user string, IdCategory int) (int, string) {
	// Check if the user is an admin
	isAdmin, issue := database.IsAdmin(user)
	if !isAdmin {
		// Return a bad request status and issue message if the user is not an admin
		return http.StatusBadRequest, issue
	}

	// Delete the category from the database
	result, err := database.DeleteCategory(IdCategory)
	if err != nil {
		// Return a bad request status and error message if the database delete fails
		return http.StatusBadRequest, "Error trying to DELETE a Category with ID > " + strconv.Itoa(int(result)) + " . \nError > " + err.Error()
	}

	// Return a success status and the ID of the deleted category
	return http.StatusOK, "{ IdCategory: " + strconv.Itoa(int(result)) + " }"
}
