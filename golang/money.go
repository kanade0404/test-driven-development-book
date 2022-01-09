package main

type Dollar struct {
	amount int
	//Dollar (amount int) {}
}

func (d *Dollar) times(multiplier int) Dollar {
	return Dollar{d.amount * multiplier}
}

func (d *Dollar) equals(object interface{}) bool {
	dollar := object.(Dollar)
	return d.amount == dollar.amount
}
