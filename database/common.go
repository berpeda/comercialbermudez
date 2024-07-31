package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/berpeda/comercialbermudez/models"
	secretsmanager "github.com/berpeda/comercialbermudez/secretsManager"
	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var Database *sql.DB

// ReadSecret retrieves the secret credentials for the database from the secrets manager.
func ReadScecret() error {
	SecretModel, err = secretsmanager.GetSecrets(os.Getenv("SecretName"))
	return err
}

// DatabaseConnect establishes a connection to the database using credentials from the secrets manager.
func DatabaseConnect() error {
	// Open a new connection to the database
	Database, err = sql.Open("mysql", ConnectionString(SecretModel))
	if err != nil {
		fmt.Println("The connection to the database failed -> ", err.Error())
		return err
	}

	// Test the connection to the database
	err = Database.Ping()
	if err != nil {
		fmt.Println("Ping got an error -> ", err)
	}

	fmt.Println("The connection is successful!")
	return nil
}

// ConnectionString constructs the connection string for the database using the provided secret credentials.
func ConnectionString(keys models.SecretRDSJson) string {
	var dbUser string = keys.Username
	var authToken string = keys.Password
	var dbEndpoint string = keys.Host
	var dbName string = "comercialbermudez"

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
}

// IsAdmin checks if a user with the given UUID has admin privileges.
func IsAdmin(userUUID string) (bool, string) {
	fmt.Println("Checking IsAdmin...")

	// Establish a connection to the database
	err := DatabaseConnect()
	if err != nil {
		return false, err.Error()
	}
	defer Database.Close()

	// Query to check if the user has admin privileges
	query := "SELECT 1 FROM Usuarios WHERE UUID_usuario = ? AND Rol = 0"
	fmt.Println(query)

	res, err := Database.Query(query, userUUID)
	if err != nil {
		return false, err.Error()
	}

	defer res.Close()

	var value string
	res.Next()
	res.Scan(&value)

	if value == "1" {
		fmt.Println("The user is an admin!")
		return true, "The user is an admin!"
	}

	fmt.Println("The user is not an admin!")
	return false, "The user is not an admin!"
}
