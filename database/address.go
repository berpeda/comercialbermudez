package database

import (
	"fmt"

	"github.com/berpeda/comercialbermudez/models"
)

// SelectAddress retrieves the addresses for a specific user
func SelectAddress(idUser string) ([]models.Address, error) {
	fmt.Println("Select a single Address function starts...")

	var userAddress []models.Address

	// Connect to the database
	err := DatabaseConnect()
	if err != nil {
		return userAddress, err
	}
	defer Database.Close()

	// Query to select addresses for a specific user
	query := "SELECT * FROM Direcciones WHERE UUID_usuario = ?"
	result, err := Database.Query(query, idUser)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return userAddress, err
	}
	defer result.Close()

	// Iterate over the results and scan into Address struct
	for result.Next() {
		var address models.Address
		err2 := result.Scan(&address.IdAddress,
			&address.UUIDUser,
			&address.NameAddress,
			&address.CityAddress,
			&address.StateAddress,
			&address.PhoneAddress,
			&address.PostalCodeAddress)
		if err2 != nil {
			fmt.Println("result.Scan is having issues...")
			return userAddress, err2
		}
		userAddress = append(userAddress, address)
	}

	fmt.Printf("Address selected successfully.")

	return userAddress, nil
}

// SelectAllAddress retrieves all addresses from the database
func SelectAllAddress() ([]models.Address, error) {
	fmt.Println("Select all Address function starts...")

	var addresses []models.Address

	// Connect to the database
	err := DatabaseConnect()
	if err != nil {
		return addresses, err
	}
	defer Database.Close()

	// Query to select all addresses
	query := "SELECT * FROM Direcciones"
	result, err := Database.Query(query)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return addresses, err
	}
	defer result.Close()

	// Iterate over the results and scan into Address struct
	for result.Next() {
		var address models.Address
		err = result.Scan(&address.IdAddress,
			&address.UUIDUser,
			&address.NameAddress,
			&address.CityAddress,
			&address.StateAddress,
			&address.PhoneAddress,
			&address.PostalCodeAddress)
		if err != nil {
			fmt.Println("Unable to Scan all the addresses > " + err.Error())
			panic(err)
		}
		addresses = append(addresses, address)
	}

	if err = result.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Addresses Selected successfully!")

	return addresses, nil
}

// InsertAddress inserts a new address into the database for a specific user
func InsertAddress(address models.Address, idUser string) (int64, error) {
	fmt.Println("Insert Address function starts...")

	// Connect to the database
	err := DatabaseConnect()
	if err != nil {
		return 0, err
	}
	defer Database.Close()

	// Query to insert a new address
	query := "INSERT INTO Direcciones (UUID_usuario, Nombre, Poblacion, Provincia, Telefono, Codigo_postal) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := Database.Exec(query, idUser,
		address.NameAddress, address.CityAddress,
		address.StateAddress, address.PhoneAddress, address.PostalCodeAddress)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return 0, err
	}

	// Retrieve the number of rows affected and the last inserted ID
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error retrieving the number of rows affected > ", err.Error())
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error retrieving the number of rows affected > ", err.Error())
	}

	fmt.Printf("Address inserted successfully.\nIndex inserted > %d\n The row(s) affected > %d", lastInsertId, rowsAffected)

	return lastInsertId, nil
}

// UpdateAddress updates an existing address in the database
func UpdateAddress(address models.Address, idAddress int) (models.Address, error) {
	fmt.Println("Update address is starting...")

	// Connect to the database
	err := DatabaseConnect()
	if err != nil {
		return address, err
	}
	defer Database.Close()

	// Query to update an address
	query := "UPDATE Direcciones SET UUID_usuario = ?, Nombre = ?, Poblacion = ?, Provincia = ?, Telefono = ?, Codigo_postal = ? WHERE Id_direccion = ?"
	_, err = Database.Exec(query, address.UUIDUser, address.NameAddress, address.CityAddress, address.StateAddress, address.PhoneAddress, address.PostalCodeAddress, idAddress)
	if err != nil {
		fmt.Println("Error with the UPDATE query > ", err.Error())
		return address, err
	}

	// Query to select the updated address
	query = "SELECT * FROM Direcciones WHERE Id_direccion = ?"
	result, err2 := Database.Query(query, idAddress)
	if err2 != nil {
		fmt.Println("Error with the SELECT query > ", err2.Error())
		return address, err
	}
	defer result.Close()

	result.Next()
	err = result.Scan(&address.IdAddress, &address.UUIDUser, &address.NameAddress, &address.CityAddress, &address.StateAddress, &address.PhoneAddress, &address.PostalCodeAddress)
	if err != nil {
		return address, err
	}

	fmt.Println("The address has been updated successfully!")
	return address, nil
}

// DeleteAddress deletes an address from the database
func DeleteAddress(idAddress int) (int64, error) {
	fmt.Println("Delete Address is starting...")

	// Connect to the database
	err := DatabaseConnect()
	if err != nil {
		return 0, err
	}
	defer Database.Close()

	// Query to delete an address
	query := "DELETE FROM Direcciones WHERE Id_direccion = ?"
	result, err := Database.Exec(query, idAddress)
	if err != nil {
		fmt.Println("There is an error in query > " + err.Error())
		return 0, err
	}

	// Retrieve the number of rows affected and the last inserted ID
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error retrieving the number of rows affected > ", err.Error())
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error retrieving the number of rows affected > ", err.Error())
	}

	fmt.Printf("Address deleted successfully.\nIndex deleted > %d\n The row(s) affected > %d", lastInsertId, rowsAffected)

	return int64(idAddress), nil
}
