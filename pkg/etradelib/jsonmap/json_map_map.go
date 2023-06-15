package jsonmap

type MapFn func(key string, value interface{}) (string, interface{})

func (m JsonMap) Map(fn MapFn) JsonMap {
	return mapMap(m, fn)
}

func mapMap(original map[string]interface{}, fn MapFn) map[string]interface{} {
	newMap := make(map[string]interface{})
	for k, v := range original {
		// Apply the mapping function to get a transformed key and value.
		newKey, newVal := fn(k, v)
		switch t := newVal.(type) {
		case map[string]interface{}:
			// If the new value is a map, recursively apply this map function to it
			// before adding the result to the new map with the new key.
			newMap[newKey] = mapMap(t, fn)
		case []interface{}:
			// If the new value is a slice, walk it, looking for any maps.
			newSlice := make([]interface{}, len(t))
			for i, val := range t {
				if valMap, ok := val.(map[string]interface{}); ok {
					// If the slice value is a map, recursively apply this map
					// function to it before adding the result to the new slice.
					newSlice[i] = mapMap(valMap, fn)
				} else {
					// The slice value is not a map, so just add it to the new slice.
					newSlice[i] = val
				}
			}
			newMap[newKey] = newSlice
		default:
			// The new value is not a map, so just add it to the new map with the new key.
			newMap[newKey] = newVal
		}
	}
	return newMap
}
