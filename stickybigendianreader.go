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
	buffer [9]byte
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

// ReadBool Reads a boolean
func (e *StickyBigEndianReader) ReadBool() bool {
	return e.ReadUint8() != 0
}

// ReadInt8 Reads a 8 bit int as a int8 using the underlying io.Reader
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

// ReadInt16 Reads a 16 bit int as a int16 using the underlying io.Reader
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

// ReadInt24 Reads a 24 bit int as a int32 using the underlying io.Reader
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianReader) ReadInt24() int32 {
	if e.Err != nil {
		return 0
	}
	n, err := io.ReadFull(e.Reader, e.buffer[:3])
	e.Count += int64(n)
	if err != nil {
		e.Err = err
		return 0
	}
	return int32(uint32(e.buffer[2]) | uint32(e.buffer[1])<<8 | uint32(e.buffer[0])<<16)
}

// ReadInt32 Reads a 32 bit int as a int32 using the underlying io.Reader
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

// ReadInt40 Reads a 40 bit int as a int64 using the underlying io.Reader
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianReader) ReadInt40() int64 {
	if e.Err != nil {
		return 0
	}
	n, err := io.ReadFull(e.Reader, e.buffer[:5])
	e.Count += int64(n)
	if err != nil {
		e.Err = err
		return 0
	}
	return int64(uint64(e.buffer[4]) | uint64(e.buffer[3])<<8 | uint64(e.buffer[2])<<16 | uint64(e.buffer[1])<<24 | uint64(e.buffer[0])<<32)
}

// ReadInt48 Reads a 48 bit int as a int64 using the underlying io.Reader
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianReader) ReadInt48() int64 {
	if e.Err != nil {
		return 0
	}
	n, err := io.ReadFull(e.Reader, e.buffer[:6])
	e.Count += int64(n)
	if err != nil {
		e.Err = err
		return 0
	}
	return int64(uint64(e.buffer[5]) | uint64(e.buffer[4])<<8 | uint64(e.buffer[3])<<16 | uint64(e.buffer[2])<<24 | uint64(e.buffer[1])<<32 | uint64(e.buffer[0])<<40)
}

// ReadInt56 Reads a 56 bit int as a int64 using the underlying io.Reader
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianReader) ReadInt56() int64 {
	if e.Err != nil {
		return 0
	}
	n, err := io.ReadFull(e.Reader, e.buffer[:7])
	e.Count += int64(n)
	if err != nil {
		e.Err = err
		return 0
	}
	return int64(uint64(e.buffer[6]) | uint64(e.buffer[5])<<8 | uint64(e.buffer[4])<<16 | uint64(e.buffer[3])<<24 | uint64(e.buffer[2])<<32 | uint64(e.buffer[1])<<40 | uint64(e.buffer[0])<<48)
}

// ReadInt64 Reads a 64 bit int as a int64 using the underlying io.Reader
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

// ReadUint8 Reads a 8 bit uint as a uint8 using the underlying io.Reader
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

// ReadUint16 Reads a 16 bit uint as a uint16 using the underlying io.Reader
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

// ReadUint24 Reads a 24 bit uint as a uint32 using the underlying io.Reader
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianReader) ReadUint24() uint32 {
	if e.Err != nil {
		return 0
	}
	n, err := io.ReadFull(e.Reader, e.buffer[:3])
	e.Count += int64(n)
	if err != nil {
		e.Err = err
		return 0
	}
	return uint32(e.buffer[2]) | uint32(e.buffer[1])<<8 | uint32(e.buffer[0])<<16
}

// ReadUint32 Reads a 32 bit uint as a uint32 using the underlying io.Reader
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

// ReadUint40 Reads a 40 bit uint as a uint64 using the underlying io.Reader
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianReader) ReadUint40() uint64 {
	if e.Err != nil {
		return 0
	}
	n, err := io.ReadFull(e.Reader, e.buffer[:5])
	e.Count += int64(n)
	if err != nil {
		e.Err = err
		return 0
	}
	return uint64(e.buffer[4]) | uint64(e.buffer[3])<<8 | uint64(e.buffer[2])<<16 | uint64(e.buffer[1])<<24 | uint64(e.buffer[0])<<32
}

// ReadUint48 Reads a 48 bit uint as a uint64 using the underlying io.Reader
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianReader) ReadUint48() uint64 {
	if e.Err != nil {
		return 0
	}
	n, err := io.ReadFull(e.Reader, e.buffer[:6])
	e.Count += int64(n)
	if err != nil {
		e.Err = err
		return 0
	}
	return uint64(e.buffer[5]) | uint64(e.buffer[4])<<8 | uint64(e.buffer[3])<<16 | uint64(e.buffer[2])<<24 | uint64(e.buffer[1])<<32 | uint64(e.buffer[0])<<40
}

// ReadUint56 Reads a 56 bit uint as a uint64 using the underlying io.Reader
// Any errors and the running byte read count are stored instead or returned
func (e *StickyBigEndianReader) ReadUint56() uint64 {
	if e.Err != nil {
		return 0
	}
	n, err := io.ReadFull(e.Reader, e.buffer[:7])
	e.Count += int64(n)
	if err != nil {
		e.Err = err
		return 0
	}
	return uint64(e.buffer[6]) | uint64(e.buffer[5])<<8 | uint64(e.buffer[4])<<16 | uint64(e.buffer[3])<<24 | uint64(e.buffer[2])<<32 | uint64(e.buffer[1])<<40 | uint64(e.buffer[0])<<48
}

// ReadUint64 Reads a 64 bit uint as a uint64 using the underlying io.Reader
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

// ReadFloat32 Reads a 32 bit float as a float32 using the underlying io.Reader
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

// ReadFloat64 Reads a 64 bit float as a float64 using the underlying io.Reader
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

// ReadBytes Reads a []byte
func (e *StickyBigEndianReader) ReadBytes(size int) []byte {
	if e.Err != nil {
		return nil
	}
	buf := make([]byte, size)
	var n int
	n, e.Err = io.ReadFull(e.Reader, buf)
	e.Count += int64(n)
	return buf[:n]
}

// ReadBytesX Reads the length of the Bytes, using ReadUintX and then Reads the bytes
func (e *StickyBigEndianReader) ReadBytesX() []byte {
	return e.ReadBytes(int(e.ReadUintX()))
}

// ReadBytes8 Reads the length of the Bytes, using ReadUint8 and then Reads the bytes
func (e *StickyBigEndianReader) ReadBytes8() []byte {
	return e.ReadBytes(int(e.ReadUint8()))
}

// ReadBytes16 Reads the length of the Bytes, using ReadUint16 and then Reads the bytes
func (e *StickyBigEndianReader) ReadBytes16() []byte {
	return e.ReadBytes(int(e.ReadUint16()))
}

// ReadBytes24 Reads the length of the Bytes, using ReadUint24 and then Reads the bytes
func (e *StickyBigEndianReader) ReadBytes24() []byte {
	return e.ReadBytes(int(e.ReadUint24()))
}

// ReadBytes32 Reads the length of the Bytes, using ReadUint32 and then Reads the bytes
func (e *StickyBigEndianReader) ReadBytes32() []byte {
	return e.ReadBytes(int(e.ReadUint32()))
}

// ReadBytes40 Reads the length of the Bytes, using ReadUint40 and then Reads the bytes
func (e *StickyBigEndianReader) ReadBytes40() []byte {
	return e.ReadBytes(int(e.ReadUint40()))
}

// ReadBytes48 Reads the length of the Bytes, using ReadUint48 and then Reads the bytes
func (e *StickyBigEndianReader) ReadBytes48() []byte {
	return e.ReadBytes(int(e.ReadUint48()))
}

// ReadBytes56 Reads the length of the Bytes, using ReadUint56 and then Reads the bytes
func (e *StickyBigEndianReader) ReadBytes56() []byte {
	return e.ReadBytes(int(e.ReadUint56()))
}

// ReadBytes64 Reads the length of the Bytes, using ReadUint64 and then Reads the bytes
func (e *StickyBigEndianReader) ReadBytes64() []byte {
	return e.ReadBytes(int(e.ReadUint64()))
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

// ReadStringX Reads the length of the String, using ReadUintX and then Reads the bytes
func (e *StickyBigEndianReader) ReadStringX() string {
	return e.ReadString(int(e.ReadUintX()))
}

// ReadString8 Reads the length of the String, using ReadUint8 and then Reads the bytes
func (e *StickyBigEndianReader) ReadString8() string {
	return e.ReadString(int(e.ReadUint8()))
}

// ReadString16 Reads the length of the String, using ReadUint16 and then Reads the bytes
func (e *StickyBigEndianReader) ReadString16() string {
	return e.ReadString(int(e.ReadUint16()))
}

// ReadString24 Reads the length of the String, using ReadUint24 and then Reads the bytes
func (e *StickyBigEndianReader) ReadString24() string {
	return e.ReadString(int(e.ReadUint24()))
}

// ReadString32 Reads the length of the String, using ReadUint32 and then Reads the bytes
func (e *StickyBigEndianReader) ReadString32() string {
	return e.ReadString(int(e.ReadUint32()))
}

// ReadString40 Reads the length of the String, using ReadUint40 and then Reads the bytes
func (e *StickyBigEndianReader) ReadString40() string {
	return e.ReadString(int(e.ReadUint40()))
}

// ReadString48 Reads the length of the String, using ReadUint48 and then Reads the bytes
func (e *StickyBigEndianReader) ReadString48() string {
	return e.ReadString(int(e.ReadUint48()))
}

// ReadString56 Reads the length of the String, using ReadUint56 and then Reads the bytes
func (e *StickyBigEndianReader) ReadString56() string {
	return e.ReadString(int(e.ReadUint56()))
}

// ReadString64 Reads the length of the String, using ReadUint64 and then Reads the bytes
func (e *StickyBigEndianReader) ReadString64() string {
	return e.ReadString(int(e.ReadUint64()))
}

// ReadString0 Reads the bytes of the string until a 0 byte is read
func (e *StickyBigEndianReader) ReadString0() string {
	var d []byte
	for {
		p := e.ReadUint8()
		if p == 0 {
			break
		}
		d = append(d, p)
	}
	return string(d)
}
