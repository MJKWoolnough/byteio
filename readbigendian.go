package byteio

import (
	"io"
	"math"
)

// BigEndianReader is a wrapper around a io.Reader that can be used to read
// types in the big endian format
type BigEndianReader struct {
	io.Reader
}

// ReadUint8 will read a single byte from the reader and return it as a uint8
func (b BigEndianReader) ReadUint8() (uint8, int, error) {
	bytes := pool.Get().(*[8]byte)
	n, err := io.ReadFull(b.Reader, bytes[:1])
	if err != nil {
		return 0, n, err
	}
	v := uint8(bytes[0])
	pool.Put(bytes)
	return v, n, nil
}

// ReadInt8 will read a single byte from the reader and return it as an int8
func (b BigEndianReader) ReadInt8() (int8, int, error) {
	u, n, err := b.ReadUint8()
	return int8(u), n, err
}

// ReadUint16 will read two bytes from the reader, in big endian format, and
// return it as a uint16
func (b BigEndianReader) ReadUint16() (uint16, int, error) {
	bytes := pool.Get().(*[8]byte)
	n, err := io.ReadFull(b.Reader, bytes[:2])
	if err != nil {
		return 0, n, err
	}
	v := uint16(bytes[1]) | uint16(bytes[0])<<8
	pool.Put(bytes)
	return v, n, nil
}

// ReadInt16 will read two bytes from the reader, in big endian format, and
// return it as an int16
func (b BigEndianReader) ReadInt16() (int16, int, error) {
	u, n, err := b.ReadUint16()
	return int16(u), n, err
}

// ReadUint32 will read four bytes from the reader, in big endian format, and
// return it as a uint32
func (b BigEndianReader) ReadUint32() (uint32, int, error) {
	bytes := pool.Get().(*[8]byte)
	n, err := io.ReadFull(b.Reader, bytes[:4])
	if err != nil {
		return 0, n, err
	}
	v := uint32(bytes[3]) | uint32(bytes[2])<<8 | uint32(bytes[1])<<16 | uint32(bytes[0])<<24
	pool.Put(bytes)
	return v, n, nil
}

// ReadInt32 will read four bytes from the reader, in big endian format, and
// return it as an int32
func (b BigEndianReader) ReadInt32() (int32, int, error) {
	u, n, err := b.ReadUint32()
	return int32(u), n, err
}

// ReadUint64 will read eight bytes from the reader, in big endian format, and
// return it as a uint64
func (b BigEndianReader) ReadUint64() (uint64, int, error) {
	bytes := pool.Get().(*[8]byte)
	n, err := io.ReadFull(b.Reader, bytes[:8])
	if err != nil {
		return 0, n, err
	}
	v := uint64(bytes[7]) | uint64(bytes[6])<<8 | uint64(bytes[5])<<16 | uint64(bytes[4])<<24 | uint64(bytes[3])<<32 | uint64(bytes[2])<<40 | uint64(bytes[1])<<48 | uint64(bytes[0])<<56
	pool.Put(bytes)
	return v, n, nil
}

// ReadInt64 will read eight bytes from the reader, in big endian format, and
// return it as an int64
func (b BigEndianReader) ReadInt64() (int64, int, error) {
	u, n, err := b.ReadUint64()
	return int64(u), n, err
}

// ReadFloat32 will read four bytes from the reader, in big endian format, and
// return it as a float32
func (b BigEndianReader) ReadFloat32() (float32, int, error) {
	u, n, err := b.ReadUint32()
	return math.Float32frombits(u), n, err
}

// ReadFloat64 will read eight bytes from the reader, in big endian format, and
// return it as a float64
func (b BigEndianReader) ReadFloat64() (float64, int, error) {
	u, n, err := b.ReadUint64()
	return math.Float64frombits(u), n, err
}
