package jsonmap

// SliceMapFn is the signature for a callback that the JsonMap.Map() function
// recursively applies to slice values. Callback implementations may mutate
// the value by returning a new one. Callbacks may also drop a value from a
// slice by returning false for the "keepValue" return value.
//
// As mapping proceeds, the elementPath parameter reflects the path of the
// current element's parent. The elementPath can be a mix of strings and
// ints, where strings represent map keys and ints represent slice indices.
// If the current slice is located within a slice (either directly or as a
// nested element of an element within a slice), the ancestorSliceIndex
// parameter will be >= 0 to represent the index of the current slice within
// the nearest slice (otherwise it will be < 0). This can be used, for example,
// to set a slice's value to something reflecting its position in the slice.
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
type SliceMapFn func(elementPath []interface{}, ancestorSliceIndex int, index int, value interface{}) (
	newValue interface{}, keepValue bool,
)

// Map recursively applies mapping functions to a map and returns the new,
// transformed map. See MapMapFn and SliceMapFn for explanations of the mapping
// function parameters. Either or both may be nil, but note that slices may
// contain maps, which is why JsonSlice.Map() allows you to specify mapping
// functions for both slices and maps.
func (s *JsonSlice) Map(mapMapFn MapMapFn, sliceMapFn SliceMapFn) JsonSlice {
	return mapSlice(*s, make([]interface{}, 0), -1, mapMapFn, sliceMapFn)
}

func mapSlice(
	original []interface{}, elementPath []interface{}, ancestorSliceIndex int, mapMapFn MapMapFn, sliceMapFn SliceMapFn,
) JsonSlice {
	newSlice := make(JsonSlice, 0, len(original))
	for sliceIndex, sliceVal := range original {
		// Update the path to include the current slice index
		currentElementPath := append(elementPath, sliceIndex)

		// Apply the mapping function to get a transformed value.
		newSliceVal := sliceVal
		keepValue := true
		if sliceMapFn != nil {
			newSliceVal, keepValue = sliceMapFn(currentElementPath, ancestorSliceIndex, sliceIndex, sliceVal)
		}
		if keepValue {
			switch newSliceValTyped := newSliceVal.(type) {
			case JsonMap:
				// If the slice value is a JsonMap, apply the map function to
				// it before adding the result to the new slice.
				newSlice = append(
					newSlice, mapMap(newSliceValTyped, currentElementPath, sliceIndex, mapMapFn, sliceMapFn),
				)
			case map[string]interface{}:
				// If the slice value is a map, recursively apply this map
				// function to it before adding the result to the new slice.
				// This will also change the new value's type from
				// map[string]interface{} to JsonMap.
				newSlice = append(
					newSlice, mapMap(newSliceValTyped, currentElementPath, sliceIndex, mapMapFn, sliceMapFn),
				)
			case JsonSlice:
				// If the slice value is a JsonSlice, recursively apply this
				// map function to it before adding the result to the new
				// slice.
				newSlice = append(
					newSlice, mapSlice(newSliceValTyped, currentElementPath, sliceIndex, mapMapFn, sliceMapFn),
				)
			case []interface{}:
				// If the slice value is a slice, recursively apply this map
				// function to it before adding the result to the new slice.
				// This will also change the new value's type from
				// []interface{} to JsonSlice.
				newSlice = append(
					newSlice, mapSlice(newSliceValTyped, currentElementPath, sliceIndex, mapMapFn, sliceMapFn),
				)
			default:
				// The slice value is not a map or slice, so just add it to
				// the new slice.
				newSlice = append(newSlice, newSliceValTyped)
			}
		}
	}
	return newSlice
}
