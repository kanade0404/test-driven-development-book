package main

type Franc struct {
	Money
}

func NewFranc(amount int) Money {
	return NewMoney(amount, "CHF")
}
func (f *Franc) Times(multiplier int) Money {
	return NewFranc(f.amount * multiplier)
}
