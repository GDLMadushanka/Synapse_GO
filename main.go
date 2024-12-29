package main

import (
	"net/http"
	"synapse/configurationcontext"
	"synapse/consolelogger"
	"synapse/deployer"
	"synapse/dispatcher"
	"time"
)

func main() {
	start := time.Now()
	consolelogger.PrintWelcomeMessage()
	router := dispatcher.Router{}

	// Configuration context to hold all artifacts
	confContext := configurationcontext.ConfigurationContext{}
	deployer.DeployEndpoints(&router, &confContext)
	deployer.DeployAPIs(&router, &confContext)
	elapsed := time.Since(start)
	consolelogger.InfoLog("Server started in " + elapsed.String())
	err := http.ListenAndServe(":8080", &router)
	if err != nil {
		consolelogger.ErrorLog("Error starting server: " + err.Error())
	}
}
