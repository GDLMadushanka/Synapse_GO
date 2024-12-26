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

func DeployAPIs(r *dispatcher.Router) {

	files, err := os.ReadDir("Deploy")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		xmlFile, err := os.Open("Deploy/" + file.Name())
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

		// process the API
		for _, resource := range api.Resources {
			r.AddRoute(resource.Methods, api.Context+resource.URITemplate, resource.DispatchResource)
			resource.InSequence.SetFileName(api.FileName)
		}

		consolelogger.InfoLog("API " + api.Name + " deployed successfully")
	}
}
