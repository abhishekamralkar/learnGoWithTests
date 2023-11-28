package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSumArray(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := SumArray(numbers)
	want := 15

	if want != got {
		t.Errorf("Expected %d but got %d: ", want, got)
	}
}

func TestSumSlice(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := SumSlice(numbers)
	want := 15

	if want != got {
		t.Errorf("Expected %d but got %d: ", want, got)
	}
}

func TestSum(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := SumArray(numbers)
		want := 15

		if want != got {
			t.Errorf("Expected %d but want %d", want, got)
		}
	})

	t.Run("Collection if any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := SumSlice(numbers)
		want := 15

		if want != got {
			t.Errorf("Expected %d but want %d", want, got)
		}
	})

}

func TestAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v but got %v", want, got)
	}
}
