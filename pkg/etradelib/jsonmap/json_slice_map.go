package jsonmap

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
// function parameters. Either or both may be nil, but note that slices may contain
// maps, which is why JsonSlice.Map() allows you to specify mapping functions for
// both slices and maps.
func (s *JsonSlice) Map(mapMapFn MapMapFn, sliceMapFn SliceMapFn) JsonSlice {
	return mapSlice(*s, -1, mapMapFn, sliceMapFn)
}

func mapSlice(original []interface{}, parentSliceIndex int, mapMapFn MapMapFn, sliceMapFn SliceMapFn) JsonSlice {
	newSlice := make(JsonSlice, len(original))
	for sliceIndex, sliceVal := range original {
		// Apply the mapping function to get a transformed value.
		newSliceVal := sliceVal
		if sliceMapFn != nil {
			newSliceVal = sliceMapFn(parentSliceIndex, sliceIndex, sliceVal)
		}
		switch newSliceValTyped := newSliceVal.(type) {
		case JsonMap:
			// If the slice value is a JsonMap, apply the map function to it
			// before adding the result to the new slice.
			newSlice[sliceIndex] = mapMap(newSliceValTyped, sliceIndex, mapMapFn, sliceMapFn)
		case map[string]interface{}:
			// If the slice value is a map, recursively apply this map
			// function to it before adding the result to the new slice.
			// This will also change the new value's type from
			// map[string]interface{} to JsonMap.
			newSlice[sliceIndex] = mapMap(newSliceValTyped, sliceIndex, mapMapFn, sliceMapFn)
		case JsonSlice:
			// If the slice value is a JsonSlice, recursively apply this map
			// function to it before adding the result to the new slice.
			newSlice[sliceIndex] = mapSlice(newSliceValTyped, sliceIndex, mapMapFn, sliceMapFn)
		case []interface{}:
			// If the slice value is a slice, recursively apply this map
			// function to it before adding the result to the new slice.
			// This will also change the new value's type from
			// []interface{} to JsonSlice.
			newSlice[sliceIndex] = mapSlice(newSliceValTyped, sliceIndex, mapMapFn, sliceMapFn)
		default:
			// The slice value is not a map, so just add it to the new slice.
			newSlice[sliceIndex] = newSliceValTyped
		}
	}
	return newSlice
}
