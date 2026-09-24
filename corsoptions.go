package ssehandler

import (
	"net/http"

	"github.com/thomas-osgood/ssehandler/internal/utils"
)

/*
normalize each element, remove empty strings and remove
duplicates from the ExposeHeaders slice.
*/
func (co *CorsOptions) cleanExposeHeaders() {
	if co == nil || len(co.ExposeHeaders) < 1 {
		return
	}

	var idx int

	co.ExposeHeaders = utils.UniqueSlice(co.ExposeHeaders, false)

	// correct header casing using CanonicalHeaderKey from the http package.
	for idx = range co.ExposeHeaders {
		co.ExposeHeaders[idx] = http.CanonicalHeaderKey(co.ExposeHeaders[idx])
	}
}

/*
normalize each element, remove empty strings, remove duplicates
and remove invalid methods in the AllowedMethods slice.
*/
func (co *CorsOptions) cleanMethods() {
	if co == nil || len(co.AllowedMethods) < 1 {
		return
	}

	var curMethod string
	var exists bool
	var finalSlice []string
	var validMethods = map[string]struct{}{
		http.MethodGet:     {},
		http.MethodHead:    {},
		http.MethodPost:    {},
		http.MethodPut:     {},
		http.MethodPatch:   {},
		http.MethodDelete:  {},
		http.MethodConnect: {},
		http.MethodOptions: {},
		http.MethodTrace:   {},
	}

	co.AllowedMethods = utils.UniqueSlice(co.AllowedMethods, true)

	for _, curMethod = range co.AllowedMethods {
		if _, exists = validMethods[curMethod]; exists {
			finalSlice = append(finalSlice, curMethod)
		}
	}

	co.AllowedMethods = finalSlice
}

/*
normalize (trim + lower) each element in the Origins slice
and remove the empty strings.
*/
func (co *CorsOptions) cleanOrigins() {
	if co == nil || len(co.Origins) < 1 {
		return
	}

	co.Origins = utils.UniqueSlice(co.Origins, false)
}

/*
go through all string slices and clean them.
*/
func (co *CorsOptions) cleanSlices() {
	co.cleanExposeHeaders()
	co.cleanMethods()
	co.cleanOrigins()
}
