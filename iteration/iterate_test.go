package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("repeats", func(t *testing.T) {
		got := Repeat("a", 9)
		want := "aaaaaaaaa"

		if got != want {
			t.Errorf("\ngot:  %v\nwant: %v", got, want)
		}
	})

	t.Run("repeats n times", func(t *testing.T) {
		got := Repeat("b", 5)
		want := "bbbbb"

		if got != want {
			t.Errorf("\ngot:  %v\nwant: %v", got, want)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("s", 4)
	fmt.Println(repeated)
	// Output: ssss
}
