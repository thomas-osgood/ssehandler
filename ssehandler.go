package ssehandler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	sseconst "github.com/thomas-osgood/ssehandler/internal/constants"
	ssemsg "github.com/thomas-osgood/ssehandler/internal/messages"
)

// function designed to be used as the function passed to the http.HandleFunc()
// function when setting up a server. this will continually transmit data to the
// client as it is read from the SSEChannel.
//
// when the client disconnects, this function will detect it, close the channel,
// remove the map entry and return.
//
// references:
//
// https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events
//
// https://www.w3schools.com/html/html5_serversentevents.asp
//
// https://medium.com/@rian.eka.cahya/server-sent-event-sse-with-go-10592d9c2aa1
//
// https://www.kelche.co/blog/go/server-sent-events/
func (sh *SSEHandler) EndpointFunc(w http.ResponseWriter, r *http.Request) {
	var ctx context.Context = context.Background()
	var comms SSEChannel = make(SSEChannel)
	var err error
	var flusher http.Flusher
	var id uuid.UUID
	var ok bool
	var val SSEMessage

	// generate a random ID to attach to the current client.
	// this will be the client's key in the map and be associated
	// with the channel used to communicate with the client.
	id, err = sh.generateID(0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// add the communications channel to the clients map. this
	// allows for transmission of the same data to multiple clients.
	sh.clients[id] = comms

	// make sure to call the cleanup logic when this function exits.
	defer sh.cleanupClient(ctx, id)

	// set the necessary SSE headers.
	sh.setHeaders(w)

	// define the object that will flush the data (ie: push it to the client).
	flusher, ok = w.(http.Flusher)
	if !ok {
		log.Printf(ssemsg.ERR_UNSUPPORTED)
		return
	}

	// continually transmit the updates to the client until the client
	// disconnects or the context defined above is done.
	//
	// look at the ".Done()" documentation example code for more info
	// on why this block is structured the way it is.
	for {
		select {
		case <-r.Context().Done():
			return
		case <-ctx.Done():
			return
		case val, ok = <-comms:

			// if the channel has been closed, return.
			if !ok {
				return
			}

			// transmit the message to the client.
			//
			// first, sepecify the event type.
			//
			// second, specify the data that will be transmitted.
			//
			// note: it is recommended that the data to be transmitted is
			// in JSON format. this will help avoid transmission errors.
			fmt.Fprintf(w, sseconst.DATA_TYPE_FORMAT, val.Event)
			fmt.Fprintf(w, sseconst.DATA_EVENTID_FORMAT, val.Id)
			fmt.Fprintf(w, sseconst.DATA_TRANSMIT_FORMAT, val.Data)
			flusher.Flush()
		}
	}
}
