package basic

import (
	"fmt"
	"log"
	"log/slog"
	"math/rand"
)

var functions = []func(){
	logicValue,
	complexValue,
	stringValue,
	variable,
	constValue,
	arrayValue,
	sliceValue,
	subSliceValue,
	stringSliceValue,
	mapValue,
	pointerValue,
}

func runBasic(logger *slog.Logger) {
	logger.Debug("Start basic")
	defer logger.Debug("End basic")

	for i := 0; i < len(functions); i++ {
		functions[i]()
	}
}

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

func arrayValue() {
	var a [3]int
	fmt.Println(a)
	fmt.Println(a[1])

	a[1] = 42
	fmt.Println(a)
	fmt.Println(a[1])

	i := a[1]
	fmt.Println(i)

	b := [3]int{2, 4, 6}
	fmt.Println(b)

	b = [...]int{7, 8, 9}
	fmt.Println(len(b))
	fmt.Println(b[len(b)-1])

	c := [...]int{}
	fmt.Println(c)
	fmt.Println(len(c))
}

func sliceValue() {
	n := make([]int, 3)
	fmt.Println(n)
	fmt.Println(len(n))

	n[0] = 8
	n[1] = 16
	n[2] = 32

	fmt.Println(n)

	m := []int{1}
	fmt.Println(m)

	m = append(m, 2)
	fmt.Println(m)

	m = append(m, 3, 4)
	fmt.Println(m)

	m = append(m, m...)
	fmt.Println(m)
}

func subSliceValue() {
	s0 := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(s0)

	s1 := s0[:4]
	fmt.Println(s1)

	s2 := s0[3:]
	fmt.Println(s2)

	s0[3] = 42
	fmt.Println(s0)
	fmt.Println(s1)
	fmt.Println(s2)
}

func stringSliceValue() {
	s := "foö"
	r := []rune(s)
	b := []byte(s)

	fmt.Printf("%7T %v\n", s, s)
	fmt.Printf("%7T %v\n", r, r)
	fmt.Printf("%7T %v\n", b, b)
}

func mapValue() {
	freezing := make(map[string]float32)
	freezing["celsius"] = 0.0
	freezing["fahrenheit"] = 32.0
	freezing["kelvin"] = 273.2

	fmt.Println(freezing["kelvin"])
	fmt.Println(len(freezing))

	delete(freezing, "kelvin")
	fmt.Println(len(freezing))

	freezing = map[string]float32{
		"celsius":    0.0,
		"fahrenheit": 32.0,
		"kelvin":     273.2,
	}

	foo := freezing["no-such-key"]
	fmt.Println(foo)

	newton, ok := freezing["newton"]
	fmt.Println(newton)
	fmt.Println(ok)
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

// func controlStructures() {
//
// }
