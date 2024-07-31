package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/berpeda/comercialbermudez/auth"
	"github.com/berpeda/comercialbermudez/routers"
)

// Handlers processes incoming requests based on the path, method, and headers.
// It routes the request to the appropriate function based on the path and HTTP method.
func Handlers(path, method, body string, header map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Processing " + path + " > " + method)

	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id) // Convert path parameter 'id' to integer

	isOk, statusCode, user := ValidAuthorization(path, method, header)
	if !isOk {
		return statusCode, user
	}

	fmt.Println("path[0:5] = " + path[0:5])

	// Route requests based on the initial path segment
	switch path[0:5] {
	case "/user":
		return UserActions(body, path, method, user, id, request)
	case "/prod":
		return ProductsActions(body, path, method, user, idn, request)
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

// ValidAuthorization checks if the request has a valid authorization token.
// It allows access to public routes and validates the token for protected routes.
func ValidAuthorization(path, method string, headers map[string]string) (bool, int, string) {
	// Allow public access for GET requests to products and categories
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

// UserActions handles user-related actions based on the HTTP method.
// It performs operations such as retrieving or updating user information.
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

// ProductsActions handles product-related actions based on the HTTP method.
// It supports operations like retrieving, creating, updating, or deleting products.
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

// CategoryActions handles category-related actions based on the HTTP method.
// It supports retrieving, creating, updating, or deleting categories.
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

// OrderActions handles order-related actions based on the HTTP method.
// It supports retrieving, creating, or deleting orders.
func OrderActions(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "GET":
		return routers.GetOrders(user, request)
	case "POST":
		return routers.PostOrder(user, body)
	case "DELETE":
		return routers.DeleteOrder(user, id)
	}
	return http.StatusBadRequest, "Method Invalid"
}

// AddressActions handles address-related actions based on the HTTP method.
// It supports retrieving, creating, updating, or deleting addresses.
func AddressActions(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println(path)
	switch method {
	case "GET":
		if path == "/address/me" {
			return routers.GetAddress(user)
		} else if path == "/address" {
			return routers.GetAllAddress(user)
		}
	case "POST":
		return routers.PostAddress(user, body)
	case "PUT":
		return routers.PutAddress(user, body, id)
	case "DELETE":
		return routers.DeleteAddress(user, id)
	}
	return http.StatusBadRequest, "Method Invalid"
}

// ProviderActions handles provider-related actions based on the HTTP method.
// It supports retrieving, creating, updating, or deleting providers.
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
