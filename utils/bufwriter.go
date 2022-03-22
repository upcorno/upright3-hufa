package utils

//这个bufwriter与标准库bufio特别像.唯一的区别是func (b *Writer) Write(p []byte) (nn int, err error)方法,标准库中,当len(b.buf)<b.n+len(p)<2*len(b.buf)时,会遗留b.n+len(p)-len(b.buf)数量的内容在b.buf中,而本库中,任何时候,只要len(b.buf)<b.n+len(p),所有内容会全部刷盘.

import (
	"io"
)

const (
	defaultBufSize = 4096
)

// BufWriter implements buffering for an io.BufWriter object.
// If an error occurs writing to a BufWriter, no more data will be
// accepted and all subsequent writes, and Flush, will return the error.
// After all data has been written, the client should call the
// Flush method to guarantee all data has been forwarded to
// the underlying io.BufWriter.
type BufWriter struct {
	n   int
	err error
	buf []byte
	wr  io.Writer
}

// NewBufWriterSize returns a new Writer whose buffer has at least the specified
// size. If the argument io.Writer is already a Writer with large enough
// size, it returns the underlying Writer.
func NewBufWriterSize(w io.Writer, size int) *BufWriter {
	// Is it already a Writer?
	b, ok := w.(*BufWriter)
	if ok && len(b.buf) >= size {
		return b
	}
	if size <= 0 {
		size = defaultBufSize
	}
	return &BufWriter{
		buf: make([]byte, size),
		wr:  w,
	}
}

// NewBufWriter returns a new Writer whose buffer has the default size.
func NewBufWriter(w io.Writer) *BufWriter {
	return NewBufWriterSize(w, defaultBufSize)
}

// Size returns the size of the underlying buffer in bytes.
func (b *BufWriter) Size() int { return len(b.buf) }

// Reset discards any unflushed buffered data, clears any error, and
// resets b to write its output to w.
func (b *BufWriter) Reset(w io.Writer) {
	b.err = nil
	b.n = 0
	b.wr = w
}

// Flush writes any buffered data to the underlying io.Writer.
func (b *BufWriter) Flush() error {
	if b.err != nil {
		return b.err
	}
	if b.n == 0 {
		return nil
	}
	n, err := b.wr.Write(b.buf[0:b.n])
	if n < b.n && err == nil {
		err = io.ErrShortWrite
	}
	if err != nil {
		if n > 0 && n < b.n {
			copy(b.buf[0:b.n-n], b.buf[n:b.n])
		}
		b.n -= n
		b.err = err
		return err
	}
	b.n = 0
	return nil
}

// Available returns how many bytes are unused in the buffer.
func (b *BufWriter) Available() int { return len(b.buf) - b.n }

// Buffered returns the number of bytes that have been written into the current buffer.
func (b *BufWriter) Buffered() int { return b.n }

// Write writes the contents of p into the buffer.
// It returns the number of bytes written.
// If nn < len(p), it also returns an error explaining
// why the write is short.
func (b *BufWriter) Write(p []byte) (nn int, err error) {
	if len(p) > b.Available() && b.err == nil {
		if b.Buffered() != 0 {
			tmp := make([]byte, b.n+len(p))
			copy(tmp[:b.n], b.buf[:b.n])
			copy(tmp[b.n:], p)
			b.Reset(b.wr)
			nn, b.err = b.wr.Write(tmp)
		} else {
			nn, b.err = b.wr.Write(p)
		}
		if b.err == nil {
			return nn, nil
		}
	}
	if b.err != nil {
		return nn, b.err
	}
	n := copy(b.buf[b.n:], p)
	b.n += n
	nn += n
	return nn, nil
}
