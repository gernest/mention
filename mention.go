// Package mention provides function for parsing twitter like mentions and hashtags
package mention

import (
	"strings"
	"unicode"
)

// GetTags returns a slice of tags, that is all characters after rune char up to occurance of space
// or another occurance of rune char. Additionally you can provide a coma separated unicode characters to
// be used as terminating sequence.
func GetTags(char rune, str string, terminator ...rune) (tags []string) {
	parts := split(str, terminator) // split on terminators

	for i, _ := range parts {
		// get index number of our char in this part. If none exists will be -1 value
		x := strings.IndexRune(parts[i], char)
		if x == -1 {
			// our char is not found inside the string, therefore no tag here
			// set to empty string so it can be removed later from the slice
			parts[i] = ""
		} else if x == 0 {
			// our char is the first character, so no need to slice anything
			// before it. Aka do nothing
		} else {
			// make sure our char is not in the middle of a word by checking
			// the previous character is a space of some kind.
			if unicode.IsSpace([]rune(parts[i])[x-1]) {
				parts[i] = parts[i][strings.IndexRune(parts[i], char):]
			} else {
				// our char is in the middle of the word. Ignore this instance,
				// set to empty string for later removal.
				parts[i] = ""
			}
		}
		// trim our char from the beginning of the string. (include
		// repeat/multiple chars)
		parts[i] = strings.TrimLeft(parts[i], string(char))
	}

	// unique-ify our slice and drop empty strings
	return uniquify(parts)
}

// Splits a string based on a slice of runes
func split(s string, separators []rune) []string {
	f := func(r rune) bool {
		for _, s := range separators {
			if r == s {
				return true
			}
		}
		return false
	}
	return strings.FieldsFunc(s, f)
}

// Ensures the given slice of strings are unique and that none are empty
// strings
func uniquify(in []string) (out []string) {
	for _, i := range in {
		if i == "" {
			continue
		}
		for _, o := range out {
			if i == o {
				continue
			}
		}
		out = append(out, i)
	}
	return
}
