package jsonmap

// MapMapFn is the signature for a callback that the JsonMap.Map() function
// recursively applies to map keys/values. Callback implementations may mutate
// the key and value parameters and return new ones. As mapping proceeds, if
// the current map is located within a slice, the parentSliceIndex parameter
// will be >= 0 and represent the index of the current map within the parent
// slice (otherwise it will be < 0). This can be used, for example, to set a
// map's key or value to something reflecting its position in the parent slice.
type MapMapFn func(parentSliceIndex int, key string, value interface{}) (string, interface{})

// SliceMapFn is the signature for a callback that the JsonMap.Map() function
// recursively applies to slice values. Callback implementations may mutate
// the value parameter and return a new one. As mapping proceeds, if
// the current value is located within a slice, the parentSliceIndex parameter
// will be >= 0 and represent the index of the current value within the parent
// slice (otherwise it will be < 0). This can be used, for example, to set a
// value to something reflecting its position in the parent slice.
type SliceMapFn func(parentSliceIndex int, index int, value interface{}) interface{}

// Map recursively applies mapping functions to a map and returns the new,
// transformed map. See MapMapFn and SliceMapFn for explanations of the mapping
// function parameters. Either or both may be nil.
func (m JsonMap) Map(mapMapFn MapMapFn, sliceMapFn SliceMapFn) JsonMap {
	return mapMap(m, -1, mapMapFn, sliceMapFn)
}

func mapMap(original map[string]interface{}, parentSliceIndex int, mapMapFn MapMapFn, sliceMapFn SliceMapFn) JsonMap {
	newMap := make(JsonMap)
	for mapKey, mapValue := range original {
		// Apply the mapping function to get a transformed key and value.
		newMapKey, newMapVal := mapKey, mapValue
		if mapMapFn != nil {
			newMapKey, newMapVal = mapMapFn(parentSliceIndex, mapKey, mapValue)
		}
		switch newMapValTyped := newMapVal.(type) {
		case JsonMap:
			// If the new value is a Json map, recursively apply this map function to it
			// before adding the result to the new map with the new key.
			newMap[newMapKey] = mapMap(newMapValTyped, parentSliceIndex, mapMapFn, sliceMapFn)
		case map[string]interface{}:
			// If the new value is a map, recursively apply this map function to it
			// before adding the result to the new map with the new key. This will also
			// change the new value's type from map[string]interface{} to JsonMap.
			newMap[newMapKey] = mapMap(newMapValTyped, parentSliceIndex, mapMapFn, sliceMapFn)
		case []interface{}:
			// If the new value is a slice, apply the slice map function to it
			// before adding the result to the new map with the new key.
			newMap[newMapKey] = mapSlice(newMapValTyped, parentSliceIndex, mapMapFn, sliceMapFn)
		default:
			// The new value is not a map, so just add it to the new map with the new key.
			newMap[newMapKey] = newMapVal
		}
	}
	return newMap
}

func mapSlice(original []interface{}, parentSliceIndex int, mapMapFn MapMapFn, sliceMapFn SliceMapFn) []interface{} {
	newSlice := make([]interface{}, len(original))
	for sliceIndex, sliceVal := range original {
		// Apply the mapping function to get a transformed value.
		newSliceVal := sliceVal
		if sliceMapFn != nil {
			newSliceVal = sliceMapFn(parentSliceIndex, sliceIndex, sliceVal)
		}
		switch newSliceValTyped := newSliceVal.(type) {
		case JsonMap:
			// If the slice value is a Json map, apply the map function to it
			// before adding the result to the new slice.
			newSlice[sliceIndex] = mapMap(newSliceValTyped, sliceIndex, mapMapFn, sliceMapFn)
		case map[string]interface{}:
			// If the slice value is a map, recursively apply this map
			// function to it before adding the result to the new slice.
			// This will also change the new value's type from
			// map[string]interface{} to JsonMap.
			newSlice[sliceIndex] = mapMap(newSliceValTyped, sliceIndex, mapMapFn, sliceMapFn)
		case []interface{}:
			// If the slice value is a slice, recursively apply this map
			// function to it before adding the result to the new slice.
			newSlice[sliceIndex] = mapSlice(newSliceValTyped, sliceIndex, mapMapFn, sliceMapFn)
		default:
			// The slice value is not a map, so just add it to the new slice.
			newSlice[sliceIndex] = newSliceValTyped
		}
	}
	return newSlice
}
