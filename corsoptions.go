package ssehandler

import "strings"

/*
normalize (trim + lower) each element in the Origins slice
and remove the empty strings.
*/
func (co *CorsOptions) cleanOrigins() {
	if co == nil || len(co.Origins) < 1 {
		return
	}

	var n int = 0
	var o string

	for _, o = range co.Origins {
		o = strings.ToLower(strings.TrimSpace(o))
		if len(o) > 0 {
			co.Origins[n] = o
			n++
		}
	}

	co.Origins = co.Origins[:n]
}
