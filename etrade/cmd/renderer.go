package cmd

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type Renderer interface {
	Render(jsonMap jsonmap.JsonMap, descriptors []RenderDescriptor) error
}

type TransformerFn func(value interface{}) interface{}

type RenderDescriptor struct {
	ObjectPath        string
	ValueHeaders      []string
	ValuePaths        []string
	ValueTransformers []TransformerFn
	DefaultValue      string
	SpaceAfter        bool
}
