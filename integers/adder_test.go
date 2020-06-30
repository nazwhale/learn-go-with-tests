package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(2, 2)
	actual := 4

	if got != actual {
		t.Errorf("\ngot:  %v\nactual: %v", got, actual)
	}
}

func ExampleAdd() {
	sum := Add(6, 1)
	fmt.Println(sum)
	// Output: 7
}
