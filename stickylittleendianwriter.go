package byteio

// File automatically generated with ./gen.sh

import (
	"io"
	"math"
)

// StickyLittleEndianWriter wraps a io.Writer to provide methods
// to make it easier to Write fundamental types
type StickyLittleEndianWriter struct {
	io.Writer
	buffer [8]byte
	Err    error
	Count  int64
}

// Write implements the io.Writer interface
func (e *StickyLittleEndianWriter) Write(p []byte) (int, error) {
	if e.Err != nil {
		return 0, e.Err
	}
	var n int
	n, e.Err = e.Writer.Write(p)
	e.Count += int64(n)
	return n, e.Err
}

// WriteInt8 Writes a int8 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyLittleEndianWriter) WriteInt8(d int8) {
	if e.Err != nil {
		return
	}
	e.buffer[0] = byte(d)
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:1])
	e.Count += int64(n)
}

// WriteInt16 Writes a int16 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyLittleEndianWriter) WriteInt16(d int16) {
	if e.Err != nil {
		return
	}
	c := uint64(d)
	e.buffer = [8]byte{
		byte(c),
		byte(c >> 8),
		byte(c >> 16),
		byte(c >> 24),
		byte(c >> 32),
		byte(c >> 40),
		byte(c >> 48),
		byte(c >> 56),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:2])
	e.Count += int64(n)
}

// WriteInt32 Writes a int32 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyLittleEndianWriter) WriteInt32(d int32) {
	if e.Err != nil {
		return
	}
	c := uint64(d)
	e.buffer = [8]byte{
		byte(c),
		byte(c >> 8),
		byte(c >> 16),
		byte(c >> 24),
		byte(c >> 32),
		byte(c >> 40),
		byte(c >> 48),
		byte(c >> 56),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:4])
	e.Count += int64(n)
}

// WriteInt64 Writes a int64 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyLittleEndianWriter) WriteInt64(d int64) {
	if e.Err != nil {
		return
	}
	c := uint64(d)
	e.buffer = [8]byte{
		byte(c),
		byte(c >> 8),
		byte(c >> 16),
		byte(c >> 24),
		byte(c >> 32),
		byte(c >> 40),
		byte(c >> 48),
		byte(c >> 56),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:8])
	e.Count += int64(n)
}

// WriteUint8 Writes a uint8 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyLittleEndianWriter) WriteUint8(d uint8) {
	if e.Err != nil {
		return
	}
	e.buffer[0] = d
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:1])
	e.Count += int64(n)
}

// WriteUint16 Writes a uint16 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyLittleEndianWriter) WriteUint16(d uint16) {
	if e.Err != nil {
		return
	}
	c := uint64(d)
	e.buffer = [8]byte{
		byte(c),
		byte(c >> 8),
		byte(c >> 16),
		byte(c >> 24),
		byte(c >> 32),
		byte(c >> 40),
		byte(c >> 48),
		byte(c >> 56),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:2])
	e.Count += int64(n)
}

// WriteUint32 Writes a uint32 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyLittleEndianWriter) WriteUint32(d uint32) {
	if e.Err != nil {
		return
	}
	c := uint64(d)
	e.buffer = [8]byte{
		byte(c),
		byte(c >> 8),
		byte(c >> 16),
		byte(c >> 24),
		byte(c >> 32),
		byte(c >> 40),
		byte(c >> 48),
		byte(c >> 56),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:4])
	e.Count += int64(n)
}

// WriteUint64 Writes a uint64 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyLittleEndianWriter) WriteUint64(d uint64) {
	if e.Err != nil {
		return
	}
	e.buffer = [8]byte{
		byte(d),
		byte(d >> 8),
		byte(d >> 16),
		byte(d >> 24),
		byte(d >> 32),
		byte(d >> 40),
		byte(d >> 48),
		byte(d >> 56),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:8])
	e.Count += int64(n)
}

// WriteFloat32 Writes a float32 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyLittleEndianWriter) WriteFloat32(d float32) {
	if e.Err != nil {
		return
	}
	c := math.Float32bits(d)
	e.buffer = [8]byte{
		byte(c),
		byte(c >> 8),
		byte(c >> 16),
		byte(c >> 24),
		byte(c >> 32),
		byte(c >> 40),
		byte(c >> 48),
		byte(c >> 56),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:4])
	e.Count += int64(n)
}

// WriteFloat64 Writes a float64 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyLittleEndianWriter) WriteFloat64(d float64) {
	if e.Err != nil {
		return
	}
	c := math.Float64bits(d)
	e.buffer = [8]byte{
		byte(c),
		byte(c >> 8),
		byte(c >> 16),
		byte(c >> 24),
		byte(c >> 32),
		byte(c >> 40),
		byte(c >> 48),
		byte(c >> 56),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:8])
	e.Count += int64(n)
}
