package byteio

// File automatically generated with ./gen.sh.

// StickyReader will wrap an EndianReader and count all byte handled and errors
// received.
// Byte counts and errors will not be returned from any method (except Read so
// it still counts as an io.Reader), but can be retrieved from this type.
// All methods will be a no-op after an error has been returned, returning 0,
// unless that error is cleared on the type.
type StickyReader struct {
	Reader EndianReader
	Err    error
	Count  int64
}

// Read will do a simple byte read from the underlying io.Reader.
func (s *StickyReader) Read(b []byte) (int, error) {
	if s.Err != nil {
		return 0, s.Err
	}

	n, err := s.Reader.Read(b)
	s.Err = err
	s.Count += int64(n)

	return n, err
}

// Read will Read a byte read with the underlying io.Reader.
func (s *StickyReader) ReadByte() (byte, error) {
	if s.Err != nil {
		return 0, s.Err
	}

	b, err := s.Reader.ReadByte()
	s.Err = err
	s.Count++

	return b, err
}

// Read will Read a boolean with the underlying io.Reader.
func (s *StickyReader) ReadBool() bool {
	if s.Err != nil {
		return false
	}

	b, n, err := s.Reader.ReadBool()
	s.Err = err
	s.Count += int64(n)

	return b
}

// ReadInt8 will read a int8 using the underlying reader.
func (s *StickyReader) ReadInt8() int8 {
	if s.Err != nil {
		return 0
	}

	i, n, err := s.Reader.ReadInt8()
	s.Err = err
	s.Count += int64(n)

	return i
}

// ReadInt16 will read a int16 using the underlying reader.
func (s *StickyReader) ReadInt16() int16 {
	if s.Err != nil {
		return 0
	}

	i, n, err := s.Reader.ReadInt16()
	s.Err = err
	s.Count += int64(n)

	return i
}

// ReadInt24 will read a int24 using the underlying reader.
func (s *StickyReader) ReadInt24() int32 {
	if s.Err != nil {
		return 0
	}

	i, n, err := s.Reader.ReadInt24()
	s.Err = err
	s.Count += int64(n)

	return i
}

// ReadInt32 will read a int32 using the underlying reader.
func (s *StickyReader) ReadInt32() int32 {
	if s.Err != nil {
		return 0
	}

	i, n, err := s.Reader.ReadInt32()
	s.Err = err
	s.Count += int64(n)

	return i
}

// ReadInt40 will read a int40 using the underlying reader.
func (s *StickyReader) ReadInt40() int64 {
	if s.Err != nil {
		return 0
	}

	i, n, err := s.Reader.ReadInt40()
	s.Err = err
	s.Count += int64(n)

	return i
}

// ReadInt48 will read a int48 using the underlying reader.
func (s *StickyReader) ReadInt48() int64 {
	if s.Err != nil {
		return 0
	}

	i, n, err := s.Reader.ReadInt48()
	s.Err = err
	s.Count += int64(n)

	return i
}

// ReadInt56 will read a int56 using the underlying reader.
func (s *StickyReader) ReadInt56() int64 {
	if s.Err != nil {
		return 0
	}

	i, n, err := s.Reader.ReadInt56()
	s.Err = err
	s.Count += int64(n)

	return i
}

// ReadInt64 will read a int64 using the underlying reader.
func (s *StickyReader) ReadInt64() int64 {
	if s.Err != nil {
		return 0
	}

	i, n, err := s.Reader.ReadInt64()
	s.Err = err
	s.Count += int64(n)

	return i
}

// ReadUint8 will read a uint8 using the underlying reader.
func (s *StickyReader) ReadUint8() uint8 {
	if s.Err != nil {
		return 0
	}

	i, n, err := s.Reader.ReadUint8()
	s.Err = err
	s.Count += int64(n)

	return i
}

// ReadUint16 will read a uint16 using the underlying reader.
func (s *StickyReader) ReadUint16() uint16 {
	if s.Err != nil {
		return 0
	}

	i, n, err := s.Reader.ReadUint16()
	s.Err = err
	s.Count += int64(n)

	return i
}

// ReadUint24 will read a uint24 using the underlying reader.
func (s *StickyReader) ReadUint24() uint32 {
	if s.Err != nil {
		return 0
	}

	i, n, err := s.Reader.ReadUint24()
	s.Err = err
	s.Count += int64(n)

	return i
}

// ReadUint32 will read a uint32 using the underlying reader.
func (s *StickyReader) ReadUint32() uint32 {
	if s.Err != nil {
		return 0
	}

	i, n, err := s.Reader.ReadUint32()
	s.Err = err
	s.Count += int64(n)

	return i
}

// ReadUint40 will read a uint40 using the underlying reader.
func (s *StickyReader) ReadUint40() uint64 {
	if s.Err != nil {
		return 0
	}

	i, n, err := s.Reader.ReadUint40()
	s.Err = err
	s.Count += int64(n)

	return i
}

// ReadUint48 will read a uint48 using the underlying reader.
func (s *StickyReader) ReadUint48() uint64 {
	if s.Err != nil {
		return 0
	}

	i, n, err := s.Reader.ReadUint48()
	s.Err = err
	s.Count += int64(n)

	return i
}

// ReadUint56 will read a uint56 using the underlying reader.
func (s *StickyReader) ReadUint56() uint64 {
	if s.Err != nil {
		return 0
	}

	i, n, err := s.Reader.ReadUint56()
	s.Err = err
	s.Count += int64(n)

	return i
}

// ReadUint64 will read a uint64 using the underlying reader.
func (s *StickyReader) ReadUint64() uint64 {
	if s.Err != nil {
		return 0
	}

	i, n, err := s.Reader.ReadUint64()
	s.Err = err
	s.Count += int64(n)

	return i
}

// ReadFloat32 will read a float32 using the underlying reader.
func (s *StickyReader) ReadFloat32() float32 {
	if s.Err != nil {
		return 0
	}

	i, n, err := s.Reader.ReadFloat32()
	s.Err = err
	s.Count += int64(n)

	return i
}

// ReadFloat64 will read a float64 using the underlying reader.
func (s *StickyReader) ReadFloat64() float64 {
	if s.Err != nil {
		return 0
	}

	i, n, err := s.Reader.ReadFloat64()
	s.Err = err
	s.Count += int64(n)

	return i
}

// ReadIntX will read a 64-bit var-int using the underlying ReadIntX method.
func (s *StickyReader) ReadIntX() int64 {
	if s.Err != nil {
		return 0
	}

	i, n, err := s.Reader.ReadIntX()
	s.Err = err
	s.Count += int64(n)

	return i
}

// ReadUintX will read a 64-bit var-uint using the underlying ReadUintX method.
func (s *StickyReader) ReadUintX() uint64 {
	if s.Err != nil {
		return 0
	}

	i, n, err := s.Reader.ReadUintX()
	s.Err = err
	s.Count += int64(n)

	return i
}

// ReadBytes will read a bytes using the underlying ReadBytes method.
func (s *StickyReader) ReadBytes(p int) []byte {
	if s.Err != nil {
		return nil
	}

	r, n, err := s.Reader.ReadBytes(p)
	s.Err = err
	s.Count += int64(n)

	return r
}

// ReadBytesX will read a bytes using the underlying ReadBytesX method.
func (s *StickyReader) ReadBytesX() []byte {
	if s.Err != nil {
		return nil
	}

	r, n, err := s.Reader.ReadBytesX()
	s.Err = err
	s.Count += int64(n)

	return r
}

// ReadBytes8 will read a bytes using the underlying ReadBytes8 method.
func (s *StickyReader) ReadBytes8() []byte {
	if s.Err != nil {
		return nil
	}

	r, n, err := s.Reader.ReadBytes8()
	s.Err = err
	s.Count += int64(n)

	return r
}

// ReadBytes16 will read a bytes using the underlying ReadBytes16 method.
func (s *StickyReader) ReadBytes16() []byte {
	if s.Err != nil {
		return nil
	}

	r, n, err := s.Reader.ReadBytes16()
	s.Err = err
	s.Count += int64(n)

	return r
}

// ReadBytes24 will read a bytes using the underlying ReadBytes24 method.
func (s *StickyReader) ReadBytes24() []byte {
	if s.Err != nil {
		return nil
	}

	r, n, err := s.Reader.ReadBytes24()
	s.Err = err
	s.Count += int64(n)

	return r
}

// ReadBytes32 will read a bytes using the underlying ReadBytes32 method.
func (s *StickyReader) ReadBytes32() []byte {
	if s.Err != nil {
		return nil
	}

	r, n, err := s.Reader.ReadBytes32()
	s.Err = err
	s.Count += int64(n)

	return r
}

// ReadBytes40 will read a bytes using the underlying ReadBytes40 method.
func (s *StickyReader) ReadBytes40() []byte {
	if s.Err != nil {
		return nil
	}

	r, n, err := s.Reader.ReadBytes40()
	s.Err = err
	s.Count += int64(n)

	return r
}

// ReadBytes48 will read a bytes using the underlying ReadBytes48 method.
func (s *StickyReader) ReadBytes48() []byte {
	if s.Err != nil {
		return nil
	}

	r, n, err := s.Reader.ReadBytes48()
	s.Err = err
	s.Count += int64(n)

	return r
}

// ReadBytes56 will read a bytes using the underlying ReadBytes56 method.
func (s *StickyReader) ReadBytes56() []byte {
	if s.Err != nil {
		return nil
	}

	r, n, err := s.Reader.ReadBytes56()
	s.Err = err
	s.Count += int64(n)

	return r
}

// ReadBytes64 will read a bytes using the underlying ReadBytes64 method.
func (s *StickyReader) ReadBytes64() []byte {
	if s.Err != nil {
		return nil
	}

	r, n, err := s.Reader.ReadBytes64()
	s.Err = err
	s.Count += int64(n)

	return r
}

// ReadString will read a string using the underlying ReadString method.
func (s *StickyReader) ReadString(p int) string {
	if s.Err != nil {
		return ""
	}

	r, n, err := s.Reader.ReadString(p)
	s.Err = err
	s.Count += int64(n)

	return r
}

// ReadString0 will read a string using the underlying ReadString0 method.
func (s *StickyReader) ReadString0() string {
	if s.Err != nil {
		return ""
	}

	r, n, err := s.Reader.ReadString0()
	s.Err = err
	s.Count += int64(n)

	return r
}

// ReadStringX will read a string using the underlying ReadStringX method.
func (s *StickyReader) ReadStringX() string {
	if s.Err != nil {
		return ""
	}

	r, n, err := s.Reader.ReadStringX()
	s.Err = err
	s.Count += int64(n)

	return r
}

// ReadString8 will read a string using the underlying ReadString8 method.
func (s *StickyReader) ReadString8() string {
	if s.Err != nil {
		return ""
	}

	r, n, err := s.Reader.ReadString8()
	s.Err = err
	s.Count += int64(n)

	return r
}

// ReadString16 will read a string using the underlying ReadString16 method.
func (s *StickyReader) ReadString16() string {
	if s.Err != nil {
		return ""
	}

	r, n, err := s.Reader.ReadString16()
	s.Err = err
	s.Count += int64(n)

	return r
}

// ReadString24 will read a string using the underlying ReadString24 method.
func (s *StickyReader) ReadString24() string {
	if s.Err != nil {
		return ""
	}

	r, n, err := s.Reader.ReadString24()
	s.Err = err
	s.Count += int64(n)

	return r
}

// ReadString32 will read a string using the underlying ReadString32 method.
func (s *StickyReader) ReadString32() string {
	if s.Err != nil {
		return ""
	}

	r, n, err := s.Reader.ReadString32()
	s.Err = err
	s.Count += int64(n)

	return r
}

// ReadString40 will read a string using the underlying ReadString40 method.
func (s *StickyReader) ReadString40() string {
	if s.Err != nil {
		return ""
	}

	r, n, err := s.Reader.ReadString40()
	s.Err = err
	s.Count += int64(n)

	return r
}

// ReadString48 will read a string using the underlying ReadString48 method.
func (s *StickyReader) ReadString48() string {
	if s.Err != nil {
		return ""
	}

	r, n, err := s.Reader.ReadString48()
	s.Err = err
	s.Count += int64(n)

	return r
}

// ReadString56 will read a string using the underlying ReadString56 method.
func (s *StickyReader) ReadString56() string {
	if s.Err != nil {
		return ""
	}

	r, n, err := s.Reader.ReadString56()
	s.Err = err
	s.Count += int64(n)

	return r
}

// ReadString64 will read a string using the underlying ReadString64 method.
func (s *StickyReader) ReadString64() string {
	if s.Err != nil {
		return ""
	}

	r, n, err := s.Reader.ReadString64()
	s.Err = err
	s.Count += int64(n)

	return r
}

// StickyWriter will wrap an EndianWriter and count all byte handled and errors
// received.
// Byte counts and errors will not be returned from any method (except Write so
// it still counts as an io.Writer), but can be retrieved from this type.
// All methods will be a no-op after an error has been returned, returning 0,
// unless that error is cleared on the type.
type StickyWriter struct {
	Writer EndianWriter
	Err    error
	Count  int64
}

// Write will do a simple byte write from the underlying io.Writer.
func (s *StickyWriter) Write(b []byte) (int, error) {
	if s.Err != nil {
		return 0, s.Err
	}

	n, err := s.Writer.Write(b)
	s.Err = err
	s.Count += int64(n)

	return n, err
}

// Write will Write a byte write with the underlying io.Writer.
func (s *StickyWriter) WriteByte(b byte) error {
	if s.Err != nil {
		return s.Err
	}

	err := s.Writer.WriteByte(b)
	s.Err = err
	s.Count++

	return err
}

// Write will Write a boolean with the underlying io.Writer.
func (s *StickyWriter) WriteBool(b bool) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteBool(b)
	s.Err = err
	s.Count += int64(n)
}

// WriteInt8 will write a int8 using the underlying writer.
func (s *StickyWriter) WriteInt8(i int8) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteInt8(i)
	s.Err = err
	s.Count += int64(n)
}

// WriteInt16 will write a int16 using the underlying writer.
func (s *StickyWriter) WriteInt16(i int16) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteInt16(i)
	s.Err = err
	s.Count += int64(n)
}

// WriteInt24 will write a int24 using the underlying writer.
func (s *StickyWriter) WriteInt24(i int32) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteInt24(i)
	s.Err = err
	s.Count += int64(n)
}

// WriteInt32 will write a int32 using the underlying writer.
func (s *StickyWriter) WriteInt32(i int32) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteInt32(i)
	s.Err = err
	s.Count += int64(n)
}

// WriteInt40 will write a int40 using the underlying writer.
func (s *StickyWriter) WriteInt40(i int64) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteInt40(i)
	s.Err = err
	s.Count += int64(n)
}

// WriteInt48 will write a int48 using the underlying writer.
func (s *StickyWriter) WriteInt48(i int64) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteInt48(i)
	s.Err = err
	s.Count += int64(n)
}

// WriteInt56 will write a int56 using the underlying writer.
func (s *StickyWriter) WriteInt56(i int64) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteInt56(i)
	s.Err = err
	s.Count += int64(n)
}

// WriteInt64 will write a int64 using the underlying writer.
func (s *StickyWriter) WriteInt64(i int64) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteInt64(i)
	s.Err = err
	s.Count += int64(n)
}

// WriteUint8 will write a uint8 using the underlying writer.
func (s *StickyWriter) WriteUint8(i uint8) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteUint8(i)
	s.Err = err
	s.Count += int64(n)
}

// WriteUint16 will write a uint16 using the underlying writer.
func (s *StickyWriter) WriteUint16(i uint16) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteUint16(i)
	s.Err = err
	s.Count += int64(n)
}

// WriteUint24 will write a uint24 using the underlying writer.
func (s *StickyWriter) WriteUint24(i uint32) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteUint24(i)
	s.Err = err
	s.Count += int64(n)
}

// WriteUint32 will write a uint32 using the underlying writer.
func (s *StickyWriter) WriteUint32(i uint32) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteUint32(i)
	s.Err = err
	s.Count += int64(n)
}

// WriteUint40 will write a uint40 using the underlying writer.
func (s *StickyWriter) WriteUint40(i uint64) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteUint40(i)
	s.Err = err
	s.Count += int64(n)
}

// WriteUint48 will write a uint48 using the underlying writer.
func (s *StickyWriter) WriteUint48(i uint64) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteUint48(i)
	s.Err = err
	s.Count += int64(n)
}

// WriteUint56 will write a uint56 using the underlying writer.
func (s *StickyWriter) WriteUint56(i uint64) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteUint56(i)
	s.Err = err
	s.Count += int64(n)
}

// WriteUint64 will write a uint64 using the underlying writer.
func (s *StickyWriter) WriteUint64(i uint64) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteUint64(i)
	s.Err = err
	s.Count += int64(n)
}

// WriteFloat32 will write a float32 using the underlying writer.
func (s *StickyWriter) WriteFloat32(i float32) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteFloat32(i)
	s.Err = err
	s.Count += int64(n)
}

// WriteFloat64 will write a float64 using the underlying writer.
func (s *StickyWriter) WriteFloat64(i float64) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteFloat64(i)
	s.Err = err
	s.Count += int64(n)
}

// WriteIntX will write a 64-bit var-int using the underlying WriteIntX method.
func (s *StickyWriter) WriteIntX(i int64) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteIntX(i)
	s.Err = err
	s.Count += int64(n)
}

// WriteUintX will write a 64-bit var-uint using the underlying WriteUintX method.
func (s *StickyWriter) WriteUintX(i uint64) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteUintX(i)
	s.Err = err
	s.Count += int64(n)
}

// WriteBytes will write a bytes using the underlying WriteBytes method.
func (s *StickyWriter) WriteBytes(p []byte) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteBytes(p)
	s.Err = err
	s.Count += int64(n)
}

// WriteBytesX will write a bytes using the underlying WriteBytesX method.
func (s *StickyWriter) WriteBytesX(p []byte) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteBytesX(p)
	s.Err = err
	s.Count += int64(n)
}

// WriteBytes8 will write a bytes using the underlying WriteBytes8 method.
func (s *StickyWriter) WriteBytes8(p []byte) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteBytes8(p)
	s.Err = err
	s.Count += int64(n)
}

// WriteBytes16 will write a bytes using the underlying WriteBytes16 method.
func (s *StickyWriter) WriteBytes16(p []byte) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteBytes16(p)
	s.Err = err
	s.Count += int64(n)
}

// WriteBytes24 will write a bytes using the underlying WriteBytes24 method.
func (s *StickyWriter) WriteBytes24(p []byte) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteBytes24(p)
	s.Err = err
	s.Count += int64(n)
}

// WriteBytes32 will write a bytes using the underlying WriteBytes32 method.
func (s *StickyWriter) WriteBytes32(p []byte) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteBytes32(p)
	s.Err = err
	s.Count += int64(n)
}

// WriteBytes40 will write a bytes using the underlying WriteBytes40 method.
func (s *StickyWriter) WriteBytes40(p []byte) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteBytes40(p)
	s.Err = err
	s.Count += int64(n)
}

// WriteBytes48 will write a bytes using the underlying WriteBytes48 method.
func (s *StickyWriter) WriteBytes48(p []byte) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteBytes48(p)
	s.Err = err
	s.Count += int64(n)
}

// WriteBytes56 will write a bytes using the underlying WriteBytes56 method.
func (s *StickyWriter) WriteBytes56(p []byte) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteBytes56(p)
	s.Err = err
	s.Count += int64(n)
}

// WriteBytes64 will write a bytes using the underlying WriteBytes64 method.
func (s *StickyWriter) WriteBytes64(p []byte) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteBytes64(p)
	s.Err = err
	s.Count += int64(n)
}

// WriteString will write a string using the underlying WriteString method.
func (s *StickyWriter) WriteString(p string) (int, error) {
	if s.Err != nil {
		return 0, s.Err
	}

	n, err := s.Writer.WriteString(p)
	s.Err = err
	s.Count += int64(n)

	return n, err
}

// WriteString0 will write a string using the underlying WriteString0 method.
func (s *StickyWriter) WriteString0(p string) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteString0(p)
	s.Err = err
	s.Count += int64(n)
}

// WriteStringX will write a string using the underlying WriteStringX method.
func (s *StickyWriter) WriteStringX(p string) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteStringX(p)
	s.Err = err
	s.Count += int64(n)
}

// WriteString8 will write a string using the underlying WriteString8 method.
func (s *StickyWriter) WriteString8(p string) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteString8(p)
	s.Err = err
	s.Count += int64(n)
}

// WriteString16 will write a string using the underlying WriteString16 method.
func (s *StickyWriter) WriteString16(p string) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteString16(p)
	s.Err = err
	s.Count += int64(n)
}

// WriteString24 will write a string using the underlying WriteString24 method.
func (s *StickyWriter) WriteString24(p string) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteString24(p)
	s.Err = err
	s.Count += int64(n)
}

// WriteString32 will write a string using the underlying WriteString32 method.
func (s *StickyWriter) WriteString32(p string) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteString32(p)
	s.Err = err
	s.Count += int64(n)
}

// WriteString40 will write a string using the underlying WriteString40 method.
func (s *StickyWriter) WriteString40(p string) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteString40(p)
	s.Err = err
	s.Count += int64(n)
}

// WriteString48 will write a string using the underlying WriteString48 method.
func (s *StickyWriter) WriteString48(p string) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteString48(p)
	s.Err = err
	s.Count += int64(n)
}

// WriteString56 will write a string using the underlying WriteString56 method.
func (s *StickyWriter) WriteString56(p string) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteString56(p)
	s.Err = err
	s.Count += int64(n)
}

// WriteString64 will write a string using the underlying WriteString64 method.
func (s *StickyWriter) WriteString64(p string) {
	if s.Err != nil {
		return
	}

	n, err := s.Writer.WriteString64(p)
	s.Err = err
	s.Count += int64(n)
}
