package main

import (
	"fmt"
	"math"
	"runtime"

	// "github.com/dev-khalid/golang-intro/interfaces"

	// "time"
	// _ "github.com/dev-khalid/golang-intro/functions-deep-dive"
	// _ "github.com/dev-khalid/golang-intro/http"
	advancedbankingsystem "github.com/dev-khalid/golang-intro/exercises/advanced-banking-system"
)

func CountAndPrint(x, y int) (int, string) {
	result := x + y

	message := fmt.Sprintf("The sum of %d and %d is %d", x, y, result)

	return result, message
}

func split(sum int) (x, y int) { 
	x = sum * 4 / 9
	y = sum - x

	return
}

func forLoopExample(limit int) {
	for ;limit > 0; limit-- {
		fmt.Println("Looping:", limit)
	}
}

func whileLoopExample(limit int) {
	for limit > 0 {
		fmt.Println("Looping:", limit)
		limit--
	}
}

func pow(x, y, lim float64) float64 { 
	if v:= math.Pow(x,y); v < lim {  // Like for loop, initializer, then condition.
		return v
	}
	return lim
}


func switchExample() { 
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s\n", os)
	}
}

func main() {
	// Run the pointer demo automatically
	// fmt.Println("=== Running Pointer Demo ===")
	// pointer.RunPointerDemo()
	// fmt.Println("=== End of Pointer Demo ===")
	

	// fmt.Println("Hello There!", rand.Intn(10))
	// fmt.Println(math.Pi)

	// _, m := CountAndPrint(3, 4)
	// fmt.Println(m)
	// fmt.Println(split(28))
	// forLoopExample(5)
	// whileLoopExample(5)
	// fmt.Println(
	// 	pow(3, 2, 10),
	// 	pow(3, 3, 26),
	// )
	// switchExample()
	// defer fmt.Println("Hello, World!") // Defer works like stack (LIFO). It will be executed at the end of the function.
	// defer fmt.Println("Goodbye!")
	// fmt.Println(time.Now().UTC())


	// fmt.Println("Running goroutine code.")
	// goroutine.GoRoutineMain()

	// fmt.Println("Running interfaces code.")
	// interfaces.Interfaces()

	// server.ServerV2()
	advancedbankingsystem.AdvancedBanking()
}
