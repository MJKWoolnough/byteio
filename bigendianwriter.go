package byteio

// File automatically generated with ./gen.sh

import (
	"io"
	"math"
)

// BigEndianWriter wraps a io.Writer to provide methods
// to make it easier to Write fundamental types
type BigEndianWriter struct {
	io.Writer
	buffer [8]byte
}

// WriteInt8 Writes a int8 using the underlying io.Writer
func (e *BigEndianWriter) WriteInt8(d int8) (int, error) {
	e.buffer[0] = byte(d)
	return e.Writer.Write(e.buffer[:1])
}

// WriteInt16 Writes a int16 using the underlying io.Writer
func (e *BigEndianWriter) WriteInt16(d int16) (int, error) {
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
	return e.Writer.Write(e.buffer[6:])
}

// WriteInt32 Writes a int32 using the underlying io.Writer
func (e *BigEndianWriter) WriteInt32(d int32) (int, error) {
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
	return e.Writer.Write(e.buffer[4:])
}

// WriteInt64 Writes a int64 using the underlying io.Writer
func (e *BigEndianWriter) WriteInt64(d int64) (int, error) {
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
	return e.Writer.Write(e.buffer[0:])
}

// WriteUint8 Writes a uint8 using the underlying io.Writer
func (e *BigEndianWriter) WriteUint8(d uint8) (int, error) {
	e.buffer[0] = d
	return e.Writer.Write(e.buffer[:1])
}

// WriteUint16 Writes a uint16 using the underlying io.Writer
func (e *BigEndianWriter) WriteUint16(d uint16) (int, error) {
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
	return e.Writer.Write(e.buffer[6:])
}

// WriteUint32 Writes a uint32 using the underlying io.Writer
func (e *BigEndianWriter) WriteUint32(d uint32) (int, error) {
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
	return e.Writer.Write(e.buffer[4:])
}

// WriteUint64 Writes a uint64 using the underlying io.Writer
func (e *BigEndianWriter) WriteUint64(d uint64) (int, error) {
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
	return e.Writer.Write(e.buffer[0:])
}

// WriteFloat32 Writes a float32 using the underlying io.Writer
func (e *BigEndianWriter) WriteFloat32(d float32) (int, error) {
	c := math.Float32bits(d)
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
	return e.Writer.Write(e.buffer[4:])
}

// WriteFloat64 Writes a float64 using the underlying io.Writer
func (e *BigEndianWriter) WriteFloat64(d float64) (int, error) {
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
	return e.Writer.Write(e.buffer[0:])
}

// WriteString Writes a string
func (e *BigEndianWriter) WriteString(str string) (int, error) {
	return io.WriteString(e.Writer, str)
}

// WriteStringX Writes the length of the string, using ReadUintX and then reads the bytes of the string
func (e *BigEndianWriter) WriteStringX(str string) (int, error) {
	n, err := e.WriteUintX(uint64(len(str)))
	if err != nil {
		return n, err
	}
	m, err := e.WriteString(str)
	return n + m, err
}

// WriteString8 Writes the length of the string, using ReadUint8 and then reads the bytes of the string
func (e *BigEndianWriter) WriteString8(str string) (int, error) {
	n, err := e.WriteUint8(uint8(len(str)))
	if err != nil {
		return n, err
	}
	m, err := e.WriteString(str)
	return n + m, err
}

// WriteString16 Writes the length of the string, using ReadUint16 and then reads the bytes of the string
func (e *BigEndianWriter) WriteString16(str string) (int, error) {
	n, err := e.WriteUint16(uint16(len(str)))
	if err != nil {
		return n, err
	}
	m, err := e.WriteString(str)
	return n + m, err
}

// WriteString32 Writes the length of the string, using ReadUint32 and then reads the bytes of the string
func (e *BigEndianWriter) WriteString32(str string) (int, error) {
	n, err := e.WriteUint32(uint32(len(str)))
	if err != nil {
		return n, err
	}
	m, err := e.WriteString(str)
	return n + m, err
}

// WriteString64 Writes the length of the string, using ReadUint64 and then reads the bytes of the string
func (e *BigEndianWriter) WriteString64(str string) (int, error) {
	n, err := e.WriteUint64(uint64(len(str)))
	if err != nil {
		return n, err
	}
	m, err := e.WriteString(str)
	return n + m, err
}
