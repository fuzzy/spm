package main

import (
	"crypto/tls"
	"fmt"
	"github.com/dutchcoders/goftp"
	"io"
	"net/http"
	"os"
	"strings"
)

var Insecure = false
var Quiet = false
var Uri = ""
var Out = ""

/*
 Function: ErrChk(error)
 Arguments: e
 Notes: It calls os.Exit(1) if triggered.
 Author: Mike 'Fuzzy' Partin
*/
func ErrChk(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

/*
 Function: HttpGet(string, *os.File)
 Arguments:
*/
func HttpGet(u string) {
	// create our http response object
	resp, err := http.Get(u)
	ErrChk(err)
	defer resp.Body.Close()

	// create our output file
	out, err := os.Create(Out)
	ErrChk(err)
	defer out.Close()

	// and download the file
	_, err = io.Copy(out, resp.Body)
	ErrChk(err)
}

/*

 */
func HttpsGet(u string, i bool) {
	// create our https support
	tr := &http.Transport{}
	if !i {
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
	resp, err := client.Get(u)
	ErrChk(err)
	defer resp.Body.Close()

	// create our output file
	out, err := os.Create(Out)
	ErrChk(err)
	defer out.Close()

	// and dload it
	_, err = io.Copy(out, resp.Body)
	ErrChk(err)
}

func parseFtpUri(u string) map[string]string {
	s := strings.Split(Uri, "/")
	ftpData := make(map[string]string)
	// if we have a ':' we at the least have a port definition
	if strings.Contains(s[2], ":") {
		// additionally if we have an '@' at this point, we have, at the least login info
		if strings.Contains(s[2], "@") {
			c := strings.Split(s[2], ":")
			a := strings.Split(c[0], "@")
			ftpData["Host"] = c[1]
			ftpData["User"] = a[0]
			ftpData["Pswd"] = a[1]
			if len(c) == 3 {
				ftpData["Port"] = c[2]
			} else {
				ftpData["Port"] = "21"
			}
		} else {
			ftpData["User"] = "ftp"
			ftpData["Pswd"] = "user@domain.tld"
			// well we still have a port anyway.
			ftpData["Host"] = strings.Split(s[2], ":")[0]
			ftpData["Port"] = strings.Split(s[2], ":")[1]
		}
	} else {
		ftpData["Host"] = s[2]
		ftpData["Port"] = "21"
		ftpData["User"] = "ftp"
		ftpData["Pswd"] = "user@domain.tld"
	}
	ftpData["Path"] = strings.Join(s[3:len(s)-1], "/")
	ftpData["File"] = s[len(s)-1]
	ftpData["Oput"] = Out
	return ftpData
}

func FtpGet(s string) {
	data := parseFtpUri(s)
	var ftp *goftp.FTP
	var err error

	fmt.Println(data)

	ftp, err = goftp.Connect(fmt.Sprintf("%s:%s", data["Host"], data["Port"]))
	ErrChk(err)
	defer ftp.Close()

	ErrChk(ftp.Login(data["User"], data["Pswd"]))
	ErrChk(ftp.Cwd(data["Path"]))

	// create our output file
	out, err := os.Create(Out)
	ErrChk(err)
	defer out.Close()

	//ErrChk(ftp.Retr)
}

func main() {
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
		// Create our output filehandle
		if Out == "" {
			Out = "/tmp"
		}
		if f, _ := os.Stat(Out); f.IsDir() {
			data := strings.Split(Uri, "/")
			turi := fmt.Sprintf("%s/%s", Out, data[len(data)-1])
			Out = turi
		}

		// And do the dirty work
		if Uri[:5] == "ftp:/" {
			FtpGet(Uri)
		} else {
			if Uri[:5] == "http:" {
				HttpGet(Uri)
			} else if Uri[:5] == "https" {
				HttpsGet(Uri, Insecure)
			}
		}
	}
	if Quiet {
		fmt.Println("fuck off")
	}
}
