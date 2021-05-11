package main

import "fmt"

type A struct {
	Val int
}

func main() {
	// Value of type A
	a := A{
		Val: 10,
	}

	fmt.Println(a.Val)

	// setting by reference does nothing to the original struct as the value in the original memory address is not changed.
	a.SetByValue(50)
	a.SetByReference(100)

	fmt.Println(a.Val)
	fmt.Println(&a)

	// Slices and maps are some of the reference types in golang, meaning they are mutated no matter if a pointer is used or not
	slice := make([]int, 5)
	//m := make(map[int]int)

	fmt.Println(slice)
	SetSlice(slice, 1, 100)
	fmt.Println(slice)
}

// Value is copied
func (a A) SetByValue(val int) A {
	a.Val = val
	return a
}

// Passing in by reference will mutate the underlying value in the address space
func (a *A) SetByReference(val int) {
	a.Val = val
}

// Slices are always passed by reference so no pointers needed
func SetSlice(s []int, i, val int) {
	s[i] = val
}