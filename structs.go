package ssehandler

type SSEHandler struct {
	clients       SSEChannelMap
	customHeaders HeaderMap
}

type SSEHandlerOption struct {
	// map of channels that will be used to transmit
	// information to the SSE endpoint.
	Clients SSEChannelMap
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
