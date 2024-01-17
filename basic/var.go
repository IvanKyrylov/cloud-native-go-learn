package basic

import (
	"fmt"
	"log"
	"math/rand"
)

func logicValue() {
	and := true && false
	fmt.Println(and)

	or := true || false
	fmt.Println(or)

	not := !true
	fmt.Println(not)
}

func complexValue() {
	var x complex64 = 3.1415i
	fmt.Println(x)
}

func stringValue() {
	basicString := "Hello\nworld!\n"
	fmt.Print(basicString)

	lowLvlString := `Hello
world!`
	fmt.Print(lowLvlString)
}

func variable() {
	// var foo int = 42
	var foo, bar = 42, 1302 // var foo, bar int = 42, 1302
	// var foo = 42
	var d, f, s = true, 2.3, "four"
	fmt.Println(foo, bar, d, f, s)

	percent := rand.Float64() * 100.0 //nolint:gosec
	x, y := 0, 2

	fmt.Println(percent, x, y)

	var i int
	var fZero float64
	var b bool
	var sEmpty string

	fmt.Printf("int: %d\n", i)
	fmt.Printf("float: %f\n", fZero)
	fmt.Printf("bool: %t\n", b)
	fmt.Printf("string: %s\n", sEmpty)

	str := "world"
	_, err := fmt.Printf("Hello %s\n", str)
	if err != nil {
		log.Fatal(err)
	}

	// import _ "github.com/lib/pq"
}

const language string = "Go"

var favorite = true // var favorite bool = true

func constValue() {
	const text = "Does %s rule? %t!"
	var output = fmt.Sprintf(text, language, favorite)
	fmt.Println(output)
}

func pointerValue() {
	var a = 10 // var a int = 10

	var p = &a // var p *int = &a

	fmt.Println(p)
	fmt.Println(*p)

	*p = 20
	fmt.Println(a)

	var n *int
	var x, y int

	fmt.Println(n)
	fmt.Println(n == nil)

	fmt.Println(x == y)
	fmt.Println(&x == &x) //nolint:gocritic,staticcheck
	fmt.Println(&x == &y)
	fmt.Println(&x == nil) //nolint:staticcheck
}
