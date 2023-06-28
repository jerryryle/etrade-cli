package cmd

type globalFlags struct {
	customerId     string
	debug          bool
	outputFileName string
	outputFormat   enumFlagValue[outputFormat]
}

type outputFormat int

const (
	outputFormatCsv = iota
	outputFormatJson
	outputFormatJsonPretty
)
