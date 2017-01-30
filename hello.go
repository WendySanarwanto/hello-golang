package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

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
	y := [...]string{"a", "b", "c", "d"}
	for i := range y {
		fmt.Println("Array item", i, "is", y[i])
	}
	fmt.Println()

	x := map[string]string{"France": "Paris", "Japan": "Tokyo", "Korea": "Seoul", "USA": "Washington DC"}
	for key, val := range x {
		fmt.Println("Capital of", key, "is", val)
	}
	fmt.Println()
}

func switchControlStructure() {
	// Switch on integer
	w := 5
	switch w {
	case 4:
		fmt.Println("w is 4")
	case 5:
		fmt.Println("w is 5")
	case 6:
		fmt.Println("w is 6")
	default:
		fmt.Println("default")
	}
	fmt.Println()

	// Swith on integer with multiple cases
	u := 4
	switch u {
	case 1, 3, 5, 7, 9:
		fmt.Println(u, "is odd")
	case 0, 2, 4, 6, 8:
		fmt.Println(u, "is even")
	default:
		fmt.Println(u, "does not fall within 0-9 values")
	}
	fmt.Println()

	// Switch on string
	v := 6
	switch v {
	case 4:
		fmt.Println(v, "was <= 4")
		fallthrough
	case 5:
		fmt.Println(v, "was <= 5")
		fallthrough
	case 6:
		fmt.Println(v, "was <= 6")
		fallthrough
	case 7:
		fmt.Println(v, "was <= 7")
		fallthrough
	case 8:
		fmt.Println(v, "was <= 8")
	default:
		fmt.Println("default case")
	}
	fmt.Println()
}

//--------------- Defer demo ----------------------------------------------------------------------

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

//--------------------------- Multiple Return values demo ---------------------------------------

func SumProductDiff(a, b int) (int, int, int) {
	return a + b, a * b, a - b
}

func multipleReturnValues() {
	a := 4
	b := 3
	sum, prod, diff := SumProductDiff(a, b)
	fmt.Println("The sum of", a, "and", b, "is", sum)
	fmt.Println("The product of", a, "and", b, "is", prod)
	fmt.Println("The diff of", a, "and", b, "is", diff, "\n\r")
}

func MySqrt(f float64) (float64, error) {
	// return error when the argument is invalid
	if f < 1 {
		return float64(math.NaN()), errors.New("Unable to sqrt the value.")
	}

	return math.Sqrt(f), nil
}

func MySqrtWithNamedReturns(f float64) (ret float64, err error) {
	if f < 1 {
		ret = float64(math.NaN())
		err = errors.New("Unable to sqrt the value.")
	} else {
		ret = math.Sqrt(f)
		//err is not assigned, so it gets default value nil
	}

	//automatically return the named return variables ret and err
	return
}

func printMySqrtResults(result float64, err error) {
	if err != nil {
		fmt.Println("Error! Return values are", result, "and error:", err)
	} else {
		fmt.Println("It's ok! Return values are", result, "and error:", err)
	}
}

func multipleReturnValuesWithErrorHandling() {
	val := float64(-1)
	fmt.Print("First try with", val, ": ")
	ret1, err1 := MySqrt(val)
	printMySqrtResults(ret1, err1)

	val = 9
	fmt.Print("Second try with", val, ": ")
	ret2, err2 := MySqrt(val)
	printMySqrtResults(ret2, err2)
	fmt.Println()

	val = 81
	fmt.Print("Third try with", val, ": ")
	ret3, err3 := MySqrtWithNamedReturns(val)
	printMySqrtResults(ret3, err3)
	fmt.Println()
}

// ------------------- OOP - Declaring struct and its members -------------------------------------

// In C#, these would be like:
//
// public class Rectangle {
//	  length int;
//	  width int;
//	  public Name string;
//
//	  public int GetRectangleArea() {
//		return length * width;
//	  }
// }

type Rectangle struct {
	length, width int
	Name          string
}

func (r Rectangle) GetRectangleArea() int {
	return r.width * r.length
}

func structsInGo() {
	rect1 := Rectangle{10, 5, "Rectangle_1"} // instantiate & initialize values in order they are defined in struct
	area := rect1.GetRectangleArea()
	fmt.Println("Rectangle rect1 is:", rect1, ", area of Rectangle_1:", area)

	rect2 := Rectangle{length: 4, width: 5, Name: "Rectangle_2"} // instantiate & initialize values by variable name in any order
	area = rect2.GetRectangleArea()
	fmt.Println("Rectangle rect2 is:", rect2, ", area:", area)

	rect3 := new(Rectangle) // Instantiate a new object using new. The result will be pointer to the object.
	(*rect3).width = 6      // Accessing the object's property using * operand
	rect3.length = 5        // //set value using . notation - same as previous.  There is no -> operator like in c++. Go automatically converts
	rect3.Name = "Rectangle_3"
	area = rect3.GetRectangleArea()
	fmt.Println("Rectangle rect3 as address is:", rect3, " as value is:", *rect3, ", area:", area)

	fmt.Println()
}

// ------------------- OOP - Anonymous field -------------------------------------

type RAM struct {
	numOfSlot int
}

type Motherboard struct {
	RAM     // Anynomous field
	chipset string
}

func anonymousFieldsInStruct() {
	mobo := Motherboard{RAM{4}, "Intel Z170"}
	fmt.Println("mobo has", mobo.RAM, "RAM slots. The chipset is", mobo.chipset)
	fmt.Println()
}

// ------------------- OOP - Inheritance -------------------------------------

type Computer struct {
	Ram   int
	Cpu   string
	Os    string
	Model string
}

func (computer Computer) PrintModelSummary() {
	fmt.Println("Model:", computer.Model, ",RAM:", computer.Ram, "GB, CPU:", computer.Cpu, ", OS:", computer.Os)
}

type Notebook struct {
	Computer
	ScreenSize int
}

func inheritanceInStruct() {
	asusRogGr8 := Computer{16, "Intel i7", "Windows 10", "ASUS ROG GR8 Gaming Desktop PC"}
	macbook := Notebook{Computer{8, "Intel i5", "Mac OSX Sierra", "Apple Macbook Air 11"}, 11}

	asusRogGr8.PrintModelSummary()
	macbook.PrintModelSummary()

	fmt.Println()
}

// ------------------- OOP - Interface ---------------------------------------
type Picturable interface {
	takePicture() string
}

type Dialable interface {
	call() string
}

type Swipable interface {
	swipeLeft() string
	swipeRight() string
}

// ------------------- OOP - Multiple Inheritances ---------------------------
type Camera struct{}

func (_ Camera) takePicture() string { // //not using the type, so discard it by putting a _
	return "* Click * - A picture is taken"
}

type Phone struct{}

func (_ Phone) call() string {
	return "* Ring ring * - Connecting to the dialed a number"
}

type SmartPhone struct {
	Camera
	Phone
	Model string
}

func (_ SmartPhone) swipeLeft() string {
	return " Screen Menu is sliding to left side"
}

func (_ SmartPhone) swipeRight() string {
	return " Screen Menu is sliding to right side"
}

func multipleInheritanceWithInterface() {
	samsungGalaxy := new(SmartPhone) // Instantiate smartphone object
	picturable := Picturable(samsungGalaxy) // Get Picturable interface from the smartphone object
	dialable := Dialable(samsungGalaxy) // Get Diallable interface from the smartphone object
	swipable := Swipable(samsungGalaxy) // Get Swipable interface from the smartphone object
	samsungGalaxy.Model = "Samsung Galaxy S7"
	fmt.Println("Press camera button on my", samsungGalaxy.Model, ":", picturable.takePicture()) // Calling takePicture method from Picturable interface
	fmt.Println("Press dial buttons on my", samsungGalaxy.Model, ":", dialable.call()) // Calling call method from Dialable interface
	fmt.Println()
	fmt.Println("Swipe the display of my", samsungGalaxy.Model, "to righthand side:", swipable.swipeRight()) // Calling swipeRight method from Swipable interface
	fmt.Println("Swipe the display of my", samsungGalaxy.Model, "to lefthand side:", swipable.swipeLeft()) // Calling swipeLeft method from Swipable interface
	fmt.Println()
}

// ------------------- OOP - Polymorphism ---------------------------
type GamingDevice interface {
	GetPlatform() string
	GetControllers() string
}

type Console struct{
	platform, controllers string
}

func (console Console) GetPlatform() string {
	return console.platform
}

func (console Console) GetControllers() string {
	return console.controllers
}

type PC struct {}

func (pc PC) GetPlatform() string {
	return "Microsoft Windows PC"
}

func (pc PC) GetControllers() string {
	return "Keyboard & Mouse"
}

func polymorphismInGo() {
	xBone := Console{"Microsoft XBox One", "XBox gamepad controllers"}
	ps4 := Console{"Sony Playstation 4", "PS4 gamepad controllers"}
	pc := new (PC)
	myGamingDevices := [...] GamingDevice {xBone, ps4, pc}

	fmt.Println("Common popular gaming devices in market nowday:")
	i := 1
	for gd := range myGamingDevices {
		myGamingDevice := myGamingDevices[gd]
		fmt.Println(i, ". Platform:", myGamingDevice.GetPlatform(), ", controllers:", myGamingDevice.GetControllers() )
		i++
	}
	fmt.Println()
}

// ------------------- Routine ---------------------------------------

type Animal interface {
	doRun(distanceInMeter int) string
}

type Mamal struct {
	race, name string
	speedInKilometerPerHour int
}

func (m Mamal) doRun(distanceInMeter int) string {
	speedInMeterPerSec := m.speedInKilometerPerHour * 1000 / 3600
	timeInSecs := distanceInMeter / speedInMeterPerSec
	time.Sleep(timeInSecs 1e9)

}

func animalRaces() {
	bugsBunny := Mamal {"bunny", "bugsBunny", 120}
	sylverster := Mamal {"cat", "Sylvester", 110}
	
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

	multipleReturnValues()

	multipleReturnValuesWithErrorHandling()

	structsInGo()

	anonymousFieldsInStruct()

	inheritanceInStruct()

	multipleInheritanceWithInterface()

	polymorphismInGo()
}
