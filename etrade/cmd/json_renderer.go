package cmd

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"os"
)

type jsonRenderer struct {
	outputFile *os.File
	pretty     bool
}

func (j *jsonRenderer) Render(jsonMap jsonmap.JsonMap, _ []RenderDescriptor) error {
	return jsonMap.ToIoWriter(j.outputFile, j.pretty, false)
}

func (j *jsonRenderer) Close() error {
	return j.outputFile.Close()
}
