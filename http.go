package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendJSON(w http.ResponseWriter, jsonObject interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// Marshalling might fail, in which case we should return a 500 with the
	// actual error.
	bs, err := json.Marshal(jsonObject)
	if err != nil {
		SendJSONError(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	fmt.Fprintf(w, "%s\n", bs)
	return nil
}

func RecvJSON(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var data map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		SendJSONError(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	return data, nil
}

func SendJSONError(w http.ResponseWriter, err string, code int) {
	bs, _ := json.Marshal(map[string]string{"error": err})
	http.Error(w, string(bs), code)
}
