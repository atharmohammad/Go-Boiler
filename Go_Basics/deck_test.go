package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected 52 but got %v", len(d))
	}
}

func TestNewDeckFromFile(t *testing.T) {
	os.Remove("_testfile")
	d := newDeck()
	d.save("_testing")
	new_d := read("_testing")
	if len(new_d) != len(d) {
		t.Errorf("deck is not correctly saved")
	}
	os.Remove("_testfile")
}
