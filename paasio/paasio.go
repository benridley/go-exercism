package paasio

import (
	"io"
	"sync"
)

type ioReadWriter interface {
	io.Reader
	io.Writer
}

// PaasWriter implements the WriteCount interface
type PaasWriter struct {
	writer io.Writer
	bytes  *int64
	nops   *int
	mu     *sync.Mutex
}

// PaasReader implements the ReadCount Interface
type PaasReader struct {
	reader io.Reader
	bytes  *int64
	nops   *int
	mu     *sync.Mutex
}

// PaasReadWriter implements both the ReadCount and WriteCount interface
type PaasReadWriter struct {
	pr ReadCounter
	pw WriteCounter
}

// WriteCount returns the number of bytes and operations a PaasWriter has completed
func (pw PaasWriter) WriteCount() (n int64, nops int) {
	return *pw.bytes, *pw.nops
}

// NewWriteCounter creates a WriteCounter (PaasWriter) from a generic writer
func NewWriteCounter(w io.Writer) WriteCounter {
	b := int64(0)
	n := 0
	mu := new(sync.Mutex)
	return PaasWriter{
		writer: w,
		bytes:  &b,
		nops:   &n,
		mu:     mu,
	}
}

// Write data to the underlying stream
func (pw PaasWriter) Write(p []byte) (n int, err error) {
	n, err = pw.writer.Write(p)
	if err == nil {
		pw.mu.Lock()
		defer pw.mu.Unlock()
		*pw.nops++
		*pw.bytes += int64(n)
	}
	return n, err
}

// NewReadCounter creates a read counter from the underlying io reader.
func NewReadCounter(r io.Reader) ReadCounter {
	b := int64(0)
	n := 0
	mu := new(sync.Mutex)
	return PaasReader{
		reader: r,
		bytes:  &b,
		nops:   &n,
		mu:     mu,
	}
}

func (pr PaasReader) Read(p []byte) (n int, err error) {
	n, err = pr.reader.Read(p)
	if err == nil {
		pr.mu.Lock()
		defer pr.mu.Unlock()
		*pr.nops++
		*pr.bytes += int64(n)
	}
	return n, err
}

// ReadCount returns the number of bytes and operations a PaasReader has completed.
func (pr PaasReader) ReadCount() (int64, int) {
	pr.mu.Lock()
	defer pr.mu.Unlock()
	return *pr.bytes, *pr.nops
}

// NewReadWriteCounter creates a read/write counter from an underlying read/write stream.
func NewReadWriteCounter(rw ioReadWriter) ReadWriteCounter {
	return PaasReadWriter{
		pr: NewReadCounter(rw),
		pw: NewWriteCounter(rw),
	}
}

func (rw PaasReadWriter) Read(p []byte) (n int, err error) {
	n, err = rw.pr.Read(p)
	return n, err
}

func (rw PaasReadWriter) Write(p []byte) (n int, err error) {
	n, err = rw.pw.Write(p)
	return n, err
}

// ReadCount returns the number of reads a readwrite counter has performed.
func (rw PaasReadWriter) ReadCount() (n int64, nops int) {
	return rw.pr.ReadCount()
}

// WriteCount returns the number of writes a readwrite counter has performed.
func (rw PaasReadWriter) WriteCount() (n int64, nops int) {
	return rw.pw.WriteCount()
}
