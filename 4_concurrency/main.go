package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main()  {
	Example4()
}

// WaitGroup
// WaitGroups allow proccess to wait until other processes are complete
func Example1() {
	wg := sync.WaitGroup{}
	list := []string{"Hi!", "Hello!", "GoodBye!"}

	// Defines the number of processes that the waitgroup should wait for
	wg.Add(len(list))
	for _, val := range list {
		go func(val string) {
			// defer is a special keyword that allows another function to be called after the function has been completed
			defer wg.Done()
			fmt.Println(val)
		}(val)
	}

	wg.Wait()
	fmt.Println("Complete")
}

// Mutexes
// Mutexes are used for limiting access to resources. This prevents race conditions when multiple process tries to read/write
// from the same resource
// In this example a counter is being used by 101 proccess where each increments the counter
func Example2() {
	counter := 0
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			// comment out the unlock and lock to see what happens when not using mutexes
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			counter +=1
			fmt.Println(counter)
		}()
	}

	wg.Wait()
	fmt.Println("Complete!")
}

// Channels
func Example3() {
	wg := sync.WaitGroup{}
	wg.Add(4)

	errChan := make(chan error)
	// quit channel
	quit := make(chan struct{})

	// Gorountine to handle errors from other processes
	go func() {
		// All errors from other processes
		// Receiving error from the channel

		// Make sure to always return from your goroutines if they are not needed anymore
		// Without exiting goroutine when its not need (goroutine leak):
		//for err := range errChan {
		//	fmt.Println(err)
		//}

		// With Exiting
		for {
			select {
			case err := <-errChan:
				fmt.Println(err)
			case <- quit:
				fmt.Println("Ending goroutine")
				return
			}
		}
	}()

	// Creates 4 processes that divide numbers
	for i := 0; i < 4; i++ {
		go func(i int) {
			defer wg.Done()
			//Intn returns, as an int, a non-negative pseudo-random number in [0,n) from the default Source. It panics if n <= 0.
			randNum := rand.Intn(3)

			// If it tries to divide by 0 then it is sent to the error channel
			if randNum == 0 {
				// Sending error into the channel
				errChan <- errors.New("cannot divide by zero")
				return
			}

			fmt.Println(i/ randNum)
		}(i)
	}

	wg.Wait()

	quit <- struct{}{}
	// time.Sleep is used in this example because we want to be able to simulate closing the error handling goroutine
	// before the main goroutine completes
	time.Sleep(3*time.Second)
	fmt.Println("Completed!")
}

func Example4() {
	// Mocking server calls
	servers := []string{"server1", "server2", "server3"}

	// Call function waits a random amount of time and returns a number
	call := func(server string) int {
		randNum := rand.Intn(5)
		fmt.Println(randNum)
		time.Sleep(time.Duration(randNum) * time.Second)

		return randNum
	}

	wg := sync.WaitGroup{}

	wg.Add(len(servers))

	//mu := sync.Mutex{}
	var list []int
	for _, server := range servers {
		go func() {
			defer wg.Done()
			response := call(server)
			//mu.Lock()
			//defer mu.Unlock()
			list = append(list, response)
		}()
	}

	// inputs 1, 2, 2
	// process 1 reads [] from list
	// process 1 writes 1 to [] resulting in [1]
	// process 2 and 3 reads [1] from the list
	// process 2 writes 2 to [1] results in [1, 2]
	// process 3 writes 2 to [1] result in [1, 2]
	// results in [1, 2]

	wg.Wait()
	fmt.Println(list)
}