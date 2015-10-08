# byteio
--
    import "github.com/MJKWoolnough/byteio"

Package byteio helps with writing number types in both big and little endian formats

## Usage

#### type BigEndianReader

```go
type BigEndianReader struct {
	io.Reader
}
```

BigEndianReader is a wrapper around a io.Reader that can be used to read types
in the big endian format

#### func (BigEndianReader) ReadFloat32

```go
func (b BigEndianReader) ReadFloat32() (float32, int, error)
```
ReadFloat32 will read four bytes from the reader, in big endian format, and
return it as a float32

#### func (BigEndianReader) ReadFloat64

```go
func (b BigEndianReader) ReadFloat64() (float64, int, error)
```
ReadFloat64 will read eight bytes from the reader, in big endian format, and
return it as a float64

#### func (BigEndianReader) ReadInt16

```go
func (b BigEndianReader) ReadInt16() (int16, int, error)
```
ReadInt16 will read two bytes from the reader, in big endian format, and return
it as an int16

#### func (BigEndianReader) ReadInt32

```go
func (b BigEndianReader) ReadInt32() (int32, int, error)
```
ReadInt32 will read four bytes from the reader, in big endian format, and return
it as an int32

#### func (BigEndianReader) ReadInt64

```go
func (b BigEndianReader) ReadInt64() (int64, int, error)
```
ReadInt64 will read eight bytes from the reader, in big endian format, and
return it as an int64

#### func (BigEndianReader) ReadInt8

```go
func (b BigEndianReader) ReadInt8() (int8, int, error)
```
ReadInt8 will read a single byte from the reader and return it as an int8

#### func (BigEndianReader) ReadUint16

```go
func (b BigEndianReader) ReadUint16() (uint16, int, error)
```
ReadUint16 will read two bytes from the reader, in big endian format, and return
it as a uint16

#### func (BigEndianReader) ReadUint32

```go
func (b BigEndianReader) ReadUint32() (uint32, int, error)
```
ReadUint32 will read four bytes from the reader, in big endian format, and
return it as a uint32

#### func (BigEndianReader) ReadUint64

```go
func (b BigEndianReader) ReadUint64() (uint64, int, error)
```
ReadUint64 will read eight bytes from the reader, in big endian format, and
return it as a uint64

#### func (BigEndianReader) ReadUint8

```go
func (b BigEndianReader) ReadUint8() (uint8, int, error)
```
ReadUint8 will read a single byte from the reader and return it as a uint8

#### type BigEndianWriter

```go
type BigEndianWriter struct {
	io.Writer
}
```

BigEndianWriter is a wrapper around a io.Writer that can be used to write types
in the big endian format

#### func (*BigEndianWriter) WriteFloat32

```go
func (b *BigEndianWriter) WriteFloat32(d float32) (int, error)
```
WriteFloat32 will write the given float32 to the writer, in big endian format

#### func (*BigEndianWriter) WriteFloat64

```go
func (b *BigEndianWriter) WriteFloat64(d float64) (int, error)
```
WriteFloat64 will write the given float64 to the writer, in big endian format

#### func (*BigEndianWriter) WriteInt16

```go
func (b *BigEndianWriter) WriteInt16(d int16) (int, error)
```
WriteInt16 will write the given int16 to the writer, in big endian format

#### func (*BigEndianWriter) WriteInt32

```go
func (b *BigEndianWriter) WriteInt32(d int32) (int, error)
```
WriteInt32 will write the given int32 to the writer, in big endian format

#### func (*BigEndianWriter) WriteInt64

```go
func (b *BigEndianWriter) WriteInt64(d int64) (int, error)
```
WriteInt64 will write the given int64 to the writer, in big endian format

#### func (*BigEndianWriter) WriteInt8

```go
func (b *BigEndianWriter) WriteInt8(d int8) (int, error)
```
WriteInt8 will write the given int8 to the writer

#### func (*BigEndianWriter) WriteUint16

```go
func (b *BigEndianWriter) WriteUint16(d uint16) (int, error)
```
WriteUint16 will write the given uint16 to the writer, in big endian format

#### func (*BigEndianWriter) WriteUint32

```go
func (b *BigEndianWriter) WriteUint32(d uint32) (int, error)
```
WriteUint32 will write the given uint32 to the writer, in big endian format

#### func (*BigEndianWriter) WriteUint64

```go
func (b *BigEndianWriter) WriteUint64(d uint64) (int, error)
```
WriteUint64 will write the given uint64 to the writer, in big endian format

#### func (*BigEndianWriter) WriteUint8

```go
func (b *BigEndianWriter) WriteUint8(d uint8) (int, error)
```
WriteUint8 will write the given uint8 to the writer

#### type EndianReader

```go
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
```

EndianReader is an interface that reads various types with a particular
endianness

#### type EndianWriter

```go
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
```

EndianWriter is an interface that writes various types with a particular
endianness

#### type LittleEndianReader

```go
type LittleEndianReader struct {
	io.Reader
}
```

LittleEndianReader is a wrapper around a io.Reader that can be used to read
types in the little endian format

#### func (LittleEndianReader) ReadFloat32

```go
func (l LittleEndianReader) ReadFloat32() (float32, int, error)
```
ReadFloat32 will read four bytes from the reader, in little endian format, and
return it as a float32

#### func (LittleEndianReader) ReadFloat64

```go
func (l LittleEndianReader) ReadFloat64() (float64, int, error)
```
ReadFloat64 will read eight bytes from the reader, in little endian format, and
return it as a float64

#### func (LittleEndianReader) ReadInt16

```go
func (l LittleEndianReader) ReadInt16() (int16, int, error)
```
ReadInt16 will read two bytes from the reader, in little endian format, and
return it as an int16

#### func (LittleEndianReader) ReadInt32

```go
func (l LittleEndianReader) ReadInt32() (int32, int, error)
```
ReadInt32 will read four bytes from the reader, in little endian format, and
return it as an int32

#### func (LittleEndianReader) ReadInt64

```go
func (l LittleEndianReader) ReadInt64() (int64, int, error)
```
ReadInt64 will read eight bytes from the reader, in little endian format, and
return it as an int64

#### func (LittleEndianReader) ReadInt8

```go
func (l LittleEndianReader) ReadInt8() (int8, int, error)
```
ReadInt8 will read a single byte from the reader, in little endian format, and
return it as an int8

#### func (LittleEndianReader) ReadUint16

```go
func (l LittleEndianReader) ReadUint16() (uint16, int, error)
```
ReadUint16 will read two bytes from the reader, in little endian format, and
return it as a uint16

#### func (LittleEndianReader) ReadUint32

```go
func (l LittleEndianReader) ReadUint32() (uint32, int, error)
```
ReadUint32 will read four bytes from the reader, in little endian format, and
return it as a uint32

#### func (LittleEndianReader) ReadUint64

```go
func (l LittleEndianReader) ReadUint64() (uint64, int, error)
```
ReadUint64 will read eight bytes from the reader, in little endian format, and
return it as a uint64

#### func (LittleEndianReader) ReadUint8

```go
func (l LittleEndianReader) ReadUint8() (uint8, int, error)
```
ReadUint8 will read a single byte from the reader, in little endian format, and
return it as a uint8

#### type LittleEndianWriter

```go
type LittleEndianWriter struct {
	io.Writer
}
```

LittleEndianWriter is a wrapper around a io.Writer that can be used to write
types in the little endian format

#### func (*LittleEndianWriter) WriteFloat32

```go
func (l *LittleEndianWriter) WriteFloat32(d float32) (int, error)
```
WriteFloat32 will write the given float32 to the writer, in little endian format

#### func (*LittleEndianWriter) WriteFloat64

```go
func (l *LittleEndianWriter) WriteFloat64(d float64) (int, error)
```
WriteFloat64 will write the given float64 to the writer, in little endian format

#### func (*LittleEndianWriter) WriteInt16

```go
func (l *LittleEndianWriter) WriteInt16(d int16) (int, error)
```
WriteInt16 will write the given int16 to the writer, in little endian format

#### func (*LittleEndianWriter) WriteInt32

```go
func (l *LittleEndianWriter) WriteInt32(d int32) (int, error)
```
WriteInt32 will write the given int32 to the writer, in little endian format

#### func (*LittleEndianWriter) WriteInt64

```go
func (l *LittleEndianWriter) WriteInt64(d int64) (int, error)
```
WriteInt64 will write the given int64 to the writer, in little endian format

#### func (*LittleEndianWriter) WriteInt8

```go
func (l *LittleEndianWriter) WriteInt8(d int8) (int, error)
```
WriteInt8 will write the given int8 to the writer

#### func (*LittleEndianWriter) WriteUint16

```go
func (l *LittleEndianWriter) WriteUint16(d uint16) (int, error)
```
WriteUint16 will write the given uint16 to the writer, in little endian format

#### func (*LittleEndianWriter) WriteUint32

```go
func (l *LittleEndianWriter) WriteUint32(d uint32) (int, error)
```
WriteUint32 will write the given uint32 to the writer, in little endian format

#### func (*LittleEndianWriter) WriteUint64

```go
func (l *LittleEndianWriter) WriteUint64(d uint64) (int, error)
```
WriteUint64 will write the given uint64 to the writer, in little endian format

#### func (*LittleEndianWriter) WriteUint8

```go
func (l *LittleEndianWriter) WriteUint8(d uint8) (int, error)
```
WriteUint8 will write the given uint8 to the writer

#### type StickyReader

```go
type StickyReader struct {
	Reader EndianReader
	Err    error
	Count  int64
}
```

StickyReader will wrap an EndianReader and record all bytes read and errors
received. Byte counts and errors will not be returned from any method (except
Read so it still counts as an io.Reader), but can be retrieved from this type.
All methods will be a no-op after an error has been returned, unless that error
is cleared on the type

#### func (*StickyReader) Read

```go
func (s *StickyReader) Read(b []byte) (int, error)
```
Read will do a simple byte read from the underlying io.Reader.

#### func (*StickyReader) ReadFloat32

```go
func (s *StickyReader) ReadFloat32() float32
```
ReadFloat32 will read a float32 from the underlying reader

#### func (*StickyReader) ReadFloat64

```go
func (s *StickyReader) ReadFloat64() float64
```
ReadFloat64 will read a float64 from the underlying reader

#### func (*StickyReader) ReadInt16

```go
func (s *StickyReader) ReadInt16() int16
```
ReadInt16 will read a int16 from the underlying reader

#### func (*StickyReader) ReadInt32

```go
func (s *StickyReader) ReadInt32() int32
```
ReadInt32 will read a int32 from the underlying reader

#### func (*StickyReader) ReadInt64

```go
func (s *StickyReader) ReadInt64() int64
```
ReadInt64 will read a int64 from the underlying reader

#### func (*StickyReader) ReadInt8

```go
func (s *StickyReader) ReadInt8() int8
```
ReadInt8 will read a int8 from the underlying reader

#### func (*StickyReader) ReadUint16

```go
func (s *StickyReader) ReadUint16() uint16
```
ReadUint16 will read a uint16 from the underlying reader

#### func (*StickyReader) ReadUint32

```go
func (s *StickyReader) ReadUint32() uint32
```
ReadUint32 will read a uint32 from the underlying reader

#### func (*StickyReader) ReadUint64

```go
func (s *StickyReader) ReadUint64() uint64
```
ReadUint64 will read a uint64 from the underlying reader

#### func (*StickyReader) ReadUint8

```go
func (s *StickyReader) ReadUint8() uint8
```
ReadUint8 will read a uint8 from the underlying reader

#### type StickyWriter

```go
type StickyWriter struct {
	Writer EndianWriter
	Err    error
	Count  int64
}
```

StickyWriter will wrap an EndianWriter and record all bytes written and errors
received. Byte counts and errors will not be returned from any method (except
Write, so it still counts as an io.Writer), but can be retrieved from this type.
All methods will be a no-op after an error has been returned, unless that error
is cleared on the type

#### func (*StickyWriter) Write

```go
func (s *StickyWriter) Write(p []byte) (int, error)
```
Write will do a simple byte write to the underlying io.Writer.

#### func (*StickyWriter) WriteFloat32

```go
func (s *StickyWriter) WriteFloat32(i float32)
```
WriteFloat32 will write a float32 to the underlying writer

#### func (*StickyWriter) WriteFloat64

```go
func (s *StickyWriter) WriteFloat64(i float64)
```
WriteFloat64 will write a float64 to the underlying writer

#### func (*StickyWriter) WriteInt16

```go
func (s *StickyWriter) WriteInt16(i int16)
```
WriteInt16 will write a int16 to the underlying writer

#### func (*StickyWriter) WriteInt32

```go
func (s *StickyWriter) WriteInt32(i int32)
```
WriteInt32 will write a int32 to the underlying writer

#### func (*StickyWriter) WriteInt64

```go
func (s *StickyWriter) WriteInt64(i int64)
```
WriteInt64 will write a int64 to the underlying writer

#### func (*StickyWriter) WriteInt8

```go
func (s *StickyWriter) WriteInt8(i int8)
```
WriteInt8 will write a int8 to the underlying writer

#### func (*StickyWriter) WriteUint16

```go
func (s *StickyWriter) WriteUint16(i uint16)
```
WriteUint16 will write a uint16 to the underlying writer

#### func (*StickyWriter) WriteUint32

```go
func (s *StickyWriter) WriteUint32(i uint32)
```
WriteUint32 will write a uint32 to the underlying writer

#### func (*StickyWriter) WriteUint64

```go
func (s *StickyWriter) WriteUint64(i uint64)
```
WriteUint64 will write a uint64 to the underlying writer

#### func (*StickyWriter) WriteUint8

```go
func (s *StickyWriter) WriteUint8(i uint8)
```
WriteUint8 will write a uint8 to the underlying writer
