package byteio

// File automatically generated with ./gen.sh.

import (
	"io"
	"math"
)

// StickyLittleEndianReader wraps a io.Reader to provide methods
// to make it easier to Read fundamental types.
type StickyLittleEndianReader struct {
	io.Reader
	buffer [9]byte
	Err    error
	Count  int64
}

// GetError returns any error received.
func (e *StickyLittleEndianReader) GetError() error {
	return e.Err
}

// GetCount returns the number of bytes read.
func (e *StickyLittleEndianReader) GetCount() int64 {
	return e.Count
}

// Read implements the io.Reader interface.
func (e *StickyLittleEndianReader) Read(p []byte) (int, error) {
	if e.Err != nil {
		return 0, e.Err
	}

	n, err := io.ReadFull(e.Reader, p)
	e.Err = err
	e.Count += int64(n)

	return n, e.Err
}

// ReadBool Reads a boolean.
func (e *StickyLittleEndianReader) ReadBool() bool {
	return e.ReadUint8() != 0
}

// ReadInt8 Reads a 8 bit int as a int8 using the underlying io.Reader.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianReader) ReadInt8() int8 {
	if e.Err != nil {
		return 0
	}

	n, err := io.ReadFull(e.Reader, e.buffer[:1])
	e.Count += int64(n)

	if err != nil {
		e.Err = err

		return 0
	}

	d := e.buffer[0]

	return int8(d)
}

// ReadInt16 Reads a 16 bit int as a int16 using the underlying io.Reader.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianReader) ReadInt16() int16 {
	if e.Err != nil {
		return 0
	}

	n, err := io.ReadFull(e.Reader, e.buffer[:2])
	e.Count += int64(n)

	if err != nil {
		e.Err = err

		return 0
	}

	d := uint16(e.buffer[0]) | uint16(e.buffer[1])<<8

	return int16(d)
}

// ReadInt24 Reads a 24 bit int as a int32 using the underlying io.Reader.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianReader) ReadInt24() int32 {
	if e.Err != nil {
		return 0
	}

	n, err := io.ReadFull(e.Reader, e.buffer[:3])
	e.Count += int64(n)

	if err != nil {
		e.Err = err

		return 0
	}

	d := uint32(e.buffer[0]) | uint32(e.buffer[1])<<8 | uint32(e.buffer[2])<<16

	if d >= 0x00800000 {
		d |= 0xff000000
	}

	return int32(d)
}

// ReadInt32 Reads a 32 bit int as a int32 using the underlying io.Reader.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianReader) ReadInt32() int32 {
	if e.Err != nil {
		return 0
	}

	n, err := io.ReadFull(e.Reader, e.buffer[:4])
	e.Count += int64(n)

	if err != nil {
		e.Err = err

		return 0
	}

	d := uint32(e.buffer[0]) | uint32(e.buffer[1])<<8 | uint32(e.buffer[2])<<16 | uint32(e.buffer[3])<<24

	return int32(d)
}

// ReadInt40 Reads a 40 bit int as a int64 using the underlying io.Reader.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianReader) ReadInt40() int64 {
	if e.Err != nil {
		return 0
	}

	n, err := io.ReadFull(e.Reader, e.buffer[:5])
	e.Count += int64(n)

	if err != nil {
		e.Err = err

		return 0
	}

	d := uint64(e.buffer[0]) | uint64(e.buffer[1])<<8 | uint64(e.buffer[2])<<16 | uint64(e.buffer[3])<<24 | uint64(e.buffer[4])<<32

	if d >= 0x0000008000000000 {
		d |= 0xffffff0000000000
	}

	return int64(d)
}

// ReadInt48 Reads a 48 bit int as a int64 using the underlying io.Reader.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianReader) ReadInt48() int64 {
	if e.Err != nil {
		return 0
	}

	n, err := io.ReadFull(e.Reader, e.buffer[:6])
	e.Count += int64(n)

	if err != nil {
		e.Err = err

		return 0
	}

	d := uint64(e.buffer[0]) | uint64(e.buffer[1])<<8 | uint64(e.buffer[2])<<16 | uint64(e.buffer[3])<<24 | uint64(e.buffer[4])<<32 | uint64(e.buffer[5])<<40

	if d >= 0x0000800000000000 {
		d |= 0xffff000000000000
	}

	return int64(d)
}

// ReadInt56 Reads a 56 bit int as a int64 using the underlying io.Reader.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianReader) ReadInt56() int64 {
	if e.Err != nil {
		return 0
	}

	n, err := io.ReadFull(e.Reader, e.buffer[:7])
	e.Count += int64(n)

	if err != nil {
		e.Err = err

		return 0
	}

	d := uint64(e.buffer[0]) | uint64(e.buffer[1])<<8 | uint64(e.buffer[2])<<16 | uint64(e.buffer[3])<<24 | uint64(e.buffer[4])<<32 | uint64(e.buffer[5])<<40 | uint64(e.buffer[6])<<48

	if d >= 0x0080000000000000 {
		d |= 0xff00000000000000
	}

	return int64(d)
}

// ReadInt64 Reads a 64 bit int as a int64 using the underlying io.Reader.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianReader) ReadInt64() int64 {
	if e.Err != nil {
		return 0
	}

	n, err := io.ReadFull(e.Reader, e.buffer[:8])
	e.Count += int64(n)

	if err != nil {
		e.Err = err

		return 0
	}

	d := uint64(e.buffer[0]) | uint64(e.buffer[1])<<8 | uint64(e.buffer[2])<<16 | uint64(e.buffer[3])<<24 | uint64(e.buffer[4])<<32 | uint64(e.buffer[5])<<40 | uint64(e.buffer[6])<<48 | uint64(e.buffer[7])<<56

	return int64(d)
}

// ReadUint8 Reads a 8 bit uint as a uint8 using the underlying io.Reader.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianReader) ReadUint8() uint8 {
	if e.Err != nil {
		return 0
	}

	n, err := io.ReadFull(e.Reader, e.buffer[:1])
	e.Count += int64(n)

	if err != nil {
		e.Err = err

		return 0
	}

	d := e.buffer[0]

	return d
}

// ReadUint16 Reads a 16 bit uint as a uint16 using the underlying io.Reader.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianReader) ReadUint16() uint16 {
	if e.Err != nil {
		return 0
	}

	n, err := io.ReadFull(e.Reader, e.buffer[:2])
	e.Count += int64(n)

	if err != nil {
		e.Err = err

		return 0
	}

	d := uint16(e.buffer[0]) | uint16(e.buffer[1])<<8

	return d
}

// ReadUint24 Reads a 24 bit uint as a uint32 using the underlying io.Reader.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianReader) ReadUint24() uint32 {
	if e.Err != nil {
		return 0
	}

	n, err := io.ReadFull(e.Reader, e.buffer[:3])
	e.Count += int64(n)

	if err != nil {
		e.Err = err

		return 0
	}

	d := uint32(e.buffer[0]) | uint32(e.buffer[1])<<8 | uint32(e.buffer[2])<<16

	return d
}

// ReadUint32 Reads a 32 bit uint as a uint32 using the underlying io.Reader.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianReader) ReadUint32() uint32 {
	if e.Err != nil {
		return 0
	}

	n, err := io.ReadFull(e.Reader, e.buffer[:4])
	e.Count += int64(n)

	if err != nil {
		e.Err = err

		return 0
	}

	d := uint32(e.buffer[0]) | uint32(e.buffer[1])<<8 | uint32(e.buffer[2])<<16 | uint32(e.buffer[3])<<24

	return d
}

// ReadUint40 Reads a 40 bit uint as a uint64 using the underlying io.Reader.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianReader) ReadUint40() uint64 {
	if e.Err != nil {
		return 0
	}

	n, err := io.ReadFull(e.Reader, e.buffer[:5])
	e.Count += int64(n)

	if err != nil {
		e.Err = err

		return 0
	}

	d := uint64(e.buffer[0]) | uint64(e.buffer[1])<<8 | uint64(e.buffer[2])<<16 | uint64(e.buffer[3])<<24 | uint64(e.buffer[4])<<32

	return d
}

// ReadUint48 Reads a 48 bit uint as a uint64 using the underlying io.Reader.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianReader) ReadUint48() uint64 {
	if e.Err != nil {
		return 0
	}

	n, err := io.ReadFull(e.Reader, e.buffer[:6])
	e.Count += int64(n)

	if err != nil {
		e.Err = err

		return 0
	}

	d := uint64(e.buffer[0]) | uint64(e.buffer[1])<<8 | uint64(e.buffer[2])<<16 | uint64(e.buffer[3])<<24 | uint64(e.buffer[4])<<32 | uint64(e.buffer[5])<<40

	return d
}

// ReadUint56 Reads a 56 bit uint as a uint64 using the underlying io.Reader.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianReader) ReadUint56() uint64 {
	if e.Err != nil {
		return 0
	}

	n, err := io.ReadFull(e.Reader, e.buffer[:7])
	e.Count += int64(n)

	if err != nil {
		e.Err = err

		return 0
	}

	d := uint64(e.buffer[0]) | uint64(e.buffer[1])<<8 | uint64(e.buffer[2])<<16 | uint64(e.buffer[3])<<24 | uint64(e.buffer[4])<<32 | uint64(e.buffer[5])<<40 | uint64(e.buffer[6])<<48

	return d
}

// ReadUint64 Reads a 64 bit uint as a uint64 using the underlying io.Reader.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianReader) ReadUint64() uint64 {
	if e.Err != nil {
		return 0
	}

	n, err := io.ReadFull(e.Reader, e.buffer[:8])
	e.Count += int64(n)

	if err != nil {
		e.Err = err

		return 0
	}

	d := uint64(e.buffer[0]) | uint64(e.buffer[1])<<8 | uint64(e.buffer[2])<<16 | uint64(e.buffer[3])<<24 | uint64(e.buffer[4])<<32 | uint64(e.buffer[5])<<40 | uint64(e.buffer[6])<<48 | uint64(e.buffer[7])<<56

	return d
}

// ReadFloat32 Reads a 32 bit float as a float32 using the underlying io.Reader.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianReader) ReadFloat32() float32 {
	if e.Err != nil {
		return 0
	}

	n, err := io.ReadFull(e.Reader, e.buffer[:4])
	e.Count += int64(n)

	if err != nil {
		e.Err = err

		return 0
	}

	d := math.Float32frombits(uint32(e.buffer[0]) | uint32(e.buffer[1])<<8 | uint32(e.buffer[2])<<16 | uint32(e.buffer[3])<<24)

	return d
}

// ReadFloat64 Reads a 64 bit float as a float64 using the underlying io.Reader.
// Any errors and the running byte read count are stored instead of being returned.
func (e *StickyLittleEndianReader) ReadFloat64() float64 {
	if e.Err != nil {
		return 0
	}

	n, err := io.ReadFull(e.Reader, e.buffer[:8])
	e.Count += int64(n)

	if err != nil {
		e.Err = err

		return 0
	}

	d := math.Float64frombits(uint64(e.buffer[0]) | uint64(e.buffer[1])<<8 | uint64(e.buffer[2])<<16 | uint64(e.buffer[3])<<24 | uint64(e.buffer[4])<<32 | uint64(e.buffer[5])<<40 | uint64(e.buffer[6])<<48 | uint64(e.buffer[7])<<56)

	return d
}

// ReadBytes Reads a []byte.
func (e *StickyLittleEndianReader) ReadBytes(size int) []byte {
	if e.Err != nil {
		return nil
	}

	buf := make([]byte, size)

	var n int
	n, e.Err = io.ReadFull(e.Reader, buf)
	e.Count += int64(n)

	return buf[:n]
}

// ReadBytesX Reads the length of the Bytes, using ReadUintX and then Reads the bytes.
func (e *StickyLittleEndianReader) ReadBytesX() []byte {
	return e.ReadBytes(int(e.ReadUintX()))
}

// ReadBytes8 Reads the length of the Bytes, using ReadUint8 and then Reads the bytes.
func (e *StickyLittleEndianReader) ReadBytes8() []byte {
	return e.ReadBytes(int(e.ReadUint8()))
}

// ReadBytes16 Reads the length of the Bytes, using ReadUint16 and then Reads the bytes.
func (e *StickyLittleEndianReader) ReadBytes16() []byte {
	return e.ReadBytes(int(e.ReadUint16()))
}

// ReadBytes24 Reads the length of the Bytes, using ReadUint24 and then Reads the bytes.
func (e *StickyLittleEndianReader) ReadBytes24() []byte {
	return e.ReadBytes(int(e.ReadUint24()))
}

// ReadBytes32 Reads the length of the Bytes, using ReadUint32 and then Reads the bytes.
func (e *StickyLittleEndianReader) ReadBytes32() []byte {
	return e.ReadBytes(int(e.ReadUint32()))
}

// ReadBytes40 Reads the length of the Bytes, using ReadUint40 and then Reads the bytes.
func (e *StickyLittleEndianReader) ReadBytes40() []byte {
	return e.ReadBytes(int(e.ReadUint40()))
}

// ReadBytes48 Reads the length of the Bytes, using ReadUint48 and then Reads the bytes.
func (e *StickyLittleEndianReader) ReadBytes48() []byte {
	return e.ReadBytes(int(e.ReadUint48()))
}

// ReadBytes56 Reads the length of the Bytes, using ReadUint56 and then Reads the bytes.
func (e *StickyLittleEndianReader) ReadBytes56() []byte {
	return e.ReadBytes(int(e.ReadUint56()))
}

// ReadBytes64 Reads the length of the Bytes, using ReadUint64 and then Reads the bytes.
func (e *StickyLittleEndianReader) ReadBytes64() []byte {
	return e.ReadBytes(int(e.ReadUint64()))
}

// ReadString Reads a string.
func (e *StickyLittleEndianReader) ReadString(size int) string {
	if e.Err != nil {
		return ""
	}

	buf := make([]byte, size)

	var n int
	n, e.Err = io.ReadFull(e.Reader, buf)
	e.Count += int64(n)

	return string(buf[:n])
}

// ReadStringX Reads the length of the String, using ReadUintX and then Reads the bytes.
func (e *StickyLittleEndianReader) ReadStringX() string {
	return e.ReadString(int(e.ReadUintX()))
}

// ReadString8 Reads the length of the String, using ReadUint8 and then Reads the bytes.
func (e *StickyLittleEndianReader) ReadString8() string {
	return e.ReadString(int(e.ReadUint8()))
}

// ReadString16 Reads the length of the String, using ReadUint16 and then Reads the bytes.
func (e *StickyLittleEndianReader) ReadString16() string {
	return e.ReadString(int(e.ReadUint16()))
}

// ReadString24 Reads the length of the String, using ReadUint24 and then Reads the bytes.
func (e *StickyLittleEndianReader) ReadString24() string {
	return e.ReadString(int(e.ReadUint24()))
}

// ReadString32 Reads the length of the String, using ReadUint32 and then Reads the bytes.
func (e *StickyLittleEndianReader) ReadString32() string {
	return e.ReadString(int(e.ReadUint32()))
}

// ReadString40 Reads the length of the String, using ReadUint40 and then Reads the bytes.
func (e *StickyLittleEndianReader) ReadString40() string {
	return e.ReadString(int(e.ReadUint40()))
}

// ReadString48 Reads the length of the String, using ReadUint48 and then Reads the bytes.
func (e *StickyLittleEndianReader) ReadString48() string {
	return e.ReadString(int(e.ReadUint48()))
}

// ReadString56 Reads the length of the String, using ReadUint56 and then Reads the bytes.
func (e *StickyLittleEndianReader) ReadString56() string {
	return e.ReadString(int(e.ReadUint56()))
}

// ReadString64 Reads the length of the String, using ReadUint64 and then Reads the bytes.
func (e *StickyLittleEndianReader) ReadString64() string {
	return e.ReadString(int(e.ReadUint64()))
}

// ReadString0 Reads the bytes of the string until a 0 byte is read.
func (e *StickyLittleEndianReader) ReadString0() string {
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
