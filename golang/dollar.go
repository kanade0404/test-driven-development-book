package main

type Dollar struct {
	Money
}

func NewDollar(amount int) Money {
	return NewMoney(amount, "USD")
}

func (d *Dollar) NewDollar(multiplier int) Money {
	return NewDollar(d.amount * multiplier)
}
