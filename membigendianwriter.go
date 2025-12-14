package byteio

// File automatically generated with ./gen.sh.

import "math"

// MemBigEndianWriter is a byte slice that has methods
// to make it easier to Write fundamental types.
type MemBigEndianWriter []byte

// WriteBool Writes a boolean.
func (e *MemBigEndianWriter) WriteBool(b bool) {
	if b {
		e.WriteUint8(1)
	} else {
		e.WriteUint8(0)
	}
}

// WriteInt8 Writes a 8 bit int as a int8 to the byte slice.
func (e *MemBigEndianWriter) WriteInt8(d int8) {
	c := uint8(d)
	*e = append(*e,
		byte(c),
	)
}

// WriteInt16 Writes a 16 bit int as a int16 to the byte slice.
func (e *MemBigEndianWriter) WriteInt16(d int16) {
	c := uint16(d)
	*e = append(*e,
		byte(c>>8),
		byte(c),
	)
}

// WriteInt24 Writes a 24 bit int as a int32 to the byte slice.
func (e *MemBigEndianWriter) WriteInt24(d int32) {
	c := uint32(d)
	*e = append(*e,
		byte(c>>16),
		byte(c>>8),
		byte(c),
	)
}

// WriteInt32 Writes a 32 bit int as a int32 to the byte slice.
func (e *MemBigEndianWriter) WriteInt32(d int32) {
	c := uint32(d)
	*e = append(*e,
		byte(c>>24),
		byte(c>>16),
		byte(c>>8),
		byte(c),
	)
}

// WriteInt40 Writes a 40 bit int as a int64 to the byte slice.
func (e *MemBigEndianWriter) WriteInt40(d int64) {
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
func (e *MemBigEndianWriter) WriteInt48(d int64) {
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
func (e *MemBigEndianWriter) WriteInt56(d int64) {
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
func (e *MemBigEndianWriter) WriteInt64(d int64) {
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
func (e *MemBigEndianWriter) WriteUint8(d uint8) {
	*e = append(*e,
		byte(d),
	)
}

// WriteUint16 Writes a 16 bit uint as a uint16 to the byte slice.
func (e *MemBigEndianWriter) WriteUint16(d uint16) {
	*e = append(*e,
		byte(d>>8),
		byte(d),
	)
}

// WriteUint24 Writes a 24 bit uint as a uint32 to the byte slice.
func (e *MemBigEndianWriter) WriteUint24(d uint32) {
	*e = append(*e,
		byte(d>>16),
		byte(d>>8),
		byte(d),
	)
}

// WriteUint32 Writes a 32 bit uint as a uint32 to the byte slice.
func (e *MemBigEndianWriter) WriteUint32(d uint32) {
	*e = append(*e,
		byte(d>>24),
		byte(d>>16),
		byte(d>>8),
		byte(d),
	)
}

// WriteUint40 Writes a 40 bit uint as a uint64 to the byte slice.
func (e *MemBigEndianWriter) WriteUint40(d uint64) {
	*e = append(*e,
		byte(d>>32),
		byte(d>>24),
		byte(d>>16),
		byte(d>>8),
		byte(d),
	)
}

// WriteUint48 Writes a 48 bit uint as a uint64 to the byte slice.
func (e *MemBigEndianWriter) WriteUint48(d uint64) {
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
func (e *MemBigEndianWriter) WriteUint56(d uint64) {
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
func (e *MemBigEndianWriter) WriteUint64(d uint64) {
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
func (e *MemBigEndianWriter) WriteFloat32(d float32) {
	c := math.Float32bits(d)
	*e = append(*e,
		byte(c>>24),
		byte(c>>16),
		byte(c>>8),
		byte(c),
	)
}

// WriteFloat64 Writes a 64 bit float as a float64 to the byte slice.
func (e *MemBigEndianWriter) WriteFloat64(d float64) {
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
func (e *MemBigEndianWriter) WriteBytes(d []byte) {
	*e = append(*e, d...)
}

// MemBigEndianWriter implements io.Writer.
func (e *MemBigEndianWriter) Write(p []byte) (int, error) {
	*e = append(*e, p...)

	return len(p), nil
}

// WriteBytesX Writes the length of the Bytes, using WriteUintX and then Writes the bytes.
func (e *MemBigEndianWriter) WriteBytesX(p []byte) {
	e.WriteUintX(uint64(len(p)))
	e.WriteBytes(p)
}

// WriteBytes8 Writes the length of the Bytes, using WriteUint8 and then Writes the bytes.
func (e *MemBigEndianWriter) WriteBytes8(p []byte) {
	e.WriteUint8(uint8(len(p)))
	e.WriteBytes(p)
}

// WriteBytes16 Writes the length of the Bytes, using WriteUint16 and then Writes the bytes.
func (e *MemBigEndianWriter) WriteBytes16(p []byte) {
	e.WriteUint16(uint16(len(p)))
	e.WriteBytes(p)
}

// WriteBytes24 Writes the length of the Bytes, using WriteUint24 and then Writes the bytes.
func (e *MemBigEndianWriter) WriteBytes24(p []byte) {
	e.WriteUint24(uint32(len(p)))
	e.WriteBytes(p)
}

// WriteBytes32 Writes the length of the Bytes, using WriteUint32 and then Writes the bytes.
func (e *MemBigEndianWriter) WriteBytes32(p []byte) {
	e.WriteUint32(uint32(len(p)))
	e.WriteBytes(p)
}

// WriteBytes40 Writes the length of the Bytes, using WriteUint40 and then Writes the bytes.
func (e *MemBigEndianWriter) WriteBytes40(p []byte) {
	e.WriteUint40(uint64(len(p)))
	e.WriteBytes(p)
}

// WriteBytes48 Writes the length of the Bytes, using WriteUint48 and then Writes the bytes.
func (e *MemBigEndianWriter) WriteBytes48(p []byte) {
	e.WriteUint48(uint64(len(p)))
	e.WriteBytes(p)
}

// WriteBytes56 Writes the length of the Bytes, using WriteUint56 and then Writes the bytes.
func (e *MemBigEndianWriter) WriteBytes56(p []byte) {
	e.WriteUint56(uint64(len(p)))
	e.WriteBytes(p)
}

// WriteBytes64 Writes the length of the Bytes, using WriteUint64 and then Writes the bytes.
func (e *MemBigEndianWriter) WriteBytes64(p []byte) {
	e.WriteUint64(uint64(len(p)))
	e.WriteBytes(p)
}

// WriteString Writes a string.
func (e *MemBigEndianWriter) WriteString(d string) (int, error) {
	*e = append(*e, d...)

	return len(d), nil
}

// WriteStringX Writes the length of the String, using WriteUintX and then Writes the bytes.
func (e *MemBigEndianWriter) WriteStringX(p string) {
	e.WriteUintX(uint64(len(p)))
	e.WriteString(p)
}

// WriteString8 Writes the length of the String, using WriteUint8 and then Writes the bytes.
func (e *MemBigEndianWriter) WriteString8(p string) {
	e.WriteUint8(uint8(len(p)))
	e.WriteString(p)
}

// WriteString16 Writes the length of the String, using WriteUint16 and then Writes the bytes.
func (e *MemBigEndianWriter) WriteString16(p string) {
	e.WriteUint16(uint16(len(p)))
	e.WriteString(p)
}

// WriteString24 Writes the length of the String, using WriteUint24 and then Writes the bytes.
func (e *MemBigEndianWriter) WriteString24(p string) {
	e.WriteUint24(uint32(len(p)))
	e.WriteString(p)
}

// WriteString32 Writes the length of the String, using WriteUint32 and then Writes the bytes.
func (e *MemBigEndianWriter) WriteString32(p string) {
	e.WriteUint32(uint32(len(p)))
	e.WriteString(p)
}

// WriteString40 Writes the length of the String, using WriteUint40 and then Writes the bytes.
func (e *MemBigEndianWriter) WriteString40(p string) {
	e.WriteUint40(uint64(len(p)))
	e.WriteString(p)
}

// WriteString48 Writes the length of the String, using WriteUint48 and then Writes the bytes.
func (e *MemBigEndianWriter) WriteString48(p string) {
	e.WriteUint48(uint64(len(p)))
	e.WriteString(p)
}

// WriteString56 Writes the length of the String, using WriteUint56 and then Writes the bytes.
func (e *MemBigEndianWriter) WriteString56(p string) {
	e.WriteUint56(uint64(len(p)))
	e.WriteString(p)
}

// WriteString64 Writes the length of the String, using WriteUint64 and then Writes the bytes.
func (e *MemBigEndianWriter) WriteString64(p string) {
	e.WriteUint64(uint64(len(p)))
	e.WriteString(p)
}

// WriteString0 Writes the bytes of the string ending with a 0 byte.
func (e *MemBigEndianWriter) WriteString0(p string) {
	e.WriteString(p)
	e.WriteUint8(0)
}
