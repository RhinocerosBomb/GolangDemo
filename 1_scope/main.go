package main

import (
	"GolangDemo/1_scope/inner"
	"fmt"
	"math/rand"
)

func main () {
	public := inner.NewPublic()
	private := inner.NewPrivate()

	fmt.Println("Random number is:",rand.Intn(20))
	fmt.Println()
	public.PubVal = 1
	fmt.Println(public.PubVal)

	private.PubVal = 1
	fmt.Println(private.PubVal)


	public = public.SetPriv(8)
	public.PubVal = 8
	fmt.Println(public)
}