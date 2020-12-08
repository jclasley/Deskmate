package config

type Config struct {
	SlackAPI string
}

func LoadConfig() {
	config := Config{SlackAPI: "test"}
	js, err := json.Marshal(config)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
