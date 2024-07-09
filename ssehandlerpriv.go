package ssehandler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	sseconst "github.com/thomas-osgood/ssehandler/internal/constants"
	ssemsg "github.com/thomas-osgood/ssehandler/internal/messages"
)

// function designed to execute the logic for when a client
// disconnects from the SSE endpoint.
func (sh *SSEHandler) cleanupClient(ctx context.Context, clientid uuid.UUID) {
	close(sh.clients[clientid])
	delete(sh.clients, clientid)
	ctx.Done()
}

// function designed to check whether a client with the given
// id already exists in the clients map.
func (sh *SSEHandler) clientIdExists(id uuid.UUID) (exists bool) {
	_, exists = sh.clients[id]
	return exists
}

// function designed to generate a unique id for a client using a-zA-Z0-9.
func (sh *SSEHandler) generateID(attempt int) (id uuid.UUID, err error) {

	id, err = uuid.NewUUID()
	if err != nil {
		return uuid.Nil, err
	}

	// if the id already exists in the map, attempt to
	// generate another id.
	//
	// if the maximum number of attempts has been reached
	// return an error.
	//
	// if the maximum number of attempts has not yet been
	// reached, the attempt number will be incremented and
	// this function will be called recursively.
	if sh.clientIdExists(id) {
		if attempt >= sseconst.GENERATE_ATTEMPT_MAX {
			return uuid.Nil, fmt.Errorf(ssemsg.ERR_GENERATEID_MAXATTEMPTS)
		}

		attempt++
		return sh.generateID(attempt)
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

	// set the headers necessary for the server-sent-events to work.
	w.Header().Set(sseconst.HEADER_ACALLOW_NAM, sseconst.HEADER_ACALLOW_VAL)
	w.Header().Set(sseconst.HEADER_ACEXPOSE_NAM, sseconst.HEADER_ACEXPOSE_VAL)
	w.Header().Set(sseconst.HEADER_ACCELBUFFER_NAM, sseconst.HEADER_ACCELBUFFER_VAL)
	w.Header().Set(sseconst.HEADER_CONTENTTYPE_NAM, sseconst.HEADER_CONTENTTYPE_VAL)
	w.Header().Set(sseconst.HEADER_CACHE_NAM, sseconst.HEADER_CACHE_VAL)

	// set the user-defined custom headers.
	for headerName, headerValue = range sh.customHeaders {
		w.Header().Set(headerName, headerValue)
	}

}
