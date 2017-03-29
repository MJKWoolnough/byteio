package byteio

import (
	"io"
	"math"
)

// BigEndianWriter is a wrapper around a io.Writer that can be used to write
// types in the big endian format
type BigEndianWriter struct {
	io.Writer
	buffer [8]byte
}

// WriteUint8 will write the given uint8 to the writer
func (b *BigEndianWriter) WriteUint8(d uint8) (int, error) {
	b.buffer[0] = byte(d)
	return b.Writer.Write(b.buffer[:1])
}

// WriteInt8 will write the given int8 to the writer
func (b *BigEndianWriter) WriteInt8(d int8) (int, error) {
	return b.WriteUint8(uint8(d))
}

// WriteUint16 will write the given uint16 to the writer, in big endian format
func (b *BigEndianWriter) WriteUint16(d uint16) (int, error) {
	b.buffer[0] = byte(d >> 8)
	b.buffer[1] = byte(d)
	return b.Writer.Write(b.buffer[:2])
}

// WriteInt16 will write the given int16 to the writer, in big endian format
func (b *BigEndianWriter) WriteInt16(d int16) (int, error) {
	return b.WriteUint16(uint16(d))
}

// WriteUint32 will write the given uint32 to the writer, in big endian format
func (b *BigEndianWriter) WriteUint32(d uint32) (int, error) {
	b.buffer = [8]byte{
		byte(d >> 24),
		byte(d >> 16),
		byte(d >> 8),
		byte(d),
	}
	return b.Writer.Write(b.buffer[:4])
}

// WriteInt32 will write the given int32 to the writer, in big endian format
func (b *BigEndianWriter) WriteInt32(d int32) (int, error) {
	return b.WriteUint32(uint32(d))
}

// WriteUint64 will write the given uint64 to the writer, in big endian format
func (b *BigEndianWriter) WriteUint64(d uint64) (int, error) {
	b.buffer = [8]byte{
		byte(d >> 56),
		byte(d >> 48),
		byte(d >> 40),
		byte(d >> 32),
		byte(d >> 24),
		byte(d >> 16),
		byte(d >> 8),
		byte(d),
	}
	return b.Writer.Write(b.buffer[:8])
}

// WriteInt64 will write the given int64 to the writer, in big endian format
func (b *BigEndianWriter) WriteInt64(d int64) (int, error) {
	return b.WriteUint64(uint64(d))
}

// WriteFloat32 will write the given float32 to the writer, in big endian format
func (b *BigEndianWriter) WriteFloat32(d float32) (int, error) {
	return b.WriteUint32(math.Float32bits(d))
}

// WriteFloat64 will write the given float64 to the writer, in big endian format
func (b *BigEndianWriter) WriteFloat64(d float64) (int, error) {
	return b.WriteUint64(math.Float64bits(d))
}
