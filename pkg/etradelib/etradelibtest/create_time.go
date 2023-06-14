package etradelibtest

import "time"

// These functions return pointers to time variables. This is useful for initializing struct members of type
// *time.Time. Go doesn't allow you to do `&time.Date()` because the returned time.Time value isn't guaranteed to have
// a known memory location. For example, it might exist solely in a register for the return call. But creating a local
// variable and returning its address will ensure that the time.Time object is created on the heap. These helper
// functions do just that.

func CreateTime(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) *time.Time {
	theTime := time.Date(year, month, day, hour, min, sec, nsec, loc)
	return &theTime
}

func CreateTimeFromString(layout, value string) *time.Time {
	theTime, err := time.Parse(layout, value)
	if err != nil {
		return nil
	}
	return &theTime
}

func CreateUnixTime(sec int64, nsec int64) *time.Time {
	theTime := time.Unix(sec, nsec).UTC()
	return &theTime
}
