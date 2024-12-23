package mediators

import (
	"fmt"
	"test/consolelogger"
	"test/synapsecontext"
)

type PayloadMediator struct {
	Category string `xml:"category,attr"`
	Message  string `xml:"message"`
}

func (l PayloadMediator) Execute(context *synapsecontext.SynapseContext) bool {
	switch l.Category {
	case "DEBUG":
		consolelogger.DebugLog(l.Message)
	case "INFO":
		consolelogger.InfoLog(l.Message)
	}
	fmt.Printf("Logging: %s\n", l.Message)
	return true
}
