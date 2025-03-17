package ArraysnSlices

import (
	"testing"
)

func TestSum(t *testing.T){
	// nums:=[5]int{1,2,3,4,5}
	// got:=Sum(nums)
	// want:=15
	// if got != want {
	// 	t.Errorf("got %d want %d given, %v", got, want, nums) //%v is used to print the array
	// }
	t.Run("collection of 5 nums", func(t *testing.T){
		nums:=[]int{1,2,3,4,5}
		got:=Sum(nums)
		want:=15
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, nums)
		}
	})
	t.Run("collection of any size", func(t *testing.T){
		nums:=[]int{1,2,3}
		got:=Sum(nums)
		want:=6
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, nums)
		}
	})
}
