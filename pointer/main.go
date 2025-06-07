package pointer

import (
	// "encoding/json"
	"fmt"
)

type Hobby struct { 
	Name string
	Level int
}

type User struct { 
	Name string
	Age int
	Hobby Hobby
}

type Vertex struct { 
	X, Y float32
}
// RunPointerDemo runs the pointer demonstration
func RunPointerDemo() {
	fmt.Println("Hello pointer!");

	// var a *int
	// i := 42
	// a = &i
	// fmt.Println(*a)
	// *a = 21; 
	// fmt.Println(i)

	var user User;
	user.Name = "John Doe"
	user.Age = 30
	user.Hobby = Hobby{
		Name:  "Programming",
		Level: 5,
	}

	// Convert user to JSON and print
	// userJSON, _ := json.MarshalIndent(user, "", "  ")
	
	// fmt.Println("User object in JSON format:")
	// fmt.Println(string(userJSON))

	v := Vertex{3, 4}
	p := &v // p now points to v, similar to var p *Vertex = &v
	p.X = 22 // (*p).X = 22 // This is the same as p.X = 22, 
	// This is a language feature in GO, that allows to access without dereferencing.
	fmt.Println(p.X)

	var pets [2]string; 
	pets[0] = "Dog"
	pets[1] = "Cat"
	fmt.Println(pets)

	var primes = []int {1,2,3,5,7,11}
	// fmt.Println(primes)
	primes = append(primes,13,17,19,23,29)
	// fmt.Println(primes)

	// fmt.Println("Pet's length and capacity:", len(pets), cap(pets)); 
	// fmt.Println("Primes length and capacity:", len(primes), cap(primes))

	TicTacToe()
}

// To dereference the value from a pointer we do *<pointer_variable>
// To grab the address of a variable we do &<variable_name>
// Vertex{} is similar to 