package pipe

/*
 Author: Mike 'Fuzzy' Partin
 Copyright: (c) 2016-2018
 Email: fuzzy@fumanchu.org
 License: See LICENSE.md for details
*/

import (
	"github.com/fuzzy/spm/gout"
	"time"
)

// Object definitions

type MeteredPipe struct {
	IoObject   interface{}
	BytesIn    uint64
	BytesOut   uint64
	BytesTotal uint64
	Started    int64
	LastUpdate int64
}

func (m *MeteredPipe) output() gout.String {
	// we need this for scope
	var remain_n float64
	var remain_s gout.String
	var retval gout.String

	// First the math bits
	elapsed_n := uint64(time.Now().Unix() - m.Started)
	speed_n := (m.BytesIn / elapsed_n)
	if m.BytesTotal > 0 {
		remain_n = ((float64(m.BytesTotal) - float64(m.BytesIn)) / float64(speed_n))
	}

	// Now the string bits
	elapsed_s := gout.HumanTime(elapsed_n)
	speed_s := gout.HumanSize(speed_n)
	if m.BytesTotal > 0 {
		remain_s = gout.HumanTime(uint64(remain_n))
	} else {
		remain_s = ""
	}

	// now build our output string
	if m.BytesTotal > 0 {
		retval = StrAppend(retval, remain_s)
		retval = StrAppend(retval, " ")
		retval = StrAppend(retval, gout.ProgressBar(float64(m.BytesIn), float64(m.BytesTotal)))
		retval = StrAppend(retval, " ")
	}
	retval = StrAppend(retval, gout.HumanSize(m.BytesIn))
	retval = StrAppend(retval, " in ")
	retval = StrAppend(retval, elapsed_s)
	retval = StrAppend(retval, " @ ")
	retval = StrAppend(retval, speed_s)
	retval = StrAppend(retval, "/sec")

	// and finally hand it back
	return retval
}

func (m *MeteredPipe) Read(b []byte) (n int, e error) {
	if v, ok := m.IoObject.(PipeReader); ok {
		n, e := v.Read(b)
		m.BytesIn += uint64(n)
		if (time.Now().Unix() - m.Started) >= 1 {
			if e != nil && e.Error() == "EOF" {
				gout.Println(string(m.output()))
				return n, e
			}
			if time.Now().Unix()-m.LastUpdate >= 1 {
				gout.Printr(string(m.output()))
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
		m.BytesOut += uint64(n)
		if e != nil {
			gout.Println(string(m.output()))
		} else {
			gout.Printr(string(m.output()))
		}
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
	retv.BytesTotal = uint64(s)
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
