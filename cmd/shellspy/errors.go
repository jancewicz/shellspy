package main

import "net/http"

func (app *application) internalSeverError(w http.ResponseWriter, r *http.Request, err error) {
	writeJSONError(w, http.StatusInternalServerError, "the server encountered problem")
}
