package jsonmap

import (
	"errors"
	"fmt"
)

func pathSet(root interface{}, path string, value interface{}) error {
	pathElements, err := pathParse(path)
	if err != nil {
		return err
	}
	if len(pathElements) < 1 {
		return fmt.Errorf("cannot set value: the path %s must have at least one key or index element", path)
	}

	// Start with the root object
	var currentObject = root
	// We'll build up the current path as we traverse path elements to help produce better error messages.
	var currentPath = ""

	// Walk all but the last path element
	for i := 0; i < len(pathElements)-1; i++ {
		// Look ahead to see if the next path element is a slice index
		// The i+1 access will always be safe here since we're in a loop
		// that ends with len(pathElements)-1
		_, nextIsSliceIndex := pathElements[i+1].(int)

		switch pathElement := pathElements[i].(type) {
		case string: // the current path element is a map key
			var previousPath string
			currentPath, previousPath = pathUpdateForKeyElement(currentPath, pathElement)

			// Try to traverse the current object as a map
			if currentObject, err = pathSetAttemptMapTraversal(
				currentObject, pathElement, nextIsSliceIndex, currentPath, previousPath,
			); err != nil {
				return err
			}
		case int: // the current path element is a slice index
			var previousPath string
			currentPath, previousPath = pathUpdateForIndexElement(currentPath, pathElement)

			// Try to traverse the current object as a slice
			if currentObject, err = pathSetAttemptSliceTraversal(
				currentObject, pathElement, nextIsSliceIndex, currentPath, previousPath,
			); err != nil {
				return err
			}
		default:
			return errors.New("cannot set value: internal error evaluating path elements")
		}
	}

	// The final path element will be the map key or slice index for which to
	// set the new value.
	switch pathElement := pathElements[len(pathElements)-1].(type) {
	case string: // the final path element is a map key
		var previousPath string
		currentPath, previousPath = pathUpdateForKeyElement(currentPath, pathElement)

		// If the object that we're currently holding from the path traversal
		// is a map, then we can set the new value for the key.
		if currentMap, ok := currentObject.(JsonMap); ok {
			currentMap[pathElement] = value
		} else {
			return fmt.Errorf("cannot set value: cannot access %s because %s is not a map", currentPath, previousPath)
		}
	case int: // the final path element is a slice index
		var previousPath string
		currentPath, previousPath = pathUpdateForIndexElement(currentPath, pathElement)

		// If the object that we're currently holding from the path traversal
		// is a slice, then we can set the new value for the index.
		if currentSlice, ok := currentObject.(JsonSlice); ok {
			if pathElement < 0 || pathElement >= len(currentSlice) {
				return fmt.Errorf("cannot set value: slice index %d out of bounds at path %s", pathElement, currentPath)
			}
			currentSlice[pathElement] = value
		} else {
			return fmt.Errorf("cannot set value: cannot access %s because %s is not a slice", currentPath, previousPath)
		}
	default:
		return errors.New("cannot set value: internal error evaluating path elements")
	}
	return nil
}

func pathSetAttemptMapTraversal(
	currentObject interface{}, key string, nextIsSliceIndex bool, currentPath string, previousPath string,
) (interface{}, error) {
	if currentMap, ok := currentObject.(JsonMap); ok {
		nextObject, found := currentMap[key]
		// The key doesn't yet exist in the current map
		if !found {
			// If the next path element is a slice index, then the key should
			// hold a slice. We can't create slice objects here because we
			// don't know what size they should be. So, we'll return failure.
			if nextIsSliceIndex {
				return nil, fmt.Errorf("cannot set value: %s does not exist", currentPath)
			}
			// If the next path element is not a slice index, then the next
			// element must be a key. This means the current key should hold
			// a map. We'll create a new map and assign it to the key in the
			// current map as well as make it the next object.
			nextObject = JsonMap{}
			currentMap[key] = nextObject
		}
		return nextObject, nil
	}
	// Return failure because the current object is not a map
	return nil, fmt.Errorf("cannot set value: cannot access %s because %s is not a map", currentPath, previousPath)
}

func pathSetAttemptSliceTraversal(
	currentObject interface{}, index int, nextIsSliceIndex bool, currentPath string, previousPath string,
) (interface{}, error) {
	if currentSlice, ok := currentObject.(JsonSlice); ok {
		if index < 0 || index >= len(currentSlice) {
			return nil, fmt.Errorf(
				"cannot set value: slice index %d out of bounds at path %s for slice of length %d", index, currentPath,
				len(currentSlice),
			)
		}
		nextObject := currentSlice[index]
		if !nextIsSliceIndex {
			// If currentSlice[index] should hold a map but doesn't, replace
			// it with a new map.
			if _, ok := nextObject.(JsonMap); !ok {
				nextObject = JsonMap{}
				currentSlice[index] = nextObject
			}
		}
		return nextObject, nil
	}
	// Return failure because the current object is not a slice
	return nil, fmt.Errorf("cannot set value: cannot access %s because %s is not a slice", currentPath, previousPath)
}
