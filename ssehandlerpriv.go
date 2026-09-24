package ssehandler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
	sseconst "github.com/thomas-osgood/ssehandler/internal/constants"
)

// function designed to execute the logic for when a client
// disconnects from the SSE endpoint.
func (sh *SSEHandler) cleanupClient(clientid uuid.UUID) {
	sh.mu.Lock()
	defer sh.mu.Unlock()
	close(sh.clients[clientid])
	delete(sh.clients, clientid)
}

// function designed to generate a unique id for a client using a-zA-Z0-9.
//
// update 2026-09-04:
// removed pre-existing check for client existence because of extremely small
// potential of collisions. the check added unnecssary overhead.
func (sh *SSEHandler) generateID() (id uuid.UUID, err error) {

	id, err = uuid.NewRandom()
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

// function designed to set the headers required for successful
// server sent event initialization.
//
// this will set any custom headers specified by the user during
// the SSEHandler's initialization.
func (sh *SSEHandler) setHeaders(w http.ResponseWriter) {
	var headerName string
	var headerValue string
	var corsOrigins string = strings.Join(sh.corssettings.Origins, " ")
	var exposeHeaders string = strings.Join(sh.corssettings.ExposeHeaders, ", ")

	// set the headers necessary for the server-sent-events to work.
	w.Header().Set(sseconst.HEADER_ACALLOW_NAM, corsOrigins)
	w.Header().Set(sseconst.HEADER_ACEXPOSE_NAM, exposeHeaders)
	w.Header().Set(sseconst.HEADER_ACCELBUFFER_NAM, sseconst.HEADER_ACCELBUFFER_VAL)
	w.Header().Set(sseconst.HEADER_CONTENTTYPE_NAM, sseconst.HEADER_CONTENTTYPE_VAL)
	w.Header().Set(sseconst.HEADER_CACHE_NAM, sseconst.HEADER_CACHE_VAL)

	// set CORS headers if needed
	w.Header().Set(sseconst.HEADER_ACCREDS_NAM, strconv.FormatBool(sh.corssettings.AllowCredentials))
	if sh.corssettings.MaxAge > 0 {
		w.Header().Set(sseconst.HEADER_ACMAXAGE_NAM, strconv.Itoa(sh.corssettings.MaxAge))
	}

	if len(sh.corssettings.AllowedMethods) > 0 {
		w.Header().Set(sseconst.HEADER_ACMETHODS_NAM, strings.Join(sh.corssettings.AllowedMethods, ", "))
	}

	// set the user-defined custom headers.
	for headerName, headerValue = range sh.customHeaders {
		w.Header().Set(headerName, headerValue)
	}

}
