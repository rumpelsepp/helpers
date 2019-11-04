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
		return err
	}
	fmt.Fprintf(w, "%s\n", bs)
	return nil
}

func RecvJSON(r *http.Request, data interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		return err
	}
	return nil
}

func SendJSONError(w http.ResponseWriter, err string, code int) {
	bs, _ := json.Marshal(map[string]string{"error": err})
	http.Error(w, string(bs), code)
}
