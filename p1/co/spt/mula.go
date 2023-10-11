/*
This is testing out multiple animations with go routines
*/

package main

import (
	"fmt"
	"sync"
)

// Using go routines wait groups to run multiple spinners and see how they function

func Mula() {

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			AniMain()
		}()
	}
	wg.Wait()
	fmt.Println("Ran 3 times")
}

func Mula2() {
	var wg sync.WaitGroup
	done := make(chan struct{})

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			AniMain()
			done <- struct{}{} // Signal completion
		}()
	}

	go func() {
		wg.Wait()      // Wait for all goroutines to finish
		close(done)   // Close the channel to unblock the main goroutine
	}()

	for range done {
		// Wait until all goroutines finish
	}

	fmt.Println("Finished executing 3 instances of AniMain()")
}

// This one uses mutexes , lets see if this works
/*
This works, but even though 3 instances are called , each function is forced to go to completion
so it is sort of doing it sequentially
- You can make 2 different animations with different messages to indicate whats going on
*/
func Mula3() {
	var wg sync.WaitGroup
	mutex := sync.Mutex{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			AniMain()
			mutex.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Finished executing 3 instances of AniMain()")
}