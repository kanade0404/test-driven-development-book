package main

type Dollar struct {
	Money
}

func NewDollar(amount int) Money {
	return Money{amount: amount}
}

func (d *Dollar) Times(multiplier int) Money {
	return Money{d.amount * multiplier}
}
