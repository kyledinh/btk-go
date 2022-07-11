package moxutil

import (
	"bytes"
	"encoding/json"

	"github.com/bitly/go-simplejson"
)

// NO DEPENDENCIES
func PrettyJsonBytes(ba []byte) []byte {
	var buf bytes.Buffer
	err := json.Indent(&buf, ba, "", "    ")
	if err == nil {
		return buf.Bytes()
	}
	return ba
}

// USES go-simplejson
func SimplePrettyJson(ba []byte) []byte {
	jsonData, err := simplejson.NewJson(ba)
	if err == nil {
		prettyBytes, err := jsonData.EncodePretty()
		if err == nil {
			return prettyBytes
		}
	}
	return ba
}
