package main

import (
	"fmt"
	"some-method-go/with_slice"
)

func main() {

	ss := []int{11, 12, 2, 0, 1, 10}
	index := 6

	t := with_slice.RemoveIndex(ss, index)
	fmt.Println(t)

}
