package artifacts

type Endpoint struct {
	Name        string      `xml:"name,attr"`
	EndpointUrl EndpointUrl `xml:"http"`
	FileName    string
}

type EndpointUrl struct {
	Method string `xml:"method,attr"`
	URL    string `xml:"uri-template,attr"`
}
