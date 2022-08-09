package byteio

// File automatically generated with ./gen.sh

import (
	"io"
	"math"
)

// LittleEndianWriter wraps a io.Writer to provide methods
// to make it easier to Write fundamental types
type LittleEndianWriter struct {
	io.Writer
	buffer [9]byte
}

// WriteBool Writes a boolean
func (e *LittleEndianWriter) WriteBool(b bool) (int, error) {
	if b {
		return e.WriteUint8(1)
	} else {
		return e.WriteUint8(0)
	}
}

// WriteInt8 Writes a 8 bit int as a int8 using the underlying io.Writer
func (e *LittleEndianWriter) WriteInt8(d int8) (int, error) {
	e.buffer[0] = byte(d)
	return e.Writer.Write(e.buffer[:1])
}

// WriteInt16 Writes a 16 bit int as a int16 using the underlying io.Writer
func (e *LittleEndianWriter) WriteInt16(d int16) (int, error) {
	c := uint16(d)
	e.buffer = [9]byte{
		byte(c),
		byte(c >> 8),
	}
	return e.Writer.Write(e.buffer[:2])
}

// WriteInt24 Writes a 24 bit int as a int32 using the underlying io.Writer
func (e *LittleEndianWriter) WriteInt24(d int32) (int, error) {
	c := uint32(d)
	e.buffer = [9]byte{
		byte(c),
		byte(c >> 8),
		byte(c >> 16),
	}
	return e.Writer.Write(e.buffer[:3])
}

// WriteInt32 Writes a 32 bit int as a int32 using the underlying io.Writer
func (e *LittleEndianWriter) WriteInt32(d int32) (int, error) {
	c := uint32(d)
	e.buffer = [9]byte{
		byte(c),
		byte(c >> 8),
		byte(c >> 16),
		byte(c >> 24),
	}
	return e.Writer.Write(e.buffer[:4])
}

// WriteInt40 Writes a 40 bit int as a int64 using the underlying io.Writer
func (e *LittleEndianWriter) WriteInt40(d int64) (int, error) {
	c := uint64(d)
	e.buffer = [9]byte{
		byte(c),
		byte(c >> 8),
		byte(c >> 16),
		byte(c >> 24),
		byte(c >> 32),
	}
	return e.Writer.Write(e.buffer[:5])
}

// WriteInt48 Writes a 48 bit int as a int64 using the underlying io.Writer
func (e *LittleEndianWriter) WriteInt48(d int64) (int, error) {
	c := uint64(d)
	e.buffer = [9]byte{
		byte(c),
		byte(c >> 8),
		byte(c >> 16),
		byte(c >> 24),
		byte(c >> 32),
		byte(c >> 40),
	}
	return e.Writer.Write(e.buffer[:6])
}

// WriteInt56 Writes a 56 bit int as a int64 using the underlying io.Writer
func (e *LittleEndianWriter) WriteInt56(d int64) (int, error) {
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
	return e.Writer.Write(e.buffer[:7])
}

// WriteInt64 Writes a 64 bit int as a int64 using the underlying io.Writer
func (e *LittleEndianWriter) WriteInt64(d int64) (int, error) {
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
	return e.Writer.Write(e.buffer[:8])
}

// WriteUint8 Writes a 8 bit uint as a uint8 using the underlying io.Writer
func (e *LittleEndianWriter) WriteUint8(d uint8) (int, error) {
	e.buffer[0] = d
	return e.Writer.Write(e.buffer[:1])
}

// WriteUint16 Writes a 16 bit uint as a uint16 using the underlying io.Writer
func (e *LittleEndianWriter) WriteUint16(d uint16) (int, error) {
	e.buffer = [9]byte{
		byte(d),
		byte(d >> 8),
	}
	return e.Writer.Write(e.buffer[:2])
}

// WriteUint24 Writes a 24 bit uint as a uint32 using the underlying io.Writer
func (e *LittleEndianWriter) WriteUint24(d uint32) (int, error) {
	e.buffer = [9]byte{
		byte(d),
		byte(d >> 8),
		byte(d >> 16),
	}
	return e.Writer.Write(e.buffer[:3])
}

// WriteUint32 Writes a 32 bit uint as a uint32 using the underlying io.Writer
func (e *LittleEndianWriter) WriteUint32(d uint32) (int, error) {
	e.buffer = [9]byte{
		byte(d),
		byte(d >> 8),
		byte(d >> 16),
		byte(d >> 24),
	}
	return e.Writer.Write(e.buffer[:4])
}

// WriteUint40 Writes a 40 bit uint as a uint64 using the underlying io.Writer
func (e *LittleEndianWriter) WriteUint40(d uint64) (int, error) {
	e.buffer = [9]byte{
		byte(d),
		byte(d >> 8),
		byte(d >> 16),
		byte(d >> 24),
		byte(d >> 32),
	}
	return e.Writer.Write(e.buffer[:5])
}

// WriteUint48 Writes a 48 bit uint as a uint64 using the underlying io.Writer
func (e *LittleEndianWriter) WriteUint48(d uint64) (int, error) {
	e.buffer = [9]byte{
		byte(d),
		byte(d >> 8),
		byte(d >> 16),
		byte(d >> 24),
		byte(d >> 32),
		byte(d >> 40),
	}
	return e.Writer.Write(e.buffer[:6])
}

// WriteUint56 Writes a 56 bit uint as a uint64 using the underlying io.Writer
func (e *LittleEndianWriter) WriteUint56(d uint64) (int, error) {
	e.buffer = [9]byte{
		byte(d),
		byte(d >> 8),
		byte(d >> 16),
		byte(d >> 24),
		byte(d >> 32),
		byte(d >> 40),
		byte(d >> 48),
	}
	return e.Writer.Write(e.buffer[:7])
}

// WriteUint64 Writes a 64 bit uint as a uint64 using the underlying io.Writer
func (e *LittleEndianWriter) WriteUint64(d uint64) (int, error) {
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
	return e.Writer.Write(e.buffer[:8])
}

// WriteFloat32 Writes a 32 bit float as a float32 using the underlying io.Writer
func (e *LittleEndianWriter) WriteFloat32(d float32) (int, error) {
	c := math.Float32bits(d)
	e.buffer = [9]byte{
		byte(c),
		byte(c >> 8),
		byte(c >> 16),
		byte(c >> 24),
	}
	return e.Writer.Write(e.buffer[:4])
}

// WriteFloat64 Writes a 64 bit float as a float64 using the underlying io.Writer
func (e *LittleEndianWriter) WriteFloat64(d float64) (int, error) {
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
	return e.Writer.Write(e.buffer[:8])
}

// WriteBytesX Writes the length of the Bytes, using ReadUintX and then Writes the bytes
func (e *LittleEndianWriter) WriteBytesX(p []byte) (int, error) {
	n, err := e.WriteUintX(uint64(len(p)))
	if err != nil {
		return n, err
	}
	m, err := e.Write(p)
	return n + m, err
}

// WriteBytes8 Writes the length of the Bytes, using ReadUint8 and then Writes the bytes
func (e *LittleEndianWriter) WriteBytes8(p []byte) (int, error) {
	n, err := e.WriteUint8(uint8(len(p)))
	if err != nil {
		return n, err
	}
	m, err := e.Write(p)
	return n + m, err
}

// WriteBytes16 Writes the length of the Bytes, using ReadUint16 and then Writes the bytes
func (e *LittleEndianWriter) WriteBytes16(p []byte) (int, error) {
	n, err := e.WriteUint16(uint16(len(p)))
	if err != nil {
		return n, err
	}
	m, err := e.Write(p)
	return n + m, err
}

// WriteBytes24 Writes the length of the Bytes, using ReadUint24 and then Writes the bytes
func (e *LittleEndianWriter) WriteBytes24(p []byte) (int, error) {
	n, err := e.WriteUint24(uint32(len(p)))
	if err != nil {
		return n, err
	}
	m, err := e.Write(p)
	return n + m, err
}

// WriteBytes32 Writes the length of the Bytes, using ReadUint32 and then Writes the bytes
func (e *LittleEndianWriter) WriteBytes32(p []byte) (int, error) {
	n, err := e.WriteUint32(uint32(len(p)))
	if err != nil {
		return n, err
	}
	m, err := e.Write(p)
	return n + m, err
}

// WriteBytes40 Writes the length of the Bytes, using ReadUint40 and then Writes the bytes
func (e *LittleEndianWriter) WriteBytes40(p []byte) (int, error) {
	n, err := e.WriteUint40(uint64(len(p)))
	if err != nil {
		return n, err
	}
	m, err := e.Write(p)
	return n + m, err
}

// WriteBytes48 Writes the length of the Bytes, using ReadUint48 and then Writes the bytes
func (e *LittleEndianWriter) WriteBytes48(p []byte) (int, error) {
	n, err := e.WriteUint48(uint64(len(p)))
	if err != nil {
		return n, err
	}
	m, err := e.Write(p)
	return n + m, err
}

// WriteBytes56 Writes the length of the Bytes, using ReadUint56 and then Writes the bytes
func (e *LittleEndianWriter) WriteBytes56(p []byte) (int, error) {
	n, err := e.WriteUint56(uint64(len(p)))
	if err != nil {
		return n, err
	}
	m, err := e.Write(p)
	return n + m, err
}

// WriteBytes64 Writes the length of the Bytes, using ReadUint64 and then Writes the bytes
func (e *LittleEndianWriter) WriteBytes64(p []byte) (int, error) {
	n, err := e.WriteUint64(uint64(len(p)))
	if err != nil {
		return n, err
	}
	m, err := e.Write(p)
	return n + m, err
}

// WriteString Writes a string
func (e *LittleEndianWriter) WriteString(d string) (int, error) {
	return io.WriteString(e.Writer, d)
}

// WriteStringX Writes the length of the String, using ReadUintX and then Writes the bytes
func (e *LittleEndianWriter) WriteStringX(p string) (int, error) {
	n, err := e.WriteUintX(uint64(len(p)))
	if err != nil {
		return n, err
	}
	m, err := e.WriteString(p)
	return n + m, err
}

// WriteString8 Writes the length of the String, using ReadUint8 and then Writes the bytes
func (e *LittleEndianWriter) WriteString8(p string) (int, error) {
	n, err := e.WriteUint8(uint8(len(p)))
	if err != nil {
		return n, err
	}
	m, err := e.WriteString(p)
	return n + m, err
}

// WriteString16 Writes the length of the String, using ReadUint16 and then Writes the bytes
func (e *LittleEndianWriter) WriteString16(p string) (int, error) {
	n, err := e.WriteUint16(uint16(len(p)))
	if err != nil {
		return n, err
	}
	m, err := e.WriteString(p)
	return n + m, err
}

// WriteString24 Writes the length of the String, using ReadUint24 and then Writes the bytes
func (e *LittleEndianWriter) WriteString24(p string) (int, error) {
	n, err := e.WriteUint24(uint32(len(p)))
	if err != nil {
		return n, err
	}
	m, err := e.WriteString(p)
	return n + m, err
}

// WriteString32 Writes the length of the String, using ReadUint32 and then Writes the bytes
func (e *LittleEndianWriter) WriteString32(p string) (int, error) {
	n, err := e.WriteUint32(uint32(len(p)))
	if err != nil {
		return n, err
	}
	m, err := e.WriteString(p)
	return n + m, err
}

// WriteString40 Writes the length of the String, using ReadUint40 and then Writes the bytes
func (e *LittleEndianWriter) WriteString40(p string) (int, error) {
	n, err := e.WriteUint40(uint64(len(p)))
	if err != nil {
		return n, err
	}
	m, err := e.WriteString(p)
	return n + m, err
}

// WriteString48 Writes the length of the String, using ReadUint48 and then Writes the bytes
func (e *LittleEndianWriter) WriteString48(p string) (int, error) {
	n, err := e.WriteUint48(uint64(len(p)))
	if err != nil {
		return n, err
	}
	m, err := e.WriteString(p)
	return n + m, err
}

// WriteString56 Writes the length of the String, using ReadUint56 and then Writes the bytes
func (e *LittleEndianWriter) WriteString56(p string) (int, error) {
	n, err := e.WriteUint56(uint64(len(p)))
	if err != nil {
		return n, err
	}
	m, err := e.WriteString(p)
	return n + m, err
}

// WriteString64 Writes the length of the String, using ReadUint64 and then Writes the bytes
func (e *LittleEndianWriter) WriteString64(p string) (int, error) {
	n, err := e.WriteUint64(uint64(len(p)))
	if err != nil {
		return n, err
	}
	m, err := e.WriteString(p)
	return n + m, err
}

// WriteString0 Writes the bytes of the string ending with a 0 byte
func (e *LittleEndianWriter) WriteString0(p string) (int, error) {
	n, err := e.WriteString(p)
	if err == nil {
		var m int
		m, err = e.WriteUint8(0)
		n += m
	}
	return n, err
}
