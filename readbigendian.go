package byteio

import (
	"io"
	"math"
)

// BigEndianReader is a wrapper around a io.Reader that can be used to read
// types in the big endian format
type BigEndianReader struct {
	io.Reader
	buffer [8]byte
}

// ReadUint8 will read a single byte from the reader and return it as a uint8
func (b *BigEndianReader) ReadUint8() (uint8, int, error) {
	n, err := io.ReadFull(b.Reader, b.buffer[:1])
	if err != nil {
		return 0, n, err
	}
	return uint8(b.buffer[0]), 1, nil
}

// ReadInt8 will read a single byte from the reader and return it as an int8
func (b *BigEndianReader) ReadInt8() (int8, int, error) {
	u, n, err := b.ReadUint8()
	return int8(u), n, err
}

// ReadUint16 will read two b.buffer from the reader, in big endian format, and
// return it as a uint16
func (b *BigEndianReader) ReadUint16() (uint16, int, error) {
	n, err := io.ReadFull(b.Reader, b.buffer[:2])
	if err != nil {
		return 0, n, err
	}
	return uint16(b.buffer[1]) | uint16(b.buffer[0])<<8, 2, nil
}

// ReadInt16 will read two b.buffer from the reader, in big endian format, and
// return it as an int16
func (b *BigEndianReader) ReadInt16() (int16, int, error) {
	u, n, err := b.ReadUint16()
	return int16(u), n, err
}

// ReadUint32 will read four b.buffer from the reader, in big endian format, and
// return it as a uint32
func (b *BigEndianReader) ReadUint32() (uint32, int, error) {
	n, err := io.ReadFull(b.Reader, b.buffer[:4])
	if err != nil {
		return 0, n, err
	}
	return uint32(b.buffer[3]) | uint32(b.buffer[2])<<8 | uint32(b.buffer[1])<<16 | uint32(b.buffer[0])<<24, 4, nil
}

// ReadInt32 will read four b.buffer from the reader, in big endian format, and
// return it as an int32
func (b *BigEndianReader) ReadInt32() (int32, int, error) {
	u, n, err := b.ReadUint32()
	return int32(u), n, err
}

// ReadUint64 will read eight b.buffer from the reader, in big endian format, and
// return it as a uint64
func (b *BigEndianReader) ReadUint64() (uint64, int, error) {
	n, err := io.ReadFull(b.Reader, b.buffer[:8])
	if err != nil {
		return 0, n, err
	}
	return uint64(b.buffer[7]) | uint64(b.buffer[6])<<8 | uint64(b.buffer[5])<<16 | uint64(b.buffer[4])<<24 | uint64(b.buffer[3])<<32 | uint64(b.buffer[2])<<40 | uint64(b.buffer[1])<<48 | uint64(b.buffer[0])<<56, 8, nil
}

// ReadInt64 will read eight b.buffer from the reader, in big endian format, and
// return it as an int64
func (b *BigEndianReader) ReadInt64() (int64, int, error) {
	u, n, err := b.ReadUint64()
	return int64(u), n, err
}

// ReadFloat32 will read four b.buffer from the reader, in big endian format, and
// return it as a float32
func (b *BigEndianReader) ReadFloat32() (float32, int, error) {
	u, n, err := b.ReadUint32()
	return math.Float32frombits(u), n, err
}

// ReadFloat64 will read eight b.buffer from the reader, in big endian format, and
// return it as a float64
func (b *BigEndianReader) ReadFloat64() (float64, int, error) {
	u, n, err := b.ReadUint64()
	return math.Float64frombits(u), n, err
}
