package postion

import (
	"github.com/ta4g/ta4g/data/interval/trade/constants/order_type"
	"github.com/ta4g/ta4g/data/interval/trade/trade_record"
	"time"
)

// Trade represents a collection of the items that are purchased or sold in a single batch
type Trade struct {
	order_type.OrderType `csv:"order_type" avro:"order_type" json:"order_type"`
	// UnixTime the order was placed, for back-testing we will assume all postion are filled immediately.
	UnixTime int64 `csv:"time" avro:"time" json:"time"`
	// OrderItems are all of the items that are purchased or sold
	OrderItems []*trade_record.TradeRecord `csv:"items" avro:"items" json:"items"`
}

func NewTrade(t time.Time, items ...*trade_record.TradeRecord) *Trade {
	output := &Trade{
		UnixTime:   t.Unix(),
		OrderItems: make([]*trade_record.TradeRecord, 0, len(items)),
	}
	output.OrderItems = append(output.OrderItems, items...)
	return output
}

func (s *Trade) Append(item *trade_record.TradeRecord) {
	s.OrderItems = append(s.OrderItems, item)
}

func (s *Trade) AddItem(index int, item *trade_record.TradeRecord) {
	items := make([]*trade_record.TradeRecord, 0, len(s.OrderItems)+1)
	items = append(items, s.OrderItems[0:index]...)
	items = append(items, item)
	if index < len(s.OrderItems) {
		items = append(items, s.OrderItems[index:]...)
	}
	s.OrderItems = items
}

func (s *Trade) RemoveItem(index int) {
	items := make([]*trade_record.TradeRecord, 0, len(s.OrderItems)-1)
	items = append(items, s.OrderItems[0:index]...)
	if index < len(s.OrderItems) {
		items = append(items, s.OrderItems[index+1:]...)
	}
	s.OrderItems = items
}

func (s *Trade) Clone() *Trade {
	return NewTrade(
		time.Unix(s.UnixTime, 0),
		s.OrderItems...,
	)
}

