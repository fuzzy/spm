package gout

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"syscall"
	"unsafe"
)

type String string

func (i String) Black() String {
	return String(fmt.Sprintf("\033[30m%s\033[0m", i))
}

func (i String) Red() String {
	return String(fmt.Sprintf("\033[31m%s\033[0m", i))
}

func (i String) Green() String {
	return String(fmt.Sprintf("\033[32m%s\033[0m", i))
}

func (i String) Yellow() String {
	return String(fmt.Sprintf("\033[33m%s\033[0m", i))
}

func (i String) Blue() String {
	return String(fmt.Sprintf("\033[34m%s\033[0m", i))
}

func (i String) Purple() String {
	return String(fmt.Sprintf("\033[35m%s\033[0m", i))
}

func (i String) Cyan() String {
	return String(fmt.Sprintf("\033[36m%s\033[0m", i))
}

func (i String) White() String {
	return String(fmt.Sprintf("\033[37m%s\033[0m", i))
}

func (i String) Bold() String {
	return String(fmt.Sprintf("\033[1m%s\033[0m", i))
}

func (i String) Underline() String {
	return String(fmt.Sprintf("\033[4m%s\033[0m", i))
}

func (i String) Blink() String {
	return String(fmt.Sprintf("\033[5m%s\033[0m", i))
}

func (i String) Reverse() String {
	return String(fmt.Sprintf("\033[7m%s\033[0m", i))
}

func (i String) Conceal() String {
	return String(fmt.Sprintf("\033[8m%s\033[0m", i))
}

// Get terminal width
type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func getWidth() uint {
	ws := &winsize{}
	retv, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))
	if int(retv) == -1 {
		panic(err)
	}
	return uint(ws.Col)
}

// Humanize # of bytes into readable strings
func HumanSize(n uint64) String {
	labels := []string{"B", "K", "M", "G", "T", "P", "E"}
	for pow := 0; pow < len(labels); pow++ {
		if pow == 0 {
			if n < 1024 {
				return String(labels[pow]).Yellow()
			}
		} else if pow == 1 {
			if float64(n) > 1024 && float64(n) <= float64(math.Pow(1024, float64(pow+1))) {
				return String(fmt.Sprintf("%.02f%s",
						(float64(n) / float64(1024)),
						String(labels[pow]).Yellow()))
			}
		} else {
			if float64(n) > math.Pow(1024, float64(pow)) && float64(n) <= math.Pow(1024, float64(pow+1)) {
				return String(fmt.Sprintf("%.02f%s",
						(float64(n) / float64(math.Pow(1024, float64(pow)))),
						String(labels[pow]).Yellow()))
			}
		}
	}
	return String("")
}

// Humanize # of seconds into readable strings
func HumanTime(n uint64) String {
	sc := uint64(1)
	mn := (sc * uint64(60))
	hr := (mn * uint64(60))
	if n > 1 && n <= mn {
		return String(fmt.Sprintf("%02d%s", n, String("s").Yellow()))
	} else if n > mn && n <= hr {
		return String(fmt.Sprintf("%02d%s%02d%s",
				(n / mn),
				String("m").Yellow(),
				(n % mn),
				String("s").Yellow()))
	} else {
		return String("")
	}
}

func ProgressBar(c float64, t float64) String {
	p := ((c / t) * float64(100))
	d := int(p / float64(5))
	if p > 0.01 {
		return String(fmt.Sprintf("%s%s%s%s",
			String("[").Yellow(),
			String(strings.Repeat("#", d)).Bold().Red(),
			strings.Repeat(" ", (20 - d)),
			String("]").Yellow()))
	} else {
		return String(fmt.Sprintf("%s%s%s",
			String("[").Yellow(),
			strings.Repeat(" ", 20),
			String("]").Yellow()))
	}	
}

func Printr(s string) {
	r, _ := regexp.Compile("\\033[[0-9;]+m")
	l := len(r.ReplaceAllString(s, ""))
	w := getWidth()
	
	fmt.Printf("%s%s\r", s, strings.Repeat(" ", int(w - uint(l))))
}

func Println(s string) {
	r, _ := regexp.Compile("\\033[[0-9;]+m")
	l := len(r.ReplaceAllString(s, ""))
	w := getWidth()
	
	fmt.Println(s, strings.Repeat(" ", int(w - uint(l+2))))
}
