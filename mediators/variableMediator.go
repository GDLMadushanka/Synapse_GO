package mediators

import (
	"test/synapsecontext"
)

type VariableMediator struct {
	Name  string `xml:"name,attr"`
	Type  string `xml:"type,attr"`
	Value string `xml:"value,attr"`
}

func (l VariableMediator) Execute(context *synapsecontext.SynapseContext) bool {
	context.Properties[l.Name] = l.Value
	return true
}
