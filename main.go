package main

import (
	"net/http"
	"synapse/consolelogger"
	"synapse/deployer"
	"synapse/dispatcher"
	"time"
)

func main() {
	start := time.Now()
	consolelogger.PrintWelcomeMessage()
	r := dispatcher.Router{}
	deployer.DeployAPIs(&r)
	elapsed := time.Since(start)
	consolelogger.InfoLog("Server started in " + elapsed.String())
	err := http.ListenAndServe(":8080", &r)
	if err != nil {
		consolelogger.ErrorLog("Error starting server: " + err.Error())
	}
}
