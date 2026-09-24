package ssehandler

import (
	"fmt"
	"net/http"
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
		Clients: nil,
		CorsSettings: CorsOptions{
			MaxAge: sseconst.DEFAULT_MAXAGE,
		},
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

	// clean all CORS string slices
	defaults.CorsSettings.cleanSlices()

	// if no CORS Origins have been specified, default to "*"
	if len(defaults.CorsSettings.Origins) < 1 {
		defaults.CorsSettings.Origins = []string{sseconst.HEADER_ACALLOW_VAL}
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

// signal to the browswer that credentials are allowed in
// a CORS request.
//
// Credentials include cookies, Transport Layer Security (TLS) client certificates,
// or authentication headers containing a username and password. By default, these
// credentials are not sent in cross-origin requests, and doing so can make a site
// vulnerable to Cross-Site Request Forgery (CSRF) attacks.
func WithCORSAllowCreds() SSEHandlerOptFunc {
	return func(so *SSEHandlerOption) error {
		so.CorsSettings.AllowCredentials = true
		return nil
	}
}

// add an allowed header to the list of allowed CORS expose headers.
func WithCORSExposeHeader(header string) SSEHandlerOptFunc {
	return func(so *SSEHandlerOption) error {
		header = strings.TrimSpace(header)
		if len(header) < 1 {
			return fmt.Errorf(ssemsg.ERR_EMPTY_HEADER)
		}
		so.CorsSettings.ExposeHeaders = append(so.CorsSettings.ExposeHeaders, http.CanonicalHeaderKey(header))
		return nil
	}
}

// set the allowed expose headers to the string slice passed in.
func WithCORSExposeHeaders(headers []string) SSEHandlerOptFunc {
	return func(so *SSEHandlerOption) error {
		so.CorsSettings.ExposeHeaders = slices.Clone(headers)
		return nil
	}
}

// add an allowed method to the list of allowed CORS methods.
func WithCORSMethod(method string) SSEHandlerOptFunc {
	return func(so *SSEHandlerOption) error {
		method = strings.TrimSpace(method)
		if len(method) < 1 {
			return fmt.Errorf(ssemsg.ERR_EMPTY_METHOD)
		}
		so.CorsSettings.AllowedMethods = append(so.CorsSettings.AllowedMethods, strings.ToUpper(method))
		return nil
	}
}

// set the allowed methods to the string slice passed in.
func WithCORSMethods(methods []string) SSEHandlerOptFunc {
	return func(so *SSEHandlerOption) error {
		so.CorsSettings.AllowedMethods = slices.Clone(methods)
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
		so.CorsSettings.Origins = slices.Clone(origins)
		return nil
	}
}

// set any custom headers that are desired when setting up the SSE endpoint.
// these will be set by the server upon client connection.
func WithCustomHeaders(customHeaders HeaderMap) SSEHandlerOptFunc {
	return func(so *SSEHandlerOption) error {
		if customHeaders == nil {
			return fmt.Errorf(ssemsg.ERR_EMPTY_MAP)
		}
		so.CustomHeaders = customHeaders
		return nil
	}
}

// set the max age for the preflight cache.
func WithMaxAge(maxAge int) SSEHandlerOptFunc {
	return func(so *SSEHandlerOption) error {
		if maxAge < 1 {
			return fmt.Errorf(ssemsg.ERR_MAXAGE_NOTPOS)
		}

		so.CorsSettings.MaxAge = maxAge
		return nil
	}
}
