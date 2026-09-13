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
	Origins []string
}
