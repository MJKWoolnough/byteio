package byteio

// File automatically generated with ./gen.sh

import (
	"io"
	"math"
)

type BigEndianReader struct {
	io.Reader
	buffer [8]byte
}

func (e *BigEndianReader) ReadInt8() (int8, int, error) {
	n, err := io.ReadFull(e.Reader, e.buffer[:1])
	if err != nil {
		return 0, n, err
	}
	return int8(e.buffer[0]), 1, nil
}

func (e *BigEndianReader) ReadInt16() (int16, int, error) {
	n, err := io.ReadFull(e.Reader, e.buffer[:2])
	if err != nil {
		return 0, n, err
	}
	return int16(uint16(e.buffer[1]) | uint16(e.buffer[0])<<8), 2, nil
}

func (e *BigEndianReader) ReadInt32() (int32, int, error) {
	n, err := io.ReadFull(e.Reader, e.buffer[:4])
	if err != nil {
		return 0, n, err
	}
	return int32(uint32(e.buffer[3]) | uint32(e.buffer[2])<<8 | uint32(e.buffer[1])<<16 | uint32(e.buffer[0])<<24), 4, nil
}

func (e *BigEndianReader) ReadInt64() (int64, int, error) {
	n, err := io.ReadFull(e.Reader, e.buffer[:8])
	if err != nil {
		return 0, n, err
	}
	return int64(uint64(e.buffer[7]) | uint64(e.buffer[6])<<8 | uint64(e.buffer[5])<<16 | uint64(e.buffer[4])<<24 | uint64(e.buffer[3])<<32 | uint64(e.buffer[2])<<40 | uint64(e.buffer[1])<<48 | uint64(e.buffer[0])<<56), 8, nil
}

func (e *BigEndianReader) ReadUint8() (uint8, int, error) {
	n, err := io.ReadFull(e.Reader, e.buffer[:1])
	if err != nil {
		return 0, n, err
	}
	return e.buffer[0], 1, nil
}

func (e *BigEndianReader) ReadUint16() (uint16, int, error) {
	n, err := io.ReadFull(e.Reader, e.buffer[:2])
	if err != nil {
		return 0, n, err
	}
	return uint16(e.buffer[1]) | uint16(e.buffer[0])<<8, 2, nil
}

func (e *BigEndianReader) ReadUint32() (uint32, int, error) {
	n, err := io.ReadFull(e.Reader, e.buffer[:4])
	if err != nil {
		return 0, n, err
	}
	return uint32(e.buffer[3]) | uint32(e.buffer[2])<<8 | uint32(e.buffer[1])<<16 | uint32(e.buffer[0])<<24, 4, nil
}

func (e *BigEndianReader) ReadUint64() (uint64, int, error) {
	n, err := io.ReadFull(e.Reader, e.buffer[:8])
	if err != nil {
		return 0, n, err
	}
	return uint64(e.buffer[7]) | uint64(e.buffer[6])<<8 | uint64(e.buffer[5])<<16 | uint64(e.buffer[4])<<24 | uint64(e.buffer[3])<<32 | uint64(e.buffer[2])<<40 | uint64(e.buffer[1])<<48 | uint64(e.buffer[0])<<56, 8, nil
}

func (e *BigEndianReader) ReadFloat32() (float32, int, error) {
	n, err := io.ReadFull(e.Reader, e.buffer[:4])
	if err != nil {
		return 0, n, err
	}
	return math.Float32frombits(uint32(e.buffer[3]) | uint32(e.buffer[2])<<8 | uint32(e.buffer[1])<<16 | uint32(e.buffer[0])<<24), 4, nil
}

func (e *BigEndianReader) ReadFloat64() (float64, int, error) {
	n, err := io.ReadFull(e.Reader, e.buffer[:8])
	if err != nil {
		return 0, n, err
	}
	return math.Float64frombits(uint64(e.buffer[7]) | uint64(e.buffer[6])<<8 | uint64(e.buffer[5])<<16 | uint64(e.buffer[4])<<24 | uint64(e.buffer[3])<<32 | uint64(e.buffer[2])<<40 | uint64(e.buffer[1])<<48 | uint64(e.buffer[0])<<56), 8, nil
}
