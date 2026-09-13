package ssehandler

import (
	"fmt"
	"slices"
	"strings"

	sseconst "github.com/thomas-osgood/ssehandler/internal/constants"
	ssemsg "github.com/thomas-osgood/ssehandler/internal/messages"
)

// function designed to create, initialize and return an instance
// of an SSEHandler object.
func NewSSEHandler(opts ...SSEHandlerOptFunc) (ssehandle *SSEHandler, err error) {
	var curopt SSEHandlerOptFunc
	var defaults SSEHandlerOption = SSEHandlerOption{
		Clients:       nil,
		CorsSettings:  CorsOptions{},
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

	// if no CORS Origins have been specified, default to "*"
	defaults.CorsSettings.cleanOrigins()
	if len(defaults.CorsSettings.Origins) < 1 {
		defaults.CorsSettings.Origins = slices.Insert(defaults.CorsSettings.Origins, 0, sseconst.HEADER_ACALLOW_VAL)
	}

	// assign the user-specified values to the SSEHandler to return.
	ssehandle = new(SSEHandler)
	ssehandle.clients = defaults.Clients
	ssehandle.corssettings = defaults.CorsSettings
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

// add an allowed origin to the list of CORS Origins.
func WithCORSOrigin(origin string) SSEHandlerOptFunc {
	return func(so *SSEHandlerOption) error {
		origin = strings.TrimSpace(origin)
		if len(origin) < 1 {
			return fmt.Errorf(ssemsg.ERR_EMPTY_ORIGIN)
		}
		so.CorsSettings.Origins = append(so.CorsSettings.Origins, origin)
		return nil
	}
}

// set the allowed origins to the string slice passed in.
func WithCORSOrigins(origins []string) SSEHandlerOptFunc {
	return func(so *SSEHandlerOption) error {
		so.CorsSettings.Origins = origins
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
