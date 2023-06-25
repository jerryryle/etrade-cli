package cmd

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type Renderer interface {
	Render(jsonMap jsonmap.JsonMap, descriptors []RenderDescriptor) error
	Close() error
}

type TransformerFn func(value interface{}) interface{}

type RenderDescriptor struct {
	ObjectPath   string
	Values       []RenderValue
	SubObjects   []RenderDescriptor
	DefaultValue string
	SpaceAfter   bool
}

type RenderValue struct {
	Header      string
	Path        string
	Transformer TransformerFn
}
