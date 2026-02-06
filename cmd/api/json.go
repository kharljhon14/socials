package main

import (
	"encoding/json"
	"maps"
	"net/http"
)

type envelope map[string]any

func writeJSON(
	w http.ResponseWriter,
	status int,
	data envelope,
	headers http.Header,
) error {

	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	maps.Insert(w.Header(), maps.All(headers))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	w.Write(js)
	return nil
}

func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func writeJSONError(w http.ResponseWriter, status int, message string) error {

	return writeJSON(w, status, envelope{"error": message}, nil)
}
