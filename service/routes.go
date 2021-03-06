// Package service implements the gohstd web service, providing remote access
// to the command history supplied by gohst clients
package service

import (
	"github.com/gorilla/mux"
	"github.com/kardianos/osext"
	gohst "github.com/warreq/gohstd/common"
	"log"
	"net/http"
	"path"
	"path/filepath"
)

// Route represents a URI route for the server to support. A definition of a
// Route declares support for an HTTP endpoint and maps that URI to the
// appropriate HTTP handler.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var userRepo gohst.UserRepo
var commandRepo gohst.CommandRepo

// NewRouter constructs a *mux.Router based on the routes defined in this
// package, which can then be passed to the net/http server.
func NewRouter(cmd gohst.CommandRepo, user gohst.UserRepo) *mux.Router {
	userRepo = user
	commandRepo = cmd
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler

		// Note that handler func wrapping like this means that the last one applied
		// is the outer-most wrapper and executes first
		handler = route.HandlerFunc
		handler = StandardHeader(handler)
		handler = Auth(handler, route.Name)
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	// Serve the web application on the root path
	execPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}

	appPath := filepath.FromSlash(path.Join(execPath, "webapp"))

	router.PathPrefix("/").Handler(http.FileServer(http.Dir(appPath)))
	return router
}

var routes = Routes{
	Route{
		"UserRegister",
		"POST",
		"/api/users/register",
		UserRegister,
	},
	Route{
		"UserLogin",
		"GET",
		"/api/users/login",
		UserLogin,
	},
	Route{
		"UserShow",
		"GET",
		"/api/users/{username}",
		UserShow,
	},
	Route{
		"CommandIndex",
		"GET",
		"/api/users/{username}/commands",
		CommandIndex,
	},
	Route{
		"CommandCreate",
		"POST",
		"/api/users/{username}/commands",
		CommandCreate,
	},
	Route{
		"CommandTagCreate",
		"POST",
		"/api/users/{username}/commands/{commandId}/tags",
		CommandTagCreate,
	},
	Route{
		"UserTagShow",
		"GET",
		"/api/users/{username}/tags",
		UserTagShow,
	},
	Route{
		"UserTagRename",
		"PUT",
		"/api/users/{username}/tags/{tag}",
		UserTagRename,
	},
	Route{
		"UserTagDelete",
		"DELETE",
		"/api/users/{username}/tags/{tag}",
		UserTagDelete,
	},
	Route{
		"CommandTagDelete",
		"DELETE",
		"/api/users/{username}/commands/{commandId/tags/{tag}",
		CommandTagDelete,
	},
	Route{
		"CommandDelete",
		"DELETE",
		"/api/users/{username}/commands/{commandId}",
		CommandDelete,
	},
}
