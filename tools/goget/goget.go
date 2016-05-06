package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func ErrChk(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func main() {
	var Insecure = false
	var Quiet = false
	var Uri = ""
	var Out = ""
	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			if os.Args[i] == "-h" || os.Args[i] == "-help" || os.Args[i] == "--help" {
				fmt.Printf("Usage: %s <-h|-q|-i> [URI] <output file/dir>\n", os.Args[0])
				fmt.Printf("%2s|%-9s|%-10s Show this help screen\n", "-h", "-help", "--help")
				fmt.Printf("%2s|%-9s|%-10s Suppress output\n", "-q", "-quiet", "--quiet")
				fmt.Printf("%2s|%-9s|%-10s Do not validate certificates (https)\n", "-i", "-insecure", "--insecure")
				os.Exit(1)
			}
			if os.Args[i] == "-q" || os.Args[i] == "-quiet" || os.Args[i] == "--quiet" {
				Quiet = true
			} else if os.Args[i] == "-i" || os.Args[i] == "-insecure" || os.Args[i] == "--insecure" {
				Insecure = true
			}
			if len(os.Args[i]) >= 5 {
				if os.Args[i][:4] == "ftp:" && Uri == "" {
					Uri = os.Args[i]
				} else if os.Args[i][:5] == "http:" || os.Args[i][:5] == "https" && Uri == "" {
					Uri = os.Args[i]
				} else if len(Out) == 0 {
					f, e := os.Stat(os.Args[i])
					if e == nil {
						if f.IsDir() {
							Out = fmt.Sprintf("%s/", os.Args[i])
						} else {
							Out = fmt.Sprintf("%s", os.Args[i])
						}
					} else if e != nil {
						fmt.Println("I am an error!")
					}
				}
			}
		}
		if !Quiet {
			// Create our output filehandle
			if f, _ := os.Stat(Out); f.IsDir() {
				data := strings.Split(Uri, "/")
				turi := fmt.Sprintf("%s/%s", Out, data[len(data)-1])
				Out = turi
			}
			out, err := os.Create(Out)
			ErrChk(err)
			defer out.Close()

			// And do the dirty work
			if Uri[:5] == "ftp:/" {
				fmt.Println("FTP Support is being worked on")
			} else {
				if Uri[:5] == "http:" {
					// create our http response object
					resp, err := http.Get(Uri)
					ErrChk(err)
					defer resp.Body.Close()
					// and download the file
					_, err = io.Copy(out, resp.Body)
					ErrChk(err)
				} else if Uri[:5] == "https" {
					// create our https support
					tr := &http.Transport{}
					if !Insecure {
						tr = &http.Transport{
							TLSClientConfig:    &tls.Config{},
							DisableCompression: true,
						}
					} else {
						tr = &http.Transport{
							TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
							DisableCompression: true,
						}
					}
					client := &http.Client{Transport: tr}
					resp, err := client.Get(Uri)
					ErrChk(err)
					defer resp.Body.Close()
					// and dload it
					_, err = io.Copy(out, resp.Body)
					ErrChk(err)
				}
			}
		}
	}
}
