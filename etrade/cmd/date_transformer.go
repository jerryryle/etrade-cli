package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func dateTimeTransformer(value interface{}) interface{} {
	if timeValue, err := getValueAsTime(value); err == nil {
		return timeValue.Format(time.DateTime)
	} else {
		return value
	}
}

func dateTransformer(value interface{}) interface{} {
	if timeValue, err := getValueAsTime(value); err == nil {
		return timeValue.Format(time.DateOnly)
	} else {
		return value
	}
}

func getValueAsTime(value interface{}) (*time.Time, error) {
	var timeValue time.Time

	switch t := value.(type) {
	case json.Number:
		location, err := time.LoadLocation("America/New_York")
		if err != nil {
			return nil, err
		}
		if unixTimeInMs, err := t.Int64(); err == nil {
			timeValue = time.Unix(unixTimeInMs/1000, 0).In(location)
		}
	case string:
		if unixTime, err := parseETradeTimeString(t); err == nil {
			timeValue = *unixTime
		} else {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("cannot parse date: unknown data type")
	}
	return &timeValue, nil
}

func parseETradeTimeString(timeString string) (*time.Time, error) {
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		return nil, err
	}

	// Try parsing as Unix timestamp first
	if unixTimeInMs, err := strconv.ParseInt(timeString, 10, 64); err == nil {
		unixTime := time.Unix(unixTimeInMs/1000, 0).In(location)
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
