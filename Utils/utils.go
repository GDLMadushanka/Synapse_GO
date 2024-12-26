package utils

import (
	"encoding/json"
	"encoding/xml"
)

func IsValidJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

func IsValidXML(s string) bool {
	return xml.Unmarshal([]byte(s), new(interface{})) == nil
}
