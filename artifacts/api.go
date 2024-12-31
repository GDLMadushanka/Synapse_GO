package artifacts

import (
	"net/http"
	"synapse/synapsecontext"
)

type Resource struct {
	Methods       string   `xml:"methods,attr"`
	URITemplate   string   `xml:"uri-template,attr"`
	InSequence    Sequence `xml:"inSequence"`
	FaultSequence Sequence `xml:"faultSequence"`
}

type API struct {
	Context   string     `xml:"context,attr"`
	Name      string     `xml:"name,attr"`
	Resources []Resource `xml:"resource"`
	FileName  string
}

func (resource *Resource) DispatchResource(w http.ResponseWriter, request *http.Request) {
	// Read transport headers
	var headers = make(map[string]string)
	for name, values := range request.Header {
		headers[name] = values[0]
	}

	// Create the mssage context
	var context = synapsecontext.SynapseContext{
		Request:    request,
		Response:   w,
		Properties: make(map[string]string),
		Message:    synapsecontext.Message{},
		Headers:    headers,
	}
	resource.InSequence.Execute(&context)
}
