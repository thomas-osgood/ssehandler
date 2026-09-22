package ssehandler

import (
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

	co.AllowedMethods = utils.UniqueSlice(co.AllowedMethods)
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
