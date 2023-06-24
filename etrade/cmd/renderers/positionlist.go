package renderers

import (
	"encoding/csv"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"os"
)

func PositionListRenderText(outputFile *os.File, positionList etradelib.ETradePositionList) error {
	writer := csv.NewWriter(outputFile)
	err := writer.Write(quickViewHeader)
	if err != nil {
		return err
	}

	for _, position := range positionList.GetAllPositions() {
		positionMap := position.GetJsonMap()
		err = writer.Write(getQuickViewValues(positionMap))
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}

var quickViewHeader = []string{
	"Symbol",
	"Last Price $", "Change $", "Change %",
	"Quantity", "Price Paid $", "Day's Gain $", "Total Gain $", "Total Gain %", "Value $",
}

var quickViewPaths = []string{
	".product.symbol",
	".quick.lastTrade", ".quick.change", ".quick.changePct",
	".quantity", ".pricePaid", ".daysGain", ".totalGain", ".totalGainPct", ".marketValue",
}

func getQuickViewValues(positionMap jsonmap.JsonMap) []string {
	values := make([]string, 0, len(quickViewPaths))
	for _, path := range quickViewPaths {
		value := positionMap.GetValueAtPathWithDefault(path, "")
		values = append(values, fmt.Sprintf("%v", value))
	}
	return values
}
