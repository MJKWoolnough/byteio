package byteio

import (
	"io"
	"math"
)

// LittleEndianReader is a wrapper around a io.Reader that can be used to read
// types in the little endian format
type LittleEndianReader struct {
	io.Reader
	buffer [8]byte
}

// ReadUint8 will read a single byte from the reader, in little endian format,
// and return it as a uint8
func (l *LittleEndianReader) ReadUint8() (uint8, int, error) {
	n, err := io.ReadFull(l.Reader, l.buffer[:1])
	if err != nil {
		return 0, n, err
	}
	v := uint8(l.buffer[0])
	return v, 1, nil
}

// ReadInt8 will read a single byte from the reader, in little endian format,
// and return it as an int8
func (l *LittleEndianReader) ReadInt8() (int8, int, error) {
	u, n, err := l.ReadUint8()
	return int8(u), n, err
}

// ReadUint16 will read two bytes from the reader, in little endian format, and
// return it as a uint16
func (l *LittleEndianReader) ReadUint16() (uint16, int, error) {
	n, err := io.ReadFull(l.Reader, l.buffer[:2])
	if err != nil {
		return 0, n, err
	}
	return uint16(l.buffer[0]) | uint16(l.buffer[1])<<8, 2, nil
}

// ReadInt16 will read two bytes from the reader, in little endian format, and
// return it as an int16
func (l *LittleEndianReader) ReadInt16() (int16, int, error) {
	u, n, err := l.ReadUint16()
	return int16(u), n, err
}

// ReadUint32 will read four bytes from the reader, in little endian format, and
// return it as a uint32
func (l *LittleEndianReader) ReadUint32() (uint32, int, error) {
	n, err := io.ReadFull(l.Reader, l.buffer[:4])
	if err != nil {
		return 0, n, err
	}
	return uint32(l.buffer[0]) | uint32(l.buffer[1])<<8 | uint32(l.buffer[2])<<16 | uint32(l.buffer[3])<<24, 4, nil
}

// ReadInt32 will read four bytes from the reader, in little endian format, and
// return it as an int32
func (l *LittleEndianReader) ReadInt32() (int32, int, error) {
	u, n, err := l.ReadUint32()
	return int32(u), n, err
}

// ReadUint64 will read eight bytes from the reader, in little endian format,
// and return it as a uint64
func (l *LittleEndianReader) ReadUint64() (uint64, int, error) {
	n, err := io.ReadFull(l.Reader, l.buffer[:8])
	if err != nil {
		return 0, n, err
	}
	return uint64(l.buffer[0]) | uint64(l.buffer[1])<<8 | uint64(l.buffer[2])<<16 | uint64(l.buffer[3])<<24 | uint64(l.buffer[4])<<32 | uint64(l.buffer[5])<<40 | uint64(l.buffer[6])<<48 | uint64(l.buffer[7])<<56, 8, nil
}

// ReadInt64 will read eight bytes from the reader, in little endian format, and
// return it as an int64
func (l *LittleEndianReader) ReadInt64() (int64, int, error) {
	u, n, err := l.ReadUint64()
	return int64(u), n, err
}

// ReadFloat32 will read four bytes from the reader, in little endian format,
// and return it as a float32
func (l *LittleEndianReader) ReadFloat32() (float32, int, error) {
	u, n, err := l.ReadUint32()
	return math.Float32frombits(u), n, err
}

// ReadFloat64 will read eight bytes from the reader, in little endian format,
// and return it as a float64
func (l *LittleEndianReader) ReadFloat64() (float64, int, error) {
	u, n, err := l.ReadUint64()
	return math.Float64frombits(u), n, err
}
