package ssehandler

import "sync"

type SSEHandler struct {
	clients       SSEChannelMap
	corssettings  CorsOptions
	customHeaders HeaderMap
	mu            sync.Mutex
}

type SSEHandlerOption struct {
	// map of channels that will be used to transmit
	// information to the SSE endpoint.
	Clients SSEChannelMap
	// CORS options for the SSEHandler.
	CorsSettings CorsOptions
	// user-defined custom headers that will be set by
	// the sever when an SSE connection gets established.
	CustomHeaders HeaderMap
}

type SSEMessage struct {
	// name of the event that will be transmitted.
	Event string `json:"event" xml:"event"`
	// id assigned to the event type. (optional)
	Id string `json:"id" xml:"id"`
	// data to be transmitted to the client.
	//
	// for best results, this should be a JSON object.
	Data string `json:"data" xml:"data"`
}

/*
struct defining the various CORS options that can be
applied/used by the SSEHandler.

this is a structure internal to the SSEHandler package
and is meant to be set during the New() call.
*/
type CorsOptions struct {
	// tells browsers whether the server allows credentials to be included
	// in cross-origin HTTP requests.
	AllowCredentials bool
	// indicate the HTTP headers that can be used during the actual request.
	// This is required if the preflight request contains Access-Control-Request-Headers.
	AllowHeaders []string
	// specifies one or more HTTP request methods allowed
	// when accessing a resource in response to a preflight
	// request.
	AllowedMethods []string
	// list of which response headers should be made
	// available to scripts running in the browser in
	// response to a cross-origin request.
	ExposeHeaders []string
	// indicates how long the results of a preflight
	// request (that is, the information contained in the
	// Access-Control-Allow-Methods and Access-Control-Allow-Headers
	// headers) can be cached.
	MaxAge int
	// list of allowed origins for CORS
	Origins []string
}
