package pipe

/*
 Author: Mike 'Fuzzy' Partin
 Copyright: (c) 2016-2018
 Email: fuzzy@fumanchu.org
 License: See LICENSE.md for details
*/

import (
	"fmt"
	"github.com/fuzzy/spm/gout"
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
	IoObject   interface{}
	BytesIn    int
	BytesOut   int
	BytesTotal int64
	Started    int64
	LastUpdate int64
}

func (m *MeteredPipe) Read(b []byte) (n int, e error) {
	if v, ok := m.IoObject.(PipeReader); ok {
		n, e := v.Read(b)
		m.BytesIn += n
		if (time.Now().Unix() - m.Started) >= 1 {
			totalTime := uint64(time.Now().Unix() - m.Started)
			avgSpeed := uint64((uint64(m.BytesIn) / totalTime))
			percentDone := (float64(m.BytesIn) / float64(m.BytesTotal)) * float64(100)

			if e != nil && e.Error() == "EOF" {
				gout.Println(fmt.Sprintf("%6.02f%% %s %s in %s @ %s/sec",
					percentDone,
					gout.ProgressBar(float64(m.BytesIn), float64(m.BytesTotal)),
					gout.HumanSize(uint64(m.BytesIn)),
					gout.HumanTime(totalTime),
					gout.HumanSize(avgSpeed)))
				return n, e
			}
			if time.Now().Unix() - m.LastUpdate >= 1 {
				gout.Printr(fmt.Sprintf("%6.02f%% %s %s in %s @ %s/sec",
					percentDone,
					gout.ProgressBar(float64(m.BytesIn), float64(m.BytesTotal)),
					gout.HumanSize(uint64(m.BytesIn)),
					gout.HumanTime(totalTime),
					gout.HumanSize(avgSpeed)))
				m.LastUpdate = time.Now().Unix()
			}
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

func NewMeteredPipe(o interface{}, s int64) *MeteredPipe {
	retv := &MeteredPipe{}
	retv.BytesTotal = s
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
