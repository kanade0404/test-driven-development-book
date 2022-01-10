package main

import (
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
		assert.True(t, dollar.Equals(NewDollar(5)))
		assert.False(t, dollar.Equals(NewDollar(6)))
	})
}

func TestFrancMultiplication(t *testing.T) {
	t.Run("$5 * 2 = $10", func(t *testing.T) {
		five := NewFranc(5)
		assert.Equal(t, NewFranc(10), five.Times(2))
		assert.Equal(t, NewFranc(15), five.Times(3))
	})
}
