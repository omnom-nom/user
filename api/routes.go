package api

import (
	"fmt"
	"net/http"

	"github.com/omnom-nom/apiserver"
)

const (
	Apiv1 = "v1"
	ApiServiceType = "user"
)

var v1Prefix = fmt.Sprintf("%s/%s", Apiv1, ApiServiceType)
var routes = map[string][]apiserver.Route{
	v1Prefix: {
		{ Name: "HealthCheck",  Method: http.MethodGet,         Path: "healthcheck",            Handler: HealthCheck},
		{ Name: "CreateUser",	Method: http.MethodPost,	Path: "create",		Handler: CreateUser},
	//	{ Name: "DeleteUser",	Method: http.MethodDelete,	Path: "{email}",	Handler: DeleteUser},
	//	{ Name: "GetUser",	Method: http.MethodGet,		Path: "{email}",	Handler: GetUser},
	},
}
