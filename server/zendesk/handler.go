package zendesk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ConnectHandler(w http.ResponseWriter, r *http.Request) {
	payload := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		fmt.Println(err.Error())
	}
	SetConfig()
	Connect(payload["url"].(string))
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(active)
	if err != nil {
		fmt.Println("Error marshalling JSON for config")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
