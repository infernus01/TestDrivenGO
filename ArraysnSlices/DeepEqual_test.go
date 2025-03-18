package ArraysnSlices

import (
	"reflect"
	"testing"
)

// func TestSumAll(t *testing.T) {
// 	got := SumAll([]int{1, 2}, []int{0, 9})
// 	want := []int{3, 9}

// 	if !reflect.DeepEqual(got, want) { //!= will not work here as we are comparing slices, it will throw error
// 		t.Errorf("got %v want %v", got, want)
// 	}
// }

func TestSumAllTails(t *testing.T) {
	t.Run("sum the tails of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}