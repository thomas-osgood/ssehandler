package ssehandler

import (
	"net/http"

	"github.com/thomas-osgood/ssehandler/internal/utils"
)

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

	co.AllowedMethods = utils.UniqueSlice(co.AllowedMethods)

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

	co.Origins = utils.UniqueSlice(co.Origins)
}

/*
go through all string slices and clean them.
*/
func (co *CorsOptions) cleanSlices() {
	co.cleanMethods()
	co.cleanOrigins()
}
