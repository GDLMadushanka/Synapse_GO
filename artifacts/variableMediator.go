package artifacts

import (
	"synapse/synapsecontext"
)

type VariableMediator struct {
	Name     string `xml:"name,attr"`
	Type     string `xml:"type,attr"`
	Value    string `xml:"value,attr"`
	LineNo   int
	FileName string
}

func (l *VariableMediator) Execute(context *synapsecontext.SynapseContext) bool {
	context.Properties[l.Name] = l.Value
	return true
}

func (l *VariableMediator) SetFileName(fileName string) {
	l.FileName = fileName
}
