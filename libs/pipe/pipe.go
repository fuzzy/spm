package pipe

/*
 Author: Mike 'Fuzzy' Partin
 Copyright: (c) 2016-2018
 Email: fuzzy@fumanchu.org
 License: See LICENSE.md for details
*/

import (
	"fmt"
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

/*
 * Object:  MeterReader
 * Methods: Read(b []byte) (int, error)
 */
type MeterReader struct {
	IoObject PipeReader
	BytesIn  int
	Started  int64
}

func (m *MeterReader) Read(b []byte) (n int, e error) {
	rv, re := m.IoObject.Read(b)
	m.BytesIn += rv
	if re != nil && re.Error() == "EOF" {
		fmt.Printf("Read %10d bytes in %10d seconds\n", m.BytesIn, (time.Now().Unix() - m.Started))
		return rv, re
	}
	fmt.Printf("Read %10d bytes in %10d seconds\r", m.BytesIn, (time.Now().Unix() - m.Started))
	return rv, re
}

/*
 * Object:  MeterReader
 * Methods: Read(b []byte) (int, error)
 */
type MeterWriter struct {
	IoObject PipeWriter
	BytesOut int
	Started  int64
}

func (m *MeterWriter) Write(b []byte) (n int, e error) {
	rv, re := m.IoObject.Write(b)
	m.BytesOut += rv
	if re != nil {
		fmt.Printf("Wrote %10d bytes in %10d seconds\n", m.BytesOut, (time.Now().Unix() - m.Started))
		return rv, re
	}
	fmt.Printf("Wrote %10d bytes in %10d seconds\r", m.BytesOut, (time.Now().Unix() - m.Started))
	return rv, re
}

/*
 * Object:  MeterReader
 * Methods: Read(b []byte) (int, error)
 */
type MeterReadWriter struct {
	IoObject PipeReadWriter
	BytesIn  int
	BytesOut int
	Started  int64
}

func (m *MeterReadWriter) Read(b []byte) (n int, e error) {
	rv, re := m.IoObject.Read(b)
	m.BytesIn += rv
	if re != nil && re.Error() == "EOF" {
		fmt.Printf("Read %10d bytes in %10d seconds\n", m.BytesIn, (time.Now().Unix() - m.Started))
		return rv, re
	}
	fmt.Printf("Read %10d bytes in %10d seconds\r", m.BytesIn, (time.Now().Unix() - m.Started))
	return rv, re
}

func (m *MeterReadWriter) Write(b []byte) (n int, e error) {
	rv, re := m.IoObject.Write(b)
	m.BytesOut += rv
	if re != nil {
		fmt.Printf("Wrote %10d bytes in %10d seconds\n", m.BytesOut, (time.Now().Unix() - m.Started))
		return rv, re
	}
	fmt.Printf("Wrote %10d bytes in %10d seconds\r", m.BytesOut, (time.Now().Unix() - m.Started))
	return rv, re
}

/*
 * Object:  MeterReadCloser
 * Methods: Read(b []byte) (int, error)
 *          Close() error
 */
type MeterReadCloser struct {
	IoObject PipeReadCloser
	BytesIn  int
	Started  int64
}

func (m *MeterReadCloser) Read(b []byte) (n int, e error) {
	rv, re := m.IoObject.Read(b)
	m.BytesIn += rv
	if re != nil && re.Error() == "EOF" {
		fmt.Printf("Read %10d bytes in %10d seconds @ %10d bytes/sec\n",
			m.BytesIn,
			(time.Now().Unix() - m.Started),
			(int64(m.BytesIn) / (time.Now().Unix() - m.Started)))
		return rv, re
	}
	fmt.Printf("Read %10d bytes in %10d seconds\r", m.BytesIn, (time.Now().Unix() - m.Started))
	return rv, re
}

func (m *MeterReadCloser) Close() error {
	return m.IoObject.Close()
}

func NewMeterReadCloser(o PipeReadCloser) *MeterReadCloser {
	return &MeterReadCloser{o, 0, time.Now().Unix()}
}

/*
 * Object:  MeterWriteCloser
 * Methods: Write(b []byte) (int, error)
 *          Close() error
 */
type MeterWriteCloser struct {
	IoObject PipeWriteCloser
	BytesOut int
	Started  int64
}

/*
 */
type MeterReadWriteCloser struct {
	IoObject PipeReadWriteCloser
	BytesIn  int
	BytesOut int
	Started  int64
}
