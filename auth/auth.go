package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// TokenJSON represents the structure of the JWT payload that we are interested in.
type TokenJSON struct {
	Sub       string // Subject
	Event_Id  string // Event ID
	Token_use string // Token usage
	Scope     string // Scope
	Auth_time int    // Authentication time
	Iss       string // Issuer
	Exp       int    // Expiry time
	Iat       int    // Issued at time
	Client_id string // Client ID
	Username  string // Username
}

// TokenValidation validates the provided JWT token.
// It checks if the token structure is valid, decodes it, and verifies its expiry time.
func TokenValidation(token string) (bool, error, string) {
	// Split the token into its components (header, payload, signature)
	substring := strings.Split(token, ".")

	// Ensure the token has exactly three parts
	if len(substring) != 3 {
		fmt.Println("Invalid token, length < 3")
		return false, nil, "Invalid token"
	}

	// Decode the base64-encoded payload part of the token
	userInfo, err := base64.StdEncoding.DecodeString(substring[1])
	if err != nil {
		fmt.Println("Token cannot be decoded", err.Error())
		return false, err, err.Error()
	}

	// Unmarshal the JSON payload into the TokenJSON struct
	var nToken TokenJSON
	err = json.Unmarshal(userInfo, &nToken)
	if err != nil {
		fmt.Println("Cannot decode JSON structure", err.Error())
		return false, err, err.Error()
	}

	// Get the current time and the token's expiry time
	now := time.Now()
	timeU := time.Unix(int64(nToken.Exp), 0)

	// Check if the token has expired
	if timeU.Before(now) {
		fmt.Println("Expired time token > " + timeU.String())
		fmt.Println("The token has expired")
		return false, err, "The token has expired"
	}

	// Token is valid
	return true, nil, string(nToken.Username)
}
