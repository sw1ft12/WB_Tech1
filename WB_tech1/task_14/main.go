package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}
*/

func main() {
	var a interface{}
	{
		var v int
		a = v
		fmt.Printf("%s\n", reflect.TypeOf(a).String())
	}

	{
		var v string
		a = v
		fmt.Printf("%s\n", reflect.TypeOf(a).String())
	}

	{
		var v bool
		a = v
		fmt.Printf("%s\n", reflect.TypeOf(a).String())
	}

	{
		var v chan int
		a = v
		fmt.Printf("%s\n", reflect.TypeOf(a).String())
	}

}
