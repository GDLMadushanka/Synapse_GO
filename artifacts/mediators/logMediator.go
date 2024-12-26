package mediators

import (
	"synapse/consolelogger"
	"synapse/synapsecontext"
)

type LogMediator struct {
	Category string `xml:"category,attr"`
	Message  string `xml:"message"`
	LineNo   int
	FileName string
}

func (l *LogMediator) Execute(context *synapsecontext.SynapseContext) bool {
	switch l.Category {
	case "DEBUG":
		consolelogger.DebugLog(l.Message)
	case "INFO":
		consolelogger.InfoLog(l.Message)
	default:
		consolelogger.MediatorErrorLog("Log", l.FileName, l.LineNo, "Invalid log category")
	}
	return true
}

func (l *LogMediator) SetFileName(fileName string) {
	l.FileName = fileName
}
