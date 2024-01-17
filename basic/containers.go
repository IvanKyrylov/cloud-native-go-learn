package basic

import "fmt"

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

	t := []byte{1_000_000: 0}
	fmt.Println("t", len(t), cap(t))
}
