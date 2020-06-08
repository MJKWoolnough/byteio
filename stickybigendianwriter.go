package byteio

// File automatically generated with ./gen.sh

import (
	"io"
	"math"
)

// StickyBigEndianWriter wraps a io.Writer to provide methods
// to make it easier to Write fundamental types
type StickyBigEndianWriter struct {
	io.Writer
	buffer [9]byte
	Err    error
	Count  int64
}

// Write implements the io.Writer interface
func (e *StickyBigEndianWriter) Write(p []byte) (int, error) {
	if e.Err != nil {
		return 0, e.Err
	}
	var n int
	n, e.Err = e.Writer.Write(p)
	e.Count += int64(n)
	return n, e.Err
}

// WriteBool Writes a boolean
func (e *StickyBigEndianWriter) WriteBool(b bool) {
	if b {
		e.WriteUint8(1)
	} else {
		e.WriteUint8(0)
	}
}

// WriteInt8 Writes a 8 bit int as a int8 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianWriter) WriteInt8(d int8) {
	if e.Err != nil {
		return
	}
	e.buffer[0] = byte(d)
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:1])
	e.Count += int64(n)
}

// WriteInt16 Writes a 16 bit int as a int16 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianWriter) WriteInt16(d int16) {
	if e.Err != nil {
		return
	}
	c := uint16(d)
	e.buffer = [9]byte{
		byte(c >> 8),
		byte(c),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:2])
	e.Count += int64(n)
}

// WriteInt24 Writes a 24 bit int as a int32 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianWriter) WriteInt24(d int32) {
	if e.Err != nil {
		return
	}
	c := uint32(d)
	e.buffer = [9]byte{
		byte(c >> 16),
		byte(c >> 8),
		byte(c),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:3])
	e.Count += int64(n)
}

// WriteInt32 Writes a 32 bit int as a int32 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianWriter) WriteInt32(d int32) {
	if e.Err != nil {
		return
	}
	c := uint32(d)
	e.buffer = [9]byte{
		byte(c >> 24),
		byte(c >> 16),
		byte(c >> 8),
		byte(c),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:4])
	e.Count += int64(n)
}

// WriteInt40 Writes a 40 bit int as a int64 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianWriter) WriteInt40(d int64) {
	if e.Err != nil {
		return
	}
	c := uint64(d)
	e.buffer = [9]byte{
		byte(c >> 32),
		byte(c >> 24),
		byte(c >> 16),
		byte(c >> 8),
		byte(c),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:5])
	e.Count += int64(n)
}

// WriteInt48 Writes a 48 bit int as a int64 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianWriter) WriteInt48(d int64) {
	if e.Err != nil {
		return
	}
	c := uint64(d)
	e.buffer = [9]byte{
		byte(c >> 40),
		byte(c >> 32),
		byte(c >> 24),
		byte(c >> 16),
		byte(c >> 8),
		byte(c),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:6])
	e.Count += int64(n)
}

// WriteInt56 Writes a 56 bit int as a int64 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianWriter) WriteInt56(d int64) {
	if e.Err != nil {
		return
	}
	c := uint64(d)
	e.buffer = [9]byte{
		byte(c >> 48),
		byte(c >> 40),
		byte(c >> 32),
		byte(c >> 24),
		byte(c >> 16),
		byte(c >> 8),
		byte(c),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:7])
	e.Count += int64(n)
}

// WriteInt64 Writes a 64 bit int as a int64 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianWriter) WriteInt64(d int64) {
	if e.Err != nil {
		return
	}
	c := uint64(d)
	e.buffer = [9]byte{
		byte(c >> 56),
		byte(c >> 48),
		byte(c >> 40),
		byte(c >> 32),
		byte(c >> 24),
		byte(c >> 16),
		byte(c >> 8),
		byte(c),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:8])
	e.Count += int64(n)
}

// WriteUint8 Writes a 8 bit uint as a uint8 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianWriter) WriteUint8(d uint8) {
	if e.Err != nil {
		return
	}
	e.buffer[0] = d
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:1])
	e.Count += int64(n)
}

// WriteUint16 Writes a 16 bit uint as a uint16 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianWriter) WriteUint16(d uint16) {
	if e.Err != nil {
		return
	}
	e.buffer = [9]byte{
		byte(d >> 8),
		byte(d),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:2])
	e.Count += int64(n)
}

// WriteUint24 Writes a 24 bit uint as a uint32 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianWriter) WriteUint24(d uint32) {
	if e.Err != nil {
		return
	}
	e.buffer = [9]byte{
		byte(d >> 16),
		byte(d >> 8),
		byte(d),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:3])
	e.Count += int64(n)
}

// WriteUint32 Writes a 32 bit uint as a uint32 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianWriter) WriteUint32(d uint32) {
	if e.Err != nil {
		return
	}
	e.buffer = [9]byte{
		byte(d >> 24),
		byte(d >> 16),
		byte(d >> 8),
		byte(d),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:4])
	e.Count += int64(n)
}

// WriteUint40 Writes a 40 bit uint as a uint64 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianWriter) WriteUint40(d uint64) {
	if e.Err != nil {
		return
	}
	e.buffer = [9]byte{
		byte(d >> 32),
		byte(d >> 24),
		byte(d >> 16),
		byte(d >> 8),
		byte(d),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:5])
	e.Count += int64(n)
}

// WriteUint48 Writes a 48 bit uint as a uint64 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianWriter) WriteUint48(d uint64) {
	if e.Err != nil {
		return
	}
	e.buffer = [9]byte{
		byte(d >> 40),
		byte(d >> 32),
		byte(d >> 24),
		byte(d >> 16),
		byte(d >> 8),
		byte(d),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:6])
	e.Count += int64(n)
}

// WriteUint56 Writes a 56 bit uint as a uint64 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianWriter) WriteUint56(d uint64) {
	if e.Err != nil {
		return
	}
	e.buffer = [9]byte{
		byte(d >> 48),
		byte(d >> 40),
		byte(d >> 32),
		byte(d >> 24),
		byte(d >> 16),
		byte(d >> 8),
		byte(d),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:7])
	e.Count += int64(n)
}

// WriteUint64 Writes a 64 bit uint as a uint64 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianWriter) WriteUint64(d uint64) {
	if e.Err != nil {
		return
	}
	e.buffer = [9]byte{
		byte(d >> 56),
		byte(d >> 48),
		byte(d >> 40),
		byte(d >> 32),
		byte(d >> 24),
		byte(d >> 16),
		byte(d >> 8),
		byte(d),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:8])
	e.Count += int64(n)
}

// WriteFloat32 Writes a 32 bit float as a float32 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianWriter) WriteFloat32(d float32) {
	if e.Err != nil {
		return
	}
	c := math.Float32bits(d)
	e.buffer = [9]byte{
		byte(c >> 24),
		byte(c >> 16),
		byte(c >> 8),
		byte(c),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:4])
	e.Count += int64(n)
}

// WriteFloat64 Writes a 64 bit float as a float64 using the underlying io.Writer
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianWriter) WriteFloat64(d float64) {
	if e.Err != nil {
		return
	}
	c := math.Float64bits(d)
	e.buffer = [9]byte{
		byte(c >> 56),
		byte(c >> 48),
		byte(c >> 40),
		byte(c >> 32),
		byte(c >> 24),
		byte(c >> 16),
		byte(c >> 8),
		byte(c),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:8])
	e.Count += int64(n)
}

// WriteBytes Writes a []byte
func (e *StickyBigEndianWriter) WriteBytes(d []byte) (int, error) {
	if e.Err != nil {
		return 0, e.Err
	}
	var n int
	n, e.Err = e.Write(d)
	e.Count += int64(n)
	return n, e.Err
}

// WriteBytesX Writes the length of the Bytes, using ReadUintX and then Writes the bytes
func (e *StickyBigEndianWriter) WriteBytesX(p []byte) {
	e.WriteUintX(uint64(len(p)))
	e.WriteBytes(p)
}

// WriteBytes8 Writes the length of the Bytes, using ReadUint8 and then Writes the bytes
func (e *StickyBigEndianWriter) WriteBytes8(p []byte) {
	e.WriteUint8(uint8(len(p)))
	e.WriteBytes(p)
}

// WriteBytes16 Writes the length of the Bytes, using ReadUint16 and then Writes the bytes
func (e *StickyBigEndianWriter) WriteBytes16(p []byte) {
	e.WriteUint16(uint16(len(p)))
	e.WriteBytes(p)
}

// WriteBytes24 Writes the length of the Bytes, using ReadUint24 and then Writes the bytes
func (e *StickyBigEndianWriter) WriteBytes24(p []byte) {
	e.WriteUint24(uint32(len(p)))
	e.WriteBytes(p)
}

// WriteBytes32 Writes the length of the Bytes, using ReadUint32 and then Writes the bytes
func (e *StickyBigEndianWriter) WriteBytes32(p []byte) {
	e.WriteUint32(uint32(len(p)))
	e.WriteBytes(p)
}

// WriteBytes40 Writes the length of the Bytes, using ReadUint40 and then Writes the bytes
func (e *StickyBigEndianWriter) WriteBytes40(p []byte) {
	e.WriteUint40(uint64(len(p)))
	e.WriteBytes(p)
}

// WriteBytes48 Writes the length of the Bytes, using ReadUint48 and then Writes the bytes
func (e *StickyBigEndianWriter) WriteBytes48(p []byte) {
	e.WriteUint48(uint64(len(p)))
	e.WriteBytes(p)
}

// WriteBytes56 Writes the length of the Bytes, using ReadUint56 and then Writes the bytes
func (e *StickyBigEndianWriter) WriteBytes56(p []byte) {
	e.WriteUint56(uint64(len(p)))
	e.WriteBytes(p)
}

// WriteBytes64 Writes the length of the Bytes, using ReadUint64 and then Writes the bytes
func (e *StickyBigEndianWriter) WriteBytes64(p []byte) {
	e.WriteUint64(uint64(len(p)))
	e.WriteBytes(p)
}

// WriteString Writes a string
func (e *StickyBigEndianWriter) WriteString(d string) (int, error) {
	if e.Err != nil {
		return 0, e.Err
	}
	var n int
	n, e.Err = io.WriteString(e.Writer, d)
	e.Count += int64(n)
	return n, e.Err
}

// WriteStringX Writes the length of the String, using ReadUintX and then Writes the bytes
func (e *StickyBigEndianWriter) WriteStringX(p string) {
	e.WriteUintX(uint64(len(p)))
	e.WriteString(p)
}

// WriteString8 Writes the length of the String, using ReadUint8 and then Writes the bytes
func (e *StickyBigEndianWriter) WriteString8(p string) {
	e.WriteUint8(uint8(len(p)))
	e.WriteString(p)
}

// WriteString16 Writes the length of the String, using ReadUint16 and then Writes the bytes
func (e *StickyBigEndianWriter) WriteString16(p string) {
	e.WriteUint16(uint16(len(p)))
	e.WriteString(p)
}

// WriteString24 Writes the length of the String, using ReadUint24 and then Writes the bytes
func (e *StickyBigEndianWriter) WriteString24(p string) {
	e.WriteUint24(uint32(len(p)))
	e.WriteString(p)
}

// WriteString32 Writes the length of the String, using ReadUint32 and then Writes the bytes
func (e *StickyBigEndianWriter) WriteString32(p string) {
	e.WriteUint32(uint32(len(p)))
	e.WriteString(p)
}

// WriteString40 Writes the length of the String, using ReadUint40 and then Writes the bytes
func (e *StickyBigEndianWriter) WriteString40(p string) {
	e.WriteUint40(uint64(len(p)))
	e.WriteString(p)
}

// WriteString48 Writes the length of the String, using ReadUint48 and then Writes the bytes
func (e *StickyBigEndianWriter) WriteString48(p string) {
	e.WriteUint48(uint64(len(p)))
	e.WriteString(p)
}

// WriteString56 Writes the length of the String, using ReadUint56 and then Writes the bytes
func (e *StickyBigEndianWriter) WriteString56(p string) {
	e.WriteUint56(uint64(len(p)))
	e.WriteString(p)
}

// WriteString64 Writes the length of the String, using ReadUint64 and then Writes the bytes
func (e *StickyBigEndianWriter) WriteString64(p string) {
	e.WriteUint64(uint64(len(p)))
	e.WriteString(p)
}

// WriteString0 Writes the bytes of the string ending with a 0 byte
func (e *StickyBigEndianWriter) WriteString0(p string) {
	e.WriteString(p)
	e.WriteUint8(0)
}
