package bitesla_srv_exchange

func (x TradeSide) TradeSideSymbol() string {
	switch x {
	case 1:
		return "BUY"
	case 2:
		return "SELL"
	case 3:
		return "BUY_MARKET"
	case 4:
		return "SELL_MARKET"
	default:
		return "UNKNOWN"
	}
}

func (x TradeStatus) TradeStatusSymbol() string {
	return tradeStatusSymbol[x]
}

var tradeStatusSymbol = [...]string{"UNFINISH", "PART_FINISH", "FINISH", "CANCEL", "REJECT", "CANCEL_ING"}

func (m *DepthRecords) Len() int {
	return len(m.List)
}

func (m *DepthRecords) Swap(i, j int) {
	m.List[i], m.List[j] = m.List[j], m.List[i]
}

func (m *DepthRecords) Less(i, j int) bool {
	return m.List[i].Price < m.List[j].Price
}
