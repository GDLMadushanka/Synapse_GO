package main

import (
	"net/http"
	"test/consolelogger"
	"test/deployer"
	"test/dispatcher"
	"time"
)

func main() {
	start := time.Now()
	consolelogger.PrintWelcomeMessage()
	r := dispatcher.Router{}
	deployer.DeployAPIs(&r)
	elapsed := time.Since(start)
	consolelogger.InfoLog("Server started in " + elapsed.String())
	http.ListenAndServe(":8080", &r)
}
