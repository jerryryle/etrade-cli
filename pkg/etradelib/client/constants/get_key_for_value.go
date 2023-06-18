package constants

import "errors"

func getKeyForValue[M ~map[K]V, K comparable, V comparable](enumMap M, findValue V) (K, error) {
	for key, value := range enumMap {
		if value == findValue {
			return key, nil
		}
	}
	return *new(K), errors.New("invalid value")
}
