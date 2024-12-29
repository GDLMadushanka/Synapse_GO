package deployer

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"synapse/artifacts"
	"synapse/configurationcontext"
	"synapse/consolelogger"
	"synapse/dispatcher"
)

func DeployEndpoints(router *dispatcher.Router, confContext *configurationcontext.ConfigurationContext) {

	files, err := os.ReadDir("Deploy/Endpoints")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		xmlFile, err := os.Open("Deploy/Endpoints/" + file.Name())
		if err != nil {
			fmt.Println("Error opening file:", err)
			continue
		}
		defer xmlFile.Close()

		data, err := io.ReadAll(xmlFile)
		if err != nil {
			fmt.Println("Error reading file:", err)
			continue
		}

		consolelogger.InfoLog("Deploying endpoint from " + file.Name())

		var endpoint artifacts.Endpoint
		endpoint.FileName = file.Name()
		if err := xml.Unmarshal(data, &endpoint); err != nil {
			fmt.Println("Error unmarshaling XML:", err)
			continue
		}

		_, ok := confContext.ApiMap[endpoint.Name]
		if ok {
			consolelogger.ErrorLog("Endpoint " + endpoint.Name + " already deployed")
			continue
		}

		confContext.AddEndpoint(endpoint)
		consolelogger.InfoLog("Endpoint " + endpoint.Name + " deployed successfully")
	}
}
