package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"os"
)

type CsvRenderer interface {
	Render(jsonMap jsonmap.JsonMap) error
}

type csvRenderer struct {
	outputFile *os.File
	pretty     bool
}

func (c *csvRenderer) Render(jsonMap jsonmap.JsonMap, descriptors []RenderDescriptor) error {
	writer := csv.NewWriter(c.outputFile)
	defer writer.Flush()

	for _, descriptor := range descriptors {
		var object interface{} = jsonMap
		if descriptor.ObjectPath != "" {
			object = jsonMap.GetValueAtPathWithDefault(descriptor.ObjectPath, nil)
		}
		if object != nil {
			err := writer.Write(descriptor.ValueHeaders)
			switch o := object.(type) {
			case jsonmap.JsonMap:
				err = writer.Write(getValuesForPaths(o, descriptor.ValuePaths, descriptor.DefaultValue))
				if err != nil {
					return err
				}
			case jsonmap.JsonSlice:
				for i := range o {
					element, err := o.GetMap(i)
					if err != nil {
						return err
					}
					err = writer.Write(getValuesForPaths(element, descriptor.ValuePaths, descriptor.DefaultValue))
					if err != nil {
						return err
					}
				}
			}
		}
		if descriptor.SpaceAfter {
			err := writer.Write(nil)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func getValuesForPaths(jsonMap jsonmap.JsonMap, paths []string, defaultValue string) []string {
	values := make([]string, 0, len(paths))
	for _, path := range paths {
		value := jsonMap.GetValueAtPathWithDefault(path, defaultValue)
		values = append(values, fmt.Sprintf("%v", value))
	}
	return values
}
