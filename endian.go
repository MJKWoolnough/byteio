// Package byteio helps with writing binary data in both big and little endian.
package byteio // import "vimagination.zapto.org/byteio"

import "io"

// EndianReader is an interface that reads various types with a particular
// endianness.
type EndianReader interface {
	io.Reader
	io.ByteReader
	ReadBool() (bool, int, error)
	ReadUint8() (uint8, int, error)
	ReadInt8() (int8, int, error)
	ReadUint16() (uint16, int, error)
	ReadInt16() (int16, int, error)
	ReadUint24() (uint32, int, error)
	ReadInt24() (int32, int, error)
	ReadUint32() (uint32, int, error)
	ReadInt32() (int32, int, error)
	ReadUint40() (uint64, int, error)
	ReadInt40() (int64, int, error)
	ReadUint48() (uint64, int, error)
	ReadInt48() (int64, int, error)
	ReadUint56() (uint64, int, error)
	ReadInt56() (int64, int, error)
	ReadUint64() (uint64, int, error)
	ReadInt64() (int64, int, error)
	ReadFloat32() (float32, int, error)
	ReadFloat64() (float64, int, error)
	ReadUintX() (uint64, int, error)
	ReadIntX() (int64, int, error)
	ReadBytes(int) ([]byte, int, error)
	ReadBytesX() ([]byte, int, error)
	ReadBytes8() ([]byte, int, error)
	ReadBytes16() ([]byte, int, error)
	ReadBytes24() ([]byte, int, error)
	ReadBytes32() ([]byte, int, error)
	ReadBytes40() ([]byte, int, error)
	ReadBytes48() ([]byte, int, error)
	ReadBytes56() ([]byte, int, error)
	ReadBytes64() ([]byte, int, error)
	ReadString(int) (string, int, error)
	ReadStringX() (string, int, error)
	ReadString0() (string, int, error)
	ReadString8() (string, int, error)
	ReadString16() (string, int, error)
	ReadString24() (string, int, error)
	ReadString32() (string, int, error)
	ReadString40() (string, int, error)
	ReadString48() (string, int, error)
	ReadString56() (string, int, error)
	ReadString64() (string, int, error)
}

// EndianWriter is an interface that writes various types with a particular
// endianness.
type EndianWriter interface {
	io.Writer
	io.ByteWriter
	io.StringWriter
	WriteBool(bool) (int, error)
	WriteUint8(uint8) (int, error)
	WriteInt8(int8) (int, error)
	WriteUint16(uint16) (int, error)
	WriteInt16(int16) (int, error)
	WriteUint24(uint32) (int, error)
	WriteInt24(int32) (int, error)
	WriteUint32(uint32) (int, error)
	WriteInt32(int32) (int, error)
	WriteUint40(uint64) (int, error)
	WriteInt40(int64) (int, error)
	WriteUint48(uint64) (int, error)
	WriteInt48(int64) (int, error)
	WriteUint56(uint64) (int, error)
	WriteInt56(int64) (int, error)
	WriteUint64(uint64) (int, error)
	WriteInt64(int64) (int, error)
	WriteFloat32(float32) (int, error)
	WriteFloat64(float64) (int, error)
	WriteUintX(uint64) (int, error)
	WriteIntX(int64) (int, error)
	WriteBytes([]byte) (int, error)
	WriteBytesX([]byte) (int, error)
	WriteBytes8([]byte) (int, error)
	WriteBytes16([]byte) (int, error)
	WriteBytes24([]byte) (int, error)
	WriteBytes32([]byte) (int, error)
	WriteBytes40([]byte) (int, error)
	WriteBytes48([]byte) (int, error)
	WriteBytes56([]byte) (int, error)
	WriteBytes64([]byte) (int, error)
	WriteStringX(string) (int, error)
	WriteString0(string) (int, error)
	WriteString8(string) (int, error)
	WriteString16(string) (int, error)
	WriteString24(string) (int, error)
	WriteString32(string) (int, error)
	WriteString40(string) (int, error)
	WriteString48(string) (int, error)
	WriteString56(string) (int, error)
	WriteString64(string) (int, error)
}

// StickyEndianReader is an interface that reads various types with a
// particular endianness and stores the Read return values.
type StickyEndianReader interface {
	io.Reader
	io.ByteReader
	GetError() error
	GetCount() int64
	ReadBool() bool
	ReadUint8() uint8
	ReadInt8() int8
	ReadUint16() uint16
	ReadInt16() int16
	ReadUint24() uint32
	ReadInt24() int32
	ReadUint32() uint32
	ReadInt32() int32
	ReadUint40() uint64
	ReadInt40() int64
	ReadUint48() uint64
	ReadInt48() int64
	ReadUint56() uint64
	ReadInt56() int64
	ReadUint64() uint64
	ReadInt64() int64
	ReadFloat32() float32
	ReadFloat64() float64
	ReadUintX() uint64
	ReadIntX() int64
	ReadBytes(int) []byte
	ReadBytesX() []byte
	ReadBytes8() []byte
	ReadBytes16() []byte
	ReadBytes24() []byte
	ReadBytes32() []byte
	ReadBytes40() []byte
	ReadBytes48() []byte
	ReadBytes56() []byte
	ReadBytes64() []byte
	ReadString(int) string
	ReadStringX() string
	ReadString0() string
	ReadString8() string
	ReadString16() string
	ReadString24() string
	ReadString32() string
	ReadString40() string
	ReadString48() string
	ReadString56() string
	ReadString64() string
}

// StickyEndianWriter is an interface that writes various types with a
// particular endianness and stores the Write return values.
type StickyEndianWriter interface {
	io.Writer
	io.ByteWriter
	io.StringWriter
	GetError() error
	GetCount() int64
	WriteBool(bool)
	WriteUint8(uint8)
	WriteInt8(int8)
	WriteUint16(uint16)
	WriteInt16(int16)
	WriteUint24(uint32)
	WriteInt24(int32)
	WriteUint32(uint32)
	WriteInt32(int32)
	WriteUint40(uint64)
	WriteInt40(int64)
	WriteUint48(uint64)
	WriteInt48(int64)
	WriteUint56(uint64)
	WriteInt56(int64)
	WriteUint64(uint64)
	WriteInt64(int64)
	WriteFloat32(float32)
	WriteFloat64(float64)
	WriteUintX(uint64)
	WriteIntX(int64)
	WriteBytes([]byte)
	WriteBytesX([]byte)
	WriteBytes8([]byte)
	WriteBytes16([]byte)
	WriteBytes24([]byte)
	WriteBytes32([]byte)
	WriteBytes40([]byte)
	WriteBytes48([]byte)
	WriteBytes56([]byte)
	WriteBytes64([]byte)
	WriteString0(string)
	WriteStringX(string)
	WriteString8(string)
	WriteString16(string)
	WriteString24(string)
	WriteString32(string)
	WriteString40(string)
	WriteString48(string)
	WriteString56(string)
	WriteString64(string)
}

var (
	_ EndianReader       = (*BigEndianReader)(nil)
	_ EndianReader       = (*LittleEndianReader)(nil)
	_ EndianWriter       = (*BigEndianWriter)(nil)
	_ EndianWriter       = (*LittleEndianWriter)(nil)
	_ StickyEndianReader = (*StickyReader)(nil)
	_ StickyEndianReader = (*StickyBigEndianReader)(nil)
	_ StickyEndianReader = (*StickyLittleEndianReader)(nil)
	_ StickyEndianWriter = (*StickyWriter)(nil)
	_ StickyEndianWriter = (*StickyBigEndianWriter)(nil)
	_ StickyEndianWriter = (*StickyLittleEndianWriter)(nil)
)
