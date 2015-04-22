package byteio

import (
	"io"
	"math"
	"sync"
)

type LittleEndianWriter struct {
	io.Writer

	mu    sync.Mutex
	bytes [8]byte
}

func (l *LittleEndianWriter) WriteUint8(d uint8) (int, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.bytes[0] = byte(d)
	return l.Write(l.bytes[:1])
}

func (l *LittleEndianWriter) WriteInt8(d int8) (int, error) {
	return l.WriteUint8(uint8(d))
}

func (l *LittleEndianWriter) WriteUint16(d uint16) (int, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.bytes[0] = byte(d)
	l.bytes[1] = byte(d >> 8)
	return l.Write(l.bytes[:2])
}

func (l *LittleEndianWriter) WriteInt16(d int16) (int, error) {
	return l.WriteUint16(uint16(d))
}

func (l *LittleEndianWriter) WriteUint32(d uint32) (int, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.bytes[0] = byte(d)
	l.bytes[1] = byte(d >> 8)
	l.bytes[2] = byte(d >> 16)
	l.bytes[3] = byte(d >> 24)
	return l.Write(l.bytes[:4])
}

func (l *LittleEndianWriter) WriteInt32(d int32) (int, error) {
	return l.WriteUint32(uint32(d))
}

func (l *LittleEndianWriter) WriteUint64(d uint64) (int, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.bytes[0] = byte(d)
	l.bytes[1] = byte(d >> 8)
	l.bytes[2] = byte(d >> 16)
	l.bytes[3] = byte(d >> 24)
	l.bytes[4] = byte(d >> 32)
	l.bytes[5] = byte(d >> 40)
	l.bytes[6] = byte(d >> 48)
	l.bytes[7] = byte(d >> 56)
	return l.Write(l.bytes[:])
}

func (l *LittleEndianWriter) WriteInt64(d int16) (int, error) {
	return l.WriteUint64(uint64(d))
}

func (l *LittleEndianWriter) WriteFloat32(d float32) (int, error) {
	return l.WriteUint32(math.Float32bits(d))
}

func (l *LittleEndianWriter) WriteFloat64(d float64) (int, error) {
	return l.WriteUint64(math.Float64bits(d))
}
