package api

import (
	"net/http"

	"github.com/omnom-nom/apiserver"
	"github.com/omnom-nom/user/handlers"
)

const (
	Apiv1 = "v1"
	ApiServiceType = "user"
)

var routes = map[string][]apiserver.Route{
	Apiv1: {
		{ Name: "Health",	Method: http.MethodGet,         Path: ApiServiceType+"/health",         Handler: handlers.HealthCheck},
		{ Name: "GetUser",	Method: http.MethodGet,		Path: ApiServiceType,			Handler: handlers.GetUser},
		{ Name: "CreateUser",	Method: http.MethodPost,	Path: ApiServiceType+"/create",		Handler: handlers.CreateUser},
	//	{ Name: "DeleteUser",	Method: http.MethodDelete,	Path: "delete",		Handler: handlers.DeleteUser},
	},
}
