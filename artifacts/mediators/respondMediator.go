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
	fmt.Fprintf(context.Response, context.Message)
	return true
}

func (l *RespondMediator) SetFileName(fileName string) {
	l.FileName = fileName
}
