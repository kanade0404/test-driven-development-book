package main

var _ Expression = (*Sum)(nil)

type Sum struct {
	augend Expression
	addend Expression
}

func NewSum(augend Expression, addend Expression) Sum {
	return Sum{augend, addend}
}

func (s Sum) Reduce(bank Bank, to string) Money {
	amount := s.augend.Reduce(bank, to).amount + s.addend.Reduce(bank, to).amount
	return NewMoney(amount, to)
}

func (s Sum) Plus(addend Expression) Expression {
	//TODO implement me
	panic("implement me")
}
