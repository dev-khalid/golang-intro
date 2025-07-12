package goroutine

import "fmt"

func Channel() { 
	ch := make(chan int)

	go func() { 
		sum := 0; 
		for i := 0; i < 10; i++ { 
			sum += i;
		}
		ch <- sum
	}()

	fmt.Println("Waiting for the channel to receive a value...")
	fmt.Println("Received value from channel %d", <-ch)
}