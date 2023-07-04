package cmd

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDateTransformer(t *testing.T) {
	// Transforms Unix timestamp as JSON Number to Date/Time
	assert.Equal(t, "1999-01-01 04:00:00", dateTimeTransformerMs(json.Number("915181200000")))
	assert.Equal(t, "1999-01-01 04:00:00", dateTimeTransformer(json.Number("915181200")))
	assert.Equal(t, "1999-01-01", dateTransformerMs(json.Number("915181200000")))
	assert.Equal(t, "1999-01-01", dateTransformer(json.Number("915181200")))
	assert.Equal(t, "04:00:00", timeTransformerMs(json.Number("915181200000")))
	assert.Equal(t, "04:00:00", timeTransformer(json.Number("915181200")))

	// Transforms Unix timestamp as string to Date/Time
	assert.Equal(t, "1999-01-01 04:00:00", dateTimeTransformerMs("915181200000"))
	assert.Equal(t, "1999-01-01 04:00:00", dateTimeTransformer("915181200"))
	assert.Equal(t, "1999-01-01", dateTransformerMs("915181200000"))
	assert.Equal(t, "1999-01-01", dateTransformer("915181200"))
	assert.Equal(t, "04:00:00", timeTransformerMs("915181200000"))
	assert.Equal(t, "04:00:00", timeTransformer("915181200"))

	// Does Not transform zero-value Unix timestamp as JSON Number
	assert.Equal(t, json.Number("0"), dateTimeTransformerMs(json.Number("0")))
	assert.Equal(t, json.Number("0"), dateTimeTransformer(json.Number("0")))
	assert.Equal(t, json.Number("0"), dateTransformerMs(json.Number("0")))
	assert.Equal(t, json.Number("0"), dateTransformer(json.Number("0")))
	assert.Equal(t, json.Number("0"), timeTransformerMs(json.Number("0")))
	assert.Equal(t, json.Number("0"), timeTransformer(json.Number("0")))

	// Does Not transform zero-value Unix timestamp as string
	assert.Equal(t, "0", dateTimeTransformerMs("0"))
	assert.Equal(t, "0", dateTimeTransformer("0"))
	assert.Equal(t, "0", dateTransformerMs("0"))
	assert.Equal(t, "0", dateTransformer("0"))
	assert.Equal(t, "0", timeTransformerMs("0"))
	assert.Equal(t, "0", timeTransformer("0"))

	// Transforms date/time string to date/time
	assert.Equal(t, "1999-01-01 04:00:00", dateTimeTransformerMs("04:00:00 EST 01-01-1999"))
	assert.Equal(t, "1999-01-01 04:00:00", dateTimeTransformer("04:00:00 EST 01-01-1999"))
	assert.Equal(t, "1999-01-01", dateTransformerMs("04:00:00 EST 01-01-1999"))
	assert.Equal(t, "1999-01-01", dateTransformer("04:00:00 EST 01-01-1999"))
	assert.Equal(t, "04:00:00", timeTransformerMs("04:00:00 EST 01-01-1999"))
	assert.Equal(t, "04:00:00", timeTransformer("04:00:00 EST 01-01-1999"))

	// Transforms date string to date/time
	assert.Equal(t, "1999-01-01 00:00:00", dateTimeTransformerMs("01/01/1999"))
	assert.Equal(t, "1999-01-01 00:00:00", dateTimeTransformer("01/01/1999"))
	assert.Equal(t, "1999-01-01", dateTransformerMs("01/01/1999"))
	assert.Equal(t, "1999-01-01", dateTransformer("01/01/1999"))
	assert.Equal(t, "00:00:00", timeTransformerMs("01/01/1999"))
	assert.Equal(t, "00:00:00", timeTransformer("01/01/1999"))

	// Does not transform invalid time string
	assert.Equal(t, "InvalidTimeString", dateTimeTransformerMs("InvalidTimeString"))
	assert.Equal(t, "InvalidTimeString", dateTimeTransformer("InvalidTimeString"))
	assert.Equal(t, "InvalidTimeString", dateTransformerMs("InvalidTimeString"))
	assert.Equal(t, "InvalidTimeString", dateTransformer("InvalidTimeString"))
	assert.Equal(t, "InvalidTimeString", timeTransformerMs("InvalidTimeString"))
	assert.Equal(t, "InvalidTimeString", timeTransformer("InvalidTimeString"))

	// Does not transform time string with unexpected time zone
	assert.Equal(t, "04:00:00 PST 01-01-1999", dateTimeTransformerMs("04:00:00 PST 01-01-1999"))
	assert.Equal(t, "04:00:00 PST 01-01-1999", dateTimeTransformer("04:00:00 PST 01-01-1999"))
	assert.Equal(t, "04:00:00 PST 01-01-1999", dateTransformerMs("04:00:00 PST 01-01-1999"))
	assert.Equal(t, "04:00:00 PST 01-01-1999", dateTransformer("04:00:00 PST 01-01-1999"))
	assert.Equal(t, "04:00:00 PST 01-01-1999", timeTransformerMs("04:00:00 PST 01-01-1999"))
	assert.Equal(t, "04:00:00 PST 01-01-1999", timeTransformer("04:00:00 PST 01-01-1999"))

	// Does not transform invalid time string that has an expected time zone
	// and enough components to otherwise look like a valid time string
	assert.Equal(t, "Invalid EST String", dateTimeTransformerMs("Invalid EST String"))
	assert.Equal(t, "Invalid EST String", dateTimeTransformer("Invalid EST String"))
	assert.Equal(t, "Invalid EST String", dateTransformerMs("Invalid EST String"))
	assert.Equal(t, "Invalid EST String", dateTransformer("Invalid EST String"))
	assert.Equal(t, "Invalid EST String", timeTransformerMs("Invalid EST String"))
	assert.Equal(t, "Invalid EST String", timeTransformer("Invalid EST String"))

	// Does not transform invalid types
	assert.Equal(t, []int{1, 2, 3, 4}, dateTimeTransformerMs([]int{1, 2, 3, 4}))
	assert.Equal(t, []int{1, 2, 3, 4}, dateTimeTransformer([]int{1, 2, 3, 4}))
	assert.Equal(t, []int{1, 2, 3, 4}, dateTransformerMs([]int{1, 2, 3, 4}))
	assert.Equal(t, []int{1, 2, 3, 4}, dateTransformer([]int{1, 2, 3, 4}))
	assert.Equal(t, []int{1, 2, 3, 4}, timeTransformerMs([]int{1, 2, 3, 4}))
	assert.Equal(t, []int{1, 2, 3, 4}, timeTransformer([]int{1, 2, 3, 4}))

}
