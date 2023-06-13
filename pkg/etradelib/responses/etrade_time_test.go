package responses

import (
	"encoding/xml"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/etradelibtest"
	"testing"
	"time"

	// Ensure the tests have an embedded copy of tzdata in case the host system is missing this.
	_ "time/tzdata"
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
				xmlData: `<root><time>1641582247</time></root>`,
			},
			expectErr: false,
			expect:    etradelibtest.MustParseRFC3339("2022-01-07T19:04:07Z"),
		},
		{
			name: "Valid Zeroed Unix timestamp",
			args: args{
				xmlData: `<root><time>0</time></root>`,
			},
			expectErr: false,
			expect:    etradelibtest.MustParseRFC3339("1970-01-01T00:00:00Z"),
		},
		{
			name: "Invalid Unix timestamp",
			args: args{
				xmlData: `<root><time>invalid</time></root>`,
			},
			expectErr: true,
		},
		{
			name: "Valid ETrade string timestamp (EDT)",
			args: args{
				xmlData: `<root><time>16:00:00 EDT 06-20-2012</time></root>`,
			},
			expectErr: false,
			expect:    etradelibtest.MustParseRFC3339("2012-06-20T20:00:00Z"),
		},
		{
			name: "Valid ETrade string timestamp (EST)",
			args: args{
				xmlData: `<root><time>16:00:00 EST 05-20-2012</time></root>`,
			},
			expectErr: false,
			expect:    etradelibtest.MustParseRFC3339("2012-05-20T21:00:00Z"),
		},
		{
			name: "Invalid ETrade string timestamp",
			args: args{
				xmlData: `<root><time>16:00:00EST 05-20-2012</time></root>`,
			},
			expectErr: true,
		},
		{
			name: "Invalid Timezone",
			args: args{
				xmlData: `<root><time>16:00:00 MOO 05-20-2012</time></root>`,
			},
			expectErr: true,
		},
		{
			name: "Valid Empty field",
			args: args{
				xmlData: `<root><time></time></root>`,
			},
			expectErr: false,
			expect:    time.Unix(0, 0),
		},
		{
			name: "Valid ETrade Date",
			args: args{
				xmlData: `<root><time>05/20/2012</time></root>`,
			},
			expectErr: false,
			expect:    etradelibtest.MustParseRFC3339("2012-05-20T00:00:00Z"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var wrapper struct {
				ETradeTime ETradeTime `xml:"time"`
			}
			if err := xml.Unmarshal([]byte(tt.args.xmlData), &wrapper); (err != nil) != tt.expectErr {
				t.Errorf("ETradeTime.UnmarshalXML() error = %v, expectErr %v", err, tt.expectErr)
			}
			if !tt.expectErr && !tt.expect.Equal(wrapper.ETradeTime.GetTime()) {
				t.Errorf("ETradeTime.UnmarshalXML() = %v, expect %v", wrapper.ETradeTime.GetTime(), tt.expect)
			}
		})
	}
}
