package main

import (
	"errors"
	"fmt"
)

type Request struct {
	requestID string
	userID    int
	keyword   string
	sessionID int
	mode      string
}

type RequestBuilder struct {
	requestID string
	userID    int
	keyword   string
	sessionID int
	mode      string
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{}
}

func (rb *RequestBuilder) SetRequestID(requestID string) *RequestBuilder {
	rb.requestID = requestID
	return rb
}

func (rb *RequestBuilder) SetUserID(userID int) *RequestBuilder {
	rb.userID = userID
	return rb
}

func (rb *RequestBuilder) SetKeyword(keyword string) *RequestBuilder {
	rb.keyword = keyword
	return rb
}

func (rb *RequestBuilder) SetSessionID(sessionID int) *RequestBuilder {
	rb.sessionID = sessionID
	return rb
}

func (rb *RequestBuilder) SetMode(mode string) *RequestBuilder {
	rb.mode = mode
	return rb
}

// Build creates a new Request instance each time it's called
func (rb *RequestBuilder) Build() Request {
	return Request{
		requestID: rb.requestID,
		userID:    rb.userID,
		keyword:   rb.keyword,
		sessionID: rb.sessionID,
		mode:      rb.mode,
	}
}

// BuildWithValidation creates a Request with validation
func (rb *RequestBuilder) BuildWithValidation() (Request, error) {
	if rb.requestID == "" {
		return Request{}, errors.New("requestID is required")
	}
	return rb.Build(), nil
}

// Reset clears the builder for reuse
func (rb *RequestBuilder) Reset() *RequestBuilder {
	rb.requestID = ""
	rb.userID = 0
	rb.keyword = ""
	rb.sessionID = 0
	rb.mode = ""
	return rb
}

func main() {
	// Basic usage
	request := NewRequestBuilder().SetRequestID("request1").Build()
	fmt.Printf("request: %+v\n", request)

	request = NewRequestBuilder().SetMode("autosuggest").SetKeyword("query").Build()
	fmt.Printf("request: %+v\n", request)

	// Demonstrating the fix - builder reuse now works correctly
	builder := NewRequestBuilder()
	req1 := builder.SetRequestID("request1").SetUserID(100).Build()
	req2 := builder.SetRequestID("request2").SetUserID(200).Build()

	fmt.Printf("req1: %+v\n", req1) // Will have requestID="request1", userID=100
	fmt.Printf("req2: %+v\n", req2) // Will have requestID="request2", userID=200

	// Using validation
	req3, err := NewRequestBuilder().SetMode("search").BuildWithValidation()
	if err != nil {
		fmt.Printf("Validation error: %v\n", err)
	} else {
		fmt.Printf("req3: %+v\n", req3)
	}

	// Using reset for builder reuse
	builder.Reset().SetRequestID("request3").SetMode("search")
	req4 := builder.Build()
	fmt.Printf("req4: %+v\n", req4)
}
