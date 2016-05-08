package pipe

/*
 Author: Mike 'Fuzzy' Partin
 Copyright: (c) 2016-2018
 Email: fuzzy@fumanchu.org
 License: See LICENSE.md for details
*/

import (
	"git.c0d0p0s0.vpn/fuzzy/spm.git/libs/errchk"
	"os"
	"time"
	//	"io"
)

type PMeter interface {
	Write(p []byte) (int, error)
	Read(p []byte) (int, error)
	Seek(offset int64, whence int) (int64, error)
	Close() error
}

type PipeMeter struct {
	File     *os.File
	Start    time.Time
	BytesIn  int64
	BytesOut int64
	Total    int64
}

func (pm PipeMeter) Write(p []byte) (int, error) {
	pm.BytesOut += int64(len(p))
	return pm.File.Write(p)
}

func (pm PipeMeter) Read(p []byte) (int, error) {
	return pm.File.Read(p)
}

func (pm PipeMeter) Seek(offset int64, whence int) (int64, error) {
	return pm.File.Seek(offset, whence)
}

func (pm PipeMeter) Close() error {
	return pm.File.Close()
}

func NewPipeMeter(f *os.File) PipeMeter {
	fInfo, err := os.Stat(f.Name())
	errchk.ErrChk(err)
	return PipeMeter{File: f,
		BytesIn:  0,
		BytesOut: 0,
		Total:    fInfo.Size(),
		Start:    time.Now()}
}
