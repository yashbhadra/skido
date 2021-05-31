package main

import (
	"fmt"
	"testing"
)

func TestPartA(t *testing.T) {
	exec := []int{10, 203, 31, 22, 34, 3, 320, 90, 190, 25, 1, 23, 13, 11239, 123, 32, 90, 123, 12, 3, 4313, 3, 999, 999, 90}
	time, _ := PartA(exec)
	if int64(time) > 20000000000 {
		t.Error("Time Taken to execute the scripts was above the threshold")
	}
}
func TestPartB(t *testing.T) {
	exec := []int{10, 203, 31, 22, 34, 3, 320, 90, 190, 25, 1, 23, 13, 11239, 123, 32, 90, 123, 12, 3, 4313, 3, 999, 999, 90}
	time, _ := PartB(exec)
	if int64(time) > 15000000000 {
		t.Error("Time Taken to execute the scripts was above the threshold")

	}

}
func TestPartAScripts(t *testing.T) {
	exec := []int{}
	_, err := PartA(exec)
	if err == nil {
		fmt.Errorf("Execution with less scripts did not result in error")
	}
}
func TestPartBScripts(t *testing.T) {
	exec := []int{}
	_, err := PartB(exec)
	if err == nil {
		fmt.Errorf("Execution with less scripts did not result in error")
	}

}
func BenchmarkPartA(b *testing.B) {
	exec := []int{10, 203, 31, 22, 34, 3, 320, 90, 190, 25, 1, 23, 13, 11239, 123, 32, 90, 123, 12, 3, 4313, 3, 999, 999, 90}
	for i := 0; i < b.N; i++ {
		PartA(exec)
	}
}
func BenchmarkPartB(b *testing.B) {
	exec := []int{10, 203, 31, 22, 34, 3, 320, 90, 190, 25, 1, 23, 13, 11239, 123, 32, 90, 123, 12, 3, 4313, 3, 999, 999, 90}
	for i := 0; i < b.N; i++ {
		PartB(exec)
	}
}
