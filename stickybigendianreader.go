package byteio

// File automatically generated with ./gen.sh

import (
	"io"
	"math"
)

// StickyBigEndianReader wraps a io.Reader to provide methods
// to make it easier to Read fundamental types
type StickyBigEndianReader struct {
	io.Reader
	buffer [8]byte
	Err    error
	Count  int64
}

// Read implements the io.Reader interface
func (e *StickyBigEndianReader) Read(p []byte) (int, error) {
	if e.Err != nil {
		return 0, e.Err
	}
	var n int
	n, e.Err = e.Reader.Read(p)
	e.Count += int64(n)
	return n, e.Err
}

// ReadInt8 Reads a int8 using the underlying io.Reader
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianReader) ReadInt8() int8 {
	if e.Err != nil {
		return 0
	}
	n, err := io.ReadFull(e.Reader, e.buffer[:1])
	e.Count += int64(n)
	if err != nil {
		e.Err = err
		return 0
	}
	return int8(e.buffer[0])
}

// ReadInt16 Reads a int16 using the underlying io.Reader
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianReader) ReadInt16() int16 {
	if e.Err != nil {
		return 0
	}
	n, err := io.ReadFull(e.Reader, e.buffer[:2])
	e.Count += int64(n)
	if err != nil {
		e.Err = err
		return 0
	}
	return int16(uint16(e.buffer[1]) | uint16(e.buffer[0])<<8)
}

// ReadInt32 Reads a int32 using the underlying io.Reader
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianReader) ReadInt32() int32 {
	if e.Err != nil {
		return 0
	}
	n, err := io.ReadFull(e.Reader, e.buffer[:4])
	e.Count += int64(n)
	if err != nil {
		e.Err = err
		return 0
	}
	return int32(uint32(e.buffer[3]) | uint32(e.buffer[2])<<8 | uint32(e.buffer[1])<<16 | uint32(e.buffer[0])<<24)
}

// ReadInt64 Reads a int64 using the underlying io.Reader
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianReader) ReadInt64() int64 {
	if e.Err != nil {
		return 0
	}
	n, err := io.ReadFull(e.Reader, e.buffer[:8])
	e.Count += int64(n)
	if err != nil {
		e.Err = err
		return 0
	}
	return int64(uint64(e.buffer[7]) | uint64(e.buffer[6])<<8 | uint64(e.buffer[5])<<16 | uint64(e.buffer[4])<<24 | uint64(e.buffer[3])<<32 | uint64(e.buffer[2])<<40 | uint64(e.buffer[1])<<48 | uint64(e.buffer[0])<<56)
}

// ReadUint8 Reads a uint8 using the underlying io.Reader
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianReader) ReadUint8() uint8 {
	if e.Err != nil {
		return 0
	}
	n, err := io.ReadFull(e.Reader, e.buffer[:1])
	e.Count += int64(n)
	if err != nil {
		e.Err = err
		return 0
	}
	return e.buffer[0]
}

// ReadUint16 Reads a uint16 using the underlying io.Reader
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianReader) ReadUint16() uint16 {
	if e.Err != nil {
		return 0
	}
	n, err := io.ReadFull(e.Reader, e.buffer[:2])
	e.Count += int64(n)
	if err != nil {
		e.Err = err
		return 0
	}
	return uint16(e.buffer[1]) | uint16(e.buffer[0])<<8
}

// ReadUint32 Reads a uint32 using the underlying io.Reader
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianReader) ReadUint32() uint32 {
	if e.Err != nil {
		return 0
	}
	n, err := io.ReadFull(e.Reader, e.buffer[:4])
	e.Count += int64(n)
	if err != nil {
		e.Err = err
		return 0
	}
	return uint32(e.buffer[3]) | uint32(e.buffer[2])<<8 | uint32(e.buffer[1])<<16 | uint32(e.buffer[0])<<24
}

// ReadUint64 Reads a uint64 using the underlying io.Reader
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianReader) ReadUint64() uint64 {
	if e.Err != nil {
		return 0
	}
	n, err := io.ReadFull(e.Reader, e.buffer[:8])
	e.Count += int64(n)
	if err != nil {
		e.Err = err
		return 0
	}
	return uint64(e.buffer[7]) | uint64(e.buffer[6])<<8 | uint64(e.buffer[5])<<16 | uint64(e.buffer[4])<<24 | uint64(e.buffer[3])<<32 | uint64(e.buffer[2])<<40 | uint64(e.buffer[1])<<48 | uint64(e.buffer[0])<<56
}

// ReadFloat32 Reads a float32 using the underlying io.Reader
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianReader) ReadFloat32() float32 {
	if e.Err != nil {
		return 0
	}
	n, err := io.ReadFull(e.Reader, e.buffer[:4])
	e.Count += int64(n)
	if err != nil {
		e.Err = err
		return 0
	}
	return math.Float32frombits(uint32(e.buffer[3]) | uint32(e.buffer[2])<<8 | uint32(e.buffer[1])<<16 | uint32(e.buffer[0])<<24)
}

// ReadFloat64 Reads a float64 using the underlying io.Reader
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianReader) ReadFloat64() float64 {
	if e.Err != nil {
		return 0
	}
	n, err := io.ReadFull(e.Reader, e.buffer[:8])
	e.Count += int64(n)
	if err != nil {
		e.Err = err
		return 0
	}
	return math.Float64frombits(uint64(e.buffer[7]) | uint64(e.buffer[6])<<8 | uint64(e.buffer[5])<<16 | uint64(e.buffer[4])<<24 | uint64(e.buffer[3])<<32 | uint64(e.buffer[2])<<40 | uint64(e.buffer[1])<<48 | uint64(e.buffer[0])<<56)
}

// ReadString Reads a string
func (e *StickyBigEndianReader) ReadString(size int) string {
	if e.Err != nil {
		return ""
	}
	buf := make([]byte, size)
	var n int
	n, e.Err = io.ReadFull(e.Reader, buf)
	e.Count += int64(n)
	return string(buf[:n])
}

// ReadStringX Reads the length of the string, using ReadUintX and then reads the bytes of the string
func (e *StickyBigEndianReader) ReadStringX() string {
	return e.ReadString(int(e.ReadUintX()))
}

// ReadString8 Reads the length of the string, using ReadUint8 and then reads the bytes of the string
func (e *StickyBigEndianReader) ReadString8() string {
	return e.ReadString(int(e.ReadUint8()))
}

// ReadString16 Reads the length of the string, using ReadUint16 and then reads the bytes of the string
func (e *StickyBigEndianReader) ReadString16() string {
	return e.ReadString(int(e.ReadUint16()))
}

// ReadString32 Reads the length of the string, using ReadUint32 and then reads the bytes of the string
func (e *StickyBigEndianReader) ReadString32() string {
	return e.ReadString(int(e.ReadUint32()))
}

// ReadString64 Reads the length of the string, using ReadUint64 and then reads the bytes of the string
func (e *StickyBigEndianReader) ReadString64() string {
	return e.ReadString(int(e.ReadUint64()))
}
