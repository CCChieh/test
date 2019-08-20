package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

func NewRouter() *mux.Router {
	return mux.NewRouter()
}
