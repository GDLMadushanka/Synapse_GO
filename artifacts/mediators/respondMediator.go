package mediators

import (
	"fmt"
	"synapse/synapsecontext"
)

type RespondMediator struct {
	LineNo   int
	FileName string
}

func (l *RespondMediator) Execute(context *synapsecontext.SynapseContext) bool {
	// Update content type header
	context.Response.Header().Set("Content-Type", context.Headers["Content-Type"])
	fmt.Fprintf(context.Response, string(context.Message.RawPayload))
	return true
}

func (l *RespondMediator) SetFileName(fileName string) {
	l.FileName = fileName
}
