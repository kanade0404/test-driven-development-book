package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiplication(t *testing.T) {
	t.Run("$5 * 2 = $10", func(t *testing.T) {
		five := NewDollar(5)
		assert.Equal(t, NewDollar(10), five.Times(2))
		assert.Equal(t, NewDollar(15), five.Times(3))
	})
}
func TestEquality(t *testing.T) {
	t.Run("Dollar values is equality", func(t *testing.T) {
		dollar := NewDollar(5)
		franc := NewFranc(5)
		assert.True(t, dollar.Equals(NewDollar(5)))
		assert.False(t, dollar.Equals(NewDollar(6)))
		assert.False(t, franc.Equals(NewDollar(5)))
	})
}

func TestFrancMultiplication(t *testing.T) {
	t.Run("$5 * 2 = $10", func(t *testing.T) {
		five := NewFranc(5)
		assert.Equal(t, NewFranc(10), five.Times(2))
		assert.Equal(t, NewFranc(15), five.Times(3))
	})
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
		t.Run(fmt.Sprintf("%s is equal %s", td.actualLabel, td.expectedLabel), func(t *testing.T) {
			assert.Equal(t, td.expected, td.actual)
		})
	}
}
