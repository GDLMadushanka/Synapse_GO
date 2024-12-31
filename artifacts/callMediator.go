package artifacts

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"synapse/consolelogger"
	"synapse/synapsecontext"
	"time"
)

type CallMediator struct {
	Endpoint CallEndpoint `xml:"endpoint"`
	LineNo   int
	FileName string
}

type CallEndpoint struct {
	Key string `xml:"key,attr"`
}

var sharedClient = &http.Client{
	Transport: &http.Transport{
		MaxIdleConns:        1000,
		MaxConnsPerHost:     1000,
		MaxIdleConnsPerHost: 1000,
		IdleConnTimeout:     90 * time.Second,
	},
}

func (l *CallMediator) Execute(context *synapsecontext.SynapseContext) bool {

	ep, exists := GetArtifactInfoInstance().EndpointMap[l.Endpoint.Key]
	if !exists {
		consolelogger.ErrorLog("Endpoint not found: " + l.Endpoint.Key)
		return false
	}
	// Create a new HTTP request with the specified method, URL, and optional body
	req, err := http.NewRequest(strings.ToUpper(ep.EndpointUrl.Method), ep.EndpointUrl.URL, bytes.NewReader(context.Message.RawPayload))
	if err != nil {
		consolelogger.ErrorLog("failed to create request: " + err.Error())
		return false
	}
	req.Header.Set("Content-Type", context.Message.ContentType)

	// Send the request
	resp, err := sharedClient.Do(req)
	if err != nil {
		consolelogger.ErrorLog("failed to send request: " + err.Error())
		return false
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		consolelogger.ErrorLog("failed to read response body: " + err.Error())
		return false
	}
	context.Message.RawPayload = respBody
	context.Message.ContentType = resp.Header.Get("Content-Type")
	return true
}

func (l *CallMediator) SetFileName(fileName string) {
	l.FileName = fileName
}
