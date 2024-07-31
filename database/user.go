package database

import (
	"database/sql"
	"fmt"

	"github.com/berpeda/comercialbermudez/models"
)

// UpdateUser updates an existing user based on its UUID.
func UpdateUser(user models.User, idUser string) (models.User, error) {
	fmt.Println("Update User function is starting...")

	var us models.User

	// Connect to the database
	err := DatabaseConnect()
	if err != nil {
		return us, err
	}

	defer Database.Close()

	// Start building the UPDATE query
	query := "UPDATE Usuarios SET"
	params := []interface{}{}
	sets := ""

	// Conditionally add fields to be updated
	if len(user.NameUser) != 0 {
		sets += " Nombre = ?,"
		params = append(params, user.NameUser)
	}
	if len(user.SurnameUser) != 0 {
		sets += " Apellidos = ?,"
		params = append(params, user.SurnameUser)
	}

	// Remove trailing comma and append WHERE clause
	sets = sets[:len(sets)-1]
	query += sets + " WHERE UUID_usuario = ?"
	params = append(params, idUser)

	// Execute the UPDATE query
	_, err = Database.Exec(query, params...)
	if err != nil {
		fmt.Println("Error trying the UPDATE sentence.")
		fmt.Println("Sentence > ", query)
		return us, err
	}

	fmt.Println(query)

	// Retrieve the updated user details
	query = "SELECT * FROM Usuarios WHERE UUID_usuario = ?"
	result, err := Database.Query(query, idUser)
	if err != nil {
		fmt.Println("Error trying to do the SELECT sentence after the update.")
		return us, err
	}

	defer result.Close()

	// Scan the result into the user struct
	var userName sql.NullString
	var userSurname sql.NullString

	result.Next()
	err2 := result.Scan(&us.UUIDUser, &userName, &userSurname, &us.EmailUser, &us.RolUser, &us.CreatedAt)
	if err2 != nil {
		fmt.Println("Error trying to Scan the user.")
		return us, err2
	}

	us.NameUser = userName.String
	us.SurnameUser = userSurname.String

	fmt.Println("The user has been updated successfully!")
	return us, nil
}

// SelectAllUsers retrieves all users from the database.
func SelectAllUsers() (models.UserDetails, error) {
	fmt.Println("Select All Users function is starting...")

	var usersDetails models.UserDetails

	// Connect to the database
	err := DatabaseConnect()
	if err != nil {
		return usersDetails, err
	}

	defer Database.Close()

	// Query to select all users
	query := "SELECT * FROM Usuarios"
	result, err := Database.Query(query)
	if err != nil {
		return usersDetails, err
	}

	defer result.Close()

	// Scan each result into the usersDetails struct
	for result.Next() {
		var user models.User
		var userName sql.NullString
		var userSurname sql.NullString
		err2 := result.Scan(&user.UUIDUser,
			&userName,
			&userSurname,
			&user.EmailUser,
			&user.RolUser,
			&user.CreatedAt)

		if err2 != nil {
			fmt.Println("Error trying to Scan the users.")
			return usersDetails, err2
		}
		user.NameUser = userName.String
		user.SurnameUser = userSurname.String
		usersDetails.TotalUsers = append(usersDetails.TotalUsers, user)
	}

	fmt.Println("The users have been selected successfully!")
	return usersDetails, nil
}

// SelectMyUser retrieves a single user based on its UUID.
func SelectMyUser(idUser string) (models.User, error) {
	fmt.Println("Select a User function is starting...")

	var user models.User

	// Connect to the database
	err := DatabaseConnect()
	if err != nil {
		return user, err
	}

	defer Database.Close()

	// Query to select a user by its UUID
	query := "SELECT * FROM Usuarios WHERE UUID_usuario = ?"
	result, err := Database.Query(query, idUser)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return user, err
	}

	defer result.Close()

	// Scan the result into the user struct
	result.Next()

	var userName sql.NullString
	var userSurname sql.NullString
	err2 := result.Scan(&user.UUIDUser, &userName, &userSurname, &user.EmailUser, &user.RolUser, &user.CreatedAt)
	if err2 != nil {
		fmt.Println("Error trying to scan the user > ", err2.Error())
		return user, err
	}
	user.NameUser = userName.String
	user.SurnameUser = userSurname.String

	fmt.Println("The user has been selected successfully!")
	return user, nil
}
