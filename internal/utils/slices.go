package utils

import "strings"

/*
helper function designed to take a string slice, normalize every element
and remove any duplicates.
*/
func UniqueSlice(original []string, upperNorm bool) []string {

	var added map[string]struct{} = make(map[string]struct{})
	var exists bool
	var normalization func(string) string = strings.ToLower
	var n int = 0
	var o string

	// if the upper normalization flag has been set, use the
	// ToUpper method to normalize the string; otherwise, the
	// normalization method will be ToLower.
	if upperNorm {
		normalization = strings.ToUpper
	}

	for _, o = range original {
		o = normalization(strings.TrimSpace(o))
		if len(o) > 0 {
			if _, exists = added[o]; !exists {
				added[o] = struct{}{}
				original[n] = o
				n++
			}
		}
	}

	return original[:n]
}
