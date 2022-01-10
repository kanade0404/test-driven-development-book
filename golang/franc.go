package main

type Franc struct {
	Money
}

func NewFranc(amount int) Money {
	return Money{amount: amount}
}
func (f *Franc) Times(multiplier int) Money {
	return Money{f.amount * multiplier}
}
