package main

import "testing"

func TestAdd(t *testing.T) {
	got := add(5, 5)
    want := 10
    if got != want {
       t.Errorf("Sum was error got: %d want %d!", got, want)
    }
}
