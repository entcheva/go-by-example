package main

import "fmt"
import "math"
import "time"

const s string = "constant"

func main() {
	fmt.Println("hello world")
  fmt.Println("go" + "lang")

	fmt.Println("1+1=", 1+1)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	var a = "initial"
	var b, c int = 1, 2
	var e int
	f := "short"
	fmt.Println(a, b, c, e, f)

	fmt.Println(s)

	const n = 50000
	const d = 3e20 / n
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

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

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

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
}



type T struct {
	name string // name of the object
	value int // its value
}
