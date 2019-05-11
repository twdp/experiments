package main

import (
	"fmt"
	"reflect"
)

type A string

func main()  {

	var a A
	a = "a"
	fmt.Println(a)

	fmt.Println(reflect.TypeOf(a))
}