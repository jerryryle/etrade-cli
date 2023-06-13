package responses

type OptionChainResponse struct {
	OptionPairs []OptionChainPair     `xml:"OptionPair"`
	TimeStamp   ETradeTime            `xml:"timeStamp"`
	QuoteType   string                `xml:"quoteType"`
	NearPrice   float64               `xml:"nearPrice"`
	SelectedED  OptionChainSelectedED `xml:"SelectedED"`
}

type OptionChainPair struct {
	OptionCall OptionChainOptionDetails `xml:"Call"`
	OptionPut  OptionChainOptionDetails `xml:"Put"`
	PairType   string                   `xml:"pairType"`
}

type OptionChainOptionDetails struct {
	OptionCategory   string                  `xml:"optionCategory"`
	OptionRootSymbol string                  `xml:"optionRootSymbol"`
	TimeStamp        ETradeTime              `xml:"timeStamp"`
	AdjustedFlag     bool                    `xml:"adjustedFlag"`
	DisplaySymbol    string                  `xml:"displaySymbol"`
	OptionType       string                  `xml:"optionType"`
	StrikePrice      float64                 `xml:"strikePrice"`
	Symbol           string                  `xml:"symbol"`
	Bid              float64                 `xml:"bid"`
	Ask              float64                 `xml:"ask"`
	BidSize          int                     `xml:"bidSize"`
	AskSize          int                     `xml:"askSize"`
	InTheMoney       string                  `xml:"inTheMoney"`
	Volume           int                     `xml:"volume"`
	OpenInterest     int                     `xml:"openInterest"`
	NetChange        float64                 `xml:"netChange"`
	LastPrice        float64                 `xml:"lastPrice"`
	QuoteDetail      string                  `xml:"quoteDetail"`
	OsiKey           string                  `xml:"osiKey"`
	OptionGreeks     OptionChainOptionGreeks `xml:"OptionGreeks"`
}

type OptionChainOptionGreeks struct {
	Rho          float64 `xml:"rho"`
	Vega         float64 `xml:"vega"`
	Theta        float64 `xml:"theta"`
	Delta        float64 `xml:"delta"`
	Gamma        float64 `xml:"gamma"`
	Iv           float64 `xml:"iv"`
	CurrentValue bool    `xml:"currentValue"`
}

type OptionChainSelectedED struct {
	Month int32 `xml:"month"`
	Year  int32 `xml:"year"`
	Day   int32 `xml:"day"`
}
