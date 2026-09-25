package ssehandler

import (
	"net/http"
	"strings"

	"github.com/thomas-osgood/ssehandler/internal/utils"
)

/*
normalize each element, remove empty strings and remove
duplicates from the AllowHeaders slice.
*/
func (co *CorsOptions) cleanAllowHeaders() {
	if co == nil || len(co.AllowHeaders) < 1 {
		return
	}

	co.AllowHeaders = utils.ToHeaderCase(utils.UniqueSlice(co.AllowHeaders, false))
}

/*
normalize each element, remove empty strings and remove
duplicates from the ExposeHeaders slice.
*/
func (co *CorsOptions) cleanExposeHeaders() {
	if co == nil || len(co.ExposeHeaders) < 1 {
		return
	}

	co.ExposeHeaders = utils.ToHeaderCase(utils.UniqueSlice(co.ExposeHeaders, false))
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
func (co *CorsOptions) cleanOrigin() {
	if co == nil || len(co.Origin) < 1 {
		return
	}

	co.Origin = strings.TrimSpace(co.Origin)
}

/*
go through all string slices and clean them.
*/
func (co *CorsOptions) cleanSlices() {
	co.cleanAllowHeaders()
	co.cleanExposeHeaders()
	co.cleanMethods()
	co.cleanOrigin()
}
