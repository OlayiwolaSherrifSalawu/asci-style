package handlers

import "net/http"

func (e *AsciHandlers) clientError(w http.ResponseWriter, err int) {
	http.Error(w, http.StatusText(err), err)
}

func (e *AsciHandlers) serverError(w http.ResponseWriter, err error) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
