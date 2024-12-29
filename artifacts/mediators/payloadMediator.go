package mediators

import (
	"encoding/xml"
	"strings"
	"synapse/consolelogger"
	"synapse/synapsecontext"
	"synapse/utils"
)

type PayloadMediator struct {
	Description string      `xml:"description,attr"`
	MediaType   string      `xml:"media-type,attr"`
	Format      FormatBlock `xml:"format"`
	LineNo      int
	FileName    string
}

// FormatBlock captures all content inside <format>...</format> as a raw string
type FormatBlock struct {
	XMLName  xml.Name `xml:"format"`
	InnerXML string   `xml:",innerxml"`
}

func (l *PayloadMediator) Execute(context *synapsecontext.SynapseContext) bool {
	if l.Format.InnerXML == "" {
		consolelogger.ErrorLog("Error occurred while creating the payload in payload mediator. Empty payload")
		return false
	}
	payload := strings.TrimSpace(l.Format.InnerXML)
	// check Format is a valid JSON
	if l.MediaType == "json" {
		if !utils.IsValidJSON(payload) {
			consolelogger.ErrorLog("Error occurred while creating the payload in payload mediator. Invalid JSON format")
			return false
		} else {
			msg := synapsecontext.Message{RawPayload: []byte(payload), ContentType: "application/json"}
			context.Message = msg
			context.Headers["Content-Type"] = "application/json"
		}
	} else if l.MediaType == "xml" {
		if !utils.IsValidXML(payload) {
			consolelogger.ErrorLog("Error occurred while creating the payload in payload mediator. Invalid XML format")
			return false
		} else {
			msg := synapsecontext.Message{RawPayload: []byte(payload), ContentType: "application/xml"}
			context.Message = msg
			context.Headers["Content-Type"] = "application/xml"
		}
	} else {
		consolelogger.ErrorLog("Error occurred while creating the payload in payload mediator. Invalid media type")
		return false
	}
	return true
}

func (l *PayloadMediator) SetFileName(fileName string) {
	l.FileName = fileName
}
