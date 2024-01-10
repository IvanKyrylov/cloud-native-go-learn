package basic

import (
	"fmt"
	"log/slog"
)

func sumInts(m map[string]int64) int64 {
	var s int64

	for _, value := range m {
		s += value
	}

	return s
}

func sumFloats(m map[string]float64) float64 {
	var s float64

	for _, value := range m {
		s += value
	}

	return s
}

func sumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, value := range m {
		s += value
	}

	return s
}

type number interface {
	int64 | float64
}

func sumNumbers[K comparable, V number](m map[K]V) V {
	var s V
	for _, value := range m {
		s += value
	}

	return s
}

func runGeneric(logger *slog.Logger) {
	logger.Debug("Start generic")
	defer logger.Debug("End generic")

	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n", sumInts(ints), sumFloats(floats))
	fmt.Printf("Generic Sums: %v and %v\n", sumIntsOrFloats[string, int64](ints), sumIntsOrFloats[string, float64](floats))
	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n", sumIntsOrFloats(ints), sumIntsOrFloats(floats))
	fmt.Printf("Generic Sums with Constraint: %v and %v\n", sumNumbers(ints), sumNumbers(floats))
	fmt.Printf("Generic Sums with Constraint: %T and %T\n", sumNumbers(ints), sumNumbers(floats))
}
