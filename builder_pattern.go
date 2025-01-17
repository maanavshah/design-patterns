package main

import "fmt"

type Request struct {
	requestID string
	userID    int
	keyword   string
	sessionID int
	mode      string
}

type RequestBuilder struct {
	request *Request
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{
		request: &Request{},
	}
}

func (rb *RequestBuilder) SetRequestID(requestID string) *RequestBuilder {
	rb.request.requestID = requestID
	return rb
}

func (rb *RequestBuilder) SetUserID(userID int) *RequestBuilder {
	rb.request.userID = userID
	return rb
}
func (rb *RequestBuilder) SetKeyword(keyword string) *RequestBuilder {
	rb.request.keyword = keyword
	return rb
}
func (rb *RequestBuilder) SetSessionID(sessionID int) *RequestBuilder {
	rb.request.sessionID = sessionID
	return rb
}
func (rb *RequestBuilder) SetMode(mode string) *RequestBuilder {
	rb.request.mode = mode
	return rb
}

func (rb *RequestBuilder) Build() Request {
	return *rb.request
}

func main() {
	request := NewRequestBuilder().SetRequestID("requestID").Build()
	fmt.Printf("request: %+v\n", request)

	request = NewRequestBuilder().SetMode("autosuggest").SetKeyword("query").Build()
	fmt.Printf("request: %+v\n", request)

	// request = *NewRequestBuilder().request // wont work when imported from other package
	// fmt.Printf("request: %+v\n", request)
}
