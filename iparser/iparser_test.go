package iparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpense(t *testing.T) {
	cmd, err := Parse("expense 100 #sbi #cash")
	assert.Equal(t, cmd, &Command{
		Ctype:  Expense,
		Amount: &Money{100, 0},
		Tags:   []string{"sbi", "cash"},
	})
	assert.Nil(t, err)

	cmd, err = Parse("e 100 sbi #cash")
	assert.Equal(t, cmd, &Command{
		Ctype:  Expense,
		Amount: &Money{100, 0},
		Tags:   []string{"sbi", "cash"},
	})
	assert.Nil(t, err)

	cmd, err = Parse("exp 100 #sbi")
	assert.Equal(t, cmd, &Command{
		Ctype:  Expense,
		Amount: &Money{100, 0},
		Tags:   []string{"sbi"},
	})
	assert.Nil(t, err)

	cmd, err = Parse("100 #sbi")
	assert.Equal(t, cmd, &Command{
		Ctype:  Expense,
		Amount: &Money{100, 0},
		Tags:   []string{"sbi"},
	})
	assert.Nil(t, err)

	cmd, err = Parse("ex")
	assert.Nil(t, cmd)
	assert.EqualError(t, err, ErrorNeedArgs.Error())

	cmd, err = Parse("e 100")
	assert.Nil(t, cmd)
	assert.EqualError(t, err, ErrorNeedArgs.Error())

	cmd, err = Parse("e #sbi")
	assert.Nil(t, cmd)
	assert.EqualError(t, err, ErrorNeedArgs.Error())

	cmd, err = Parse("")
	assert.Nil(t, cmd)
	assert.EqualError(t, err, ErrorEmptyCmd.Error())
}

func TestIncome(t *testing.T) {
	cmd, err := Parse("income 100 #sbi #cash")
	assert.Equal(t, cmd, &Command{
		Ctype:  Income,
		Amount: &Money{100, 0},
		Tags:   []string{"sbi", "cash"},
	})
	assert.Nil(t, err)

	cmd, err = Parse("i 100 sbi #cash")
	assert.Equal(t, cmd, &Command{
		Ctype:  Income,
		Amount: &Money{100, 0},
		Tags:   []string{"sbi", "cash"},
	})
	assert.Nil(t, err)

	cmd, err = Parse("income 100 #sbi")
	assert.Equal(t, cmd, &Command{
		Ctype:  Income,
		Amount: &Money{100, 0},
		Tags:   []string{"sbi"},
	})
	assert.Nil(t, err)

	cmd, err = Parse("in")
	assert.Nil(t, cmd)
	assert.EqualError(t, err, ErrorNeedArgs.Error())

	cmd, err = Parse("inc 100")
	assert.Nil(t, cmd)
	assert.EqualError(t, err, ErrorNeedArgs.Error())

	cmd, err = Parse("i #sbi")
	assert.Nil(t, cmd)
	assert.EqualError(t, err, ErrorNeedArgs.Error())

	cmd, err = Parse("")
	assert.Nil(t, cmd)
	assert.EqualError(t, err, ErrorEmptyCmd.Error())
}

func TestTransfer(t *testing.T) {
	cmd, err := Parse("transfer 100 #sbi #cash")
	assert.Equal(t, cmd, &Command{
		Ctype:  Transfer,
		Amount: &Money{100, 0},
		Tags:   []string{"sbi", "cash"},
	})
	assert.Nil(t, err)

	cmd, err = Parse("tran 1 sbi cash")
	assert.Equal(t, cmd, &Command{
		Ctype:  Transfer,
		Amount: &Money{1, 0},
		Tags:   []string{"sbi", "cash"},
	})
	assert.Nil(t, err)

	cmd, err = Parse("t 001 sbi > cash")
	assert.Equal(t, cmd, &Command{
		Ctype:  Transfer,
		Amount: &Money{1, 0},
		Tags:   []string{"sbi", "cash"},
	})
	assert.Nil(t, err)

	cmd, err = Parse("t 100 #sbi > #cash")
	assert.Equal(t, cmd, &Command{
		Ctype:  Transfer,
		Amount: &Money{100, 0},
		Tags:   []string{"sbi", "cash"},
	})
	assert.Nil(t, err)

	cmd, err = Parse("t")
	assert.Nil(t, cmd)
	assert.EqualError(t, err, ErrorNeedArgs.Error())

	cmd, err = Parse("t 100")
	assert.Nil(t, cmd)
	assert.EqualError(t, err, ErrorNeedArgs.Error())

	cmd, err = Parse("t #sbi #cash")
	assert.Nil(t, cmd)
	assert.EqualError(t, err, ErrorNeedArgs.Error())

	cmd, err = Parse("t 100 #sbi")
	assert.Nil(t, cmd)
	assert.EqualError(t, err, ErrorNeedArgs.Error())

	cmd, err = Parse("t 100 #sbi >")
	assert.Nil(t, cmd)
	assert.EqualError(t, err, ErrorNeedArgs.Error())

	cmd, err = Parse("")
	assert.Nil(t, cmd)
	assert.EqualError(t, err, ErrorEmptyCmd.Error())
}

func TestHistory(t *testing.T) {
	cmd, err := Parse("h")
	assert.Equal(t, cmd, historyCommand)
	assert.Nil(t, err)

	cmd, err = Parse("history")
	assert.Equal(t, cmd, historyCommand)
	assert.Nil(t, err)

	cmd, err = Parse("hist")
	assert.Equal(t, cmd, historyCommand)
	assert.Nil(t, err)

	cmd, err = Parse("h m2d")
	assert.Equal(t, cmd, &Command{
		Ctype: History,
		Tags:  []string{"m2d"},
	})
	assert.Nil(t, err)

	cmd, err = Parse("h 10")
	assert.Equal(t, cmd, &Command{
		Ctype: History,
		Tags:  []string{"10"},
	})
	assert.Nil(t, err)
}

func TestBalance(t *testing.T) {
	cmd, err := Parse("balance")
	assert.Equal(t, cmd, balanceCommand)
	assert.Nil(t, err)

	cmd, err = Parse("b")
	assert.Equal(t, cmd, balanceCommand)
	assert.Nil(t, err)

	cmd, err = Parse("ba")
	assert.Equal(t, cmd, balanceCommand)
	assert.Nil(t, err)

	cmd, err = Parse("b #test #sbi #cash")
	assert.Equal(t, cmd, &Command{
		Ctype: Balance,
		Tags:  []string{"test", "sbi", "cash"},
	})
	assert.Nil(t, err)

	cmd, err = Parse("b #test sbi")
	assert.Equal(t, cmd, &Command{
		Ctype: Balance,
		Tags:  []string{"test", "sbi"},
	})
	assert.Nil(t, err)

	cmd, err = Parse("b #test")
	assert.Equal(t, cmd, &Command{
		Ctype: Balance,
		Tags:  []string{"test"},
	})
	assert.Nil(t, err)
}

func TestHelp(t *testing.T) {
	cmd, err := Parse("help")
	assert.Equal(t, cmd, helpCommand)
	assert.Nil(t, err)
}
