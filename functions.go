package ssehandler

import (
	"fmt"

	ssemsg "github.com/thomas-osgood/ssehandler/internal/messages"
)

// function designed to create, initialize and return an instance
// of an SSEHandler object.
func NewSSEHandler(opts ...SSEHandlerOptFunc) (ssehandle *SSEHandler, err error) {
	var curopt SSEHandlerOptFunc
	var defaults SSEHandlerOption = SSEHandlerOption{
		Clients:       nil,
		CustomHeaders: make(HeaderMap),
	}

	// assign the user-specified options.
	for _, curopt = range opts {
		err = curopt(&defaults)
		if err != nil {
			return nil, err
		}
	}

	// make sure the user passed in an SSEChannelMap to use to keep
	// track of the clients connecting into the SSE endpoint.
	if defaults.Clients == nil {
		return nil, fmt.Errorf(ssemsg.ERR_EMPTY_MAP)
	}

	// assign the user-specified values to the SSEHandler to return.
	ssehandle = new(SSEHandler)
	ssehandle.clients = defaults.Clients
	ssehandle.customHeaders = defaults.CustomHeaders

	return ssehandle, nil
}

// set the map to use to keep track of the clients that are connecting
// into the SSE endpoint.
func WithClientMap(clients SSEChannelMap) SSEHandlerOptFunc {
	return func(so *SSEHandlerOption) error {
		if clients == nil {
			return fmt.Errorf(ssemsg.ERR_EMPTY_MAP)
		}
		so.Clients = clients
		return nil
	}
}

// set any custom headers that are desired when setting up the SSE endpoint.
// these will be set by the server upon client connection.
func WithCustomHeaders(customHaders HeaderMap) SSEHandlerOptFunc {
	return func(so *SSEHandlerOption) error {
		if customHaders == nil {
			return fmt.Errorf(ssemsg.ERR_EMPTY_MAP)
		}
		so.CustomHeaders = customHaders
		return nil
	}
}
