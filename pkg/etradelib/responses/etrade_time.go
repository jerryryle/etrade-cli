package responses

import (
	"encoding/xml"
	"time"
)

type ETradeTime struct {
	time.Time
}

func (c *ETradeTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v int64
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}
	c.Time = time.Unix(v, 0)
	return nil
}

func (c *ETradeTime) GetTime() time.Time {
	return c.Time
}
