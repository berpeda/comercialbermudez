package database

import (
	"fmt"

	"github.com/berpeda/comercialbermudez/models"
)

func SelectProvider(idProvider int) (models.Provider, error) {
	fmt.Println("Select a single Provider starts...")

	var provider models.Provider

	err := DatabaseConnect()
	if err != nil {
		return provider, err
	}

	defer Database.Close()

	query := "SELECT * FROM Proveedores WHERE Id_proveedor = ?"

	result, err := Database.Query(query, idProvider)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return provider, err
	}

	result.Next()
	err2 := result.Scan(&provider.IdProvider,
		&provider.NameProvider,
		&provider.PhoneNumberProvider,
		&provider.EmailProvider)

	if err2 != nil {
		fmt.Println("result.Scan is having issues...")
		return provider, err2
	}

	fmt.Printf("Provider selected successfully.")

	return provider, nil
}

func SelectAllProviders() ([]models.Provider, error) {
	fmt.Println("Select all Provider starts...")

	var providers []models.Provider

	err := DatabaseConnect()
	if err != nil {
		return providers, err
	}

	defer Database.Close()

	query := "SELECT * FROM Proveedores"
	result, err := Database.Query(query)
	if err != nil {
		fmt.Println("Error with the query > ", err.Error())
		return providers, err
	}

	for result.Next() {
		var provider models.Provider
		err = result.Scan(&provider.IdProvider,
			&provider.NameProvider,
			&provider.PhoneNumberProvider,
			&provider.EmailProvider)

		if err != nil {
			fmt.Println("Unable to Scan all the providers > " + err.Error())
			panic(err)
		}
		providers = append(providers, provider)
	}

	if err = result.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Providers Selected successfully!")

	return providers, nil
}

func InsertProvider(provider models.Provider) (int64, error) {
	fmt.Println("Insert a Provider starts...")

	err := DatabaseConnect()
	if err != nil {
		return 0, err
	}

	defer Database.Close()

	query := "INSERT INTO Proveedores (Nombre, Telefono, Email) VALUES (?, ?, ?)"
	result, err := Database.Exec(query, provider.NameProvider, provider.PhoneNumberProvider, provider.EmailProvider)
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

	fmt.Printf("Provider inserted successfully.\nIndex inserted > %d\n The row(s) affected > %d", lastInsertId, rowsAffected)

	return lastInsertId, nil
}

func UpdateProvider(provider models.Provider, idProvider int) (models.Provider, error) {
	fmt.Println("Update provider is starting...")

	err := DatabaseConnect()
	if err != nil {
		return provider, err
	}

	defer Database.Close()

	query := "UPDATE Proveedores SET Nombre = ?, Telefono = ?, Email = ? WHERE Id_proveedor = ?"
	_, err = Database.Exec(query, provider.NameProvider, provider.PhoneNumberProvider, provider.EmailProvider, idProvider)
	if err != nil {
		fmt.Println("Error with the UPDATE query > ", err.Error())
		return provider, err
	}

	query = "SELECT * FROM Proveedores WHERE Id_proveedor = ?"
	result, err := Database.Query(query, idProvider)
	if err != nil {
		fmt.Println("Error with the SELECT query > ", err.Error())
		return provider, err
	}

	result.Next()
	err = result.Scan(&provider.IdProvider, &provider.NameProvider, &provider.PhoneNumberProvider, &provider.EmailProvider)
	if err != nil {
		fmt.Println("Unable to Scan all the provider > " + err.Error())
		return provider, err
	}

	fmt.Println("The provider has been updated successfully!")
	return provider, nil
}

func DeleteProvider(idProvider int) (int64, error) {
	fmt.Println("Delete provider is starting...")

	err := DatabaseConnect()
	if err != nil {
		return 0, err
	}

	defer Database.Close()

	query := "DELETE FROM Proveedores WHERE Id_proveedor = ?"
	result, err := Database.Exec(query, idProvider)
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

	fmt.Printf("Provider deleted successfully.\nIndex deleted > %d\n The row(s) affected > %d", lastInsertId, rowsAffected)

	return int64(idProvider), nil
}
