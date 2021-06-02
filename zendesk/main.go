package main

import (
	logger "github.com/tylerconlee/Deskmate/zendesk/log"
)

var (
	log = logger.Log
)

func main() {
	log.Info("Deskmate - Zendesk Assistant by Tyler Conlee")

	NewRouter()
}
