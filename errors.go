package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	code    int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.code, e.message)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "Can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 Failed", e)
		} else {
			fmt.Println("f1 Worked", r)
		}
	}

	for _, i := range []int{7, 42} {
		r, e := f2(i) //same as the if condition above, just split into two lines
		if e != nil {
			fmt.Println("f2 Failed", e)
		} else {
			fmt.Println("f2 Worked", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.code)
		fmt.Println(ae.message)
	}
}
