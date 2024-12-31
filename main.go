package main

import (
	"net/http"
	"runtime"
	"synapse/artifacts"
	"synapse/consolelogger"
	"synapse/deployer"
	"synapse/dispatcher"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	start := time.Now()
	consolelogger.PrintWelcomeMessage()
	router := dispatcher.Router{}

	// Configuration context to hold all artifacts
	artifactInfo := artifacts.GetArtifactInfoInstance()
	deployer.DeployEndpoints(&router, artifactInfo)
	deployer.DeployAPIs(&router, artifactInfo)
	elapsed := time.Since(start)
	consolelogger.InfoLog("Server started in " + elapsed.String())
	err := http.ListenAndServe(":8080", &router)
	if err != nil {
		consolelogger.ErrorLog("Error starting server: " + err.Error())
	}
}
