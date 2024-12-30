package deployer

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"synapse/artifacts"
	"synapse/consolelogger"
	"synapse/dispatcher"
)

func DeployAPIs(router *dispatcher.Router, confContext *artifacts.ArtifactInfo) {

	files, err := os.ReadDir("Deploy/APIs")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		xmlFile, err := os.Open("Deploy/APIs/" + file.Name())
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

		consolelogger.InfoLog("Deploying API from " + file.Name())

		var api artifacts.API
		api.FileName = file.Name()
		if err := xml.Unmarshal(data, &api); err != nil {
			fmt.Println("Error unmarshaling XML:", err)
			continue
		}

		// check the api already deployed
		apiAlreadyDeployed := false
		_, ok := confContext.ApiMap[api.Name]
		if ok {
			consolelogger.ErrorLog("API " + api.Name + " already deployed")
			continue
		}
		for _, deployedAPI := range confContext.ApiMap {
			if deployedAPI.Context == api.Context {
				consolelogger.ErrorLog("API " + deployedAPI.Name + " already deployed with same context : " + api.Context)
				apiAlreadyDeployed = true
			}
		}
		if apiAlreadyDeployed {
			continue
		}

		// process the API
		for _, resource := range api.Resources {
			router.AddRoute(resource.Methods, api.Context+resource.URITemplate, resource.DispatchResource)
			resource.InSequence.SetFileName(api.FileName)
		}
		confContext.AddAPI(api)
		consolelogger.InfoLog("API " + api.Name + " deployed successfully")
	}
}
