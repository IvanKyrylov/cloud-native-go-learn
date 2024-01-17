package basic

import "testing"

//
// go test --bench=. --benchmem --count 5
// goos: windows
// goarch: amd64
// pkg: cloud-native-go-learn/basic
// cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
// BenchmarkFactorial-8            54295939                21.24 ns/op            0 B/op          0 allocs/op
// BenchmarkFactorial-8            59570200                19.91 ns/op            0 B/op          0 allocs/op
// BenchmarkFactorial-8            58771383                19.89 ns/op            0 B/op          0 allocs/op
// BenchmarkFactorial-8            59473363                19.93 ns/op            0 B/op          0 allocs/op
// BenchmarkFactorial-8            57626092                20.22 ns/op            0 B/op          0 allocs/op
// BenchmarkFactorialV2-8          159582786                7.216 ns/op           0 B/op          0 allocs/op
// BenchmarkFactorialV2-8          162228236                7.292 ns/op           0 B/op          0 allocs/op
// BenchmarkFactorialV2-8          165965186                7.255 ns/op           0 B/op          0 allocs/op
// BenchmarkFactorialV2-8          159230132                8.275 ns/op           0 B/op          0 allocs/op
// BenchmarkFactorialV2-8          156095916                7.525 ns/op           0 B/op          0 allocs/op
// PASS
// ok      cloud-native-go-learn/basic      15.933s

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorial(11)
	}
}

func BenchmarkFactorialV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorialV2(11)
	}
}
