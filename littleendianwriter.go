package byteio

// File automatically generated with ./gen.sh

import (
	"io"
	"math"
)

type LittleEndianWriter struct {
	io.Writer
	buffer [8]byte
}

func (e *LittleEndianWriter) WriteInt8(d int8) (int, error) {
	e.buffer = [8]byte{
		byte(d),
	}
	return e.Writer.Write(e.buffer[:1])
}

func (e *LittleEndianWriter) WriteInt16(d int16) (int, error) {
	c := uint16(d)
	e.buffer = [8]byte{
		byte(c),
		byte(c >> 8),
	}
	return e.Writer.Write(e.buffer[:2])
}

func (e *LittleEndianWriter) WriteInt32(d int32) (int, error) {
	c := uint32(d)
	e.buffer = [8]byte{
		byte(c),
		byte(c >> 8),
		byte(c >> 16),
		byte(c >> 24),
	}
	return e.Writer.Write(e.buffer[:4])
}

func (e *LittleEndianWriter) WriteInt64(d int64) (int, error) {
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
	return e.Writer.Write(e.buffer[:8])
}

func (e *LittleEndianWriter) WriteUint8(d uint8) (int, error) {
	e.buffer = [8]byte{
		d,
	}
	return e.Writer.Write(e.buffer[:1])
}

func (e *LittleEndianWriter) WriteUint16(d uint16) (int, error) {
	e.buffer = [8]byte{
		byte(d),
		byte(d >> 8),
	}
	return e.Writer.Write(e.buffer[:2])
}

func (e *LittleEndianWriter) WriteUint32(d uint32) (int, error) {
	e.buffer = [8]byte{
		byte(d),
		byte(d >> 8),
		byte(d >> 16),
		byte(d >> 24),
	}
	return e.Writer.Write(e.buffer[:4])
}

func (e *LittleEndianWriter) WriteUint64(d uint64) (int, error) {
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
	return e.Writer.Write(e.buffer[:8])
}

func (e *LittleEndianWriter) WriteFloat32(d float32) (int, error) {
	c := math.Float32bits(d)
	e.buffer = [8]byte{
		byte(c),
		byte(c >> 8),
		byte(c >> 16),
		byte(c >> 24),
	}
	return e.Writer.Write(e.buffer[:4])
}

func (e *LittleEndianWriter) WriteFloat64(d float64) (int, error) {
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
	return e.Writer.Write(e.buffer[:8])
}
