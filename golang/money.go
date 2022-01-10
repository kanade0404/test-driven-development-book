package main

import "fmt"

var _ Moneyer = (*Money)(nil)
var _ Expression = (*Money)(nil)

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

func NewDollar(amount int) Money {
	return NewMoney(amount, "USD")
}

func NewFranc(amount int) Money {
	return NewMoney(amount, "CHF")
}

func (m Money) Currency() string {
	return m.currency
}

func (m Money) Equals(object interface{}) bool {
	money := object.(Money)
	return money.amount == m.amount && money.currency == m.currency
}

func (m Money) Plus(addend Money) Expression {
	return NewSum(NewMoney(m.amount, m.currency), addend)
}
func (m Money) Reduce(bank Bank, to string) Money {
	rate := bank.Rate(m.Currency(), to)
	return NewMoney(m.amount/rate, to)
}

func (m Money) Times(multiplier int) Money {
	return NewMoney(m.amount*multiplier, m.currency)
}
func (m Money) String() string {
	return fmt.Sprintf("Amount: %d, Currency: %s", m.amount, m.currency)
}
