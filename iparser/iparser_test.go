package iparser

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	cCommandNil = (*Command)(nil)
)

func check(message string, err error, t *testing.T) {
	if err != nil {
		t.Log(message)
		t.Fatal(err)
	}
}

func setup(t *testing.T) {
	log.Println("setup done!")
}

func tear(t *testing.T) {
	log.Println("teared down!")
}

func TestTest(t *testing.T) {
	setup(t)
	defer tear(t)

	cmd, err := Parse("transfer 100 #sbi #cash")
	check("parse returned error", err, t)
	assert.Equal(t, cmd, cCommandNil)
}
