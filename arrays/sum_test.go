package main

import (
	"testing"
	"reflect"
)

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	expected := 15

	if got != expected {
		t.Errorf("got %d expected %d given the numbers %v", got, expected, numbers )
	}
}

func TestSumOfSlice(t *testing.T) {
	numbers := []int {1, 2, 3, 4, 5, 6}

	got := SumOfSlice(numbers)
	expected := 21

	if got != expected {
		t.Errorf("got %d expected %d given the numbers %v", got, expected, numbers )
	} 
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1 ,2, 3}, []int{1, 1, 1}, []int{4, 5, 6})
	expected := []int{6, 3, 15}

	if ! reflect.DeepEqual(got, expected) {
		t.Errorf("got %v expected %v", got, expected )
	}
}
 

func TestSumAllTails(t *testing.T) {

	checkSums := func (t testing.TB, got, want []int)  {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sum of some slices", func( t *testing.T) {
		got := SumAllTails([]int{1 ,2, 3}, []int{1, 1, 1}, []int{4, 5, 6})
		expected := []int{3, 15}

		checkSums(t, got, expected)
	})

	t.Run("sum empty slice without error", func( t *testing.T) {
		got := SumAllTails([]int{}, []int{}, []int{4, 5, 6})
		expected := []int{0, 15}

		checkSums(t, got, expected)
	})
	
	t.Run("sum empty slice", func( t *testing.T) {
		got := SumAllTails([]int{})
		expected := []int{0}

		checkSums(t, got, expected)
	})
	
}