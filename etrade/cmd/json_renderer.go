package cmd

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"os"
)

type JsonRenderer interface {
	Render(jsonMap jsonmap.JsonMap) error
}

type jsonRenderer struct {
	outputFile *os.File
	pretty     bool
}

func (j *jsonRenderer) Render(jsonMap jsonmap.JsonMap) error {
	return jsonMap.ToIoWriter(j.outputFile, j.pretty)
}
