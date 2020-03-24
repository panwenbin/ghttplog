package entities

import "net/http"

type Http struct {
	Uri                 string      `bson:"uri"`
	RequestMethod       string      `bson:"request_method"`
	RequestHeader       http.Header `bson:"request_header"`
	RequestBody         string      `bson:"request_body"`
	ResponseHeader      http.Header `bson:"response_header"`
	ResponseBody        interface{} `bson:"response_body"`
	Jsonp               string      `bson:"jsonp"`
}
