package iexpense

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMoney(t *testing.T) {
	money, err := Parse("100.1")
	assert.Equal(t, money, &Money{
		val:   100,
		cents: 1,
	})
	assert.Nil(t, err)

	money, err = Parse("100.12")
	assert.Equal(t, money, &Money{
		val:   100,
		cents: 12,
	})
	assert.Nil(t, err)

	money, err = Parse("10000.165")
	assert.Equal(t, money, &Money{
		val:   10000,
		cents: 16,
	})
	assert.Nil(t, err)

	money, err = Parse("10000.1655")
	assert.Equal(t, money, &Money{
		val:   10000,
		cents: 16,
	})
	assert.Nil(t, err)

	money, err = Parse("10000")
	assert.Equal(t, money, &Money{
		val:   10000,
		cents: 0,
	})
	assert.Nil(t, err)

	money, err = Parse("10000.00000")
	assert.Equal(t, money, &Money{
		val:   10000,
		cents: 0,
	})
	assert.Nil(t, err)

	money, err = Parse("-100.12")
	assert.Equal(t, money, &Money{
		val:   -100,
		cents: -12,
	})
	assert.Nil(t, err)

	money, err = Parse("-100")
	assert.Equal(t, money, &Money{
		val:   -100,
		cents: 0,
	})
	assert.Nil(t, err)

	money, err = Parse("10000.")
	assert.Nil(t, money)
	assert.EqualError(t, err, ErrorInvalidAmount.Error())

	money, err = Parse("10000.0.0")
	assert.Nil(t, money)
	assert.EqualError(t, err, ErrorInvalidAmount.Error())

	money, err = Parse("#sbi")
	assert.Nil(t, money)
	assert.EqualError(t, err, ErrorInvalidAmount.Error())

	money, err = Parse("")
	assert.Nil(t, money)
	assert.EqualError(t, err, ErrorInvalidAmount.Error())
}

func TestMathWithMoney(t *testing.T) {
	m1, err := Parse("10.12")
	assert.Nil(t, err)
	m2, err := Parse("4.6")
	assert.Nil(t, err)
	assert.Equal(t, m1.Add(m2), &Money{
		val:   14,
		cents: 18,
	})

	m1, err = Parse("10.56")
	assert.Nil(t, err)
	m2, err = Parse("4.98")
	assert.Nil(t, err)
	assert.Equal(t, m1.Add(m2), &Money{
		val:   15,
		cents: 54,
	})

	m1, err = Parse("10.56")
	assert.Nil(t, err)
	m2, err = Parse("-4.98")
	assert.Nil(t, err)
	assert.Equal(t, m1.Add(m2), &Money{
		val:   5,
		cents: 58,
	})

	m1, err = Parse("10.98")
	assert.Nil(t, err)
	m2, err = Parse("-4.56")
	assert.Nil(t, err)
	assert.Equal(t, m1.Add(m2), &Money{
		val:   6,
		cents: 42,
	})

	m1, err = Parse("10.98")
	assert.Nil(t, err)
	m2, err = Parse("4.56")
	assert.Nil(t, err)
	assert.Equal(t, m1.Sub(m2), &Money{
		val:   6,
		cents: 42,
	})

	m1, err = Parse("10.56")
	assert.Nil(t, err)
	m2, err = Parse("4.98")
	assert.Nil(t, err)
	assert.Equal(t, m1.Sub(m2), &Money{
		val:   5,
		cents: 58,
	})
}
