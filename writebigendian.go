package byteio

import (
	"io"
	"math"
)

type BigEndianWriter struct {
	io.Writer
}

func (b *BigEndianWriter) WriteUint8(d uint8) (int, error) {
	return b.Write([]byte{
		byte(d),
	})
}

func (b *BigEndianWriter) WriteInt8(d int8) (int, error) {
	return b.WriteUint8(uint8(d))
}

func (b *BigEndianWriter) WriteUint16(d uint16) (int, error) {
	return b.Write([]byte{
		byte(d >> 8),
		byte(d),
	})
}

func (b *BigEndianWriter) WriteInt16(d int16) (int, error) {
	return b.WriteUint16(uint16(d))
}

func (b *BigEndianWriter) WriteUint32(d uint32) (int, error) {
	return b.Write([]byte{
		byte(d >> 24),
		byte(d >> 16),
		byte(d >> 8),
		byte(d),
	})
}

func (b *BigEndianWriter) WriteInt32(d int32) (int, error) {
	return b.WriteUint32(uint32(d))
}

func (b *BigEndianWriter) WriteUint64(d uint64) (int, error) {
	return b.Write([]byte{
		byte(d >> 56),
		byte(d >> 48),
		byte(d >> 40),
		byte(d >> 32),
		byte(d >> 24),
		byte(d >> 16),
		byte(d >> 8),
		byte(d),
	})
}

func (b *BigEndianWriter) WriteInt64(d int64) (int, error) {
	return b.WriteUint64(uint64(d))
}

func (b *BigEndianWriter) WriteFloat32(d float32) (int, error) {
	return b.WriteUint32(math.Float32bits(d))
}

func (b *BigEndianWriter) WriteFloat64(d float64) (int, error) {
	return b.WriteUint64(math.Float64bits(d))
}
