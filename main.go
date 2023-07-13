package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte{'1'}
	res, _ := strconv.Atoi(string(b))
	fmt.Println(res)
}
