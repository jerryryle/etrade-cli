package jsonmap

import (
	"errors"
	"fmt"
)

func pathGet(root interface{}, path string) (interface{}, error) {
	pathElements, err := pathParse(path)
	if err != nil {
		return nil, err
	}
	// Start with the root object as the current value
	var currentValue = root
	// We'll build up the current path as we traverse path elements to help produce better error messages.
	var currentPath = ""

	// Iterate over the path elements.
	for _, pathElement := range pathElements {
		switch pathElementTyped := pathElement.(type) {
		// If the path element is an int, use it to index the current element as a slice
		case int:
			lastPath := currentPath
			currentPath = currentPath + fmt.Sprintf("[%d]", pathElementTyped)
			switch currentValueTyped := currentValue.(type) {
			case JsonSlice:
				if pathElementTyped >= len(currentValueTyped) {
					return nil, fmt.Errorf(
						"cannot access %s because array index %d is greater than slice len %d", currentPath,
						pathElementTyped, len(currentValueTyped),
					)
				}
				if pathElementTyped < 0 {
					return nil, fmt.Errorf(
						"cannot access %s because array index %d is negative", currentPath, pathElementTyped,
					)
				}
				currentValue = currentValueTyped[pathElementTyped]
			default:
				return nil, fmt.Errorf("cannot access %s because %s is not an array", currentPath, lastPath)
			}
		// If the path element is a string, use it to index the current element as a map
		case string:
			lastPath := currentPath
			currentPath = currentPath + "." + pathElementTyped
			switch currentValueTyped := currentValue.(type) {
			case JsonMap:
				newValue, found := currentValueTyped[pathElementTyped]
				if !found {
					return nil, fmt.Errorf(
						"cannot access %s because key %s is not found in parent map", currentPath, pathElementTyped,
					)
				}
				currentValue = newValue
			default:
				return nil, fmt.Errorf("cannot access %s because %s is not a map", currentPath, lastPath)
			}
		default:
			return nil, errors.New("internal error evaluating path elements to get value")
		}
	}
	return currentValue, nil
}
