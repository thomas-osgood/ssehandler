package utils

import "strings"

/*
helper function designed to take a string slice, normalize every element
and remove any duplicates.
*/
func UniqueSlice(original []string) []string {

	var added map[string]struct{} = make(map[string]struct{})
	var exists bool
	var n int = 0
	var o string

	for _, o = range original {
		o = strings.ToLower(strings.TrimSpace(o))
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
