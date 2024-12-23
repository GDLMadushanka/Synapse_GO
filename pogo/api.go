package pogo

import (
	"net/http"
	"test/synapsecontext"
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
}

func (resource *Resource) DispatchResource(w http.ResponseWriter, r *http.Request) {
	// creating the message context
	var context = synapsecontext.SynapseContext{
		Request:    r,
		Response:   w,
		Properties: make(map[string]string),
	}
	resource.InSequence.Execute(&context)
}
