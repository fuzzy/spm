package main

/*
 Author: Mike 'Fuzzy' Partin
 Copyright: (c) 2016-2018
 Email: fuzzy@fumanchu.org
 License: See LICENSE.md for details
*/

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var Author = "Mike 'Fuzzy' Partin"
var Email = "fuzzy@fumanchu.org"
var Version = "0.0"

var From string
var To string
var Quiet bool
var Delete bool
var files = 0
var dirs = 0
var now = time.Now()

func walkPath(path string, f os.FileInfo, err error) error {
	if !f.IsDir() {
		os.Symlink(path, fmt.Sprintf("%s/%s", To, path[len(From):]))
		files++
	} else {
		os.Mkdir(fmt.Sprintf("%s/%s", To, path[len(From):]), 0755)
		dirs++
	}
	if !Quiet {
		if (time.Now().Unix() - now.Unix()) >= 1 {
			fmt.Printf("%s Merged %d files and %d dirs.     \r", files, dirs)
			now = time.Now()
		}
	}
	return nil
}

func main() {
	flag.StringVar(&From, "src", "DIR", "Source directory.")
	flag.StringVar(&To, "dst", "DIR", "Destination directory.")
	version := flag.Bool("version", false, "Show version information.")
	flag.BoolVar(&Quiet, "quiet", false, "Suppress output.")

	flag.Parse()

	/* Display version and exit */
	if *version {
		fmt.Printf("%s v%s by %s <%s>\n", os.Args[0][2:], Version, Author, Email)
		os.Exit(0)
	}

	if _, sErr := os.Stat(From); sErr != nil {
		fmt.Printf("Source directory: %s does not exist.\n", From)
		os.Exit(1)
	}

	if _, dErr := os.Stat(To); dErr != nil {
		fmt.Printf("Destination directory: %s does not exist.\n", To)
		os.Exit(1)
	}

	/* walk the source directory */
	filepath.Walk(From, walkPath)

	/* and clean up */
	if !Quiet {
		fmt.Printf("Merged %d files and %d dirs.\n", files, dirs)
	}
}
