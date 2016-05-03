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
	"github.com/fuzzy/gocolor"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

var Author = "Mike 'Fuzzy' Partin"
var Email = "fuzzy@fumanchu.org"
var Version = "0.2"
var From string
var To string
var Delete bool

/*
 Walk the given path (which in general will be the install root of a package)
 and build then return lists of all dirs and files contained recursively therein
*/
func walkDir(path string) ([]string, []string) {
	var dirs []string
	var files []string
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			dirs = append(dirs, path)
		} else {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return dirs, files
}

func main() {
	// set our command line arguments
	version := flag.Bool("version", false, "Show version information.")
	flag.StringVar(&From, "src", "DIR", "Source directory.")
	flag.StringVar(&To, "dst", "DIR", "Destination directory.")
	flag.BoolVar(&Delete, "delete", false, "Delete files if they exist in <dst>")
	flag.Parse()

	/* Display version and exit */
	if *version {
		fmt.Printf("%s\n", Version)
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
	dirs, files := walkDir(From)
	if !Delete {
		var fn = 0
		var dn = 0
		sort.Strings(dirs)
		for d := 0; d < len(dirs); d++ {
			if len(From)+1 < len(dirs[d]) {
				de := os.Mkdir(fmt.Sprintf("%s/%s", To, dirs[d][len(From)+1:]), 0755)
				if de == nil {
					dn++
				}
			}
		}
		for f := 0; f < len(files); f++ {
			fe := os.Symlink(files[f], fmt.Sprintf("%s/%s", To, files[f][len(From)+1:]))
			if fe == nil {
				fn++
			}
		}
		fmt.Printf("%s%s Linked %10d files and %10d directories\n",
			string(gocolor.String(">").Bold().Cyan()),
			string(gocolor.String(">").Bold().White()),
			fn,
			dn)
	} else {
		var fn = 0
		var dn = 0
		sort.Sort(sort.Reverse(sort.StringSlice(dirs)))
		for f := 0; f < len(files); f++ {
			fe := os.Remove(fmt.Sprintf("%s/%s", To, files[f][len(From)+1:]))
			if fe == nil {
				fn++
			}
		}
		for d := 0; d < len(dirs); d++ {
			if len(From)+1 < len(dirs[d]) {
				dd, _ := ioutil.ReadDir(fmt.Sprintf("%s/%s", To, dirs[d][len(From):]))
				if len(dd) < 1 {
					de := os.RemoveAll(fmt.Sprintf("%s/%s", To, dirs[d][len(From):]))
					if de == nil {
						dn++
					}
				}
			}
		}
		fmt.Printf("%s%s Deleted %10d files and %10d directories\n",
			string(gocolor.String(">").Bold().Cyan()),
			string(gocolor.String(">").Bold().White()),
			fn,
			dn)
	}
}
