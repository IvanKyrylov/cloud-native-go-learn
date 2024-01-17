package basic

import (
	"fmt"
	"os"
	"time"
)

func controlStructures() {
	forControl()
	ifControl()
	switchControl()
}

func forControl() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum++
	}

	fmt.Println(sum)

	sum, i := 0, 0

	for i < 10 {
		sum += i
		i++
	}

	fmt.Println(sum, i)

	// fmt.Println("For ever ...")
	//
	// for {
	// 	fmt.Println("... and ever")
	// }

	s := []int{2, 4, 8, 16, 32}

	for index, value := range s {
		fmt.Println(index, "->", value)
	}

	a := []int{2, 4, 8, 6, 8}
	sum = 0

	for _, value := range a {
		sum += value
	}

	fmt.Println(sum)

	m := map[int]string{
		1: "January",
		2: "February",
		3: "Match",
		4: "April",
	}

	for key, value := range m {
		fmt.Println(key, "->", value)
	}
}

func ifControl() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if _, err := os.Open("foo.txt"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("All is fine.")
	}

	_, err := os.Open("foo.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("All is fine.")
	}
}

func switchControl() {
	i := 0

	switch i % 3 {
	case 0:
		fmt.Println("Zero")
		fallthrough
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Huh?")
	}

	hour := time.Now().Hour()

	switch {
	case hour >= 5 && hour < 9:
		fmt.Println("I'm writing")
	case hour >= 9 && hour < 18:
		fmt.Println("I'm working")
	default:
		fmt.Println("I'm sleeping")
	}

	switch hourLocal := time.Now().Hour(); {
	case hourLocal >= 5 && hourLocal < 9:
		fmt.Println("I'm writing")
	case hourLocal >= 9 && hourLocal < 18:
		fmt.Println("I'm working")
	default:
		fmt.Println("I'm sleeping")
	}
}
