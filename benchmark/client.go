package benchmark

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type TopicRequest struct {
	Text string `json:"text"`
}

func postRequest(url string, topicData string) {

	// Prepare request data
	reqObj := &TopicRequest{
		Text: topicData,
	}
	reqData, _ := json.Marshal(reqObj)

	req := fasthttp.AcquireRequest()
	req.SetBody(reqData)
	req.Header.SetMethodBytes([]byte("POST"))
	req.SetRequestURIBytes([]byte(url))
	res := fasthttp.AcquireResponse()

	// Start request/response
	if err := fasthttp.Do(req, res); err != nil {
		panic("handle error")
	}

	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)
}
