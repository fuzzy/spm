package pipe

/*
 Author: Mike 'Fuzzy' Partin
 Copyright: (c) 2016-2018
 Email: fuzzy@fumanchu.org
 License: See LICENSE.md for details
*/

import (
	"fmt"
	"github.com/fuzzy/spm/libs/gout"
	"time"
)

// Interface definitions

type PipeReader interface {
	Read(b []byte) (n int, e error)
}

type PipeWriter interface {
	Write(b []byte) (n int, e error)
}

type PipeCloser interface {
	Close() error
}

type PipeReadWriter interface {
	PipeReader
	PipeWriter
}

type PipeReadCloser interface {
	PipeReader
	PipeCloser
}

type PipeWriteCloser interface {
	PipeWriter
	PipeCloser
}

type PipeReadWriteCloser interface {
	PipeReader
	PipeWriter
	PipeCloser
}

// Object definitions

type MeteredPipe struct {
	IoObject interface{}
	BytesIn  int
	BytesOut int
	Started  int64
}

func (m *MeteredPipe) Read(b []byte) (n int, e error) {
	if v, ok := m.IoObject.(PipeReader); ok {
		n, e := v.Read(b)
		m.BytesIn += n
		if (time.Now().Unix() - m.Started) >= 1 {
			totalTime := uint64(time.Now().Unix() - m.Started)
			avgSpeed := uint64((uint64(m.BytesIn) / totalTime))
			_, humanBytesIn := gout.HumanSize(uint64(m.BytesIn))
			_, humanTotalTime := gout.HumanTime(totalTime)
			_, humanAvgSpeed := gout.HumanSize(avgSpeed)
			if e != nil && e.Error() == "EOF" {
				fmt.Printf("Read %8s in %10s @ %8s/sec\n",
					humanBytesIn,
					fmt.Sprintf("%10s", humanTotalTime),
					humanAvgSpeed)
				return n, e
			}
			fmt.Printf("Read %8s in %10s @ %8s/sec                  \r",
				humanBytesIn,
				fmt.Sprintf("%10s", humanTotalTime),
				humanAvgSpeed)
		}
		return n, e
	} else {
		return n, e
	}
}

func (m *MeteredPipe) Write(b []byte) (n int, e error) {
	if v, ok := m.IoObject.(PipeWriter); ok {
		n, e := v.Write(b)
		m.BytesOut += n
		if e != nil {
			fmt.Printf("Wrote %10d bytes in %10d seconds\n",
				m.BytesOut,
				(time.Now().Unix() - m.Started))
			return n, e
		}
		fmt.Printf("Wrote %10d bytes in %10d seconds\r",
			m.BytesOut,
			(time.Now().Unix() - m.Started))
		return n, e
	} else {
		return n, e
	}
}

func (m *MeteredPipe) Close() error {
	if v, ok := m.IoObject.(PipeCloser); ok {
		return v.Close()
	} else {
		return nil
	}
}

func NewMeteredPipe(o interface{}) *MeteredPipe {
	retv := &MeteredPipe{}
	if v, ok := o.(PipeReadWriteCloser); ok {
		retv.IoObject = v
	} else if v, ok := o.(PipeWriteCloser); ok {
		retv.IoObject = v
	} else if v, ok := o.(PipeReadCloser); ok {
		retv.IoObject = v
	} else {
		retv.IoObject = v
	}
	retv.Started = time.Now().Unix()
	return retv
}
