package main

type Dollar struct {
	amount int
	//Dollar (amount int) {}
}

func (d *Dollar) times(multiplier int) {
	d.amount *= multiplier
}
