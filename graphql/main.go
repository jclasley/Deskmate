package main

import (
	logger "github.com/circleci/Deskmate/graphql/log"
)

var (
	log = logger.Log
)

func main() {
	log.Info("Deskmate - Zendesk Assistant by Tyler Conlee")

	NewRouter()
}
