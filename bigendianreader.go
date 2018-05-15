package byteio

// File automatically generated with ./gen.sh

import (
	"io"
	"math"
)

// BigEndianReader wraps a io.Reader to provide methods
// to make it easier to Read fundamental types
type BigEndianReader struct {
	io.Reader
	buffer [9]byte
}

// ReadInt8 Reads a int8 using the underlying io.Reader
func (e *BigEndianReader) ReadInt8() (int8, int, error) {
	n, err := io.ReadFull(e.Reader, e.buffer[:1])
	if err != nil {
		return 0, n, err
	}
	return int8(e.buffer[0]), 1, nil
}

// ReadInt16 Reads a int16 using the underlying io.Reader
func (e *BigEndianReader) ReadInt16() (int16, int, error) {
	n, err := io.ReadFull(e.Reader, e.buffer[:2])
	if err != nil {
		return 0, n, err
	}
	return int16(uint16(e.buffer[1]) | uint16(e.buffer[0])<<8), 2, nil
}

// ReadInt32 Reads a int32 using the underlying io.Reader
func (e *BigEndianReader) ReadInt32() (int32, int, error) {
	n, err := io.ReadFull(e.Reader, e.buffer[:4])
	if err != nil {
		return 0, n, err
	}
	return int32(uint32(e.buffer[3]) | uint32(e.buffer[2])<<8 | uint32(e.buffer[1])<<16 | uint32(e.buffer[0])<<24), 4, nil
}

// ReadInt64 Reads a int64 using the underlying io.Reader
func (e *BigEndianReader) ReadInt64() (int64, int, error) {
	n, err := io.ReadFull(e.Reader, e.buffer[:8])
	if err != nil {
		return 0, n, err
	}
	return int64(uint64(e.buffer[7]) | uint64(e.buffer[6])<<8 | uint64(e.buffer[5])<<16 | uint64(e.buffer[4])<<24 | uint64(e.buffer[3])<<32 | uint64(e.buffer[2])<<40 | uint64(e.buffer[1])<<48 | uint64(e.buffer[0])<<56), 8, nil
}

// ReadUint8 Reads a uint8 using the underlying io.Reader
func (e *BigEndianReader) ReadUint8() (uint8, int, error) {
	n, err := io.ReadFull(e.Reader, e.buffer[:1])
	if err != nil {
		return 0, n, err
	}
	return e.buffer[0], 1, nil
}

// ReadUint16 Reads a uint16 using the underlying io.Reader
func (e *BigEndianReader) ReadUint16() (uint16, int, error) {
	n, err := io.ReadFull(e.Reader, e.buffer[:2])
	if err != nil {
		return 0, n, err
	}
	return uint16(e.buffer[1]) | uint16(e.buffer[0])<<8, 2, nil
}

// ReadUint32 Reads a uint32 using the underlying io.Reader
func (e *BigEndianReader) ReadUint32() (uint32, int, error) {
	n, err := io.ReadFull(e.Reader, e.buffer[:4])
	if err != nil {
		return 0, n, err
	}
	return uint32(e.buffer[3]) | uint32(e.buffer[2])<<8 | uint32(e.buffer[1])<<16 | uint32(e.buffer[0])<<24, 4, nil
}

// ReadUint64 Reads a uint64 using the underlying io.Reader
func (e *BigEndianReader) ReadUint64() (uint64, int, error) {
	n, err := io.ReadFull(e.Reader, e.buffer[:8])
	if err != nil {
		return 0, n, err
	}
	return uint64(e.buffer[7]) | uint64(e.buffer[6])<<8 | uint64(e.buffer[5])<<16 | uint64(e.buffer[4])<<24 | uint64(e.buffer[3])<<32 | uint64(e.buffer[2])<<40 | uint64(e.buffer[1])<<48 | uint64(e.buffer[0])<<56, 8, nil
}

// ReadFloat32 Reads a float32 using the underlying io.Reader
func (e *BigEndianReader) ReadFloat32() (float32, int, error) {
	n, err := io.ReadFull(e.Reader, e.buffer[:4])
	if err != nil {
		return 0, n, err
	}
	return math.Float32frombits(uint32(e.buffer[3]) | uint32(e.buffer[2])<<8 | uint32(e.buffer[1])<<16 | uint32(e.buffer[0])<<24), 4, nil
}

// ReadFloat64 Reads a float64 using the underlying io.Reader
func (e *BigEndianReader) ReadFloat64() (float64, int, error) {
	n, err := io.ReadFull(e.Reader, e.buffer[:8])
	if err != nil {
		return 0, n, err
	}
	return math.Float64frombits(uint64(e.buffer[7]) | uint64(e.buffer[6])<<8 | uint64(e.buffer[5])<<16 | uint64(e.buffer[4])<<24 | uint64(e.buffer[3])<<32 | uint64(e.buffer[2])<<40 | uint64(e.buffer[1])<<48 | uint64(e.buffer[0])<<56), 8, nil
}

// ReadString Reads a string
func (e *BigEndianReader) ReadString(size int) (string, int, error) {
	buf := make([]byte, size)
	n, err := io.ReadFull(e, buf)
	return string(buf[:n]), n, err
}

// ReadStringX Reads the length of the string, using ReadUintX and then Reads the bytes of the string
func (e *BigEndianReader) ReadStringX() (string, int, error) {
	size, n, err := e.ReadUintX()
	if err != nil {
		return "", n, err
	}
	str, m, err := e.ReadString(int(size))
	return str, n + m, err
}

// ReadString8 Reads the length of the string, using ReadUint8 and then Reads the bytes of the string
func (e *BigEndianReader) ReadString8() (string, int, error) {
	size, n, err := e.ReadUint8()
	if err != nil {
		return "", n, err
	}
	str, m, err := e.ReadString(int(size))
	return str, n + m, err
}

// ReadString16 Reads the length of the string, using ReadUint16 and then Reads the bytes of the string
func (e *BigEndianReader) ReadString16() (string, int, error) {
	size, n, err := e.ReadUint16()
	if err != nil {
		return "", n, err
	}
	str, m, err := e.ReadString(int(size))
	return str, n + m, err
}

// ReadString32 Reads the length of the string, using ReadUint32 and then Reads the bytes of the string
func (e *BigEndianReader) ReadString32() (string, int, error) {
	size, n, err := e.ReadUint32()
	if err != nil {
		return "", n, err
	}
	str, m, err := e.ReadString(int(size))
	return str, n + m, err
}

// ReadString64 Reads the length of the string, using ReadUint64 and then Reads the bytes of the string
func (e *BigEndianReader) ReadString64() (string, int, error) {
	size, n, err := e.ReadUint64()
	if err != nil {
		return "", n, err
	}
	str, m, err := e.ReadString(int(size))
	return str, n + m, err
}
