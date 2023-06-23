package jsonmap

// MapMapFn is the signature for a callback that the JsonMap.Map() function
// recursively applies to map keys/values. Callback implementations may mutate
// the key and value parameters and return new ones.
// As mapping proceeds, the elementPath parameter reflects the path of the
// current element's parent. The elementPath can be a mix of strings and
// ints, where strings represent map keys and ints represent slice indices.
// If the current map is located within a slice (either directly or as a nested
// element of an element within a slice), the ancestorSliceIndex parameter will
// be >= 0 to represent the index of the current map within the nearest slice
// (otherwise it will be < 0). This can be used, for example, to set a
// map's key or value to something reflecting its position in the slice.
// You could extract ancestorSliceIndex from elementPath; however, it is provided
// separately as a convenience.
//
// Example:
//
//	{
//	  "Key1": [
//	    {
//	      "Key2": {
//	        "Key3: 1
//	      }
//	    },
//	    {
//	      "Key4": {
//	        "Key5: 2
//	      }
//	    }
//	  ]
//	}
//
// When MapMapFn is called for Key1, elementPath will be [] and
// ancestorSliceIndex will be -1
// When MapMapFn is called for Key3, elementPath will be ["Key1", 0, "Key2"]
// and ancestorSliceIndex will be 0
// When MapMapFn is called for Key5, elementPath will be ["Key1", 1, "Key4"]
// and ancestorSliceIndex will be 1
type MapMapFn func(elementPath []interface{}, ancestorSliceIndex int, key string, value interface{}) (
	string, interface{},
)

// Map recursively applies mapping functions to a map and returns the new,
// transformed map. See MapMapFn and SliceMapFn for explanations of the mapping
// function parameters. Either or both may be nil, but note that maps may contain
// slices, which is why JsonMap.Map() allows you to specify mapping functions for
// both maps and slices.
func (m *JsonMap) Map(mapMapFn MapMapFn, sliceMapFn SliceMapFn) JsonMap {
	return mapMap(*m, make([]interface{}, 0), -1, mapMapFn, sliceMapFn)
}

func mapMap(
	original map[string]interface{}, elementPath []interface{}, ancestorSliceIndex int, mapMapFn MapMapFn,
	sliceMapFn SliceMapFn,
) JsonMap {
	newMap := make(JsonMap)
	for mapKey, mapValue := range original {
		// Update the path to include the current map key.
		currentElementPath := append(elementPath, mapKey)

		// Apply the mapping function to get a transformed key and value.
		newMapKey, newMapVal := mapKey, mapValue
		if mapMapFn != nil {
			newMapKey, newMapVal = mapMapFn(elementPath, ancestorSliceIndex, mapKey, mapValue)
		}
		switch newMapValTyped := newMapVal.(type) {
		case JsonMap:
			// If the new value is a JsonMap, recursively apply this map function to it
			// before adding the result to the new map with the new key.
			newMap[newMapKey] = mapMap(newMapValTyped, currentElementPath, ancestorSliceIndex, mapMapFn, sliceMapFn)
		case map[string]interface{}:
			// If the new value is a map, recursively apply this map function to it
			// before adding the result to the new map with the new key. This will also
			// change the new value's type from map[string]interface{} to JsonMap.
			newMap[newMapKey] = mapMap(newMapValTyped, currentElementPath, ancestorSliceIndex, mapMapFn, sliceMapFn)
		case JsonSlice:
			// If the new value is a JsonSlice, apply the slice map function to it
			// before adding the result to the new map with the new key.
			newMap[newMapKey] = mapSlice(newMapValTyped, currentElementPath, ancestorSliceIndex, mapMapFn, sliceMapFn)
		case []interface{}:
			// If the new value is a slice, apply the slice map function to it
			// before adding the result to the new map with the new key. This will also
			// change the new value's type from []interface{} to JsonSlice.
			newMap[newMapKey] = mapSlice(newMapValTyped, currentElementPath, ancestorSliceIndex, mapMapFn, sliceMapFn)
		default:
			// The new value is not a map, so just add it to the new map with the new key.
			newMap[newMapKey] = newMapVal
		}
	}
	return newMap
}
