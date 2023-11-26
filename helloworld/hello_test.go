package main

import "testing"

func TestHello(t *testing.T) {
	got := "Hello World!"
	want := "Hello World!"

	if got != want {
		t.Errorf("Expected %s but got %s: ", got, want)
	}
}
