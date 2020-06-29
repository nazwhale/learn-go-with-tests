package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got: %q\nwant: %q", got, want)
		}
	}

	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("Sue")
		want := "Hello, Sue"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Default to 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}
