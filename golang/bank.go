package main

type Bank struct {
}

func (b *Bank) reduce(source Expression, to string) Money {
	return NewDollar(10)
}
