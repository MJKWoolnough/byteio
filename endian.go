// Package byteio helps with writing number types in both big and little endian formats
package byteio

import (
	"io"
	"sync"
)

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

var (
	_ EndianReader = BigEndianReader{}
	_ EndianReader = LittleEndianReader{}
	_ EndianWriter = BigEndianWriter{}
	_ EndianWriter = LittleEndianWriter{}
)

// pool for read/write buffers
var pool = sync.Pool{
	New: func() interface{} {
		return new([8]byte)
	},
}
