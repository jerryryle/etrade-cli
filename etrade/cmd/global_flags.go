package cmd

type GlobalFlags struct {
	customerId     string
	debug          bool
	outputFileName string
	outputFormat   enumFlagValue[OutputFormat]
}

type OutputFormat int

const (
	OutputFormatCsv = iota
	OutputFormatJson
	OutputFormatJsonPretty
)
