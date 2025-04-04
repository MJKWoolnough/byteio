package byteio

// File automatically generated with ./gen.sh.

import (
	"io"
	"math"
)

// StickyLittleEndianWriter wraps a io.Writer to provide methods
// to make it easier to Write fundamental types.
type StickyLittleEndianWriter struct {
	io.Writer
	buffer [9]byte
	Err    error
	Count  int64
}

// GetCount returns any error received.
func (e *StickyLittleEndianWriter) GetError() error {
	return e.Err
}

// GetCount returns the number of bytes written.
func (e *StickyLittleEndianWriter) GetCount() int64 {
	return e.Count
}

// Write implements the io.Writer interface.
func (e *StickyLittleEndianWriter) Write(p []byte) (int, error) {
	if e.Err != nil {
		return 0, e.Err
	}

	n, err := e.Writer.Write(p)
	e.Err = err
	e.Count += int64(n)

	return n, e.Err
}

// WriteBool Writes a boolean.
func (e *StickyLittleEndianWriter) WriteBool(b bool) {
	if b {
		e.WriteUint8(1)
	} else {
		e.WriteUint8(0)
	}
}

// WriteInt8 Writes a 8 bit int as a int8 using the underlying io.Writer.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianWriter) WriteInt8(d int8) {
	if e.Err != nil {
		return
	}

	e.buffer[0] = byte(d)

	var n int

	n, e.Err = e.Writer.Write(e.buffer[:1])
	e.Count += int64(n)
}

// WriteInt16 Writes a 16 bit int as a int16 using the underlying io.Writer.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianWriter) WriteInt16(d int16) {
	if e.Err != nil {
		return
	}

	c := uint16(d)
	e.buffer = [9]byte{
		byte(c),
		byte(c >> 8),
	}

	var n int

	n, e.Err = e.Writer.Write(e.buffer[:2])
	e.Count += int64(n)
}

// WriteInt24 Writes a 24 bit int as a int32 using the underlying io.Writer.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianWriter) WriteInt24(d int32) {
	if e.Err != nil {
		return
	}

	c := uint32(d)
	e.buffer = [9]byte{
		byte(c),
		byte(c >> 8),
		byte(c >> 16),
	}

	var n int

	n, e.Err = e.Writer.Write(e.buffer[:3])
	e.Count += int64(n)
}

// WriteInt32 Writes a 32 bit int as a int32 using the underlying io.Writer.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianWriter) WriteInt32(d int32) {
	if e.Err != nil {
		return
	}

	c := uint32(d)
	e.buffer = [9]byte{
		byte(c),
		byte(c >> 8),
		byte(c >> 16),
		byte(c >> 24),
	}

	var n int

	n, e.Err = e.Writer.Write(e.buffer[:4])
	e.Count += int64(n)
}

// WriteInt40 Writes a 40 bit int as a int64 using the underlying io.Writer.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianWriter) WriteInt40(d int64) {
	if e.Err != nil {
		return
	}

	c := uint64(d)
	e.buffer = [9]byte{
		byte(c),
		byte(c >> 8),
		byte(c >> 16),
		byte(c >> 24),
		byte(c >> 32),
	}

	var n int

	n, e.Err = e.Writer.Write(e.buffer[:5])
	e.Count += int64(n)
}

// WriteInt48 Writes a 48 bit int as a int64 using the underlying io.Writer.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianWriter) WriteInt48(d int64) {
	if e.Err != nil {
		return
	}

	c := uint64(d)
	e.buffer = [9]byte{
		byte(c),
		byte(c >> 8),
		byte(c >> 16),
		byte(c >> 24),
		byte(c >> 32),
		byte(c >> 40),
	}

	var n int

	n, e.Err = e.Writer.Write(e.buffer[:6])
	e.Count += int64(n)
}

// WriteInt56 Writes a 56 bit int as a int64 using the underlying io.Writer.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianWriter) WriteInt56(d int64) {
	if e.Err != nil {
		return
	}

	c := uint64(d)
	e.buffer = [9]byte{
		byte(c),
		byte(c >> 8),
		byte(c >> 16),
		byte(c >> 24),
		byte(c >> 32),
		byte(c >> 40),
		byte(c >> 48),
	}

	var n int

	n, e.Err = e.Writer.Write(e.buffer[:7])
	e.Count += int64(n)
}

// WriteInt64 Writes a 64 bit int as a int64 using the underlying io.Writer.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianWriter) WriteInt64(d int64) {
	if e.Err != nil {
		return
	}

	c := uint64(d)
	e.buffer = [9]byte{
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

// WriteUint8 Writes a 8 bit uint as a uint8 using the underlying io.Writer.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianWriter) WriteUint8(d uint8) {
	if e.Err != nil {
		return
	}

	e.buffer[0] = d

	var n int

	n, e.Err = e.Writer.Write(e.buffer[:1])
	e.Count += int64(n)
}

// WriteUint16 Writes a 16 bit uint as a uint16 using the underlying io.Writer.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianWriter) WriteUint16(d uint16) {
	if e.Err != nil {
		return
	}

	e.buffer = [9]byte{
		byte(d),
		byte(d >> 8),
	}

	var n int

	n, e.Err = e.Writer.Write(e.buffer[:2])
	e.Count += int64(n)
}

// WriteUint24 Writes a 24 bit uint as a uint32 using the underlying io.Writer.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianWriter) WriteUint24(d uint32) {
	if e.Err != nil {
		return
	}

	e.buffer = [9]byte{
		byte(d),
		byte(d >> 8),
		byte(d >> 16),
	}

	var n int

	n, e.Err = e.Writer.Write(e.buffer[:3])
	e.Count += int64(n)
}

// WriteUint32 Writes a 32 bit uint as a uint32 using the underlying io.Writer.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianWriter) WriteUint32(d uint32) {
	if e.Err != nil {
		return
	}

	e.buffer = [9]byte{
		byte(d),
		byte(d >> 8),
		byte(d >> 16),
		byte(d >> 24),
	}

	var n int

	n, e.Err = e.Writer.Write(e.buffer[:4])
	e.Count += int64(n)
}

// WriteUint40 Writes a 40 bit uint as a uint64 using the underlying io.Writer.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianWriter) WriteUint40(d uint64) {
	if e.Err != nil {
		return
	}

	e.buffer = [9]byte{
		byte(d),
		byte(d >> 8),
		byte(d >> 16),
		byte(d >> 24),
		byte(d >> 32),
	}

	var n int

	n, e.Err = e.Writer.Write(e.buffer[:5])
	e.Count += int64(n)
}

// WriteUint48 Writes a 48 bit uint as a uint64 using the underlying io.Writer.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianWriter) WriteUint48(d uint64) {
	if e.Err != nil {
		return
	}

	e.buffer = [9]byte{
		byte(d),
		byte(d >> 8),
		byte(d >> 16),
		byte(d >> 24),
		byte(d >> 32),
		byte(d >> 40),
	}

	var n int

	n, e.Err = e.Writer.Write(e.buffer[:6])
	e.Count += int64(n)
}

// WriteUint56 Writes a 56 bit uint as a uint64 using the underlying io.Writer.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianWriter) WriteUint56(d uint64) {
	if e.Err != nil {
		return
	}

	e.buffer = [9]byte{
		byte(d),
		byte(d >> 8),
		byte(d >> 16),
		byte(d >> 24),
		byte(d >> 32),
		byte(d >> 40),
		byte(d >> 48),
	}

	var n int

	n, e.Err = e.Writer.Write(e.buffer[:7])
	e.Count += int64(n)
}

// WriteUint64 Writes a 64 bit uint as a uint64 using the underlying io.Writer.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianWriter) WriteUint64(d uint64) {
	if e.Err != nil {
		return
	}

	e.buffer = [9]byte{
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

// WriteFloat32 Writes a 32 bit float as a float32 using the underlying io.Writer.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianWriter) WriteFloat32(d float32) {
	if e.Err != nil {
		return
	}

	c := math.Float32bits(d)
	e.buffer = [9]byte{
		byte(c),
		byte(c >> 8),
		byte(c >> 16),
		byte(c >> 24),
	}

	var n int

	n, e.Err = e.Writer.Write(e.buffer[:4])
	e.Count += int64(n)
}

// WriteFloat64 Writes a 64 bit float as a float64 using the underlying io.Writer.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianWriter) WriteFloat64(d float64) {
	if e.Err != nil {
		return
	}

	c := math.Float64bits(d)
	e.buffer = [9]byte{
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

// WriteBytes Writes a []byte.
func (e *StickyLittleEndianWriter) WriteBytes(d []byte) {
	if e.Err != nil {
		return
	}

	n, err := e.Write(d)
	e.Count += int64(n)
	e.Err = err
}

// WriteBytesX Writes the length of the Bytes, using WriteUintX and then Writes the bytes.
func (e *StickyLittleEndianWriter) WriteBytesX(p []byte) {
	e.WriteUintX(uint64(len(p)))
	e.Write(p)
}

// WriteBytes8 Writes the length of the Bytes, using WriteUint8 and then Writes the bytes.
func (e *StickyLittleEndianWriter) WriteBytes8(p []byte) {
	e.WriteUint8(uint8(len(p)))
	e.Write(p)
}

// WriteBytes16 Writes the length of the Bytes, using WriteUint16 and then Writes the bytes.
func (e *StickyLittleEndianWriter) WriteBytes16(p []byte) {
	e.WriteUint16(uint16(len(p)))
	e.Write(p)
}

// WriteBytes24 Writes the length of the Bytes, using WriteUint24 and then Writes the bytes.
func (e *StickyLittleEndianWriter) WriteBytes24(p []byte) {
	e.WriteUint24(uint32(len(p)))
	e.Write(p)
}

// WriteBytes32 Writes the length of the Bytes, using WriteUint32 and then Writes the bytes.
func (e *StickyLittleEndianWriter) WriteBytes32(p []byte) {
	e.WriteUint32(uint32(len(p)))
	e.Write(p)
}

// WriteBytes40 Writes the length of the Bytes, using WriteUint40 and then Writes the bytes.
func (e *StickyLittleEndianWriter) WriteBytes40(p []byte) {
	e.WriteUint40(uint64(len(p)))
	e.Write(p)
}

// WriteBytes48 Writes the length of the Bytes, using WriteUint48 and then Writes the bytes.
func (e *StickyLittleEndianWriter) WriteBytes48(p []byte) {
	e.WriteUint48(uint64(len(p)))
	e.Write(p)
}

// WriteBytes56 Writes the length of the Bytes, using WriteUint56 and then Writes the bytes.
func (e *StickyLittleEndianWriter) WriteBytes56(p []byte) {
	e.WriteUint56(uint64(len(p)))
	e.Write(p)
}

// WriteBytes64 Writes the length of the Bytes, using WriteUint64 and then Writes the bytes.
func (e *StickyLittleEndianWriter) WriteBytes64(p []byte) {
	e.WriteUint64(uint64(len(p)))
	e.Write(p)
}

// WriteString Writes a string.
func (e *StickyLittleEndianWriter) WriteString(d string) (int, error) {
	if e.Err != nil {
		return 0, e.Err
	}

	n, err := io.WriteString(e.Writer, d)
	e.Count += int64(n)
	e.Err = err

	return n, err
}

// WriteStringX Writes the length of the String, using WriteUintX and then Writes the bytes.
func (e *StickyLittleEndianWriter) WriteStringX(p string) {
	e.WriteUintX(uint64(len(p)))
	e.WriteString(p)
}

// WriteString8 Writes the length of the String, using WriteUint8 and then Writes the bytes.
func (e *StickyLittleEndianWriter) WriteString8(p string) {
	e.WriteUint8(uint8(len(p)))
	e.WriteString(p)
}

// WriteString16 Writes the length of the String, using WriteUint16 and then Writes the bytes.
func (e *StickyLittleEndianWriter) WriteString16(p string) {
	e.WriteUint16(uint16(len(p)))
	e.WriteString(p)
}

// WriteString24 Writes the length of the String, using WriteUint24 and then Writes the bytes.
func (e *StickyLittleEndianWriter) WriteString24(p string) {
	e.WriteUint24(uint32(len(p)))
	e.WriteString(p)
}

// WriteString32 Writes the length of the String, using WriteUint32 and then Writes the bytes.
func (e *StickyLittleEndianWriter) WriteString32(p string) {
	e.WriteUint32(uint32(len(p)))
	e.WriteString(p)
}

// WriteString40 Writes the length of the String, using WriteUint40 and then Writes the bytes.
func (e *StickyLittleEndianWriter) WriteString40(p string) {
	e.WriteUint40(uint64(len(p)))
	e.WriteString(p)
}

// WriteString48 Writes the length of the String, using WriteUint48 and then Writes the bytes.
func (e *StickyLittleEndianWriter) WriteString48(p string) {
	e.WriteUint48(uint64(len(p)))
	e.WriteString(p)
}

// WriteString56 Writes the length of the String, using WriteUint56 and then Writes the bytes.
func (e *StickyLittleEndianWriter) WriteString56(p string) {
	e.WriteUint56(uint64(len(p)))
	e.WriteString(p)
}

// WriteString64 Writes the length of the String, using WriteUint64 and then Writes the bytes.
func (e *StickyLittleEndianWriter) WriteString64(p string) {
	e.WriteUint64(uint64(len(p)))
	e.WriteString(p)
}

// WriteString0 Writes the bytes of the string ending with a 0 byte.
func (e *StickyLittleEndianWriter) WriteString0(p string) {
	e.WriteString(p)
	e.WriteUint8(0)
}
