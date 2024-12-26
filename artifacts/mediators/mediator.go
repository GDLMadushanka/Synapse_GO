package mediators

import (
	"synapse/synapsecontext"
)

type Mediator interface {
	Execute(context *synapsecontext.SynapseContext) bool
	SetFileName(fileName string)
}

// type MediatorImpl struct {
// 	FileName string
// }

// func (m *MediatorImpl) SetFileName(fileName string) {
// 	m.FileName = fileName
// }
