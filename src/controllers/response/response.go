package response

import (
	"encoding/json"
	"net/http"
)

//Do create and return a proper response to the client
func Do(w http.ResponseWriter, body interface{}, statusCode int) {
	res, _ := json.Marshal(body)
	w.WriteHeader(statusCode)
	w.Write(res)
}
