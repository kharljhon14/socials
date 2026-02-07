package main

import "net/http"

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := envelope{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}

	err := writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.statusInternalServerError(w, r, err)
	}
}
