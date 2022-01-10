package main

type Expression interface {
	Reduce(to string) Money
}
