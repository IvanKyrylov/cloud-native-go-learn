package basic

import (
	"fmt"
	"io"
)

// foo()
// go foo()

func Log(w io.Writer, message string) {
	go func() {
		fmt.Fprintln(w, message)
	}()
}

func goroutines() {}

func channels() {
	// val := 1
	// var ch chan int = make(chan int)
	// ch <- val
	// <-ch

	// ch := make(chan string)
	// go func() {
	// 	message := <-ch
	// 	fmt.Println(message)
	// 	ch <- "pong"
	// }()
	//
	// ch <- "ping"
	// fmt.Println(<-ch)

	// ch := make(chan string, 2)
	// ch <- "foo"
	// ch <- "bar"
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	// ch := make(chan string, 10)
	// ch <- "foo"
	// close(ch)
	//
	// msg, ok := <-ch
	// fmt.Printf("%q, %v\n", msg, ok)
	//
	// msg, ok = <-ch
	// fmt.Printf("%q, %v\n", msg, ok)

	// ch := make(chan string, 3)
	// ch <- "foo"
	// ch <- "bar"
	// ch <- "baz"
	// close(ch)
	//
	// for s := range ch {
	// 	fmt.Println(s)
	// }
}

func selectControl() {
	// ch1 := make(chan struct{})
	// ch2 := make(chan struct{})
	// ch3 := make(chan struct{})
	// y := struct{}{}
	// select {
	// case <-ch1:
	// 	fmt.Println("Got something")
	// case x := <-ch2:
	// 	fmt.Println(x)
	// case ch3 <- y:
	// 	fmt.Println(y)
	// default:
	// 	fmt.Println("None if the above")
	// }

	// ch := make(chan int)
	// select {
	// case m := <-ch:
	// 	fmt.Println(m)
	// case <-time.After(10 * time.Second):
	// 	fmt.Println("Timed out")
	// }
}
