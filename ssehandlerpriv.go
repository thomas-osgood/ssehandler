package ssehandler

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"net/http"

	sseconst "github.com/thomas-osgood/ssehandler/internal/constants"
	ssemsg "github.com/thomas-osgood/ssehandler/internal/messages"
)

// function designed to execute the logic for when a client
// disconnects from the SSE endpoint.
func (sh *SSEHandler) cleanupClient(ctx context.Context, clientid string) {
	close(sh.clients[clientid])
	delete(sh.clients, clientid)
	ctx.Done()
}

// function designed to check whether a client with the given
// id already exists in the clients map.
func (sh *SSEHandler) clientIdExists(id string) (exists bool) {
	_, exists = sh.clients[id]
	return exists
}

// function designed to generate a unique id for a client using a-zA-Z0-9.
func (sh *SSEHandler) generateID(attempt int) (id string, err error) {
	var baseval *big.Int
	var biglen *big.Int
	var bigmin *big.Int
	const charset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var i int
	var length int
	const maxlen int = 15
	const minlen int = 5
	var randidx *big.Int

	// validate min/max parameters
	if (minlen <= 0) || (maxlen <= 0) {
		return "", errors.New(ssemsg.ERR_MAXMIN_LEN)
	} else if minlen > maxlen {
		return "", errors.New(ssemsg.ERR_MIN_LEN)
	}

	// this is the number that will be used to generate
	// the random number. this is the difference of the
	// max value and min value because the final random
	// number will be calculated by adding the min value
	// so the number falls within the range MIN <= x <= MAX.
	baseval = big.NewInt(int64(maxlen - minlen))

	// convert the minimum value to a big.Int so it can be
	// used to adjust the randomly generated length.
	bigmin = big.NewInt(int64(minlen))

	// use the crypto/rand library to generate a length
	// for the string.
	biglen, err = rand.Int(rand.Reader, big.NewInt(baseval.Int64()))
	if err != nil {
		return "", fmt.Errorf(ssemsg.ERR_RANDSTR_LEN, err)
	}

	// adjust the generated number to fit within the range.
	biglen = biglen.Add(biglen, bigmin)

	length = int(biglen.Int64())

	id = ""

	for i = 0; i < length; i++ {
		// calculate the random index to choose. if
		// there is an error, choose index 0.
		randidx, err = rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			randidx = big.NewInt(0)
		}
		// append the char at the randomly generated
		// index to the randomly generated string.
		id = fmt.Sprintf("%s%c", id, charset[randidx.Int64()])
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
			return "", fmt.Errorf(ssemsg.ERR_GENERATEID_MAXATTEMPTS)
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
