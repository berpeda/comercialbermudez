package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/berpeda/comercialbermudez/auth"
	"github.com/berpeda/comercialbermudez/routers"
)

func Handlers(path, method, body string, header map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {

	fmt.Println("Proccessing " + path + " > " + method)

	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

	isOk, statusCode, user := ValidAuthorization(path, method, header)
	if !isOk {
		return statusCode, user
	}

	fmt.Println("path[0:5] = " + path[0:5])

	switch path[0:5] {
	case "/user":
		return UserActions(body, path, method, user, id, request)
	case "/prod":
		return ProductsActions(body, path, method, user, idn, request)
	case "/odet":
		return OrderDetailsActions(body, path, method, user, idn, request)
	case "/addr":
		return AddressActions(body, path, method, user, idn, request)
	case "/prov":
		return ProviderActions(body, path, method, user, idn, request)
	case "/cate":
		return CategoryActions(body, path, method, user, idn, request)
	case "/orde":
		return OrderActions(body, path, method, user, idn, request)
	}

	return http.StatusBadRequest, "Invalid Method"
}

func ValidAuthorization(path, method string, headers map[string]string) (bool, int, string) {

	if (path[:5] == "/prod" && method == "GET") || (path[:5] == "/cate" && method == "GET") {
		return true, http.StatusOK, ""
	}

	token := headers["authorization"]
	if len(token) == 0 {
		return false, http.StatusUnauthorized, "Token is required"
	}

	isOK, err, msg := auth.TokenValidation(token)
	if !isOK {
		if err != nil {
			fmt.Println("Error in the token: " + err.Error())
			return false, http.StatusUnauthorized, err.Error()
		}
		fmt.Println("Error in the token: " + msg)
		return false, http.StatusUnauthorized, msg
	}

	fmt.Println("Authorization OK!")
	return true, http.StatusOK, msg
}

func UserActions(body, path, method, user, id string, request events.APIGatewayV2HTTPRequest) (int, string) {
	if path == "/user/me" {
		if method == "GET" {
			return routers.GetMyUser(user)
		} else if method == "PUT" {
			return routers.PutUser(body, user)
		}
	}

	if path == "/user" {
		return routers.GetAllUsers(user, request)
	}

	return http.StatusBadRequest, "Method Invalid"
}

func ProductsActions(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "GET":
		return routers.GetProduct(request)

	case "POST":
		return routers.PostProduct(user, body)

	case "PUT":
		return routers.PutProduct(user, body, id)

	case "DELETE":
		return routers.DeleteProduct(user, id)
	}
	return http.StatusBadRequest, "Method Invalid"
}

func CategoryActions(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "GET":
		if id != 0 {
			return routers.GetCategory(id)
		}
		return routers.GetAllCategories()
	case "POST":
		return routers.PostCategory(user, body)
	case "PUT":
		if id == 0 {
			return http.StatusBadRequest, "ID is required for PUT method"
		}
		return routers.PutCategory(user, body, id)
	case "DELETE":
		if id == 0 {
			return http.StatusBadRequest, "ID is required for DELETE method"
		}
		return routers.DeleteCategory(user, id)
	}
	return http.StatusBadRequest, "Method Invalid"
}

func OrderDetailsActions(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "GET":
		if id != 0 {
			return routers.GetOrderDetail(id)
		}
		return routers.GetAllOrderDetails()
	case "POST":
		return routers.PostOrderDetail(user, body)
	case "PUT":
		return routers.PutOrderDetail(user, body, id)
	case "DELETE":
		return routers.DeleteOrderDetail(user, id)
	}
	return http.StatusBadRequest, "Method Invalid"
}

func OrderActions(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "GET":
		if id != 0 {
			return routers.GetOrder(id)
		}
		return routers.GetAllOrders()
	case "POST":
		return routers.PostOrder(user, body)
	case "PUT":
		return routers.PutOrder(user, body, id)
	case "DELETE":
		return routers.DeleteOrder(user, id)
	}
	return http.StatusBadRequest, "Method Invalid"
}

func AddressActions(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "GET":
		if id != 0 {
			return routers.GetAddress(id)
		}
		return routers.GetAllAddress()
	case "POST":
		return routers.PostAddress(user, body)
	case "PUT":
		return routers.PutAddress(user, body, id)
	case "DELETE":
		return routers.DeleteAddress(user, id)
	}
	return http.StatusBadRequest, "Method Invalid"
}

func ProviderActions(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "GET":
		if id != 0 {
			return routers.GetProvider(id)
		}
		return routers.GetAllProviders()
	case "POST":
		return routers.PostProvider(user, body)
	case "PUT":
		return routers.PutProvider(user, body, id)
	case "DELETE":
		return routers.DeleteProvider(user, id)
	}
	return http.StatusBadRequest, "Method Invalid"
}
