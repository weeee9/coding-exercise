package onlinestockspan

type StockSpanner struct {
	days  int
	stack []stock
}

type stock struct {
	price int
	day   int
}

func Constructor() StockSpanner {
	return StockSpanner{
		stack: []stock{},
	}
}

func (this *StockSpanner) Next(price int) int {
	this.days++

	for len(this.stack) > 0 {
		peak := this.stack[len(this.stack)-1]

		if peak.price > price {
			break
		}

		this.stack = this.stack[:len(this.stack)-1]
	}

	span := 0
	if len(this.stack) > 0 {
		peak := this.stack[len(this.stack)-1]
		span = this.days - peak.day
	} else {
		span = this.days
	}

	this.stack = append(this.stack, stock{price, this.days})

	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
