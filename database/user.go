package database

import (
	"database/sql"
	"fmt"

	"github.com/berpeda/comercialbermudez/models"
)

func UpdateUser(user models.User, idUser string) (models.User, error) {
	fmt.Println("Update User function is starting...")

	var us models.User

	err := DatabaseConnect()
	if err != nil {
		return us, err
	}

	defer Database.Close()

	query := "UPDATE Usuarios SET"
	params := []interface{}{}
	sets := ""

	if len(user.NameUser) != 0 {
		sets += " Nombre = ?,"
		params = append(params, user.NameUser)
	}
	if len(user.SurnameUser) != 0 {
		sets += " Apellidos = ?,"
		params = append(params, user.SurnameUser)
	}

	sets = sets[:len(sets)-1]
	query += sets + " WHERE UUID_usuario = ?"
	params = append(params, idUser)

	_, err = Database.Exec(query, params...)
	if err != nil {
		fmt.Println("Error trying the UPDATE sentece.")
		fmt.Println("Sentence > ", query)
		return us, err
	}

	fmt.Println(query)

	query = "SELECT * FROM Usuarios WHERE UUID_usuario = ?"
	result, err := Database.Query(query, idUser)
	if err != nil {
		fmt.Println("Error trying to do the SELECT sentence after the update.")
		return us, err
	}

	defer result.Close()

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

	fmt.Println("The user has been updated succesfully!")
	return us, nil
}

func SelectAllUsers() (models.UserDetails, error) {
	fmt.Println("Select All Users function is starting...")

	var usersDetails models.UserDetails

	err := DatabaseConnect()
	if err != nil {
		return usersDetails, err
	}

	defer Database.Close()

	query := "SELECT * FROM Usuarios"
	result, err := Database.Query(query)
	if err != nil {
		return usersDetails, err
	}

	defer result.Close()

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

	fmt.Println("The users has been selected succesfully!")
	return usersDetails, nil
}

func SelectMyUser(idUser string) (models.User, error) {
	fmt.Println("Select a User function is starting...")

	var user models.User

	err := DatabaseConnect()
	if err != nil {
		return user, err
	}

	defer Database.Close()

	query := "SELECT * FROM Usuarios WHERE UUID_usuario = ?"
	result, err := Database.Query(query, idUser)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return user, err
	}

	defer result.Close()

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

	fmt.Println("The user has been selected succesfully!")
	return user, nil
}
