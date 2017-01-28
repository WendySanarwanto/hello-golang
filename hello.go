package main

import "fmt"

func main() {
	// The Hello World
	fmt.Println("Hello World!")

	// declare & initialise primitive variables
	i := 5
	var j int
	fmt.Println("i is: ", i)
	fmt.Println("j is: ", j)
	fmt.Println()

	// Default values of primitive types
	// Addresses and memory location
	var k int
	fmt.Println("default int is: ", k)
	fmt.Println("address of k is: ", &k)
	fmt.Println("value at", &k, " is:", *(&k))
	fmt.Println()

	var s string
	fmt.Println("default string is: ", s)
	fmt.Println("address of s is: ", &s)
	fmt.Println("value at", &s, " is:", *(&s))
	fmt.Println()

	var f float64
	fmt.Println("default float64 is: ", f)
	fmt.Println("address of f is: ", &f)
	fmt.Println("value at", &f, " is:", *(&f))
	fmt.Println()

	var ints [3]int
	fmt.Println("default int array is: ", ints)
	fmt.Println("address of ints[0] is: ", &ints[0])
	fmt.Println("value at", &ints[0], " is:", *(&ints[0]))
	fmt.Println("address of ints[1] is: ", &ints[1])
	fmt.Println("value at", &ints[1], " is:", *(&ints[1]))
	fmt.Println("address of ints[2] is: ", &ints[2])
	fmt.Println("value at", &ints[2], " is:", *(&ints[2]))
	fmt.Println()

	var c complex64
	fmt.Println("default complex64 is: ", c)
	fmt.Println("address of c is: ", &c)
	fmt.Println("value at", &c, " is:", *(&c))
	fmt.Print("\n\r")

	// Control structure - if..else

	a, b := 4, 5
	fmt.Println("a is", a, ", b is ", b)

	if a < b {
		fmt.Println("a is less than b")
	} else if a > b {
		fmt.Println("a is greater than b")
	} else {
		fmt.Println("a is equal to b")
	}
	fmt.Println()

	// Control structure - for loop, break , continue, range
	for i := 0; i < 5; i++ {
		fmt.Println("Value of i is now:", i)
	}
	fmt.Println()

	// Multiple Index
	for i, j := 0, 1; i < 5 && j < 10; i, j = i+1, j+2 {
		fmt.Println("Value of i is now:", i, ", and j is now: ", j)
	}

	// break in for loop
	z := 0
	fmt.Println()

	for {
		fmt.Println("Value of z is:", z)
		if z >= 3 {
			break
		}
		z++
	}
	fmt.Println()

	// continue in for loop
	for i := 0; i < 8; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd value is:", i)
	}
	fmt.Println()

	// For range loop
	y := [...]string{"a","b","c","d"}
	for i := range y {
		fmt.Println("Array item", i, "is", y[i]);
	}
	fmt.Println()

	x := map[string] string {"France": "Paris", "Japan": "Tokyo", "Korea":"Seoul", "USA": "Washington DC"}
	for key, val := range x {
		fmt.Println("Capital of", key,"is", val)
	}
}
