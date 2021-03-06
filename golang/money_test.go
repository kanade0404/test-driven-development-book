package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	for _, td := range []struct {
		expected int
		multiply int
		by       int
	}{
		{
			expected: 10,
			multiply: five.amount,
			by:       2,
		},
		{
			expected: 15,
			multiply: five.amount,
			by:       3,
		},
	} {
		t.Run(fmt.Sprintf("$%d * %d = $%d", td.multiply, td.by, td.expected), func(t *testing.T) {
			assert.Equal(t, NewDollar(td.expected), five.Times(td.by))
		})
	}
}
func TestEquality(t *testing.T) {
	dollar := NewDollar(5)
	franc := NewFranc(5)
	for _, td := range []struct {
		compare  string
		compared string
		expected bool
		actual   bool
	}{
		{
			compare:  "5 dollar",
			compared: "5 dollar",
			expected: true,
			actual:   dollar.Equals(NewDollar(5)),
		},
		{
			compare:  "5 dollar",
			compared: "6 dollar",
			expected: false,
			actual:   dollar.Equals(NewDollar(6)),
		},
		{
			compare:  "5 franc",
			compared: "5 dollar",
			expected: false,
			actual:   franc.Equals(NewDollar(5)),
		},
	} {
		t.Run(fmt.Sprintf("(%s equals %s) is %t", td.compare, td.compared, td.expected), func(t *testing.T) {
			if td.expected {
				assert.True(t, td.expected, td.actual)
			} else {
				assert.False(t, td.expected, td.actual)
			}
		})
	}
}

func TestCurrency(t *testing.T) {
	dollar := NewDollar(1)
	franc := NewFranc(1)
	for _, td := range []struct {
		expected      string
		expectedLabel string
		actual        string
		actualLabel   string
	}{
		{
			expected: "USD",
			actual:   dollar.Currency(),
		},
		{
			expected: "CHF",
			actual:   franc.Currency(),
		},
	} {
		t.Run(fmt.Sprintf("%s is equals %s", td.actualLabel, td.expectedLabel), func(t *testing.T) {
			assert.Equal(t, td.expected, td.actual)
		})
	}
}

func TestSimpleAddition(t *testing.T) {
	five := NewDollar(5)
	sum := five.Plus(NewDollar(5))
	bank := Bank{}
	reduce := bank.Reduce(sum, "USD")
	assert.Equal(t, NewDollar(10), reduce)
}

func TestPlusReturnSum(t *testing.T) {
	five := NewDollar(5)
	result := five.Plus(five)
	sum := result.(Sum)
	for _, td := range []struct {
		expected Money
		actual   Expression
	}{
		{
			expected: five,
			actual:   sum.augend,
		},
		{
			expected: five,
			actual:   sum.addend,
		},
	} {
		t.Run(fmt.Sprintf("$5 + $5 returns %v", td.expected), func(t *testing.T) {
			assert.Equal(t, td.expected, td.actual)
		})
	}
}

func TestReduceSum(t *testing.T) {
	sum := NewSum(NewDollar(3), NewDollar(4))
	bank := Bank{}
	result := bank.Reduce(sum, "USD")
	assert.Equal(t, NewDollar(7), result)
}

func TestReduceMoney(t *testing.T) {
	bank := Bank{}
	dollar := NewDollar(1)
	result := bank.Reduce(dollar, "USD")
	assert.Equal(t, NewDollar(1), result)
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(NewFranc(2), "USD")
	for _, td := range []struct {
		expected Money
		actual   Money
	}{
		{
			NewDollar(1),
			result,
		},
	} {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, td.expected, td.actual)
		})
	}
}

func TestIdentityRate(t *testing.T) {
	bank := Bank{}
	assert.Equal(t, 1, bank.Rate("USD", "USD"))
}

func TestMixedAddition(t *testing.T) {
	fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(fiveBucks.Plus(tenFrancs), "USD")
	assert.Equal(t, NewDollar(10), result)
}

func TestSumPlusMoney(t *testing.T) {
	fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	sum := NewSum(fiveBucks, tenFrancs).Plus(fiveBucks)
	result := bank.Reduce(sum, "USD")
	assert.Equal(t, NewDollar(15), result)
}

func TestSumTimes(t *testing.T) {
	fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	sum := NewSum(fiveBucks, tenFrancs).Times(2)
	result := bank.Reduce(sum, "USD")
	assert.Equal(t, NewDollar(20), result)
}
