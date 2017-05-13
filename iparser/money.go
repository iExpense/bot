package iparser

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Money struct {
	val   int64
	cents int64
}

var (
	// Error Codes
	ErrorInvalidAmount = errors.New("Invalid amount specified")
)

func NewMoney(str string) (*Money, error) {
	var val, cents int64
	var valStr, centsStr string
	var negativeAmount bool
	var err error

	if len(str) <= 0 {
		return nil, ErrorInvalidAmount
	}

	if strings.HasPrefix(str, "-") {
		str = strings.TrimPrefix(str, "-")
		negativeAmount = true
	}

	if strings.Contains(str, ".") {
		tokens := strings.Split(str, ".")
		if len(tokens) != 2 {
			return nil, ErrorInvalidAmount
		}

		valStr = tokens[0]
		if len(tokens[1]) > 2 {
			for i, c := range tokens[1] {
				if i >= 2 {
					break
				}

				centsStr += string(c)
			}
		} else {
			centsStr = tokens[1]
		}
	} else {
		centsStr = "0"
		valStr = str
	}

	val, err = strconv.ParseInt(valStr, 10, 64)
	if err != nil {
		log.Printf("[WARN] Unable to parse %s (amount: %s)\n", valStr, str)
		return nil, ErrorInvalidAmount
	}

	cents, err = strconv.ParseInt(centsStr, 10, 64)
	if err != nil {
		log.Printf("[WARN] Unable to parse %s (amount: %s)\n", centsStr, str)
		return nil, ErrorInvalidAmount
	}

	// sanity check
	if cents >= 100 {
		panic("cents cannot be bigger than 100!")
	}

	if negativeAmount {
		cents = 0 - cents
		val = 0 - val
	}

	return &Money{
		val:   val,
		cents: cents,
	}, nil
}

func (m Money) String() string {
	return fmt.Sprintf("%d.%d", m.val, m.cents)
}

func (m *Money) ensure() {
	if m.cents >= 100 {
		m.val += 1
		m.cents -= 100
	}

	if m.cents < 0 && m.val > 0 {
		m.val -= 1
		m.cents += 100
	}
}

func (m1 *Money) Add(m2 *Money) *Money {
	m1.val += m2.val
	m1.cents += m2.cents
	m1.ensure()
	return m1
}

func (m1 *Money) Sub(m2 *Money) *Money {
	m1.val -= m2.val
	m1.cents -= m2.cents
	m1.ensure()
	return m1
}
