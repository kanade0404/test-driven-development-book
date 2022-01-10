package main

type Expression interface {
	Plus(addend Expression) Expression
	Reduce(bank Bank, to string) Money
	Times(multiplier int) Expression
}
