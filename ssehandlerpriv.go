package ssehandler

import (
	"net/http"

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

// function designed to check whether a client with the given
// id already exists in the clients map.
func (sh *SSEHandler) clientIdExists(id uuid.UUID) (exists bool) {
	sh.mu.Lock()
	defer sh.mu.Unlock()
	_, exists = sh.clients[id]
	return exists
}

// function designed to generate a unique id for a client using a-zA-Z0-9.
func (sh *SSEHandler) generateID() (id uuid.UUID, err error) {

	id, err = uuid.NewUUID()
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
