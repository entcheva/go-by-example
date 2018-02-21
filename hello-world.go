package main

import "fmt"
import "math"
import "time"

const s string = "constant"

func main() {

	// Println
	fmt.Println("hello world")
	fmt.Println("go" + "lang")

	fmt.Println("1+1=", 1+1)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// variables
	var a = "initial"
	var b, c int = 1, 2
	var e int
	f := "short"
	fmt.Println(a, b, c, e, f)

	fmt.Println(s)

	// constants
	const n = 50000
	const d = 3e20 / n
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

	// loops
	r := 1
	for r <= 3 {
		fmt.Println(r)
		r = r + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// conditionals
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// switch
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}

	timee := time.Now()
	switch {
	case timee.Hour() < 12:
		fmt.Println("before noon")
	default:
		fmt.Println("after noon")
	}

	whatAmI := func(hey interface{}) {
		switch thing := hey.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("don't know", thing)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hello there")

	// arrays
	var arrr [5]int
	fmt.Println("emp:", arrr)

	arrr[4] = 100
	fmt.Println("set:", arrr)
	fmt.Println("get:", arrr[4])
	fmt.Println("len:", len(arrr))

	anotherArrr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", anotherArrr)

	var twoD [2][3]int
	for p := 0; p < 2; p++ {
		for q := 0; q < 3; q++ {
			twoD[p][q] = p + q
		}
	}
	fmt.Println("2d: ", twoD)

	// slices
	slice := make([]string, 3)
	fmt.Println("emp:", slice)

	slice[0] = "a"
	slice[1] = "b"
	slice[2] = "c"
	fmt.Println("set:", slice)
	fmt.Println("get:", slice[2])
	slice = append(slice, "d")
	slice = append(slice, "e", "f")
	fmt.Println("apd:", slice)

	copiedSlice := make([]string, len(slice))
	copy(copiedSlice, slice)
	fmt.Println("cpy:", copiedSlice)

	// slice's slice operator
	sliced := slice[2:5]
	fmt.Println("sl1:", sliced)

	// Slices can be composed into multi-dimensional data structures. The length
	// of the inner slices can vary, unlike with multi-dimensional arrays.

	twoDSlice := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDSlice[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDSlice[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoDSlice)
}

type T struct {
	name  string // name of the object
	value int    // its value
}
