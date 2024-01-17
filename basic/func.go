package basic

import (
	"fmt"
	"os"
)

func add(x, y int) int {
	return x + y
}

// func fooV1(i int, j int, a string, b string) {}
// func fooV2(i, j int, a, b string)            {}

func swap(x, y string) (string, string) {
	return y, x
}

// recursion
func factorial(n int) int {
	if n < 1 {
		return 1
	}

	return n * factorial(n-1)
}

// not recursion
func factorialV2(n int) int {
	fact := 1
	for i := n; i > 0; i-- {
		fact *= i
	}

	return fact
}

func baseDefer() {
	defer fmt.Println("cruel world")
	fmt.Println("goodbye")
}

func closeFile(f *os.File) {
	if err := f.Close(); err != nil {
		fmt.Println("Error closing file:", err.Error())
	} else {
		fmt.Println("File closed successfully")
	}
}

func zeroByValue(x int) {
	_ = x
	x = 0
	_ = x
}

func zeroByReference(x *int) {
	*x = 0
}

func update(m map[string]int) {
	m["c"] = 2
}

func product(factors ...int) int {
	p := 1

	for _, n := range factors {
		p *= n
	}

	return p
}

func sumFunc(x, y int) int     { return x + y }
func productFunc(x, y int) int { return x * y }

func incrementer() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func function() {
	sum := add(10, 5)
	fmt.Println(sum)

	a, b := swap("foo", "bar")
	fmt.Println(a, b)

	fmt.Println(factorial(11))
	fmt.Println(factorialV2(11))

	baseDefer()

	file, err := os.Create("./basic/tmp/foo.out")
	defer closeFile(file)
	if err != nil {
		return
	}

	_, err = fmt.Fprintln(file, "Your mother was a hamster")
	if err != nil {
		return
	}

	fmt.Println("File written to successfully")

	defer fmt.Println("world")
	defer fmt.Println("cruel")
	defer fmt.Println("goodbye")

	x := 5

	zeroByValue(x)
	fmt.Println(x)

	zeroByReference(&x)
	fmt.Println(x)

	m := map[string]int{"a": 0, "b": 1}
	fmt.Println(m)

	update(m)
	fmt.Println(m)

	fmt.Println(product(2, 2, 2))
	v := []int{3, 3, 3}
	fmt.Println(product(v...))

	var f func(int, int) int
	f = sumFunc
	fmt.Println(f(3, 5))

	f = productFunc
	fmt.Println(f(3, 5))

	increment := incrementer()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	increment = incrementer()
	fmt.Println(increment())
}
