package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"os"
)

type csvRenderer struct {
	outputFile *os.File
	pretty     bool
}

func (c *csvRenderer) Render(jsonMap jsonmap.JsonMap, descriptors []RenderDescriptor) error {
	writer := csv.NewWriter(c.outputFile)
	defer writer.Flush()

	return renderObject(writer, jsonMap, descriptors)
}

func (c *csvRenderer) Close() error {
	return c.outputFile.Close()
}

func renderObject(writer *csv.Writer, jsonMap jsonmap.JsonMap, descriptors []RenderDescriptor) error {
	for _, descriptor := range descriptors {
		var object interface{} = jsonMap
		if descriptor.ObjectPath != "" {
			object = jsonMap.GetValueAtPathWithDefault(descriptor.ObjectPath, nil)
		}
		if object != nil {
			switch o := object.(type) {
			case jsonmap.JsonMap:
				err := writer.Write(getHeadersForRenderValues(descriptor.Values))
				err = writer.Write(
					getValuesForRenderValues(o, descriptor.Values, descriptor.DefaultValue),
				)
				if err != nil {
					return err
				}
				// Render any sub-objects
				if len(descriptor.SubObjects) > 0 {
					err = renderObject(writer, o, descriptor.SubObjects)
					if err != nil {
						return err
					}
				}
			case jsonmap.JsonSlice:
				var err error
				for i := range o {
					// Always write the header before the first element, but also repeat
					// the header for each element if there are sub-objects.
					if i == 0 || len(descriptor.SubObjects) > 0 {
						err = writer.Write(getHeadersForRenderValues(descriptor.Values))
					}
					var element jsonmap.JsonMap
					element, err = o.GetMap(i)
					if err != nil {
						return err
					}
					err = writer.Write(
						getValuesForRenderValues(element, descriptor.Values, descriptor.DefaultValue),
					)
					if err != nil {
						return err
					}
					// Render any sub-objects
					if len(descriptor.SubObjects) > 0 {
						err = renderObject(writer, element, descriptor.SubObjects)
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
	}
	return nil
}

func getValuesForRenderValues(
	jsonMap jsonmap.JsonMap, renderValues []RenderValue, defaultValue string,
) []string {
	values := make([]string, 0, len(renderValues))
	for _, renderValue := range renderValues {
		value := jsonMap.GetValueAtPathWithDefault(renderValue.Path, defaultValue)
		if renderValue.Transformer != nil {
			value = renderValue.Transformer(value)
		}
		values = append(values, fmt.Sprintf("%v", value))
	}
	return values
}

func getHeadersForRenderValues(renderValues []RenderValue) []string {
	headers := make([]string, 0, len(renderValues))
	for _, value := range renderValues {
		headers = append(headers, value.Header)
	}
	return headers
}
