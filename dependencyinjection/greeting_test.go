package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Naz")

	got := buffer.String()
	want := "Hello, Naz"

	if got != want {
		t.Errorf("\ngot:  %q\nwant: %q", got, want)
	}
}

func TestServe(t *testing.T) {
	main()
}
