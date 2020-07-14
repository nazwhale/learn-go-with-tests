package slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("size: any", func(t *testing.T) {
		numbers := []int{4, 5, 6, 7, 8}

		got := Sum(numbers)
		want := 30

		if got != want {
			t.Errorf("\ngot:  %d\nwant: %d\ninput: %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("one slice", func(t *testing.T) {
		slice1 := []int{4, 5, 6, 7, 8}

		got := SumAll(slice1)
		want := []int{30}

		if got[0] != want[0] {
			t.Errorf("\ngot:  %v\nwant: %v\ninput: %v", got, want, slice1)
		}
	})

	t.Run("two slices", func(t *testing.T) {
		slice1 := []int{4, 5, 6, 7, 8}
		slice2 := []int{1, 2, 3, 4}

		got := SumAll(slice1, slice2)
		want := []int{30, 10}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("\ngot:  %v\nwant: %v\ninput: %v", got, want, slice1)
		}
	})

	t.Run("four slices", func(t *testing.T) {
		slice1 := []int{4, 5, 6, 7, 8}
		slice2 := []int{1, 2, 3, 4}
		slice3 := []int{3, 4, 6, 3, 12}
		slice4 := []int{4, 6, 3}

		got := SumAll(slice1, slice2, slice3, slice4)
		want := []int{30, 10, 28, 13}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("\ngot:  %v\nwant: %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("\ngot:  %v\nwant: %v\n", got, want)
		}
	}

	t.Run("1 head", func(t *testing.T) {
		slice1 := []int{4, 5, 6, 7, 8}
		slice2 := []int{1, 2, 3, 4}

		got := SumAllTails(1, slice1, slice2)
		want := []int{26, 9}

		checkSums(t, got, want)
	})

	t.Run("n tails", func(t *testing.T) {
		slice1 := []int{4, 5, 6, 7, 8}
		slice2 := []int{1, 2, 3, 4}

		got := SumAllTails(2, slice1, slice2)
		want := []int{21, 7}

		checkSums(t, got, want)
	})

	t.Run("n tail smaller than heads", func(t *testing.T) {
		slice1 := []int{4, 5, 6, 7, 8}
		slice2 := []int{1}
		slice3 := []int{1, 2, 3, 4}

		got := SumAllTails(2, slice1, slice2, slice3)
		want := []int{21, 0, 7}

		checkSums(t, got, want)
	})

	t.Run("empty slice", func(t *testing.T) {
		slice1 := []int{}
		slice2 := []int{1}
		slice3 := []int{1, 2, 3, 4}

		got := SumAllTails(1, slice1, slice2, slice3)
		want := []int{0, 0, 9}

		checkSums(t, got, want)
	})
}
