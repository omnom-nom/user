package handlers

import(
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Error	string	`json:"error"`
}

func responseToJSON(httpResponseCode int, response interface{}) (int, []byte) {
	jsonResp, err := json.Marshal(response)
	if err != nil {
		log.Errorf("response data conversion to json failed: %s", err)

		return http.StatusInternalServerError, []byte{}
	}

	return httpResponseCode, jsonResp
}

// log error and write http response
func writeHTTPError(w http.ResponseWriter, httpResponseCode int, err error) {

	httpResponseCode, jsonResp := responseToJSON(httpResponseCode, ErrorResponse{Error: err.Error()})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpResponseCode)
	w.Write(jsonResp)
}

// write http response and indicate success
func writeHTTPSuccess(w http.ResponseWriter, httpResponseCode int, response interface{}) {

	httpResponseCode, jsonResp := responseToJSON(httpResponseCode, response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpResponseCode)
	w.Write(jsonResp)
}
