package byteio

import (
	"io"
	"math"
)

type BigEndianReader struct {
	io.Reader
}

func (b BigEndianReader) ReadUint8() (uint8, int, error) {
	var bytes [1]byte
	n, err := io.ReadFull(b, bytes[:])
	if err != nil {
		return 0, n, err
	}
	return uint8(bytes[0]), n, nil
}

func (b BigEndianReader) ReadInt8() (int8, int, error) {
	u, n, err := b.ReadUint8()
	return int8(u), n, err
}

func (b BigEndianReader) ReadUint16() (uint16, int, error) {
	var bytes [2]byte
	n, err := io.ReadFull(b, bytes[:])
	if err != nil {
		return 0, n, err
	}
	return uint16(bytes[1]) | uint16(bytes[0])<<8, n, nil
}

func (b BigEndianReader) ReadInt16() (int16, int, error) {
	u, n, err := b.ReadUint16()
	return int16(u), n, err
}

func (b BigEndianReader) ReadUint32() (uint32, int, error) {
	var bytes [4]byte
	n, err := io.ReadFull(b, bytes[:])
	if err != nil {
		return 0, n, err
	}
	return uint32(bytes[3]) | uint32(bytes[2])<<8 | uint32(bytes[1])<<16 | uint32(bytes[0])<<24, n, nil
}

func (b BigEndianReader) ReadInt32() (int32, int, error) {
	u, n, err := b.ReadUint32()
	return int32(u), n, err
}

func (b BigEndianReader) ReadUint64() (uint64, int, error) {
	var bytes [8]byte
	n, err := io.ReadFull(b, bytes[:])
	if err != nil {
		return 0, n, err
	}
	return uint64(bytes[7]) | uint64(bytes[6])<<8 | uint64(bytes[5])<<16 | uint64(bytes[4])<<24 | uint64(bytes[3])<<32 | uint64(bytes[2])<<40 | uint64(bytes[1])<<48 | uint64(bytes[0])<<56, n, nil
}

func (b BigEndianReader) ReadInt64() (int64, int, error) {
	u, n, err := b.ReadUint64()
	return int64(u), n, err
}

func (b BigEndianReader) ReadFloat32() (float32, int, error) {
	u, n, err := b.ReadUint32()
	return math.Float32frombits(u), n, err
}

func (b BigEndianReader) ReadFloat64() (float64, int, error) {
	u, n, err := b.ReadUint64()
	return math.Float64frombits(u), n, err
}
