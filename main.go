package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

func main() {
	testingVars()
	bitwiseOperators()
	conditions()
	loops()
	structs()
	maps()
	useArrayList()
	pointers()
	fmt.Println(multipleOutput(5))
	fmt.Println(variadics(1, 3, 5, 6))

	myClosure := closure("bye")
	fmt.Println(myClosure(" nigger1"))
	fmt.Println(myClosure(" nigger2"))
	fmt.Println(myClosure(" nigger3"))
	fmt.Println(myClosure(" nigger4"))

	unicode()
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

func bitwiseOperators() {
	var a uint = 60 // 60 = 0011 1100
	var b uint = 13
	var c uint = 0

	c = a & b
	fmt.Printf("Line 1 - Value of c is %d\n", c)

	c = a | b
	fmt.Printf("Line 2 - Value of c is %d\n", c)

	c = a ^ b
	fmt.Printf("Line 3 - Value of c is %d\n", c)

	c = a << 2
	fmt.Printf("Line 4 - Value of c is %d\n", c)

	c = a >> 2
	fmt.Printf("Line 5 - Value of c is %d\n", c)
}

func conditions() {
	var a int = 10
	if a < 20 {
		fmt.Printf("a is less than 20\n")
	}

	fmt.Printf("value of a is : %d\n", a)

	fmt.Println()
	fmt.Println()

	switch a {
	case 20:
		fmt.Printf("Value of a is 20\n")
		break
	case 10:
		fmt.Printf("Value of a is 10\n")
		break
	}
}

func loops() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var fruits = []string{"apple", "orange", "banana"}

	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}
}

func structs() {
	type Person struct {
		name   string
		age    int
		job    string
		salary int
	}

	type container struct {
		Person
		sex string
	}

	var c1 = container{
		Person: Person{
			name:   "ben",
			age:    10,
			job:    "super hero",
			salary: 0,
		},
		sex: "male",
	}

	fmt.Println(c1.name + "'s sex is " + c1.sex)

	var p1 Person
	p1.name = "John"
	p1.age = 25
	p1.job = "Developer"
	p1.salary = 5000

	fmt.Println(p1)

	p2 := Person{name: "Jane", age: 30, job: "Designer", salary: 6000}

	fmt.Println(p2)

	p3 := Person{"Jack", 35, "Manager", 7000}

	fmt.Println(p3)
}

func maps() {
	var m1 = map[string]int{"one": 1, "two": 2, "three": 3}

	fmt.Println(m1)

	m2 := make(map[string]int)

	m2["one"] = 1
	m2["two"] = 2
	m2["three"] = 3

	fmt.Println(m2)

	delete(m2, "two")

	fmt.Println(m2)

	value, exists := m2["two"]

	fmt.Println(value, exists)

	for k, v := range m1 {
		fmt.Println(k, v)
	}

	m3 := map[string]map[string]int{
		"one":   {"one": 1, "two": 2},
		"two":   {"two": 2},
		"three": {"three": 3},
	}

	fmt.Println(m3)

}

func useArrayList() {
	var list = createArraylist(int8(0))
	list.get(0)

	list.add(1)
	list.add(2)
	list.add(3)

	list.remove(2)

	fmt.Println(list.printList())

	list.add(3)
	list.add(3)
	list.removeAll(3)
	fmt.Println(list.printList())
}

func multipleOutput(x int8) (int8, int8) {
	return x, x + 1
}

func pointers() {
	var a int = 10
	var b *int = &a
	var pointer *int8 = new(int8)

	*pointer = 3

	fmt.Println(a, b, *b)
	fmt.Println(&b)
	fmt.Println(&b)
	fmt.Println(*(&a))
}

func variadics(num ...int32) int32 {
	var total int32 = 0
	for _, x := range num {
		total += x
	}
	return total
}

func closure(initialVal string) func(added string) string {
	str := initialVal

	return func(added string) string {
		str += added
		return str
	}
}

func unicode() {
	str := "Hello"

	fmt.Println(utf8.RuneCountInString(str))
}
