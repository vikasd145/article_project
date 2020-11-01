package Rotuers

import (
	"net/http"

	"github.com/vikasd145/article_project/api"

	"github.com/gorilla/context"

	"github.com/gorilla/mux"
)

//Route Struct keep the info of one path
type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

//Routes ...
type Routes []Route

//NewRouter Return new Router
func NewRouter(root string) *mux.Router {
	router := mux.NewRouter()
	for _, route := range HandlerRoute {
		router.Methods(route.Method).Path(root + route.Pattern).Name(route.Name).Handler(route.HandleFunc)
	}
	router.Use(context.ClearHandler)
	return router
}

var HandlerRoute = Routes{
	Route{"Create a article", "POST", "/article", api.CreateArticle},
	Route{"Get article by id", "GET", "/articles/", api.GetArticle},
}
