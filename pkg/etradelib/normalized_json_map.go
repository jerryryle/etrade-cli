package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"unicode"
	"unicode/utf8"
)

// NewNormalizedJsonMap returns a JsonMap representation of a JSON
// response from the ETrade API. It normalizes all keys in the returned map
// so that their first letter is lower-case. ETrade's keys are camel-cased;
// however, some keys are UpperCamelCase and some are lowerCamelCase. Because
// map keys are case-sensitive, normalizing them to lowerCamelCase helps avoid
// issues indexing the map.
//
// Completely lower-casing the keys would help with other inconsistencies in
// ETrade's camel-casing; however, doing so would hurt key readability. The
// normalization should result in mostly-consistent lowerCamelCase keys, where
// the remaining inconsistencies stem from ETrade's choice of word boundaries
// (e.g. lowerCamelCase vs lowerCamelcase).
func NewNormalizedJsonMap(responseBytes []byte) (jsonmap.JsonMap, error) {
	jMap, err := jsonmap.NewMapFromJsonBytes(responseBytes)
	if err != nil {
		return nil, err
	}
	return jMap.Map(lowerCaseKey, nil), nil
}

// lowerCaseKey returns the input key with its first letter lower-cased.
// It returns the input value untouched.
func lowerCaseKey(_ int, key string, value interface{}) (string, interface{}) {
	return lowerCaseFirstRuneInString(key), value
}

func lowerCaseFirstRuneInString(s string) string {
	firstRune, firstRuneSize := utf8.DecodeRuneInString(s)

	// If the first rune is the "Unicode replacement character" and its size
	// is zero or one, then the string is empty or incorrectly encoded. In
	// this case, just return the unmodified input string
	if firstRune == utf8.RuneError && firstRuneSize <= 1 {
		return s
	}
	// Lower case the first rune
	lc := unicode.ToLower(firstRune)

	// Optimization: If the lower-case version is the same as the original,
	// then just return the unmodified input string
	if firstRune == lc {
		return s
	}
	// Concatenate the lower-cased rune with the rest of the string.
	return string(lc) + s[firstRuneSize:]
}
