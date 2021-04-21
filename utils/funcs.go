package utils

import "strings"

func TrimSlashes(s *string) *string {

	found := false

	if strings.HasPrefix(*s, "/") {
		*s = (*s)[1:]
		found = true
	}

	if strings.HasSuffix(*s, "/") {
		*s = (*s)[:len(*s)-1]
		found = true
	}

	if found {
		return TrimSlashes(s)
	}
	return s
}
