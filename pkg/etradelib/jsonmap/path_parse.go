package jsonmap

import (
	"fmt"
	"strconv"
	"strings"
)

// pathParse attempts to parse a path into a slice consisting of any key strings
// and index integers. Leading and/or redundant path dots will be ignored.
// e.g. pathParse("foo.bar[1].moo")  -> []interface{}{"foo", "bar", 1, "moo"}
// e.g. pathParse(".foo.bar[1].moo") -> []interface{}{"foo", "bar", 1, "moo"}
// e.g. pathParse("foo..bar[1].moo") -> []interface{}{"foo", "bar", 1, "moo"}
func pathParse(path string) ([]interface{}, error) {
	pathElements := strings.Split(path, ".")
	returnPathElements := make([]interface{}, 0, len(pathElements)*2)
	for _, pathElement := range pathElements {
		if pathElement != "" {
			keyIndexPathElements, err := splitKeyIndices(pathElement)
			if err != nil {
				return nil, err
			}
			returnPathElements = append(returnPathElements, keyIndexPathElements...)
		}
	}
	return returnPathElements, nil
}

// splitKeyIndices splits a key and array indices into a slice containing the
// key string and zero or more integer indices. If the key has no indices, then
// the return slice will contain only the key.
// e.g. splitKeyIndices("test") -> []interface{}{"test"}
// e.g. splitKeyIndices("test[0]") -> []interface{}{"test", 0}
// e.g. splitKeyIndices("test[0][1]") -> []interface{}{"test", 0, 1}
// If the input string contains missing/spurious characters around the indices
// or the index values cannot be parsed as integers, then splitKeyIndices will
// return an error.
func splitKeyIndices(s string) ([]interface{}, error) {
	// Try to split the string on opening brackets.
	// e.g. s:"test[0][1]" -> keyAndIndicesSlice:[ "test", "0]", "1]" ]
	keyAndIndicesSlice := strings.Split(s, "[")
	// If the length of the resulting slice is 1 (it will never be zero--even
	// for an empty string, but we compare against <= 1 to be extra paranoid),
	// then no opening bracket was found. Just return a slice with the original
	// string since there's no index to extract.
	if len(keyAndIndicesSlice) <= 1 {
		return []interface{}{s}, nil
	}

	// Create a new slice for the key and indices
	// e.g. keyAndIndicesSlice:[ "test", "0]", "1]" ] -> keyAndIndicesSlice[0]:"test"
	returnKeyAndIndices := []interface{}{}

	// If the key isn't empty, add it to the slice.
	if keyAndIndicesSlice[0] != "" {
		returnKeyAndIndices = append(returnKeyAndIndices, keyAndIndicesSlice[0])
	}

	// Now extract index value(s)
	// e.g. keyAndIndicesSlice:[ "test", "0]", "1]" ] -> keyAndIndicesSlice[1:]:[ "0]", "1]" ]
	for _, indexString := range keyAndIndicesSlice[1:] {
		// Try to split the string on closing brackets.
		// e.g. indexString:"0]" -> indexSlice: [ "0", "" ]
		indexSlice := strings.Split(indexString, "]")
		// If the length of the resulting slice is not exactly 2, then
		// the closing bracket is missing.
		// e.g. indexString:"0" -> indexSlice: [ "0" ]
		//
		// If the 2nd element of the slice isn't an empty string, then there
		// are spurious characters after the closing bracket.
		// indexString:"0]foo" -> indexSlice: [ "0", "foo" ]
		if len(indexSlice) != 2 || indexSlice[1] != "" {
			return nil, fmt.Errorf("invalid index in key %s", s)
		}
		// Convert the string index value to an integer
		// e.g. indexSlice[0]:"0" -> indexValue: 0
		indexValue, err := strconv.Atoi(indexSlice[0])
		if err != nil {
			return nil, fmt.Errorf("invalid index in key %s", s)
		}
		// Append the index integer to the return slice
		returnKeyAndIndices = append(returnKeyAndIndices, indexValue)
	}
	return returnKeyAndIndices, nil
}
