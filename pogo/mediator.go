package pogo

import (
	"test/synapsecontext"
)

type Mediator interface {
	Execute(context *synapsecontext.SynapseContext) bool
}
