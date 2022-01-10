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
