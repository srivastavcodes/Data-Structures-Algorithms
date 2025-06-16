package main

// similar 739,

type StockStack struct {
	price int
	span  int
}

type StockSpanner struct {
	stack []StockStack
}

func NewStockSpanner() StockSpanner {
	return StockSpanner{}
}

func (sp *StockSpanner) Next(price int) int {
	top, span := len(sp.stack)-1, 1

	for len(sp.stack) > 0 && sp.stack[top].price <= price {
		span = span + sp.stack[top].span
		sp.stack = sp.stack[:top]

		top = len(sp.stack) - 1
	}
	sp.stack = append(sp.stack, StockStack{price, span})
	return span
}
