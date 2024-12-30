package synapsecontext

import (
	"net/http"
)

type SynapseContext struct {
	Properties map[string]string
	Response   http.ResponseWriter
	Request    *http.Request
	Message    Message
	Headers    map[string]string
}

type Message struct {
	RawPayload  []byte
	ContentType string
}
