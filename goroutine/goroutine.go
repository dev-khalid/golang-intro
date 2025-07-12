package goroutine

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Practicing goroutines, waitgroups and channels.
*/
func GoRoutine() { 
	fmt.Print("Hello from goroutine!\n")

	wg := sync.WaitGroup{}
	wg.Add(1) // Add a count of 1 to the waitgroup

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			sleepingFor := time.Duration(rand.Intn(1000) * int(time.Millisecond))
			fmt.Printf("First 🚀🚀 %d, sleeping for %v\n", i, sleepingFor)
			time.Sleep(sleepingFor)
		}
	}()
	
	wg.Add(1)
	go func() { 
		defer wg.Done()
		for i := 0; i < 5; i++ {
			sleepingFor := time.Duration(rand.Intn(1000) * int(time.Millisecond))
			fmt.Printf("Second 😎😎 %d, sleeping for %v\n", i, sleepingFor)
			time.Sleep(sleepingFor)
		}
	}()

	wg.Wait()
	fmt.Println("All goroutines finished executing.")
}