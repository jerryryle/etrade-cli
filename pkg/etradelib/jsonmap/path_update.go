package jsonmap

import "fmt"

func pathUpdateForKeyElement(currentPath string, key string) (string, string) {
	previousPath := currentPath
	if previousPath == "" {
		// If the last path is empty, set it to a dot to indicate the root for
		// a clearer error message if the first path element causes an error
		previousPath = "."
	}
	currentPath = fmt.Sprintf("%s.%s", currentPath, key)
	return currentPath, previousPath
}

func pathUpdateForIndexElement(currentPath string, index int) (string, string) {
	previousPath := currentPath
	if previousPath == "" {
		// If the last path is empty, set it to a dot to indicate the root for
		// a clearer error message if the first path element causes an error
		previousPath = "."
	}
	currentPath = fmt.Sprintf("%s[%d]", currentPath, index)
	return currentPath, previousPath
}
