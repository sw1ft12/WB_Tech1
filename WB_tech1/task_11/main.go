package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств
*/

func main() {
	a := map[string]struct{}{"1": {}, "2": {}, "6": {}, "7": {}}
	b := map[string]struct{}{"7": {}, "1": {}, "3": {}}
	res := make(map[string]struct{})
	for k := range a {
		if _, ok := b[k]; ok {
			res[k] = struct{}{}
		}
	}
	for k := range res {
		fmt.Printf("%s ", k)
	}
}
