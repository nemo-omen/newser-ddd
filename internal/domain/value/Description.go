package value

import (
	strip "github.com/grokify/html-strip-tags-go"
)

type Description = string

// NewDescription takes a string and returns
// a string.
// If HTML tags are present, they are stripped.
// If the string is longer that 240 chars, it
// is truncated to 240 chars with ellipses appended.
func NewDescription(d string) Description {
	if len(d) < 1 {
		return ""
	}
	dd := strip.StripTags(d)
	if len(dd) > 240 {
		return dd[:240] + "..."
	}
	return dd
}
