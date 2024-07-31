package main

import (
	"fmt"
	"math"
)

var text string = "hi"

func main() {
	fmt.Println(Fibo(45))
	testingVars()
}

func testingVars() {

	const kir = "khar"
	var name = "John"
	kiarash := 10

	fmt.Print("Hello, " + name + "!\n")
	fmt.Println(kiarash)

	var num1 int8 = 127

	fmt.Println(num1)

	var num2 int16 = 32767

	fmt.Println(num2)

	var num3 int32 = 2147483647

	fmt.Println(num3)

	var array1 [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println(array1[1])

	var array2 = [5]int{1, 2, 3, 4, 5}

	fmt.Println(array2[1])

	var array3 = [...]int{1, 2, 3, 4, 5}

	fmt.Println(array3[1])

	var slice1 = []int{1, 2, 3, 4, 5}

	fmt.Println(slice1[1])

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1 = append(slice1, 6)
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1 = append(slice1, slice1...)
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = array1[1:3]

	fmt.Println(slice2)

	var slice3 = make([]int, 5, 10)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	fmt.Println(math.Atan(1))
}
