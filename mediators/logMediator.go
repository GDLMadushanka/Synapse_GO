package mediators

import (
	"test/consolelogger"
	"test/synapsecontext"
)

type LogMediator struct {
	Category string `xml:"category,attr"`
	Message  string `xml:"message"`
}

func (l LogMediator) Execute(context *synapsecontext.SynapseContext) bool {
	switch l.Category {
	case "DEBUG":
		consolelogger.DebugLog(l.Message)
	case "INFO":
		consolelogger.InfoLog(l.Message)
	}
	return true
}
