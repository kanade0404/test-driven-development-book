package main

var _ Moneyer = (*Money)(nil)

type Moneyer interface {
	Equals(object interface{}) bool
	Times(multiplier int) Money
	Currency() string
}

type Money struct {
	amount   int
	currency string
}

func NewMoney(amount int, currency string) Money {
	return Money{amount, currency}
}

func (m *Money) Currency() string {
	return m.currency
}

func (m *Money) Equals(object interface{}) bool {
	money := object.(Money)
	return money.amount == m.amount && money.currency == m.currency
}

func (m *Money) Times(multiplier int) Money {
	return Money{m.amount * multiplier, m.currency}
}
