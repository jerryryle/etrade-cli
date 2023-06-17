package jsonmap

import (
	"encoding/json"
	"fmt"
)

func valueToString(value interface{}) (string, error) {
	switch valueTyped := value.(type) {
	case string:
		return valueTyped, nil
	case nil:
		return "", nil
	default:
		return "", fmt.Errorf("type %T is not a string", valueTyped)
	}
}

func valueToInt(value interface{}) (int64, error) {
	switch valueTyped := value.(type) {
	case int64:
		return valueTyped, nil
	case int32:
		return int64(valueTyped), nil
	case int:
		return int64(valueTyped), nil
	case json.Number:
		intVal, err := valueTyped.Int64()
		if err != nil {
			return 0, fmt.Errorf("type %T is not an int: %w", valueTyped, err)
		}
		return intVal, nil
	default:
		return 0, fmt.Errorf("type %T is not an int", valueTyped)
	}
}

func valueToFloat(value interface{}) (float64, error) {
	switch valueTyped := value.(type) {
	case float64:
		return valueTyped, nil
	case float32:
		return float64(valueTyped), nil
	case json.Number:
		floatVal, err := valueTyped.Float64()
		if err != nil {
			return 0, fmt.Errorf("type %T is not a float: %w", valueTyped, err)
		}
		return floatVal, nil
	default:
		return 0, fmt.Errorf("type %T is not a float", valueTyped)
	}
}

func valueToBool(value interface{}) (bool, error) {
	switch valueTyped := value.(type) {
	case bool:
		return valueTyped, nil
	default:
		return false, fmt.Errorf("type %T is not a bool", valueTyped)
	}
}

func valueToMap(value interface{}) (JsonMap, error) {
	switch valueTyped := value.(type) {
	case JsonMap:
		return valueTyped, nil
	case nil:
		return nil, nil
	default:
		return nil, fmt.Errorf("type %T is not a map", valueTyped)
	}
}

func valueToSlice(value interface{}) (JsonSlice, error) {
	switch valueTyped := value.(type) {
	case JsonSlice:
		return valueTyped, nil
	case nil:
		return nil, nil
	default:
		return nil, fmt.Errorf("type %T is not a slice", valueTyped)
	}
}
