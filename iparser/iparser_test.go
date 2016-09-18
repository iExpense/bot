package iparser

import (
	"log"
  "testing"
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
}
