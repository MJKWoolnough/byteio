package byteio

// File automatically generated with ./gen.sh.

import (
	"io"
	"math"
)

// MemBigEndian is a byte slice that has methods to make it easier to read and write fundamental types.
type MemBigEndian []byte

// GetError returns any error received.
func (e MemBigEndian) GetError() error {
	return nil
}

// GetCount returns the length of the slice.
func (e MemBigEndian) GetCount() int64 {
	return int64(len(e))
}

// ReadBool Reads a boolean.
func (e *MemBigEndian) ReadBool() bool {
	return e.ReadUint8() != 0
}

// ReadInt8 Reads a 8 bit int as a int8 from the byte slice.
func (e *MemBigEndian) ReadInt8() int8 {
	if len(*e) < 1 {
		return 0
	}

	d := (*e)[0]
	*e = (*e)[1:]

	return int8(d)
}

// ReadInt16 Reads a 16 bit int as a int16 from the byte slice.
func (e *MemBigEndian) ReadInt16() int16 {
	if len(*e) < 2 {
		return 0
	}

	d := uint16((*e)[1]) | uint16((*e)[0])<<8
	*e = (*e)[2:]

	return int16(d)
}

// ReadInt24 Reads a 24 bit int as a int32 from the byte slice.
func (e *MemBigEndian) ReadInt24() int32 {
	if len(*e) < 3 {
		return 0
	}

	d := uint32((*e)[2]) | uint32((*e)[1])<<8 | uint32((*e)[0])<<16

	if d >= 0x00800000 {
		d |= 0xff000000
	}

	*e = (*e)[3:]

	return int32(d)
}

// ReadInt32 Reads a 32 bit int as a int32 from the byte slice.
func (e *MemBigEndian) ReadInt32() int32 {
	if len(*e) < 4 {
		return 0
	}

	d := uint32((*e)[3]) | uint32((*e)[2])<<8 | uint32((*e)[1])<<16 | uint32((*e)[0])<<24
	*e = (*e)[4:]

	return int32(d)
}

// ReadInt40 Reads a 40 bit int as a int64 from the byte slice.
func (e *MemBigEndian) ReadInt40() int64 {
	if len(*e) < 5 {
		return 0
	}

	d := uint64((*e)[4]) | uint64((*e)[3])<<8 | uint64((*e)[2])<<16 | uint64((*e)[1])<<24 | uint64((*e)[0])<<32

	if d >= 0x0000008000000000 {
		d |= 0xffffff0000000000
	}

	*e = (*e)[5:]

	return int64(d)
}

// ReadInt48 Reads a 48 bit int as a int64 from the byte slice.
func (e *MemBigEndian) ReadInt48() int64 {
	if len(*e) < 6 {
		return 0
	}

	d := uint64((*e)[5]) | uint64((*e)[4])<<8 | uint64((*e)[3])<<16 | uint64((*e)[2])<<24 | uint64((*e)[1])<<32 | uint64((*e)[0])<<40

	if d >= 0x0000800000000000 {
		d |= 0xffff000000000000
	}

	*e = (*e)[6:]

	return int64(d)
}

// ReadInt56 Reads a 56 bit int as a int64 from the byte slice.
func (e *MemBigEndian) ReadInt56() int64 {
	if len(*e) < 7 {
		return 0
	}

	d := uint64((*e)[6]) | uint64((*e)[5])<<8 | uint64((*e)[4])<<16 | uint64((*e)[3])<<24 | uint64((*e)[2])<<32 | uint64((*e)[1])<<40 | uint64((*e)[0])<<48

	if d >= 0x0080000000000000 {
		d |= 0xff00000000000000
	}

	*e = (*e)[7:]

	return int64(d)
}

// ReadInt64 Reads a 64 bit int as a int64 from the byte slice.
func (e *MemBigEndian) ReadInt64() int64 {
	if len(*e) < 8 {
		return 0
	}

	d := uint64((*e)[7]) | uint64((*e)[6])<<8 | uint64((*e)[5])<<16 | uint64((*e)[4])<<24 | uint64((*e)[3])<<32 | uint64((*e)[2])<<40 | uint64((*e)[1])<<48 | uint64((*e)[0])<<56
	*e = (*e)[8:]

	return int64(d)
}

// ReadUint8 Reads a 8 bit uint as a uint8 from the byte slice.
func (e *MemBigEndian) ReadUint8() uint8 {
	if len(*e) < 1 {
		return 0
	}

	d := (*e)[0]
	*e = (*e)[1:]

	return d
}

// ReadUint16 Reads a 16 bit uint as a uint16 from the byte slice.
func (e *MemBigEndian) ReadUint16() uint16 {
	if len(*e) < 2 {
		return 0
	}

	d := uint16((*e)[1]) | uint16((*e)[0])<<8
	*e = (*e)[2:]

	return d
}

// ReadUint24 Reads a 24 bit uint as a uint32 from the byte slice.
func (e *MemBigEndian) ReadUint24() uint32 {
	if len(*e) < 3 {
		return 0
	}

	d := uint32((*e)[2]) | uint32((*e)[1])<<8 | uint32((*e)[0])<<16
	*e = (*e)[3:]

	return d
}

// ReadUint32 Reads a 32 bit uint as a uint32 from the byte slice.
func (e *MemBigEndian) ReadUint32() uint32 {
	if len(*e) < 4 {
		return 0
	}

	d := uint32((*e)[3]) | uint32((*e)[2])<<8 | uint32((*e)[1])<<16 | uint32((*e)[0])<<24
	*e = (*e)[4:]

	return d
}

// ReadUint40 Reads a 40 bit uint as a uint64 from the byte slice.
func (e *MemBigEndian) ReadUint40() uint64 {
	if len(*e) < 5 {
		return 0
	}

	d := uint64((*e)[4]) | uint64((*e)[3])<<8 | uint64((*e)[2])<<16 | uint64((*e)[1])<<24 | uint64((*e)[0])<<32
	*e = (*e)[5:]

	return d
}

// ReadUint48 Reads a 48 bit uint as a uint64 from the byte slice.
func (e *MemBigEndian) ReadUint48() uint64 {
	if len(*e) < 6 {
		return 0
	}

	d := uint64((*e)[5]) | uint64((*e)[4])<<8 | uint64((*e)[3])<<16 | uint64((*e)[2])<<24 | uint64((*e)[1])<<32 | uint64((*e)[0])<<40
	*e = (*e)[6:]

	return d
}

// ReadUint56 Reads a 56 bit uint as a uint64 from the byte slice.
func (e *MemBigEndian) ReadUint56() uint64 {
	if len(*e) < 7 {
		return 0
	}

	d := uint64((*e)[6]) | uint64((*e)[5])<<8 | uint64((*e)[4])<<16 | uint64((*e)[3])<<24 | uint64((*e)[2])<<32 | uint64((*e)[1])<<40 | uint64((*e)[0])<<48
	*e = (*e)[7:]

	return d
}

// ReadUint64 Reads a 64 bit uint as a uint64 from the byte slice.
func (e *MemBigEndian) ReadUint64() uint64 {
	if len(*e) < 8 {
		return 0
	}

	d := uint64((*e)[7]) | uint64((*e)[6])<<8 | uint64((*e)[5])<<16 | uint64((*e)[4])<<24 | uint64((*e)[3])<<32 | uint64((*e)[2])<<40 | uint64((*e)[1])<<48 | uint64((*e)[0])<<56
	*e = (*e)[8:]

	return d
}

// ReadFloat32 Reads a 32 bit float as a float32 from the byte slice.
func (e *MemBigEndian) ReadFloat32() float32 {
	if len(*e) < 4 {
		return 0
	}

	d := math.Float32frombits(uint32((*e)[3]) | uint32((*e)[2])<<8 | uint32((*e)[1])<<16 | uint32((*e)[0])<<24)
	*e = (*e)[4:]

	return d
}

// ReadFloat64 Reads a 64 bit float as a float64 from the byte slice.
func (e *MemBigEndian) ReadFloat64() float64 {
	if len(*e) < 8 {
		return 0
	}

	d := math.Float64frombits(uint64((*e)[7]) | uint64((*e)[6])<<8 | uint64((*e)[5])<<16 | uint64((*e)[4])<<24 | uint64((*e)[3])<<32 | uint64((*e)[2])<<40 | uint64((*e)[1])<<48 | uint64((*e)[0])<<56)
	*e = (*e)[8:]

	return d
}

// ReadBytes Reads a []byte.
func (e *MemBigEndian) ReadBytes(size int) []byte {
	if len(*e) < size {
		return nil
	}

	d := []byte((*e)[:size])
	*e = (*e)[size:]

	return d
}

// MemBigEndianReader implements io.Reader.
func (e *MemBigEndian) Read(p []byte) (int, error) {
	n := copy(p, *e)

	if n == 0 && len(p) != 0 {
		return 0, io.EOF
	}

	*e = (*e)[n:]

	return n, nil
}

// ReadBytesX Reads the length of the Bytes, using ReadUintX and then Reads the bytes.
func (e *MemBigEndian) ReadBytesX() []byte {
	return e.ReadBytes(int(e.ReadUintX()))
}

// ReadBytes8 Reads the length of the Bytes, using ReadUint8 and then Reads the bytes.
func (e *MemBigEndian) ReadBytes8() []byte {
	return e.ReadBytes(int(e.ReadUint8()))
}

// ReadBytes16 Reads the length of the Bytes, using ReadUint16 and then Reads the bytes.
func (e *MemBigEndian) ReadBytes16() []byte {
	return e.ReadBytes(int(e.ReadUint16()))
}

// ReadBytes24 Reads the length of the Bytes, using ReadUint24 and then Reads the bytes.
func (e *MemBigEndian) ReadBytes24() []byte {
	return e.ReadBytes(int(e.ReadUint24()))
}

// ReadBytes32 Reads the length of the Bytes, using ReadUint32 and then Reads the bytes.
func (e *MemBigEndian) ReadBytes32() []byte {
	return e.ReadBytes(int(e.ReadUint32()))
}

// ReadBytes40 Reads the length of the Bytes, using ReadUint40 and then Reads the bytes.
func (e *MemBigEndian) ReadBytes40() []byte {
	return e.ReadBytes(int(e.ReadUint40()))
}

// ReadBytes48 Reads the length of the Bytes, using ReadUint48 and then Reads the bytes.
func (e *MemBigEndian) ReadBytes48() []byte {
	return e.ReadBytes(int(e.ReadUint48()))
}

// ReadBytes56 Reads the length of the Bytes, using ReadUint56 and then Reads the bytes.
func (e *MemBigEndian) ReadBytes56() []byte {
	return e.ReadBytes(int(e.ReadUint56()))
}

// ReadBytes64 Reads the length of the Bytes, using ReadUint64 and then Reads the bytes.
func (e *MemBigEndian) ReadBytes64() []byte {
	return e.ReadBytes(int(e.ReadUint64()))
}

// ReadString Reads a string.
func (e *MemBigEndian) ReadString(size int) string {
	if len(*e) < size {
		return ""
	}

	d := string((*e)[:size])
	*e = (*e)[size:]

	return d
}

// ReadStringX Reads the length of the String, using ReadUintX and then Reads the bytes.
func (e *MemBigEndian) ReadStringX() string {
	return e.ReadString(int(e.ReadUintX()))
}

// ReadString8 Reads the length of the String, using ReadUint8 and then Reads the bytes.
func (e *MemBigEndian) ReadString8() string {
	return e.ReadString(int(e.ReadUint8()))
}

// ReadString16 Reads the length of the String, using ReadUint16 and then Reads the bytes.
func (e *MemBigEndian) ReadString16() string {
	return e.ReadString(int(e.ReadUint16()))
}

// ReadString24 Reads the length of the String, using ReadUint24 and then Reads the bytes.
func (e *MemBigEndian) ReadString24() string {
	return e.ReadString(int(e.ReadUint24()))
}

// ReadString32 Reads the length of the String, using ReadUint32 and then Reads the bytes.
func (e *MemBigEndian) ReadString32() string {
	return e.ReadString(int(e.ReadUint32()))
}

// ReadString40 Reads the length of the String, using ReadUint40 and then Reads the bytes.
func (e *MemBigEndian) ReadString40() string {
	return e.ReadString(int(e.ReadUint40()))
}

// ReadString48 Reads the length of the String, using ReadUint48 and then Reads the bytes.
func (e *MemBigEndian) ReadString48() string {
	return e.ReadString(int(e.ReadUint48()))
}

// ReadString56 Reads the length of the String, using ReadUint56 and then Reads the bytes.
func (e *MemBigEndian) ReadString56() string {
	return e.ReadString(int(e.ReadUint56()))
}

// ReadString64 Reads the length of the String, using ReadUint64 and then Reads the bytes.
func (e *MemBigEndian) ReadString64() string {
	return e.ReadString(int(e.ReadUint64()))
}

// ReadString0 Reads the bytes of the string until a 0 byte is read.
func (e *MemBigEndian) ReadString0() string {
	var d string

	for n, c := range *e {
		if c == 0 {
			d = string((*e)[:n])
			*e = (*e)[n+1:]

			break
		}
	}

	return d
}

// WriteBool Writes a boolean.
func (e *MemBigEndian) WriteBool(b bool) {
	if b {
		e.WriteUint8(1)
	} else {
		e.WriteUint8(0)
	}
}

// WriteInt8 Writes a 8 bit int as a int8 to the byte slice.
func (e *MemBigEndian) WriteInt8(d int8) {
	c := uint8(d)
	*e = append(*e,
		byte(c),
	)
}

// WriteInt16 Writes a 16 bit int as a int16 to the byte slice.
func (e *MemBigEndian) WriteInt16(d int16) {
	c := uint16(d)
	*e = append(*e,
		byte(c>>8),
		byte(c),
	)
}

// WriteInt24 Writes a 24 bit int as a int32 to the byte slice.
func (e *MemBigEndian) WriteInt24(d int32) {
	c := uint32(d)
	*e = append(*e,
		byte(c>>16),
		byte(c>>8),
		byte(c),
	)
}

// WriteInt32 Writes a 32 bit int as a int32 to the byte slice.
func (e *MemBigEndian) WriteInt32(d int32) {
	c := uint32(d)
	*e = append(*e,
		byte(c>>24),
		byte(c>>16),
		byte(c>>8),
		byte(c),
	)
}

// WriteInt40 Writes a 40 bit int as a int64 to the byte slice.
func (e *MemBigEndian) WriteInt40(d int64) {
	c := uint64(d)
	*e = append(*e,
		byte(c>>32),
		byte(c>>24),
		byte(c>>16),
		byte(c>>8),
		byte(c),
	)
}

// WriteInt48 Writes a 48 bit int as a int64 to the byte slice.
func (e *MemBigEndian) WriteInt48(d int64) {
	c := uint64(d)
	*e = append(*e,
		byte(c>>40),
		byte(c>>32),
		byte(c>>24),
		byte(c>>16),
		byte(c>>8),
		byte(c),
	)
}

// WriteInt56 Writes a 56 bit int as a int64 to the byte slice.
func (e *MemBigEndian) WriteInt56(d int64) {
	c := uint64(d)
	*e = append(*e,
		byte(c>>48),
		byte(c>>40),
		byte(c>>32),
		byte(c>>24),
		byte(c>>16),
		byte(c>>8),
		byte(c),
	)
}

// WriteInt64 Writes a 64 bit int as a int64 to the byte slice.
func (e *MemBigEndian) WriteInt64(d int64) {
	c := uint64(d)
	*e = append(*e,
		byte(c>>56),
		byte(c>>48),
		byte(c>>40),
		byte(c>>32),
		byte(c>>24),
		byte(c>>16),
		byte(c>>8),
		byte(c),
	)
}

// WriteUint8 Writes a 8 bit uint as a uint8 to the byte slice.
func (e *MemBigEndian) WriteUint8(d uint8) {
	*e = append(*e,
		byte(d),
	)
}

// WriteUint16 Writes a 16 bit uint as a uint16 to the byte slice.
func (e *MemBigEndian) WriteUint16(d uint16) {
	*e = append(*e,
		byte(d>>8),
		byte(d),
	)
}

// WriteUint24 Writes a 24 bit uint as a uint32 to the byte slice.
func (e *MemBigEndian) WriteUint24(d uint32) {
	*e = append(*e,
		byte(d>>16),
		byte(d>>8),
		byte(d),
	)
}

// WriteUint32 Writes a 32 bit uint as a uint32 to the byte slice.
func (e *MemBigEndian) WriteUint32(d uint32) {
	*e = append(*e,
		byte(d>>24),
		byte(d>>16),
		byte(d>>8),
		byte(d),
	)
}

// WriteUint40 Writes a 40 bit uint as a uint64 to the byte slice.
func (e *MemBigEndian) WriteUint40(d uint64) {
	*e = append(*e,
		byte(d>>32),
		byte(d>>24),
		byte(d>>16),
		byte(d>>8),
		byte(d),
	)
}

// WriteUint48 Writes a 48 bit uint as a uint64 to the byte slice.
func (e *MemBigEndian) WriteUint48(d uint64) {
	*e = append(*e,
		byte(d>>40),
		byte(d>>32),
		byte(d>>24),
		byte(d>>16),
		byte(d>>8),
		byte(d),
	)
}

// WriteUint56 Writes a 56 bit uint as a uint64 to the byte slice.
func (e *MemBigEndian) WriteUint56(d uint64) {
	*e = append(*e,
		byte(d>>48),
		byte(d>>40),
		byte(d>>32),
		byte(d>>24),
		byte(d>>16),
		byte(d>>8),
		byte(d),
	)
}

// WriteUint64 Writes a 64 bit uint as a uint64 to the byte slice.
func (e *MemBigEndian) WriteUint64(d uint64) {
	*e = append(*e,
		byte(d>>56),
		byte(d>>48),
		byte(d>>40),
		byte(d>>32),
		byte(d>>24),
		byte(d>>16),
		byte(d>>8),
		byte(d),
	)
}

// WriteFloat32 Writes a 32 bit float as a float32 to the byte slice.
func (e *MemBigEndian) WriteFloat32(d float32) {
	c := math.Float32bits(d)
	*e = append(*e,
		byte(c>>24),
		byte(c>>16),
		byte(c>>8),
		byte(c),
	)
}

// WriteFloat64 Writes a 64 bit float as a float64 to the byte slice.
func (e *MemBigEndian) WriteFloat64(d float64) {
	c := math.Float64bits(d)
	*e = append(*e,
		byte(c>>56),
		byte(c>>48),
		byte(c>>40),
		byte(c>>32),
		byte(c>>24),
		byte(c>>16),
		byte(c>>8),
		byte(c),
	)
}

// WriteBytes Writes a []byte.
func (e *MemBigEndian) WriteBytes(d []byte) {
	*e = append(*e, d...)
}

// MemBigEndianWriter implements io.Writer.
func (e *MemBigEndian) Write(p []byte) (int, error) {
	*e = append(*e, p...)

	return len(p), nil
}

// WriteBytesX Writes the length of the Bytes, using WriteUintX and then Writes the bytes.
func (e *MemBigEndian) WriteBytesX(p []byte) {
	e.WriteUintX(uint64(len(p)))
	e.WriteBytes(p)
}

// WriteBytes8 Writes the length of the Bytes, using WriteUint8 and then Writes the bytes.
func (e *MemBigEndian) WriteBytes8(p []byte) {
	e.WriteUint8(uint8(len(p)))
	e.WriteBytes(p)
}

// WriteBytes16 Writes the length of the Bytes, using WriteUint16 and then Writes the bytes.
func (e *MemBigEndian) WriteBytes16(p []byte) {
	e.WriteUint16(uint16(len(p)))
	e.WriteBytes(p)
}

// WriteBytes24 Writes the length of the Bytes, using WriteUint24 and then Writes the bytes.
func (e *MemBigEndian) WriteBytes24(p []byte) {
	e.WriteUint24(uint32(len(p)))
	e.WriteBytes(p)
}

// WriteBytes32 Writes the length of the Bytes, using WriteUint32 and then Writes the bytes.
func (e *MemBigEndian) WriteBytes32(p []byte) {
	e.WriteUint32(uint32(len(p)))
	e.WriteBytes(p)
}

// WriteBytes40 Writes the length of the Bytes, using WriteUint40 and then Writes the bytes.
func (e *MemBigEndian) WriteBytes40(p []byte) {
	e.WriteUint40(uint64(len(p)))
	e.WriteBytes(p)
}

// WriteBytes48 Writes the length of the Bytes, using WriteUint48 and then Writes the bytes.
func (e *MemBigEndian) WriteBytes48(p []byte) {
	e.WriteUint48(uint64(len(p)))
	e.WriteBytes(p)
}

// WriteBytes56 Writes the length of the Bytes, using WriteUint56 and then Writes the bytes.
func (e *MemBigEndian) WriteBytes56(p []byte) {
	e.WriteUint56(uint64(len(p)))
	e.WriteBytes(p)
}

// WriteBytes64 Writes the length of the Bytes, using WriteUint64 and then Writes the bytes.
func (e *MemBigEndian) WriteBytes64(p []byte) {
	e.WriteUint64(uint64(len(p)))
	e.WriteBytes(p)
}

// WriteString Writes a string.
func (e *MemBigEndian) WriteString(d string) (int, error) {
	*e = append(*e, d...)

	return len(d), nil
}

// WriteStringX Writes the length of the String, using WriteUintX and then Writes the bytes.
func (e *MemBigEndian) WriteStringX(p string) {
	e.WriteUintX(uint64(len(p)))
	e.WriteString(p)
}

// WriteString8 Writes the length of the String, using WriteUint8 and then Writes the bytes.
func (e *MemBigEndian) WriteString8(p string) {
	e.WriteUint8(uint8(len(p)))
	e.WriteString(p)
}

// WriteString16 Writes the length of the String, using WriteUint16 and then Writes the bytes.
func (e *MemBigEndian) WriteString16(p string) {
	e.WriteUint16(uint16(len(p)))
	e.WriteString(p)
}

// WriteString24 Writes the length of the String, using WriteUint24 and then Writes the bytes.
func (e *MemBigEndian) WriteString24(p string) {
	e.WriteUint24(uint32(len(p)))
	e.WriteString(p)
}

// WriteString32 Writes the length of the String, using WriteUint32 and then Writes the bytes.
func (e *MemBigEndian) WriteString32(p string) {
	e.WriteUint32(uint32(len(p)))
	e.WriteString(p)
}

// WriteString40 Writes the length of the String, using WriteUint40 and then Writes the bytes.
func (e *MemBigEndian) WriteString40(p string) {
	e.WriteUint40(uint64(len(p)))
	e.WriteString(p)
}

// WriteString48 Writes the length of the String, using WriteUint48 and then Writes the bytes.
func (e *MemBigEndian) WriteString48(p string) {
	e.WriteUint48(uint64(len(p)))
	e.WriteString(p)
}

// WriteString56 Writes the length of the String, using WriteUint56 and then Writes the bytes.
func (e *MemBigEndian) WriteString56(p string) {
	e.WriteUint56(uint64(len(p)))
	e.WriteString(p)
}

// WriteString64 Writes the length of the String, using WriteUint64 and then Writes the bytes.
func (e *MemBigEndian) WriteString64(p string) {
	e.WriteUint64(uint64(len(p)))
	e.WriteString(p)
}

// WriteString0 Writes the bytes of the string ending with a 0 byte.
func (e *MemBigEndian) WriteString0(p string) {
	e.WriteString(p)
	e.WriteUint8(0)
}
