package cmd

import "github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"

var sortOrderMap = map[string]enumValueWithHelp[constants.SortOrder]{
	"ascending":  {constants.SortOrderAsc, "sort in ascending order"},
	"descending": {constants.SortOrderDesc, "sort in descending order"},
}

var marketSessionMap = map[string]enumValueWithHelp[constants.MarketSession]{
	"regular":  {constants.MarketSessionRegular, "regular market session"},
	"extended": {constants.MarketSessionExtended, "extended market session"},
}
