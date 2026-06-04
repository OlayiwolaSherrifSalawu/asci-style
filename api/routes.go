package api

import "net/http"

func (a *Application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /createUser", a.UserHandler.CreateUser)
	return mux
}
