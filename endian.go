package byteio

import "io"

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
}

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
}
