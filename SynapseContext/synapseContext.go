package synapsecontext

import (
	"net/http"
)

type SynapseContext struct {
	Properties map[string]string
	Response   http.ResponseWriter
	Request    *http.Request
	Message    string
	Headers    map[string]string
}
