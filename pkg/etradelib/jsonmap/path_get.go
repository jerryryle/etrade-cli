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
	// Start with the root object
	var currentObject = root
	// We'll build up the current path as we traverse path elements to help produce better error messages.
	var currentPath = ""

	// Iterate over the path elements.
	for i := 0; i < len(pathElements); i++ {
		switch pathElement := pathElements[i].(type) {
		// If the path element is an int, use it to index the current object as a slice
		case int:
			var previousPath string
			currentPath, previousPath = pathUpdateForIndexElement(currentPath, pathElement)

			// Try to traverse the current object as a slice
			if currentObject, err = pathGetAttemptSliceTraversal(
				currentObject, pathElement, currentPath, previousPath,
			); err != nil {
				return nil, err
			}
		// If the path element is a string, use it to index the current object as a map
		case string:
			var previousPath string
			currentPath, previousPath = pathUpdateForKeyElement(currentPath, pathElement)

			// Try to traverse the current object as a map
			if currentObject, err = pathGetAttemptMapTraversal(
				currentObject, pathElement, currentPath, previousPath,
			); err != nil {
				return nil, err
			}
		default:
			return nil, errors.New("cannot get value: internal error evaluating path elements")
		}
	}
	return currentObject, nil
}

func pathGetAttemptMapTraversal(
	currentObject interface{}, key string, currentPath string, previousPath string,
) (interface{}, error) {
	if currentMap, ok := currentObject.(JsonMap); ok {
		nextObject, found := currentMap[key]
		if !found {
			return nil, fmt.Errorf(
				"cannot get value: cannot access %s because key %s is not found in parent map", currentPath, key,
			)
		}
		return nextObject, nil
	}
	// Return failure because the current object is not a map
	return nil, fmt.Errorf("cannot get value: cannot access %s because %s is not a map", currentPath, previousPath)
}

func pathGetAttemptSliceTraversal(
	currentObject interface{}, index int, currentPath string, previousPath string,
) (interface{}, error) {
	if currentSlice, ok := currentObject.(JsonSlice); ok {
		if index < 0 || index >= len(currentSlice) {
			return nil, fmt.Errorf("cannot get value: slice index %d out of bounds at path %s", index, currentPath)
		}
		nextObject := currentSlice[index]
		return nextObject, nil
	}
	// Return failure because the current object is not a slice
	return nil, fmt.Errorf("cannot get value: cannot access %s because %s is not a slice", currentPath, previousPath)
}
