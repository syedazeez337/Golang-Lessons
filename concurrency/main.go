// having multiple tasks running independently and not necessarily at the same time
// one is pausing till the other task is finished

// java takes a source code and gives it to JVM > bytecode
// run a java programme with class file through JRT > native code
// java, python, c, cpp they all have thread which depends on OS
// it puts a layer between OS and go source code
// java thread cost like 1MB and with a RAM 1 GB you can spin 1000 threads
// Go allots like 4 or 8 kb with 1 GB it can up to 250,000 goroutines

// go also has GoRuntime
// go produces native code directly according to the system arch

// scheduler > takes care of which goroutine has the priority
// CSP (communicating sequectial processes) by Tony Hoare (quick sort algorithm)
// coroutine > goroutine

package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go printMessage("goroutine 1")
	printMessage("Main function")
}

/*
func hello() {
	fmt.Println("from hello function")
}

// in go main is goroutine and it has the highest priority
func main() { // main function takes 1 second
	go hello()  // let's say this function takes 2 seconds
	time.Sleep(3)
	// synchronisation of goroutines
	// we have time pacakge we can use Sleep
	// we have an upgraded package called sync and we waitgroups
	// mutex
	// in order to overcome the constraints of waitgroups we have channels
	// we use waitgroups to remove time.Sleep
	fmt.Println("from main function")
}
*/
