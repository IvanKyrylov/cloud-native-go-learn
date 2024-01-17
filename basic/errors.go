package basic

import (
	"errors"
	"fmt"
	"os"
)

type NestedError struct {
	Message string
	Err     error
}

func (e *NestedError) Error() string {
	return fmt.Sprintf("%s\n constains: %s", e.Message, e.Err.Error())
}

func errorProcessing() {
	_, err := os.Open("somefile.txt")
	if err != nil {
		// log.Fatal(err)
		// return err

		fmt.Println(err)
	}

	e1 := errors.New("error 42")
	e2 := fmt.Errorf("error %d", 42)
	fmt.Println(e1, e2)
}
