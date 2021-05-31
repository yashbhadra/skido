package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	_ "github.com/google/uuid"
)

func main() {
	//imported uuid package for displaying go mod usage
	exec := []int{10, 203, 31, 22, 34, 3, 320, 90, 190, 25, 1, 23, 13, 11239, 123, 32, 90, 123, 12, 3, 4313, 3, 999, 999, 90}

	//Executing the PartA - Single core - Single go routine
	TimeTakenA, err := PartA(exec)
	if err == nil {
		fmt.Println("T.A.T for Part A:", TimeTakenA)
	} else {
		fmt.Println("A problem occured while executing scripts", err)
		return
	}

	//Increasing the number cores that can be used by go routines
	runtime.GOMAXPROCS(runtime.NumCPU())

	//Executing the Part B - Multiple cores - Multiple go routines
	TimeTakenB, err := PartB(exec)
	if err == nil {
		fmt.Println("T.A.T for Part B:", TimeTakenB)
	} else {
		fmt.Println("A problem occured while executing scripts", err)
		return
	}

}
func PartA(exec []int) (time.Duration, error) {
	if len(exec) < 25 {
		return time.Millisecond * 0, fmt.Errorf("Not enough scripts to run the game")
	}
	start1 := time.Now()
	for i := 0; i < len(exec); i++ {
		time.Sleep(time.Duration(exec[i]) * time.Millisecond)
	}
	return time.Since(start1), nil
}
func PartB(exec []int) (time.Duration, error) {
	if len(exec) < 25 {
		return time.Millisecond * 0, fmt.Errorf("Not enough scripts to run the game")
	}
	start2 := time.Now()
	var wg sync.WaitGroup = sync.WaitGroup{}
	wg.Add(len(exec))
	for i := 0; i < len(exec); i++ {
		go Execute(exec[i], &wg)

	}
	wg.Wait()
	return time.Since(start2), nil
}
func Execute(script int, wg *sync.WaitGroup) {
	time.Sleep(time.Duration(script) * time.Millisecond)
	wg.Done()
}
