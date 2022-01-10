package main

var _ Moneyer = (*Money)(nil)

type Moneyer interface {
	Equals(object interface{}) bool
	Times(multiplier int) Money
}

type Money struct {
	amount int
}

func (m *Money) Equals(object interface{}) bool {
	money := object.(Money)
	return money.amount == m.amount
}

func (m *Money) Times(multiplier int) Money {
	return Money{m.amount * multiplier}
}
