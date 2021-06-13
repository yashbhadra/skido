package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//imported uuid package for displaying go mod usage
	exec := []int{10, 203, 31, 22, 34, 3, 320, 90, 190, 25, 1, 23, 13, 11239, 123, 32, 90, 123, 12, 3, 4313, 3, 999, 999, 90}

	//Executing the PartA - Single core - Single go routine
	/*TimeTakenA, err := PartA(exec)
	if err == nil {
		fmt.Println("T.A.T for Part A:", TimeTakenA)
	} else {
		fmt.Println("A problem occured while executing scripts", err)
		return
	}*/

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
	wg := sync.WaitGroup{}
	wc := sync.WaitGroup{}

	// Change the limit to control the number of goroutines in workerpool.
	limit := 10
	//The channel buffer limit is also set to 10
	queue := make(chan int, 10)
	scripts := make(chan int, 10)
	wg.Add(limit)
	wc.Add(limit)
	//Creating a producer workerpool
	for i := 0; i < limit; i++ {
		go Producer(scripts, queue, &wg)
	}

	//supplying jobs to worker pool
	for i := 0; i < len(exec); i++ {
		scripts <- exec[i]
	}
	close(scripts)

	//Creating a consumer workerpool
	for i := 0; i < limit; i++ {
		go Consumer(queue, &wc)
	}

	wg.Wait()
	close(queue)
	wc.Wait()
	return time.Since(start2), nil
}
func Producer(scripts chan int, queue chan int, wg *sync.WaitGroup) {
	//If the number of script items go above the limit then it will get blocked here
	for script := range scripts {
		queue <- script
	}
	wg.Done()

}
func Consumer(queue chan int, wc *sync.WaitGroup) {
	//All the goroutines wait here to consume the script items.
	for script := range queue {
		time.Sleep(time.Duration(script) * time.Millisecond)
	}
	wc.Done()

}
