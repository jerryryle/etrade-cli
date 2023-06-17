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
			return 0, fmt.Errorf("json.Number %s cannot be parsed as an int: %w", valueTyped.String(), err)
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
			return 0, fmt.Errorf("json.Number %s cannot be parsed as a float: %w", valueTyped.String(), err)
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
		return nil, fmt.Errorf("type %T is not a JsonMap", valueTyped)
	}
}

func valueToSlice(value interface{}) (JsonSlice, error) {
	switch valueTyped := value.(type) {
	case JsonSlice:
		return valueTyped, nil
	case nil:
		return nil, nil
	default:
		return nil, fmt.Errorf("type %T is not a JsonSlice", valueTyped)
	}
}

func valueToStringSlice(value interface{}) ([]string, error) {
	slice, err := valueToSlice(value)
	if err != nil {
		return nil, err
	}
	resultSlice := make([]string, 0, len(slice))
	for _, element := range slice {
		stringElement, err := valueToString(element)
		if err != nil {
			return nil, err
		}
		resultSlice = append(resultSlice, stringElement)
	}
	return resultSlice, nil
}

func valueToIntSlice(value interface{}) ([]int64, error) {
	slice, err := valueToSlice(value)
	if err != nil {
		return nil, err
	}
	resultSlice := make([]int64, 0, len(slice))
	for _, element := range slice {
		intElement, err := valueToInt(element)
		if err != nil {
			return nil, err
		}
		resultSlice = append(resultSlice, intElement)
	}
	return resultSlice, nil
}

func valueToFloatSlice(value interface{}) ([]float64, error) {
	slice, err := valueToSlice(value)
	if err != nil {
		return nil, err
	}
	resultSlice := make([]float64, 0, len(slice))
	for _, element := range slice {
		floatElement, err := valueToFloat(element)
		if err != nil {
			return nil, err
		}
		resultSlice = append(resultSlice, floatElement)
	}
	return resultSlice, nil
}

func valueToBoolSlice(value interface{}) ([]bool, error) {
	slice, err := valueToSlice(value)
	if err != nil {
		return nil, err
	}
	resultSlice := make([]bool, 0, len(slice))
	for _, element := range slice {
		boolElement, err := valueToBool(element)
		if err != nil {
			return nil, err
		}
		resultSlice = append(resultSlice, boolElement)
	}
	return resultSlice, nil
}

func valueToMapSlice(value interface{}) ([]JsonMap, error) {
	slice, err := valueToSlice(value)
	if err != nil {
		return nil, err
	}
	resultSlice := make([]JsonMap, 0, len(slice))
	for _, element := range slice {
		mapElement, err := valueToMap(element)
		if err != nil {
			return nil, err
		}
		resultSlice = append(resultSlice, mapElement)
	}
	return resultSlice, nil
}

func valueToSliceSlice(value interface{}) ([]JsonSlice, error) {
	slice, err := valueToSlice(value)
	if err != nil {
		return nil, err
	}
	resultSlice := make([]JsonSlice, 0, len(slice))
	for _, element := range slice {
		sliceElement, err := valueToSlice(element)
		if err != nil {
			return nil, err
		}
		resultSlice = append(resultSlice, sliceElement)
	}
	return resultSlice, nil
}
