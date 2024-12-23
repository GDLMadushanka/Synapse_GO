package mediators

import (
	"fmt"
	"test/synapsecontext"
)

type RespondMediator struct {
}

func (l RespondMediator) Execute(context *synapsecontext.SynapseContext) bool {
	fmt.Fprintf(context.Response, "Hello, from Synapse!")
	return true
}
