package main

import (
	"fmt"
	"os"
)

func main() {
	splat, err := ReadSplat(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v\n", splat)
}
