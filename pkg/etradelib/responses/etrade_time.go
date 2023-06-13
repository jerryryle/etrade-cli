package responses

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ETradeTime struct {
	time.Time
}

func (et *ETradeTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var (
		v   string
		err error
	)
	err = d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}

	// Allow for an empty tag that results in an zeroed time
	if v == "" {
		et.Time = time.Unix(0, 0).UTC()
		return nil
	}

	// Try parsing as Unix timestamp first
	var unixTime int64
	if unixTime, err = strconv.ParseInt(v, 10, 64); err == nil {
		et.Time = time.Unix(unixTime, 0).UTC()
		return nil
	}

	// If that fails, try parsing as the date string that ETrade uses.
	var parsedTime time.Time
	if parsedTime, err = time.Parse("01/02/2006", v); err == nil {
		et.Time = parsedTime.UTC()
		return nil
	}

	// If that fails, try parsing as the date-time string that ETrade uses.
	components := strings.Split(v, " ")
	if len(components) < 2 {
		return errors.New(fmt.Sprintf("invalid date format: %s", v))
	}
	var location *time.Location
	switch components[1] {
	case "EDT", "EST":
		location, err = time.LoadLocation("America/New_York")
		if err != nil {
			return err
		}
	default:
		return errors.New(fmt.Sprintf("unknown timezone in date: %s, (%s is not known)", v, components[1]))
	}

	if parsedTime, err = time.ParseInLocation("15:04:05 MST 01-02-2006", v, location); err == nil {
		et.Time = parsedTime.UTC()
		return nil
	}

	// If that fails, return the parse error
	return err
}

func (et *ETradeTime) GetTime() time.Time {
	return et.Time
}
