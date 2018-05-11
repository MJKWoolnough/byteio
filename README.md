# byteio
--
    import "github.com/MJKWoolnough/byteio"

Package byteio helps with writing number types in both big and little endian
### formats

## Usage

#### type BigEndianReader

```go
type BigEndianReader struct {
	io.Reader
}
```

BigEndianReader wraps a io.Reader to provide methods to make it easier to Read
fundamental types

#### func (*BigEndianReader) ReadByte

```go
func (e *BigEndianReader) ReadByte() (byte, error)
```
ReadByte implements the io.ByteReader interface

#### func (*BigEndianReader) ReadFloat32

```go
func (e *BigEndianReader) ReadFloat32() (float32, int, error)
```
ReadFloat32 Reads a float32 using the underlying io.Reader

#### func (*BigEndianReader) ReadFloat64

```go
func (e *BigEndianReader) ReadFloat64() (float64, int, error)
```
ReadFloat64 Reads a float64 using the underlying io.Reader

#### func (*BigEndianReader) ReadInt16

```go
func (e *BigEndianReader) ReadInt16() (int16, int, error)
```
ReadInt16 Reads a int16 using the underlying io.Reader

#### func (*BigEndianReader) ReadInt32

```go
func (e *BigEndianReader) ReadInt32() (int32, int, error)
```
ReadInt32 Reads a int32 using the underlying io.Reader

#### func (*BigEndianReader) ReadInt64

```go
func (e *BigEndianReader) ReadInt64() (int64, int, error)
```
ReadInt64 Reads a int64 using the underlying io.Reader

#### func (*BigEndianReader) ReadInt8

```go
func (e *BigEndianReader) ReadInt8() (int8, int, error)
```
ReadInt8 Reads a int8 using the underlying io.Reader

#### func (*BigEndianReader) ReadUint16

```go
func (e *BigEndianReader) ReadUint16() (uint16, int, error)
```
ReadUint16 Reads a uint16 using the underlying io.Reader

#### func (*BigEndianReader) ReadUint32

```go
func (e *BigEndianReader) ReadUint32() (uint32, int, error)
```
ReadUint32 Reads a uint32 using the underlying io.Reader

#### func (*BigEndianReader) ReadUint64

```go
func (e *BigEndianReader) ReadUint64() (uint64, int, error)
```
ReadUint64 Reads a uint64 using the underlying io.Reader

#### func (*BigEndianReader) ReadUint8

```go
func (e *BigEndianReader) ReadUint8() (uint8, int, error)
```
ReadUint8 Reads a uint8 using the underlying io.Reader

#### type BigEndianWriter

```go
type BigEndianWriter struct {
	io.Writer
}
```

BigEndianWriter wraps a io.Writer to provide methods to make it easier to Write
fundamental types

#### func (*BigEndianWriter) WriteByte

```go
func (e *BigEndianWriter) WriteByte(c byte) error
```
WriteByte implements the io.ByteWriter interface

#### func (*BigEndianWriter) WriteFloat32

```go
func (e *BigEndianWriter) WriteFloat32(d float32) (int, error)
```
WriteFloat32 Writes a float32 using the underlying io.Writer

#### func (*BigEndianWriter) WriteFloat64

```go
func (e *BigEndianWriter) WriteFloat64(d float64) (int, error)
```
WriteFloat64 Writes a float64 using the underlying io.Writer

#### func (*BigEndianWriter) WriteInt16

```go
func (e *BigEndianWriter) WriteInt16(d int16) (int, error)
```
WriteInt16 Writes a int16 using the underlying io.Writer

#### func (*BigEndianWriter) WriteInt32

```go
func (e *BigEndianWriter) WriteInt32(d int32) (int, error)
```
WriteInt32 Writes a int32 using the underlying io.Writer

#### func (*BigEndianWriter) WriteInt64

```go
func (e *BigEndianWriter) WriteInt64(d int64) (int, error)
```
WriteInt64 Writes a int64 using the underlying io.Writer

#### func (*BigEndianWriter) WriteInt8

```go
func (e *BigEndianWriter) WriteInt8(d int8) (int, error)
```
WriteInt8 Writes a int8 using the underlying io.Writer

#### func (*BigEndianWriter) WriteUint16

```go
func (e *BigEndianWriter) WriteUint16(d uint16) (int, error)
```
WriteUint16 Writes a uint16 using the underlying io.Writer

#### func (*BigEndianWriter) WriteUint32

```go
func (e *BigEndianWriter) WriteUint32(d uint32) (int, error)
```
WriteUint32 Writes a uint32 using the underlying io.Writer

#### func (*BigEndianWriter) WriteUint64

```go
func (e *BigEndianWriter) WriteUint64(d uint64) (int, error)
```
WriteUint64 Writes a uint64 using the underlying io.Writer

#### func (*BigEndianWriter) WriteUint8

```go
func (e *BigEndianWriter) WriteUint8(d uint8) (int, error)
```
WriteUint8 Writes a uint8 using the underlying io.Writer

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

LittleEndianReader wraps a io.Reader to provide methods to make it easier to
Read fundamental types

#### func (*LittleEndianReader) ReadByte

```go
func (e *LittleEndianReader) ReadByte() (byte, error)
```
ReadByte implements the io.ByteReader interface

#### func (*LittleEndianReader) ReadFloat32

```go
func (e *LittleEndianReader) ReadFloat32() (float32, int, error)
```
ReadFloat32 Reads a float32 using the underlying io.Reader

#### func (*LittleEndianReader) ReadFloat64

```go
func (e *LittleEndianReader) ReadFloat64() (float64, int, error)
```
ReadFloat64 Reads a float64 using the underlying io.Reader

#### func (*LittleEndianReader) ReadInt16

```go
func (e *LittleEndianReader) ReadInt16() (int16, int, error)
```
ReadInt16 Reads a int16 using the underlying io.Reader

#### func (*LittleEndianReader) ReadInt32

```go
func (e *LittleEndianReader) ReadInt32() (int32, int, error)
```
ReadInt32 Reads a int32 using the underlying io.Reader

#### func (*LittleEndianReader) ReadInt64

```go
func (e *LittleEndianReader) ReadInt64() (int64, int, error)
```
ReadInt64 Reads a int64 using the underlying io.Reader

#### func (*LittleEndianReader) ReadInt8

```go
func (e *LittleEndianReader) ReadInt8() (int8, int, error)
```
ReadInt8 Reads a int8 using the underlying io.Reader

#### func (*LittleEndianReader) ReadUint16

```go
func (e *LittleEndianReader) ReadUint16() (uint16, int, error)
```
ReadUint16 Reads a uint16 using the underlying io.Reader

#### func (*LittleEndianReader) ReadUint32

```go
func (e *LittleEndianReader) ReadUint32() (uint32, int, error)
```
ReadUint32 Reads a uint32 using the underlying io.Reader

#### func (*LittleEndianReader) ReadUint64

```go
func (e *LittleEndianReader) ReadUint64() (uint64, int, error)
```
ReadUint64 Reads a uint64 using the underlying io.Reader

#### func (*LittleEndianReader) ReadUint8

```go
func (e *LittleEndianReader) ReadUint8() (uint8, int, error)
```
ReadUint8 Reads a uint8 using the underlying io.Reader

#### type LittleEndianWriter

```go
type LittleEndianWriter struct {
	io.Writer
}
```

LittleEndianWriter wraps a io.Writer to provide methods to make it easier to
Write fundamental types

#### func (*LittleEndianWriter) WriteByte

```go
func (e *LittleEndianWriter) WriteByte(c byte) error
```
WriteByte implements the io.ByteWriter interface

#### func (*LittleEndianWriter) WriteFloat32

```go
func (e *LittleEndianWriter) WriteFloat32(d float32) (int, error)
```
WriteFloat32 Writes a float32 using the underlying io.Writer

#### func (*LittleEndianWriter) WriteFloat64

```go
func (e *LittleEndianWriter) WriteFloat64(d float64) (int, error)
```
WriteFloat64 Writes a float64 using the underlying io.Writer

#### func (*LittleEndianWriter) WriteInt16

```go
func (e *LittleEndianWriter) WriteInt16(d int16) (int, error)
```
WriteInt16 Writes a int16 using the underlying io.Writer

#### func (*LittleEndianWriter) WriteInt32

```go
func (e *LittleEndianWriter) WriteInt32(d int32) (int, error)
```
WriteInt32 Writes a int32 using the underlying io.Writer

#### func (*LittleEndianWriter) WriteInt64

```go
func (e *LittleEndianWriter) WriteInt64(d int64) (int, error)
```
WriteInt64 Writes a int64 using the underlying io.Writer

#### func (*LittleEndianWriter) WriteInt8

```go
func (e *LittleEndianWriter) WriteInt8(d int8) (int, error)
```
WriteInt8 Writes a int8 using the underlying io.Writer

#### func (*LittleEndianWriter) WriteUint16

```go
func (e *LittleEndianWriter) WriteUint16(d uint16) (int, error)
```
WriteUint16 Writes a uint16 using the underlying io.Writer

#### func (*LittleEndianWriter) WriteUint32

```go
func (e *LittleEndianWriter) WriteUint32(d uint32) (int, error)
```
WriteUint32 Writes a uint32 using the underlying io.Writer

#### func (*LittleEndianWriter) WriteUint64

```go
func (e *LittleEndianWriter) WriteUint64(d uint64) (int, error)
```
WriteUint64 Writes a uint64 using the underlying io.Writer

#### func (*LittleEndianWriter) WriteUint8

```go
func (e *LittleEndianWriter) WriteUint8(d uint8) (int, error)
```
WriteUint8 Writes a uint8 using the underlying io.Writer

#### type StickyBigEndianReader

```go
type StickyBigEndianReader struct {
	io.Reader

	Err   error
	Count int64
}
```

StickyBigEndianReader wraps a io.Reader to provide methods to make it easier to
Read fundamental types

#### func (*StickyBigEndianReader) Read

```go
func (e *StickyBigEndianReader) Read(p []byte) (int, error)
```
Read implements the io.Reader interface

#### func (*StickyBigEndianReader) ReadByte

```go
func (e *StickyBigEndianReader) ReadByte() (byte, error)
```
ReadByte implements the io.ByteReader interface

#### func (*StickyBigEndianReader) ReadFloat32

```go
func (e *StickyBigEndianReader) ReadFloat32() float32
```
ReadFloat32 Reads a float32 using the underlying io.Reader Any errors and the
running byte read count are stored instead or returned

#### func (*StickyBigEndianReader) ReadFloat64

```go
func (e *StickyBigEndianReader) ReadFloat64() float64
```
ReadFloat64 Reads a float64 using the underlying io.Reader Any errors and the
running byte read count are stored instead or returned

#### func (*StickyBigEndianReader) ReadInt16

```go
func (e *StickyBigEndianReader) ReadInt16() int16
```
ReadInt16 Reads a int16 using the underlying io.Reader Any errors and the
running byte read count are stored instead or returned

#### func (*StickyBigEndianReader) ReadInt32

```go
func (e *StickyBigEndianReader) ReadInt32() int32
```
ReadInt32 Reads a int32 using the underlying io.Reader Any errors and the
running byte read count are stored instead or returned

#### func (*StickyBigEndianReader) ReadInt64

```go
func (e *StickyBigEndianReader) ReadInt64() int64
```
ReadInt64 Reads a int64 using the underlying io.Reader Any errors and the
running byte read count are stored instead or returned

#### func (*StickyBigEndianReader) ReadInt8

```go
func (e *StickyBigEndianReader) ReadInt8() int8
```
ReadInt8 Reads a int8 using the underlying io.Reader Any errors and the running
byte read count are stored instead or returned

#### func (*StickyBigEndianReader) ReadUint16

```go
func (e *StickyBigEndianReader) ReadUint16() uint16
```
ReadUint16 Reads a uint16 using the underlying io.Reader Any errors and the
running byte read count are stored instead or returned

#### func (*StickyBigEndianReader) ReadUint32

```go
func (e *StickyBigEndianReader) ReadUint32() uint32
```
ReadUint32 Reads a uint32 using the underlying io.Reader Any errors and the
running byte read count are stored instead or returned

#### func (*StickyBigEndianReader) ReadUint64

```go
func (e *StickyBigEndianReader) ReadUint64() uint64
```
ReadUint64 Reads a uint64 using the underlying io.Reader Any errors and the
running byte read count are stored instead or returned

#### func (*StickyBigEndianReader) ReadUint8

```go
func (e *StickyBigEndianReader) ReadUint8() uint8
```
ReadUint8 Reads a uint8 using the underlying io.Reader Any errors and the
running byte read count are stored instead or returned

#### type StickyBigEndianWriter

```go
type StickyBigEndianWriter struct {
	io.Writer

	Err   error
	Count int64
}
```

StickyBigEndianWriter wraps a io.Writer to provide methods to make it easier to
Write fundamental types

#### func (*StickyBigEndianWriter) Write

```go
func (e *StickyBigEndianWriter) Write(p []byte) (int, error)
```
Write implements the io.Writer interface

#### func (*StickyBigEndianWriter) WriteByte

```go
func (e *StickyBigEndianWriter) WriteByte(c byte) error
```
WriteByte implements the io.ByteWriter interface

#### func (*StickyBigEndianWriter) WriteFloat32

```go
func (e *StickyBigEndianWriter) WriteFloat32(d float32)
```
WriteFloat32 Writes a float32 using the underlying io.Writer Any errors and the
running byte read count are stored instead or returned

#### func (*StickyBigEndianWriter) WriteFloat64

```go
func (e *StickyBigEndianWriter) WriteFloat64(d float64)
```
WriteFloat64 Writes a float64 using the underlying io.Writer Any errors and the
running byte read count are stored instead or returned

#### func (*StickyBigEndianWriter) WriteInt16

```go
func (e *StickyBigEndianWriter) WriteInt16(d int16)
```
WriteInt16 Writes a int16 using the underlying io.Writer Any errors and the
running byte read count are stored instead or returned

#### func (*StickyBigEndianWriter) WriteInt32

```go
func (e *StickyBigEndianWriter) WriteInt32(d int32)
```
WriteInt32 Writes a int32 using the underlying io.Writer Any errors and the
running byte read count are stored instead or returned

#### func (*StickyBigEndianWriter) WriteInt64

```go
func (e *StickyBigEndianWriter) WriteInt64(d int64)
```
WriteInt64 Writes a int64 using the underlying io.Writer Any errors and the
running byte read count are stored instead or returned

#### func (*StickyBigEndianWriter) WriteInt8

```go
func (e *StickyBigEndianWriter) WriteInt8(d int8)
```
WriteInt8 Writes a int8 using the underlying io.Writer Any errors and the
running byte read count are stored instead or returned

#### func (*StickyBigEndianWriter) WriteUint16

```go
func (e *StickyBigEndianWriter) WriteUint16(d uint16)
```
WriteUint16 Writes a uint16 using the underlying io.Writer Any errors and the
running byte read count are stored instead or returned

#### func (*StickyBigEndianWriter) WriteUint32

```go
func (e *StickyBigEndianWriter) WriteUint32(d uint32)
```
WriteUint32 Writes a uint32 using the underlying io.Writer Any errors and the
running byte read count are stored instead or returned

#### func (*StickyBigEndianWriter) WriteUint64

```go
func (e *StickyBigEndianWriter) WriteUint64(d uint64)
```
WriteUint64 Writes a uint64 using the underlying io.Writer Any errors and the
running byte read count are stored instead or returned

#### func (*StickyBigEndianWriter) WriteUint8

```go
func (e *StickyBigEndianWriter) WriteUint8(d uint8)
```
WriteUint8 Writes a uint8 using the underlying io.Writer Any errors and the
running byte read count are stored instead or returned

#### type StickyEndianReader

```go
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
```

StickyEndianReader is an interface that reads various types with a particular
endianness and stores the Read return values

#### type StickyEndianWriter

```go
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
```

StickyEndianWriter is an interface that writes various types with a particular
endianness and stores the Write return values

#### type StickyLittleEndianReader

```go
type StickyLittleEndianReader struct {
	io.Reader

	Err   error
	Count int64
}
```

StickyLittleEndianReader wraps a io.Reader to provide methods to make it easier
to Read fundamental types

#### func (*StickyLittleEndianReader) Read

```go
func (e *StickyLittleEndianReader) Read(p []byte) (int, error)
```
Read implements the io.Reader interface

#### func (*StickyLittleEndianReader) ReadByte

```go
func (e *StickyLittleEndianReader) ReadByte() (byte, error)
```
ReadByte implements the io.ByteReader interface

#### func (*StickyLittleEndianReader) ReadFloat32

```go
func (e *StickyLittleEndianReader) ReadFloat32() float32
```
ReadFloat32 Reads a float32 using the underlying io.Reader Any errors and the
running byte read count are stored instead or returned

#### func (*StickyLittleEndianReader) ReadFloat64

```go
func (e *StickyLittleEndianReader) ReadFloat64() float64
```
ReadFloat64 Reads a float64 using the underlying io.Reader Any errors and the
running byte read count are stored instead or returned

#### func (*StickyLittleEndianReader) ReadInt16

```go
func (e *StickyLittleEndianReader) ReadInt16() int16
```
ReadInt16 Reads a int16 using the underlying io.Reader Any errors and the
running byte read count are stored instead or returned

#### func (*StickyLittleEndianReader) ReadInt32

```go
func (e *StickyLittleEndianReader) ReadInt32() int32
```
ReadInt32 Reads a int32 using the underlying io.Reader Any errors and the
running byte read count are stored instead or returned

#### func (*StickyLittleEndianReader) ReadInt64

```go
func (e *StickyLittleEndianReader) ReadInt64() int64
```
ReadInt64 Reads a int64 using the underlying io.Reader Any errors and the
running byte read count are stored instead or returned

#### func (*StickyLittleEndianReader) ReadInt8

```go
func (e *StickyLittleEndianReader) ReadInt8() int8
```
ReadInt8 Reads a int8 using the underlying io.Reader Any errors and the running
byte read count are stored instead or returned

#### func (*StickyLittleEndianReader) ReadUint16

```go
func (e *StickyLittleEndianReader) ReadUint16() uint16
```
ReadUint16 Reads a uint16 using the underlying io.Reader Any errors and the
running byte read count are stored instead or returned

#### func (*StickyLittleEndianReader) ReadUint32

```go
func (e *StickyLittleEndianReader) ReadUint32() uint32
```
ReadUint32 Reads a uint32 using the underlying io.Reader Any errors and the
running byte read count are stored instead or returned

#### func (*StickyLittleEndianReader) ReadUint64

```go
func (e *StickyLittleEndianReader) ReadUint64() uint64
```
ReadUint64 Reads a uint64 using the underlying io.Reader Any errors and the
running byte read count are stored instead or returned

#### func (*StickyLittleEndianReader) ReadUint8

```go
func (e *StickyLittleEndianReader) ReadUint8() uint8
```
ReadUint8 Reads a uint8 using the underlying io.Reader Any errors and the
running byte read count are stored instead or returned

#### type StickyLittleEndianWriter

```go
type StickyLittleEndianWriter struct {
	io.Writer

	Err   error
	Count int64
}
```

StickyLittleEndianWriter wraps a io.Writer to provide methods to make it easier
to Write fundamental types

#### func (*StickyLittleEndianWriter) Write

```go
func (e *StickyLittleEndianWriter) Write(p []byte) (int, error)
```
Write implements the io.Writer interface

#### func (*StickyLittleEndianWriter) WriteByte

```go
func (e *StickyLittleEndianWriter) WriteByte(c byte) error
```
WriteByte implements the io.ByteWriter interface

#### func (*StickyLittleEndianWriter) WriteFloat32

```go
func (e *StickyLittleEndianWriter) WriteFloat32(d float32)
```
WriteFloat32 Writes a float32 using the underlying io.Writer Any errors and the
running byte read count are stored instead or returned

#### func (*StickyLittleEndianWriter) WriteFloat64

```go
func (e *StickyLittleEndianWriter) WriteFloat64(d float64)
```
WriteFloat64 Writes a float64 using the underlying io.Writer Any errors and the
running byte read count are stored instead or returned

#### func (*StickyLittleEndianWriter) WriteInt16

```go
func (e *StickyLittleEndianWriter) WriteInt16(d int16)
```
WriteInt16 Writes a int16 using the underlying io.Writer Any errors and the
running byte read count are stored instead or returned

#### func (*StickyLittleEndianWriter) WriteInt32

```go
func (e *StickyLittleEndianWriter) WriteInt32(d int32)
```
WriteInt32 Writes a int32 using the underlying io.Writer Any errors and the
running byte read count are stored instead or returned

#### func (*StickyLittleEndianWriter) WriteInt64

```go
func (e *StickyLittleEndianWriter) WriteInt64(d int64)
```
WriteInt64 Writes a int64 using the underlying io.Writer Any errors and the
running byte read count are stored instead or returned

#### func (*StickyLittleEndianWriter) WriteInt8

```go
func (e *StickyLittleEndianWriter) WriteInt8(d int8)
```
WriteInt8 Writes a int8 using the underlying io.Writer Any errors and the
running byte read count are stored instead or returned

#### func (*StickyLittleEndianWriter) WriteUint16

```go
func (e *StickyLittleEndianWriter) WriteUint16(d uint16)
```
WriteUint16 Writes a uint16 using the underlying io.Writer Any errors and the
running byte read count are stored instead or returned

#### func (*StickyLittleEndianWriter) WriteUint32

```go
func (e *StickyLittleEndianWriter) WriteUint32(d uint32)
```
WriteUint32 Writes a uint32 using the underlying io.Writer Any errors and the
running byte read count are stored instead or returned

#### func (*StickyLittleEndianWriter) WriteUint64

```go
func (e *StickyLittleEndianWriter) WriteUint64(d uint64)
```
WriteUint64 Writes a uint64 using the underlying io.Writer Any errors and the
running byte read count are stored instead or returned

#### func (*StickyLittleEndianWriter) WriteUint8

```go
func (e *StickyLittleEndianWriter) WriteUint8(d uint8)
```
WriteUint8 Writes a uint8 using the underlying io.Writer Any errors and the
running byte read count are stored instead or returned

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
All methods will be a no-op after an error has been returned, returning 0,
unless that error is cleared on the type

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
