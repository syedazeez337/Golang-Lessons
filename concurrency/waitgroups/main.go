package main

import (
	"fmt"
	"sync"
)

func funcOne(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Goroutine 1")
}

func funcTwo(wg *sync.WaitGroup) {
	fmt.Println("Goroutine 2")
	wg.Done()
}

func funcThree(wg *sync.WaitGroup) {
	fmt.Println("Goroutine 3")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go funcOne(&wg)
	go funcTwo(&wg)
	go funcThree(&wg)

	wg.Wait()

	fmt.Println("Main function")
}

/*

// waitgroups in sync package
var wg sync.WaitGroup

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(2 * time.Second)
	fmt.Printf("Worker %d done\n", id)
	wg.Done()
}

func main() {

	// wg.Add(), wg.Done(), wg.Wait()

	//wg.Add(2)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

	fmt.Println("All workers completed")
}
*/
