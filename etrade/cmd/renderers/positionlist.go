package renderers

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"os"
)

func PositionListRenderJson(outputFile *os.File, positionList etradelib.ETradePositionList, pretty bool) error {
	positionMap := positionList.AsJsonMap()
	return positionMap.ToIoWriter(outputFile, pretty)
}

func PositionListRenderText(outputFile *os.File, positionList etradelib.ETradePositionList) error {
	return nil
}
