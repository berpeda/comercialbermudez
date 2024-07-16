package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TokenJSON struct {
	Sub       string
	Event_Id  string
	Token_use string
	Scope     string
	Auth_time int
	Iss       string
	Exp       int
	Iat       int
	Client_id string
	Username  string
}

func TokenValidation(token string) (bool, error, string) {
	substring := strings.Split(token, ".")

	if len(substring) != 3 {
		fmt.Println("Invalid token, length < 3")
		return false, nil, "Invalid token"
	}

	userInfo, err := base64.StdEncoding.DecodeString(substring[1])
	if err != nil {
		fmt.Println("Token cannot be decoded", err.Error())
		return false, err, err.Error()
	}

	var nToken TokenJSON
	err = json.Unmarshal(userInfo, &nToken)
	if err != nil {
		fmt.Println("Cannot decode JSON structure", err.Error())
		return false, err, err.Error()
	}

	now := time.Now()
	timeU := time.Unix(int64(nToken.Exp), 0)

	if timeU.Before(now) {
		fmt.Println("Expired time token > " + timeU.String())
		fmt.Println("The token has expired")
		return false, err, "The token has expired"
	}

	return true, nil, string(nToken.Username)
}
