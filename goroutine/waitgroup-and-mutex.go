package goroutine

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"
)


type SocialMediaPost struct { 
	Title   string
	Content string
	Author  string
	Likes   int  // Optional: nil means no likes count available
	mu  	  sync.Mutex
}

func (post *SocialMediaPost) IncrementLikes() {
	defer post.mu.Unlock()

	post.mu.Lock()
	post.Likes++
}

type FanWentCrazyPayload struct {
	Post *SocialMediaPost
	WG   *sync.WaitGroup
	TotalTime *time.Duration
	idx *int
}

func FanWentCrazy(payload *FanWentCrazyPayload) {
	post, wg, totalTime := payload.Post, payload.WG, payload.TotalTime
	defer wg.Done()
	// fmt.Printf("Working on %d...\n", *idx)

	simulateDbOperation := (time.Duration(rand.Intn(5 * 1000)) * time.Millisecond)
	*totalTime += simulateDbOperation
	time.Sleep(simulateDbOperation)

	post.IncrementLikes()

	// fmt.Printf("Simulation took %v to complete.\n", simulateDbOperation)
	// fmt.Printf("Post '%s' by %s now has %d likes!\n\n\n", post.Title, post.Author, post.Likes)
}

func (post *SocialMediaPost) PrintMetadata() {
	fmt.Printf("Title: %s\nAuthor: %s\nLikes: %d\n", post.Title, post.Author, post.Likes)
}

func (post *SocialMediaPost) PrintJson() {
	jsonData, err := json.MarshalIndent(post, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling to JSON: %v\n", err)
		return
	}
	fmt.Printf("JSON format:\n%s\n", jsonData)
}

func WaitGroup() { 

	post := SocialMediaPost {
		Title:  "My First Post",
		Content: "This is the content of my first post.",
		Author:  "John Doe",
		Likes:   0,
	}

	fmt.Printf("Post: %v\n", post)
	
	post.PrintJson()

	wg := sync.WaitGroup{}

	// Start timing from line 70
	startTime := time.Now()

	// This is the total time it could have taken if we did it without goroutines.
	totalTime := time.Duration(0)

	// NOTE: If I had to use mutex here, instead of putting it into the SocialMediaPost struct, then I would have to do this:
	// mu := sync.Mutex{} and pass it to the FanWentCrazy function, and from there I would have to pass it to the IncrementLikes function. or I could use it before the 
	// post.IncrementLikes() call.
	
	for i := 0; i < 500000; i++ {
		wg.Add(1)
		payload := &FanWentCrazyPayload{
			Post:      &post,
			WG:        &wg,
			TotalTime: &totalTime,
			idx:       &i,
		}
		go FanWentCrazy(payload)
	}

	wg.Wait()

	// End timing at line 87
	endTime := time.Now()
	executionTime := endTime.Sub(startTime)

	fmt.Printf("Execution time of the loop: %v\n", executionTime)
	fmt.Printf("Total time could have taken for all operations: %v\n", totalTime)
	fmt.Printf("Saved time by using goroutines: %v\n", totalTime-executionTime)
	post.PrintMetadata()
}