package handlers

import (
	"encoding/json"
        "net/http"

	log "github.com/sirupsen/logrus"
)


type Health struct {
        Name  string  `json:"Name"`
}


func HealthCheck(w http.ResponseWriter, r *http.Request) {
        log.Infof("Healthcheck Api invoked")
        w.Header().Set("Content-Type", "application/json")

        h := Health{Name: "user"}

        if err := json.NewEncoder(w).Encode(&h); err != nil {
                log.Infof("/health Internal Error: %s", err)
                http.Error(w, err.Error(), http.StatusInternalServerError)
        }
}
