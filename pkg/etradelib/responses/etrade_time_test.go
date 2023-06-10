package responses

import (
	"encoding/xml"
	"testing"
	"time"
)

func TestETradeTime_UnmarshalXML(t *testing.T) {
	type args struct {
		xmlData string
	}
	tests := []struct {
		name      string
		args      args
		expectErr bool
		expect    time.Time
	}{
		{
			name: "Valid Unix timestamp",
			args: args{
				xmlData: `<root><unix_epoch>1641582247</unix_epoch></root>`,
			},
			expectErr: false,
			expect:    parseTimeOrPanic("2022-01-07T19:04:07Z"),
		},
		{
			name: "Valid Zeroed Unix timestamp",
			args: args{
				xmlData: `<root><unix_epoch>0</unix_epoch></root>`,
			},
			expectErr: false,
			expect:    parseTimeOrPanic("1970-01-01T00:00:00Z"),
		},
		{
			name: "Invalid Unix timestamp",
			args: args{
				xmlData: `<root><unix_epoch>invalid</unix_epoch></root>`,
			},
			expectErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var wrapper struct {
				Time ETradeTime `xml:"unix_epoch"`
			}
			if err := xml.Unmarshal([]byte(tt.args.xmlData), &wrapper); (err != nil) != tt.expectErr {
				t.Errorf("ETradeTime.UnmarshalXML() error = %v, expectErr %v", err, tt.expectErr)
			}
			if !tt.expectErr && !tt.expect.Equal(wrapper.Time.Time.UTC()) {
				t.Errorf("ETradeTime.UnmarshalXML() = %v, expect %v", wrapper.Time.Time.UTC(), tt.expect)
			}
		})
	}
}

func parseTimeOrPanic(value string) time.Time {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		panic(err)
	}
	return t
}
