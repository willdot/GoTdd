package arraysslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Array of 5 numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := SumArray(numbers)
		want := 15

		assertArray(t, got, want, numbers)
	})

	t.Run("Array of 5 numbers but using the range function", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := SumArrayUsingRange(numbers)
		want := 15

		assertArray(t, got, want, numbers)
	})

	t.Run("Slice of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := SumSlice(numbers)
		want := 6

		assertSlice(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {

	t.Run("2 slices each with 2 numbers", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})

		want := []int{3, 9}

		assertDeepEqual(t, got, want)
	})

	t.Run("2 slices each with 2 numbers, refactored code", func(t *testing.T) {
		got := SumAllRefactored([]int{1, 2}, []int{0, 9})

		want := []int{3, 9}

		assertDeepEqual(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {

	t.Run("makes the sums of some slices", func(t *testing.T){
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9})

		want := []int{5, 9}

		assertDeepEqual(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T){
		got := SumAllTails([]int{}, []int{3, 4, 5})

		want := []int{0, 9}

		assertDeepEqual(t, got, want)
	})

}

func assertArray(t *testing.T, expected int, actual int, numbers [5]int) {
	if expected != actual {
		t.Errorf("got %d want %d given, %v", actual, expected, numbers)
	}
}

func assertSlice(t *testing.T, expected int, actual int, numbers []int) {
	if expected != actual {
		t.Errorf("got %d want %d given, %v", actual, expected, numbers)
	}
}

func assertDeepEqual(t *testing.T, expected []int, actual []int) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("got %v want %v", expected, actual)
	}
}
