package api

import (
	"fmt"
	"net/http"

	"github.com/omnom-nom/apiserver"
	"github.com/omnom-nom/user/handlers"
)

const (
	Apiv1 = "v1"
	ApiServiceType = "user"
)

var v1Prefix = fmt.Sprintf("%s/%s", Apiv1, ApiServiceType)
var routes = map[string][]apiserver.Route{
	v1Prefix: {
		{ Name: "Health",	Method: http.MethodGet,         Path: "health",         Handler: handlers.HealthCheck},
		{ Name: "CreateUser",	Method: http.MethodPost,	Path: "create",		Handler: handlers.CreateUser},
	//	{ Name: "DeleteUser",	Method: http.MethodDelete,	Path: "{email}",	Handler: DeleteUser},
	//	{ Name: "GetUser",	Method: http.MethodGet,		Path: "{email}",	Handler: GetUser},
	},
}
