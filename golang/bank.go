package main

type Bank struct {
	rates map[Pair]int
}

func NewBank() Bank {
	b := Bank{}
	b.rates = make(map[Pair]int)
	return b
}

func (b *Bank) Reduce(source Expression, to string) Money {
	return source.Reduce(*b, to)
}

func (b *Bank) AddRate(from string, to string, rate int) {
	b.rates[Pair{from, to}] = rate
}

func (b *Bank) Rate(from string, to string) int {
	if from == to {
		return 1
	}
	p := Pair{from, to}
	return b.rates[p]
}
