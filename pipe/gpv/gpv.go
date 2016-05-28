package main

import (
	"fmt"
	"github.com/fuzzy/spm/gout"
)

func main() {
	fmt.Println(gout.String("I exist.").Bold().Red())
	fmt.Println(gout.String("I think....").Red())

}
