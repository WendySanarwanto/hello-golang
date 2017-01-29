package main

import "fmt"

func defaultValuesAddressMemoryLocation() {
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
}

func ifControlStructure() {
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
}

func forLoopControlStructure() {
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
	fmt.Println()
}

func switchControlStructure() {
	// Switch on integer
	w := 5
	switch w {
		case 4: fmt.Println("w is 4")
		case 5: fmt.Println("w is 5")
		case 6: fmt.Println("w is 6")
		default: fmt.Println("default")
	}
	fmt.Println()

	// Swith on integer with multiple cases
	u := 4
	switch u {
		case 1,3,5,7,9: fmt.Println(u, "is odd")
		case 0,2,4,6,8: fmt.Println(u, "is even")
		default: fmt.Println(u, "does not fall within 0-9 values")
	}
	fmt.Println()

	// Switch on string
	v := 6
	switch v {
		case 4: fmt.Println(v, "was <= 4"); fallthrough;
		case 5: fmt.Println(v, "was <= 5"); fallthrough;
		case 6: fmt.Println(v, "was <= 6"); fallthrough;
		case 7: fmt.Println(v, "was <= 7"); fallthrough;
		case 8: fmt.Println(v, "was <= 8");
		default: fmt.Println("default case");
	}
	fmt.Println()
}

func connectToDb() {
	fmt.Println("[INFO] - ok, connected to db")
}

func disconnectFromDb() {
	fmt.Println("[INFO] - ok, disconnected from db")
	fmt.Println()	
}

func deferControlStructure() {
	connectToDb()
	fmt.Println("[DEBUG] - Defering the database disconnect.")
	defer disconnectFromDb()
	fmt.Println("[DEBUG] - Doing some DB Operations ...")
	fmt.Println("[ERROR] - Oops! some crash or network error ...")
	fmt.Println("[DEBUG] - Returning from function here!")
	return 

	// deferred function is executed here just before this method actually returns.
}

func multipleDefersControlStructure() {
	// These defered calls will be executed in LIFO order: connectToDb() -> disconnectFromDb()
	defer disconnectFromDb()
	defer connectToDb()
}

func main() {
	// The Hello World
	fmt.Println("Hello World!")

	// declare & initialise primitive variables
	i := 5
	var j int
	fmt.Println("i is: ", i)
	fmt.Println("j is: ", j)
	fmt.Println()

	defaultValuesAddressMemoryLocation()

	ifControlStructure()

	forLoopControlStructure()

	switchControlStructure()

	deferControlStructure()

	multipleDefersControlStructure()
}
