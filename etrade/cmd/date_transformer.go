package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func dateTimeTransformerMs(value interface{}) interface{} {
	if timeValue, err := getValueAsTime(value, true); err == nil {
		return timeValue.Format(time.DateTime)
	} else {
		return value
	}
}

func dateTransformerMs(value interface{}) interface{} {
	if timeValue, err := getValueAsTime(value, true); err == nil {
		return timeValue.Format(time.DateOnly)
	} else {
		return value
	}
}

func dateTimeTransformer(value interface{}) interface{} {
	if timeValue, err := getValueAsTime(value, false); err == nil {
		return timeValue.Format(time.DateTime)
	} else {
		return value
	}
}

func dateTransformer(value interface{}) interface{} {
	if timeValue, err := getValueAsTime(value, false); err == nil {
		return timeValue.Format(time.DateOnly)
	} else {
		return value
	}
}

func getValueAsTime(value interface{}, valueIsMs bool) (*time.Time, error) {
	var timeValue time.Time

	switch t := value.(type) {
	case json.Number:
		location, err := time.LoadLocation("America/New_York")
		if err != nil {
			return nil, err
		}
		if unixTimeInt, err := t.Int64(); err == nil {
			// If the time is zero, return an error so the value is just
			// passed along without trying to format it as a date/time string.
			if unixTimeInt == 0 {
				return nil, fmt.Errorf("timestamp is zero")
			}
			if valueIsMs {
				unixTimeInt /= 1000
			}
			timeValue = time.Unix(unixTimeInt, 0).In(location)
		}
	case string:
		if unixTime, err := parseETradeTimeString(t, valueIsMs); err == nil {
			timeValue = *unixTime
		} else {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("cannot parse date: unknown data type")
	}
	return &timeValue, nil
}

func parseETradeTimeString(timeString string, valueIsMs bool) (*time.Time, error) {
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		return nil, err
	}

	// Try parsing as Unix timestamp first
	if unixTimeInt, err := strconv.ParseInt(timeString, 10, 64); err == nil {
		// If the time is zero, return an error so the value is just
		// passed along without trying to format it as a date/time string.
		if unixTimeInt == 0 {
			return nil, fmt.Errorf("timestamp is zero")
		}
		if valueIsMs {
			unixTimeInt /= 1000
		}
		unixTime := time.Unix(unixTimeInt, 0).In(location)
		return &unixTime, nil
	}

	// If that fails, try parsing as the date string that ETrade uses.
	if parsedTime, err := time.Parse("01/02/2006", timeString); err == nil {
		return &parsedTime, nil
	}

	// If that fails, try parsing as the date-time string that ETrade uses.
	components := strings.Split(timeString, " ")
	if len(components) < 2 {
		return nil, errors.New(fmt.Sprintf("invalid date format: %s", timeString))
	}
	if components[1] != "EDT" && components[1] != "EST" {
		return nil, fmt.Errorf("unknown timezone in date: %s, (%s is not known)", timeString, components[1])
	}
	if parsedTime, err := time.ParseInLocation("15:04:05 MST 01-02-2006", timeString, location); err == nil {
		return &parsedTime, nil
	} else {
		// If that fails, return the parse error
		return nil, err
	}
}
