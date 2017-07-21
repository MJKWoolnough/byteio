// Package byteio helps with writing number types in both big and little endian formats
package byteio

import "io"

// EndianReader is an interface that reads various types with a particular
// endianness
type EndianReader interface {
	io.Reader
	ReadUint8() (uint8, int, error)
	ReadInt8() (int8, int, error)
	ReadUint16() (uint16, int, error)
	ReadInt16() (int16, int, error)
	ReadUint32() (uint32, int, error)
	ReadInt32() (int32, int, error)
	ReadUint64() (uint64, int, error)
	ReadInt64() (int64, int, error)
	ReadFloat32() (float32, int, error)
	ReadFloat64() (float64, int, error)
}

// EndianWriter is an interface that writes various types with a particular
// endianness
type EndianWriter interface {
	io.Writer
	WriteUint8(uint8) (int, error)
	WriteInt8(int8) (int, error)
	WriteUint16(uint16) (int, error)
	WriteInt16(int16) (int, error)
	WriteUint32(uint32) (int, error)
	WriteInt32(int32) (int, error)
	WriteUint64(uint64) (int, error)
	WriteInt64(int64) (int, error)
	WriteFloat32(float32) (int, error)
	WriteFloat64(float64) (int, error)
}

// StickyEndianReader is an interface that reads various types with a
// particular endianness and stores the Read return values
type StickyEndianReader interface {
	io.Reader
	ReadUint8() uint8
	ReadInt8() int8
	ReadUint16() uint16
	ReadInt16() int16
	ReadUint32() uint32
	ReadInt32() int32
	ReadUint64() uint64
	ReadInt64() int64
	ReadFloat32() float32
	ReadFloat64() float64
}

// StickyEndianWriter is an interface that writes various types with a
// particular endianness and stores the Write return values
type StickyEndianWriter interface {
	io.Writer
	WriteUint8(uint8)
	WriteInt8(int8)
	WriteUint16(uint16)
	WriteInt16(int16)
	WriteUint32(uint32)
	WriteInt32(int32)
	WriteUint64(uint64)
	WriteInt64(int64)
	WriteFloat32(float32)
	WriteFloat64(float64)
}

var (
	_ EndianReader       = &BigEndianReader{}
	_ EndianReader       = &LittleEndianReader{}
	_ EndianWriter       = &BigEndianWriter{}
	_ EndianWriter       = &LittleEndianWriter{}
	_ StickyEndianReader = &StickyBigEndianReader{}
	_ StickyEndianReader = &StickyLittleEndianReader{}
	_ StickyEndianWriter = &StickyBigEndianWriter{}
	_ StickyEndianWriter = &StickyLittleEndianWriter{}
)
