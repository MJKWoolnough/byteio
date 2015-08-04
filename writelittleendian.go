package byteio

import (
	"io"
	"math"
)

type LittleEndianWriter struct {
	io.Writer
}

func (l *LittleEndianWriter) WriteUint8(d uint8) (int, error) {
	return l.Write([]byte{
		byte(d),
	})
}

func (l *LittleEndianWriter) WriteInt8(d int8) (int, error) {
	return l.WriteUint8(uint8(d))
}

func (l *LittleEndianWriter) WriteUint16(d uint16) (int, error) {
	return l.Write([]byte{
		byte(d),
		byte(d >> 8),
	})
}

func (l *LittleEndianWriter) WriteInt16(d int16) (int, error) {
	return l.WriteUint16(uint16(d))
}

func (l *LittleEndianWriter) WriteUint32(d uint32) (int, error) {
	return l.Write([]byte{
		byte(d),
		byte(d >> 8),
		byte(d >> 16),
		byte(d >> 24),
	})
}

func (l *LittleEndianWriter) WriteInt32(d int32) (int, error) {
	return l.WriteUint32(uint32(d))
}

func (l *LittleEndianWriter) WriteUint64(d uint64) (int, error) {
	return l.Write([]byte{
		byte(d),
		byte(d >> 8),
		byte(d >> 16),
		byte(d >> 24),
		byte(d >> 32),
		byte(d >> 40),
		byte(d >> 48),
		byte(d >> 56),
	})
}

func (l *LittleEndianWriter) WriteInt64(d int64) (int, error) {
	return l.WriteUint64(uint64(d))
}

func (l *LittleEndianWriter) WriteFloat32(d float32) (int, error) {
	return l.WriteUint32(math.Float32bits(d))
}

func (l *LittleEndianWriter) WriteFloat64(d float64) (int, error) {
	return l.WriteUint64(math.Float64bits(d))
}
