package main

type Pair struct {
	from string
	to   string
}

func (p *Pair) Equals(object interface{}) bool {
	pair := object.(Pair)
	return p.from == pair.from && p.to == pair.to
}

func (p *Pair) HashCode() int {
	return 0
}
