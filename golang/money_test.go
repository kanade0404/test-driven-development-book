package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiplication(t *testing.T) {
	t.Run("$5 * 2 = $10", func(t *testing.T) {
		five := Dollar{5}
		product := five.times(2)
		assert.Equal(t, 10, product.amount)
		product = five.times(3)
		assert.Equal(t, 15, product.amount)
	})
}
func TestEquality(t *testing.T) {
	t.Run("Dollar values is equality", func(t *testing.T) {
		dollar := Dollar{5}
		assert.True(t, dollar.equals(Dollar{5}))
		assert.False(t, dollar.equals(Dollar{6}))
	})
}
