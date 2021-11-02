package _004_Median_Of_Two_Sorted_Arrays

import (
	"fmt"
	"testing"
)

func TestGo(t *testing.T) {
	a := []int{1, 3}
	b := []int{2, 4}
	fmt.Println(findMedianSortedArrays(a, b))
}
