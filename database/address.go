package database

import (
	"fmt"

	"github.com/berpeda/comercialbermudez/models"
)

func SelectAddress(id int) (models.Address, error) {
	fmt.Println("Select a single Address function starts...")

	var address models.Address

	err := DatabaseConnect()
	if err != nil {
		return address, err
	}

	defer Database.Close()

	query := "SELECT * FROM Direcciones WHERE Id_direccion = ?"

	result, err := Database.Query(query, id)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return address, err
	}

	result.Next()
	err2 := result.Scan(&address.IdAddress,
		&address.UUIDUser,
		&address.NameAddress,
		&address.CityAddress,
		&address.StateAddress,
		&address.PhoneAddress,
		&address.PostalCodeAddress)

	if err2 != nil {
		fmt.Println("result.Scan is having issues...")
		return address, err2
	}

	fmt.Printf("Address selected successfully.")

	return address, nil
}

func SelectAllAddress() ([]models.Address, error) {
	fmt.Println("Select all Address function starts...")

	var addresses []models.Address

	err := DatabaseConnect()
	if err != nil {
		return addresses, err
	}

	defer Database.Close()

	query := "SELECT * FROM Direcciones"
	result, err := Database.Query(query)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return addresses, err
	}

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

func InsertAddress(address models.Address) (int64, error) {
	fmt.Println("Insert Address function starts...")

	err := DatabaseConnect()
	if err != nil {
		return 0, err
	}

	defer Database.Close()

	query := "INSERT INTO Direcciones (UUID_usuario, Nombre, Poblacion, Provincia, Telefono, Codigo_postal) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := Database.Exec(query, address.UUIDUser,
		address.NameAddress, address.CityAddress,
		address.StateAddress, address.PhoneAddress, address.PostalCodeAddress)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return 0, err
	}

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

func UpdateAddress(address models.Address, idAddress int) (models.Address, error) {
	fmt.Println("Update address is starting...")

	err := DatabaseConnect()
	if err != nil {
		return address, err
	}

	defer Database.Close()

	query := "UPDATE Direcciones SET UUID_usuario = ?, Nombre = ?, Poblacion = ?, Provincia = ?, Telefono = ?, Codigo_postal = ? WHERE Id_direccion = ?"
	_, err = Database.Exec(query, address.UUIDUser, address.NameAddress, address.CityAddress, address.StateAddress, address.PhoneAddress, address.PostalCodeAddress, idAddress)
	if err != nil {
		fmt.Println("Error with the UPDATE query > ", err.Error())
		return address, err
	}

	query = "SELECT * FROM Direcciones WHERE Id_direccion = ?"
	result, err2 := Database.Query(query, idAddress)
	if err2 != nil {
		fmt.Println("Error with the SELECT query > ", err2.Error())
		return address, err
	}

	result.Next()
	err = result.Scan(&address.IdAddress, &address.UUIDUser, &address.NameAddress, &address.CityAddress, &address.StateAddress, &address.PhoneAddress, &address.PostalCodeAddress)
	if err != nil {
		return address, err
	}

	fmt.Println("The address has been updated successfully!")
	return address, nil
}

func DeleteAddress(idAddress int) (int64, error) {
	fmt.Println("Delete Address is starting...")

	err := DatabaseConnect()
	if err != nil {
		return 0, err
	}

	defer Database.Close()

	query := "DELETE FROM Direcciones WHERE Id_direccion = ?"
	result, err := Database.Exec(query, idAddress)
	if err != nil {
		fmt.Println("There is an error in query > " + err.Error())
		return 0, err
	}

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
