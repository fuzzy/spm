package errchk

/*
 Author: Mike 'Fuzzy' Partin
 Copyright: (c) 2016-2018
 Email: fuzzy@fumanchu.org
 License: See LICENSE.md for details
*/

import (
	"fmt"
	"os"
)

func ErrChk(e error) {
	if e != nil {
		fmt.Println("ERROR:", e)
		os.Exit(1)
	}
}
