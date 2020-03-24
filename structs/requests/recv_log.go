package requests

import (
	"encoding/json"
	"github.com/panwenbin/ghttplog/databases/entities"
	"net/http"
	"regexp"
	"strings"
)

type RecvLog struct {
	Method         string      `json:"method"`
	Uri            string      `json:"uri"`
	RequestHeader  http.Header `json:"request_header"`
	RequestBody    string      `json:"request_body"`
	ResponseHeader http.Header `json:"response_header"`
	ResponseBody   string      `json:"response_body"`
}

func (r RecvLog) ToHttp() entities.Http {
	h := entities.Http{
		Uri:            r.Uri,
		RequestMethod:  r.Method,
		RequestHeader:  r.RequestHeader,
		RequestBody:    r.RequestBody,
		ResponseHeader: r.ResponseHeader,
	}

	rBody := r.ResponseBody
	contentType := r.ResponseHeader.Get("Content-Type")
	if strings.Contains(contentType, "json") {
		regx := regexp.MustCompile(`^(?s)\s*(\w+)\((.*)\)`)
		matches := regx.FindStringSubmatch(r.ResponseBody)
		if len(matches) == 3 {
			h.Jsonp = matches[1]
			rBody = matches[2]
		}

		var jsonObj interface{}
		err := json.Unmarshal([]byte(rBody), &jsonObj)
		if err != nil {
			h.ResponseBody = rBody
		} else {
			h.ResponseBody = jsonObj
		}
	} else {
		h.ResponseBody = rBody
	}

	return h
}
