package zendesk

import (
	"encoding/json"
	"net/http"
)

func ConnectHandler(w http.ResponseWriter, r *http.Request) {
	payload := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Errorw("Error decoding JSON for Zendesk connect", "error", err.Error())
		return
	}
	SetConfig()
	Connect(payload["url"].(string))
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(active)
	if err != nil {
		log.Errorw("Error marshalling JSON for config", "error", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
