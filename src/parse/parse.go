package parse

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func UnmarshalJSON(r *http.Request, i interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), i); err != nil {
			return
		}
	}
}
