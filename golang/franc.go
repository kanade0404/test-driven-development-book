package main

type Franc struct {
	amount int
}

func (f *Franc) times(multiplier int) Franc {
	return Franc{f.amount * multiplier}
}

func (f *Franc) equals(object interface{}) bool {
	franc := object.(Franc)
	return f.amount == franc.amount
}
