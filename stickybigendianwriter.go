package byteio

// File automatically generated with ./gen.sh

import (
	"io"
	"math"
)

type StickyBigEndianWriter struct {
	io.Writer
	buffer [8]byte
	Err    error
	Count  int64
}

func (e *StickyBigEndianWriter) Write(p []byte) (int, error) {
	if e.Err != nil {
		return 0, e.Err
	}
	var n int
	n, e.Err = e.Writer.Write(p)
	e.Count += int64(n)
	return n, e.Err
}

func (e *StickyBigEndianWriter) WriteInt8(d int8) {
	if e.Err != nil {
		return
	}
	e.buffer = [8]byte{
		byte(d),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:1])
	e.Count += int64(n)
}

func (e *StickyBigEndianWriter) WriteInt16(d int16) {
	if e.Err != nil {
		return
	}
	c := uint16(d)
	e.buffer = [8]byte{
		byte(c >> 8),
		byte(c),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:2])
	e.Count += int64(n)
}

func (e *StickyBigEndianWriter) WriteInt32(d int32) {
	if e.Err != nil {
		return
	}
	c := uint32(d)
	e.buffer = [8]byte{
		byte(c >> 24),
		byte(c >> 16),
		byte(c >> 8),
		byte(c),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:4])
	e.Count += int64(n)
}

func (e *StickyBigEndianWriter) WriteInt64(d int64) {
	if e.Err != nil {
		return
	}
	c := uint64(d)
	e.buffer = [8]byte{
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

func (e *StickyBigEndianWriter) WriteUint8(d uint8) {
	if e.Err != nil {
		return
	}
	e.buffer = [8]byte{
		d,
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:1])
	e.Count += int64(n)
}

func (e *StickyBigEndianWriter) WriteUint16(d uint16) {
	if e.Err != nil {
		return
	}
	e.buffer = [8]byte{
		byte(d >> 8),
		byte(d),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:2])
	e.Count += int64(n)
}

func (e *StickyBigEndianWriter) WriteUint32(d uint32) {
	if e.Err != nil {
		return
	}
	e.buffer = [8]byte{
		byte(d >> 24),
		byte(d >> 16),
		byte(d >> 8),
		byte(d),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:4])
	e.Count += int64(n)
}

func (e *StickyBigEndianWriter) WriteUint64(d uint64) {
	if e.Err != nil {
		return
	}
	e.buffer = [8]byte{
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

func (e *StickyBigEndianWriter) WriteFloat32(d float32) {
	if e.Err != nil {
		return
	}
	c := math.Float32bits(d)
	e.buffer = [8]byte{
		byte(c >> 24),
		byte(c >> 16),
		byte(c >> 8),
		byte(c),
	}
	var n int
	n, e.Err = e.Writer.Write(e.buffer[:4])
	e.Count += int64(n)
}

func (e *StickyBigEndianWriter) WriteFloat64(d float64) {
	if e.Err != nil {
		return
	}
	c := math.Float64bits(d)
	e.buffer = [8]byte{
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
