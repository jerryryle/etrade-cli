package cmd

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type Renderer interface {
	Render(jsonMap jsonmap.JsonMap, descriptors []RenderDescriptor) error
}

type RenderDescriptor struct {
	ObjectPath   string
	ValueHeaders []string
	ValuePaths   []string
	SpaceAfter   bool
}
