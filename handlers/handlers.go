package handlers

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/berpeda/comercialbermudez/auth"
)

func Handlers(path, method, body string, header map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {

	fmt.Println("Proccessing " + path + " > " + method)

	id := request.PathParameters["id "]
	idn, _ := strconv.Atoi(id)

	isOk, statusCode, user := ValidAuthorization(path, method, header)
	if !isOk {
		return statusCode, user
	}

	switch path[0:4] {
	case "user":
		return ProcesoUsers(body, path, method, user, id, request)
	case "prod":
		return ProcesoProducts(body, path, method, user, idn, request)
	case "stoc":
		return ProcesoStock(body, path, method, user, idn, request)
	case "addr":
		return ProcesoAddress(body, path, method, user, idn, request)
	case "cate":
		return ProcesoCategory(body, path, method, user, idn, request)
	case "orde":
		return ProcesoOrder(body, path, method, user, idn, request)
	}

	return 400, "Invalid Method"
}

func ValidAuthorization(path, method string, headers map[string]string) (bool, int, string) {
	if (path == "product" && method == "GET") || (path == "category" && method == "GET") {
		return true, 200, ""
	}

	token := headers["authorization"]

	if len(token) == 0 {
		return false, 401, "Token is required"
	}

	isOK, err, msg := auth.TokenValidation(token)
	if !isOK {
		if err != nil {
			fmt.Println("Error in the token" + err.Error())
			return false, 401, err.Error()
		} else {
			fmt.Println("Error in the token " + msg)
			return false, 401, msg
		}
	}

	fmt.Println("Everything it's OK!")
	return true, 200, msg
}

func ProcesoUsers(body, path, method, user, id string, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}

func ProcesoProducts(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}

func ProcesoCategory(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}

func ProcesoStock(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}

func ProcesoOrder(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}

func ProcesoAddress(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}
