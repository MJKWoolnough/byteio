package byteio

// File automatically generated with ./gen.sh.

import (
	"io"
	"math"
)

// MemLittleEndianReader is a byte slice that has methods
// to make it easier to Read fundamental types.
type MemLittleEndianReader []byte

// ReadBool Reads a boolean.
func (e *MemLittleEndianReader) ReadBool() bool {
	return e.ReadUint8() != 0
}

// ReadInt8 Reads a 8 bit int as a int8 from the byte slice.
func (e *MemLittleEndianReader) ReadInt8() int8 {
	if len(*e) < 1 {
		return 0
	}

	d := (*e)[0]
	*e = (*e)[1:]

	return int8(d)
}

// ReadInt16 Reads a 16 bit int as a int16 from the byte slice.
func (e *MemLittleEndianReader) ReadInt16() int16 {
	if len(*e) < 2 {
		return 0
	}

	d := uint16((*e)[0]) | uint16((*e)[1])<<8
	*e = (*e)[2:]

	return int16(d)
}

// ReadInt24 Reads a 24 bit int as a int32 from the byte slice.
func (e *MemLittleEndianReader) ReadInt24() int32 {
	if len(*e) < 3 {
		return 0
	}

	d := uint32((*e)[0]) | uint32((*e)[1])<<8 | uint32((*e)[2])<<16

	if d >= 0x00800000 {
		d |= 0xff000000
	}

	*e = (*e)[3:]

	return int32(d)
}

// ReadInt32 Reads a 32 bit int as a int32 from the byte slice.
func (e *MemLittleEndianReader) ReadInt32() int32 {
	if len(*e) < 4 {
		return 0
	}

	d := uint32((*e)[0]) | uint32((*e)[1])<<8 | uint32((*e)[2])<<16 | uint32((*e)[3])<<24
	*e = (*e)[4:]

	return int32(d)
}

// ReadInt40 Reads a 40 bit int as a int64 from the byte slice.
func (e *MemLittleEndianReader) ReadInt40() int64 {
	if len(*e) < 5 {
		return 0
	}

	d := uint64((*e)[0]) | uint64((*e)[1])<<8 | uint64((*e)[2])<<16 | uint64((*e)[3])<<24 | uint64((*e)[4])<<32

	if d >= 0x0000008000000000 {
		d |= 0xffffff0000000000
	}

	*e = (*e)[5:]

	return int64(d)
}

// ReadInt48 Reads a 48 bit int as a int64 from the byte slice.
func (e *MemLittleEndianReader) ReadInt48() int64 {
	if len(*e) < 6 {
		return 0
	}

	d := uint64((*e)[0]) | uint64((*e)[1])<<8 | uint64((*e)[2])<<16 | uint64((*e)[3])<<24 | uint64((*e)[4])<<32 | uint64((*e)[5])<<40

	if d >= 0x0000800000000000 {
		d |= 0xffff000000000000
	}

	*e = (*e)[6:]

	return int64(d)
}

// ReadInt56 Reads a 56 bit int as a int64 from the byte slice.
func (e *MemLittleEndianReader) ReadInt56() int64 {
	if len(*e) < 7 {
		return 0
	}

	d := uint64((*e)[0]) | uint64((*e)[1])<<8 | uint64((*e)[2])<<16 | uint64((*e)[3])<<24 | uint64((*e)[4])<<32 | uint64((*e)[5])<<40 | uint64((*e)[6])<<48

	if d >= 0x0080000000000000 {
		d |= 0xff00000000000000
	}

	*e = (*e)[7:]

	return int64(d)
}

// ReadInt64 Reads a 64 bit int as a int64 from the byte slice.
func (e *MemLittleEndianReader) ReadInt64() int64 {
	if len(*e) < 8 {
		return 0
	}

	d := uint64((*e)[0]) | uint64((*e)[1])<<8 | uint64((*e)[2])<<16 | uint64((*e)[3])<<24 | uint64((*e)[4])<<32 | uint64((*e)[5])<<40 | uint64((*e)[6])<<48 | uint64((*e)[7])<<56
	*e = (*e)[8:]

	return int64(d)
}

// ReadUint8 Reads a 8 bit uint as a uint8 from the byte slice.
func (e *MemLittleEndianReader) ReadUint8() uint8 {
	if len(*e) < 1 {
		return 0
	}

	d := (*e)[0]
	*e = (*e)[1:]

	return d
}

// ReadUint16 Reads a 16 bit uint as a uint16 from the byte slice.
func (e *MemLittleEndianReader) ReadUint16() uint16 {
	if len(*e) < 2 {
		return 0
	}

	d := uint16((*e)[0]) | uint16((*e)[1])<<8
	*e = (*e)[2:]

	return d
}

// ReadUint24 Reads a 24 bit uint as a uint32 from the byte slice.
func (e *MemLittleEndianReader) ReadUint24() uint32 {
	if len(*e) < 3 {
		return 0
	}

	d := uint32((*e)[0]) | uint32((*e)[1])<<8 | uint32((*e)[2])<<16
	*e = (*e)[3:]

	return d
}

// ReadUint32 Reads a 32 bit uint as a uint32 from the byte slice.
func (e *MemLittleEndianReader) ReadUint32() uint32 {
	if len(*e) < 4 {
		return 0
	}

	d := uint32((*e)[0]) | uint32((*e)[1])<<8 | uint32((*e)[2])<<16 | uint32((*e)[3])<<24
	*e = (*e)[4:]

	return d
}

// ReadUint40 Reads a 40 bit uint as a uint64 from the byte slice.
func (e *MemLittleEndianReader) ReadUint40() uint64 {
	if len(*e) < 5 {
		return 0
	}

	d := uint64((*e)[0]) | uint64((*e)[1])<<8 | uint64((*e)[2])<<16 | uint64((*e)[3])<<24 | uint64((*e)[4])<<32
	*e = (*e)[5:]

	return d
}

// ReadUint48 Reads a 48 bit uint as a uint64 from the byte slice.
func (e *MemLittleEndianReader) ReadUint48() uint64 {
	if len(*e) < 6 {
		return 0
	}

	d := uint64((*e)[0]) | uint64((*e)[1])<<8 | uint64((*e)[2])<<16 | uint64((*e)[3])<<24 | uint64((*e)[4])<<32 | uint64((*e)[5])<<40
	*e = (*e)[6:]

	return d
}

// ReadUint56 Reads a 56 bit uint as a uint64 from the byte slice.
func (e *MemLittleEndianReader) ReadUint56() uint64 {
	if len(*e) < 7 {
		return 0
	}

	d := uint64((*e)[0]) | uint64((*e)[1])<<8 | uint64((*e)[2])<<16 | uint64((*e)[3])<<24 | uint64((*e)[4])<<32 | uint64((*e)[5])<<40 | uint64((*e)[6])<<48
	*e = (*e)[7:]

	return d
}

// ReadUint64 Reads a 64 bit uint as a uint64 from the byte slice.
func (e *MemLittleEndianReader) ReadUint64() uint64 {
	if len(*e) < 8 {
		return 0
	}

	d := uint64((*e)[0]) | uint64((*e)[1])<<8 | uint64((*e)[2])<<16 | uint64((*e)[3])<<24 | uint64((*e)[4])<<32 | uint64((*e)[5])<<40 | uint64((*e)[6])<<48 | uint64((*e)[7])<<56
	*e = (*e)[8:]

	return d
}

// ReadFloat32 Reads a 32 bit float as a float32 from the byte slice.
func (e *MemLittleEndianReader) ReadFloat32() float32 {
	if len(*e) < 4 {
		return 0
	}

	d := math.Float32frombits(uint32((*e)[0]) | uint32((*e)[1])<<8 | uint32((*e)[2])<<16 | uint32((*e)[3])<<24)
	*e = (*e)[4:]

	return d
}

// ReadFloat64 Reads a 64 bit float as a float64 from the byte slice.
func (e *MemLittleEndianReader) ReadFloat64() float64 {
	if len(*e) < 8 {
		return 0
	}

	d := math.Float64frombits(uint64((*e)[0]) | uint64((*e)[1])<<8 | uint64((*e)[2])<<16 | uint64((*e)[3])<<24 | uint64((*e)[4])<<32 | uint64((*e)[5])<<40 | uint64((*e)[6])<<48 | uint64((*e)[7])<<56)
	*e = (*e)[8:]

	return d
}

// ReadBytes Reads a []byte.
func (e *MemLittleEndianReader) ReadBytes(size int) []byte {
	if len(*e) < size {
		return nil
	}

	d := []byte((*e)[:size])
	*e = (*e)[size:]

	return d
}

// MemLittleEndianReader implements io.Reader.
func (e *MemLittleEndianReader) Read(p []byte) (int, error) {
	n := copy(p, *e)

	if n == 0 && len(p) != 0 {
		return 0, io.EOF
	}

	*e = (*e)[n:]

	return n, nil
}

// ReadBytesX Reads the length of the Bytes, using ReadUintX and then Reads the bytes.
func (e *MemLittleEndianReader) ReadBytesX() []byte {
	return e.ReadBytes(int(e.ReadUintX()))
}

// ReadBytes8 Reads the length of the Bytes, using ReadUint8 and then Reads the bytes.
func (e *MemLittleEndianReader) ReadBytes8() []byte {
	return e.ReadBytes(int(e.ReadUint8()))
}

// ReadBytes16 Reads the length of the Bytes, using ReadUint16 and then Reads the bytes.
func (e *MemLittleEndianReader) ReadBytes16() []byte {
	return e.ReadBytes(int(e.ReadUint16()))
}

// ReadBytes24 Reads the length of the Bytes, using ReadUint24 and then Reads the bytes.
func (e *MemLittleEndianReader) ReadBytes24() []byte {
	return e.ReadBytes(int(e.ReadUint24()))
}

// ReadBytes32 Reads the length of the Bytes, using ReadUint32 and then Reads the bytes.
func (e *MemLittleEndianReader) ReadBytes32() []byte {
	return e.ReadBytes(int(e.ReadUint32()))
}

// ReadBytes40 Reads the length of the Bytes, using ReadUint40 and then Reads the bytes.
func (e *MemLittleEndianReader) ReadBytes40() []byte {
	return e.ReadBytes(int(e.ReadUint40()))
}

// ReadBytes48 Reads the length of the Bytes, using ReadUint48 and then Reads the bytes.
func (e *MemLittleEndianReader) ReadBytes48() []byte {
	return e.ReadBytes(int(e.ReadUint48()))
}

// ReadBytes56 Reads the length of the Bytes, using ReadUint56 and then Reads the bytes.
func (e *MemLittleEndianReader) ReadBytes56() []byte {
	return e.ReadBytes(int(e.ReadUint56()))
}

// ReadBytes64 Reads the length of the Bytes, using ReadUint64 and then Reads the bytes.
func (e *MemLittleEndianReader) ReadBytes64() []byte {
	return e.ReadBytes(int(e.ReadUint64()))
}

// ReadString Reads a string.
func (e *MemLittleEndianReader) ReadString(size int) string {
	if len(*e) < size {
		return ""
	}

	d := string((*e)[:size])
	*e = (*e)[size:]

	return d
}

// ReadStringX Reads the length of the String, using ReadUintX and then Reads the bytes.
func (e *MemLittleEndianReader) ReadStringX() string {
	return e.ReadString(int(e.ReadUintX()))
}

// ReadString8 Reads the length of the String, using ReadUint8 and then Reads the bytes.
func (e *MemLittleEndianReader) ReadString8() string {
	return e.ReadString(int(e.ReadUint8()))
}

// ReadString16 Reads the length of the String, using ReadUint16 and then Reads the bytes.
func (e *MemLittleEndianReader) ReadString16() string {
	return e.ReadString(int(e.ReadUint16()))
}

// ReadString24 Reads the length of the String, using ReadUint24 and then Reads the bytes.
func (e *MemLittleEndianReader) ReadString24() string {
	return e.ReadString(int(e.ReadUint24()))
}

// ReadString32 Reads the length of the String, using ReadUint32 and then Reads the bytes.
func (e *MemLittleEndianReader) ReadString32() string {
	return e.ReadString(int(e.ReadUint32()))
}

// ReadString40 Reads the length of the String, using ReadUint40 and then Reads the bytes.
func (e *MemLittleEndianReader) ReadString40() string {
	return e.ReadString(int(e.ReadUint40()))
}

// ReadString48 Reads the length of the String, using ReadUint48 and then Reads the bytes.
func (e *MemLittleEndianReader) ReadString48() string {
	return e.ReadString(int(e.ReadUint48()))
}

// ReadString56 Reads the length of the String, using ReadUint56 and then Reads the bytes.
func (e *MemLittleEndianReader) ReadString56() string {
	return e.ReadString(int(e.ReadUint56()))
}

// ReadString64 Reads the length of the String, using ReadUint64 and then Reads the bytes.
func (e *MemLittleEndianReader) ReadString64() string {
	return e.ReadString(int(e.ReadUint64()))
}

// ReadString0 Reads the bytes of the string until a 0 byte is read.
func (e *MemLittleEndianReader) ReadString0() string {
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
