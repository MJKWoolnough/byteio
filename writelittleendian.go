package byteio

import (
	"io"
	"math"
)

// LittleEndianWriter is a wrapper around a io.Writer that can be used to write
// types in the little endian format
type LittleEndianWriter struct {
	io.Writer
}

// WriteUint8 will write the given uint8 to the writer
func (l LittleEndianWriter) WriteUint8(d uint8) (int, error) {
	bytes := pool.Get().(*[8]byte)
	bytes[0] = byte(d)
	n, err := l.Writer.Write(bytes[:1])
	pool.Put(bytes)
	return n, err
}

// WriteInt8 will write the given int8 to the writer
func (l LittleEndianWriter) WriteInt8(d int8) (int, error) {
	return l.WriteUint8(uint8(d))
}

// WriteUint16 will write the given uint16 to the writer, in little endian
// format
func (l LittleEndianWriter) WriteUint16(d uint16) (int, error) {
	bytes := pool.Get().(*[8]byte)
	bytes[0] = byte(d)
	bytes[1] = byte(d >> 8)
	n, err := l.Writer.Write(bytes[:2])
	pool.Put(bytes)
	return n, err
}

// WriteInt16 will write the given int16 to the writer, in little endian format
func (l LittleEndianWriter) WriteInt16(d int16) (int, error) {
	return l.WriteUint16(uint16(d))
}

// WriteUint32 will write the given uint32 to the writer, in little endian
// format
func (l LittleEndianWriter) WriteUint32(d uint32) (int, error) {
	bytes := pool.Get().(*[8]byte)
	bytes[0] = byte(d)
	bytes[1] = byte(d >> 8)
	bytes[2] = byte(d >> 16)
	bytes[3] = byte(d >> 24)
	n, err := l.Writer.Write(bytes[:4])
	pool.Put(bytes)
	return n, err
}

// WriteInt32 will write the given int32 to the writer, in little endian format
func (l LittleEndianWriter) WriteInt32(d int32) (int, error) {
	return l.WriteUint32(uint32(d))
}

// WriteUint64 will write the given uint64 to the writer, in little endian
// format
func (l LittleEndianWriter) WriteUint64(d uint64) (int, error) {
	bytes := pool.Get().(*[8]byte)
	bytes[0] = byte(d)
	bytes[1] = byte(d >> 8)
	bytes[2] = byte(d >> 16)
	bytes[3] = byte(d >> 24)
	bytes[4] = byte(d >> 32)
	bytes[5] = byte(d >> 40)
	bytes[6] = byte(d >> 48)
	bytes[7] = byte(d >> 56)
	n, err := l.Writer.Write(bytes[:8])
	pool.Put(bytes)
	return n, err
}

// WriteInt64 will write the given int64 to the writer, in little endian format
func (l LittleEndianWriter) WriteInt64(d int64) (int, error) {
	return l.WriteUint64(uint64(d))
}

// WriteFloat32 will write the given float32 to the writer, in little endian
// format
func (l LittleEndianWriter) WriteFloat32(d float32) (int, error) {
	return l.WriteUint32(math.Float32bits(d))
}

// WriteFloat64 will write the given float64 to the writer, in little endian
// format
func (l LittleEndianWriter) WriteFloat64(d float64) (int, error) {
	return l.WriteUint64(math.Float64bits(d))
}
