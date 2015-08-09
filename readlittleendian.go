package byteio

import (
	"io"
	"math"
)

// LittleEndianReader is a wrapper around a io.Reader that can be used to read
// types in the little endian format
type LittleEndianReader struct {
	io.Reader
}

// ReadUint8 will read a single byte from the reader, in little endian format,
// and return it as a uint8
func (l LittleEndianReader) ReadUint8() (uint8, int, error) {
	var bytes [1]byte
	n, err := io.ReadFull(l, bytes[:])
	if err != nil {
		return 0, n, err
	}
	return uint8(bytes[0]), n, nil
}

// ReadInt8 will read a single byte from the reader, in little endian format,
// and return it as an int8
func (l LittleEndianReader) ReadInt8() (int8, int, error) {
	u, n, err := l.ReadUint8()
	return int8(u), n, err
}

// ReadUint16 will read two bytes from the reader, in little endian format, and
// return it as a uint16
func (l LittleEndianReader) ReadUint16() (uint16, int, error) {
	var bytes [2]byte
	n, err := io.ReadFull(l, bytes[:])
	if err != nil {
		return 0, n, err
	}
	return uint16(bytes[0]) | uint16(bytes[1])<<8, n, nil
}

// ReadInt16 will read two bytes from the reader, in little endian format, and
// return it as an int16
func (l LittleEndianReader) ReadInt16() (int16, int, error) {
	u, n, err := l.ReadUint16()
	return int16(u), n, err
}

// ReadUint32 will read four bytes from the reader, in little endian format, and
// return it as a uint32
func (l LittleEndianReader) ReadUint32() (uint32, int, error) {
	var bytes [4]byte
	n, err := io.ReadFull(l, bytes[:])
	if err != nil {
		return 0, n, err
	}
	return uint32(bytes[0]) | uint32(bytes[1])<<8 | uint32(bytes[2])<<16 | uint32(bytes[3])<<24, n, nil
}

// ReadInt32 will read four bytes from the reader, in little endian format, and
// return it as an int32
func (l LittleEndianReader) ReadInt32() (int32, int, error) {
	u, n, err := l.ReadUint32()
	return int32(u), n, err
}

// ReadUint64 will read eight bytes from the reader, in little endian format,
// and return it as a uint64
func (l LittleEndianReader) ReadUint64() (uint64, int, error) {
	var bytes [8]byte
	n, err := io.ReadFull(l, bytes[:])
	if err != nil {
		return 0, n, err
	}
	return uint64(bytes[0]) | uint64(bytes[1])<<8 | uint64(bytes[2])<<16 | uint64(bytes[3])<<24 | uint64(bytes[4])<<32 | uint64(bytes[5])<<40 | uint64(bytes[6])<<48 | uint64(bytes[7])<<56, n, nil
}

// ReadInt64 will read eight bytes from the reader, in little endian format, and
// return it as an int64
func (l LittleEndianReader) ReadInt64() (int64, int, error) {
	u, n, err := l.ReadUint64()
	return int64(u), n, err
}

// ReadFloat32 will read four bytes from the reader, in little endian format,
// and return it as a float32
func (l LittleEndianReader) ReadFloat32() (float32, int, error) {
	u, n, err := l.ReadUint32()
	return math.Float32frombits(u), n, err
}

// ReadFloat64 will read eight bytes from the reader, in little endian format,
// and return it as a float64
func (l LittleEndianReader) ReadFloat64() (float64, int, error) {
	u, n, err := l.ReadUint64()
	return math.Float64frombits(u), n, err
}
