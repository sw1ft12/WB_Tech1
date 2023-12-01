package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество
*/

func main() {
	seq := []string{"cat", "cat", "dog", "cat", "tree"}
	res := make(map[string]struct{})
	for k := range seq {
		res[seq[k]] = struct{}{}
	}
	for k := range res {
		fmt.Printf("%s ", k)
	}
}
