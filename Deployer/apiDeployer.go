package deployer

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"test/dispatcher"
	"test/pogo"
)

func DeployAPIs(r *dispatcher.Router) {

	files, err := os.ReadDir("Artifacts")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		xmlFile, err := os.Open("Artifacts/" + file.Name())
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

		var api pogo.API
		if err := xml.Unmarshal(data, &api); err != nil {
			fmt.Println("Error unmarshaling XML:", err)
			continue
		}

		// process the API
		for _, resource := range api.Resources {
			r.AddRoute(resource.Methods, api.Context+resource.URITemplate, resource.DispatchResource)
		}
	}
}
