package ssehandler

import (
	"github.com/thomas-osgood/ssehandler/internal/utils"
)

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
