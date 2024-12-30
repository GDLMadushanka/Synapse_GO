package artifacts

import (
	"synapse/synapsecontext"
)

type Mediator interface {
	Execute(context *synapsecontext.SynapseContext) bool
	SetFileName(fileName string)
}
