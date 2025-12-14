package byteio

// File automatically generated with ./gen.sh.

import "math"

// MemLittleEndianWriter is a byte slice that has methods
// to make it easier to Write fundamental types.
type MemLittleEndianWriter []byte

// WriteBool Writes a boolean.
func (e *MemLittleEndianWriter) WriteBool(b bool) {
	if b {
		e.WriteUint8(1)
	} else {
		e.WriteUint8(0)
	}
}

// WriteInt8 Writes a 8 bit int as a int8 to the byte slice.
func (e *MemLittleEndianWriter) WriteInt8(d int8) {
	c := uint8(d)
	*e = append(*e,
		byte(c),
	)
}

// WriteInt16 Writes a 16 bit int as a int16 to the byte slice.
func (e *MemLittleEndianWriter) WriteInt16(d int16) {
	c := uint16(d)
	*e = append(*e,
		byte(c),
		byte(c>>8),
	)
}

// WriteInt24 Writes a 24 bit int as a int32 to the byte slice.
func (e *MemLittleEndianWriter) WriteInt24(d int32) {
	c := uint32(d)
	*e = append(*e,
		byte(c),
		byte(c>>8),
		byte(c>>16),
	)
}

// WriteInt32 Writes a 32 bit int as a int32 to the byte slice.
func (e *MemLittleEndianWriter) WriteInt32(d int32) {
	c := uint32(d)
	*e = append(*e,
		byte(c),
		byte(c>>8),
		byte(c>>16),
		byte(c>>24),
	)
}

// WriteInt40 Writes a 40 bit int as a int64 to the byte slice.
func (e *MemLittleEndianWriter) WriteInt40(d int64) {
	c := uint64(d)
	*e = append(*e,
		byte(c),
		byte(c>>8),
		byte(c>>16),
		byte(c>>24),
		byte(c>>32),
	)
}

// WriteInt48 Writes a 48 bit int as a int64 to the byte slice.
func (e *MemLittleEndianWriter) WriteInt48(d int64) {
	c := uint64(d)
	*e = append(*e,
		byte(c),
		byte(c>>8),
		byte(c>>16),
		byte(c>>24),
		byte(c>>32),
		byte(c>>40),
	)
}

// WriteInt56 Writes a 56 bit int as a int64 to the byte slice.
func (e *MemLittleEndianWriter) WriteInt56(d int64) {
	c := uint64(d)
	*e = append(*e,
		byte(c),
		byte(c>>8),
		byte(c>>16),
		byte(c>>24),
		byte(c>>32),
		byte(c>>40),
		byte(c>>48),
	)
}

// WriteInt64 Writes a 64 bit int as a int64 to the byte slice.
func (e *MemLittleEndianWriter) WriteInt64(d int64) {
	c := uint64(d)
	*e = append(*e,
		byte(c),
		byte(c>>8),
		byte(c>>16),
		byte(c>>24),
		byte(c>>32),
		byte(c>>40),
		byte(c>>48),
		byte(c>>56),
	)
}

// WriteUint8 Writes a 8 bit uint as a uint8 to the byte slice.
func (e *MemLittleEndianWriter) WriteUint8(d uint8) {
	*e = append(*e,
		byte(d),
	)
}

// WriteUint16 Writes a 16 bit uint as a uint16 to the byte slice.
func (e *MemLittleEndianWriter) WriteUint16(d uint16) {
	*e = append(*e,
		byte(d),
		byte(d>>8),
	)
}

// WriteUint24 Writes a 24 bit uint as a uint32 to the byte slice.
func (e *MemLittleEndianWriter) WriteUint24(d uint32) {
	*e = append(*e,
		byte(d),
		byte(d>>8),
		byte(d>>16),
	)
}

// WriteUint32 Writes a 32 bit uint as a uint32 to the byte slice.
func (e *MemLittleEndianWriter) WriteUint32(d uint32) {
	*e = append(*e,
		byte(d),
		byte(d>>8),
		byte(d>>16),
		byte(d>>24),
	)
}

// WriteUint40 Writes a 40 bit uint as a uint64 to the byte slice.
func (e *MemLittleEndianWriter) WriteUint40(d uint64) {
	*e = append(*e,
		byte(d),
		byte(d>>8),
		byte(d>>16),
		byte(d>>24),
		byte(d>>32),
	)
}

// WriteUint48 Writes a 48 bit uint as a uint64 to the byte slice.
func (e *MemLittleEndianWriter) WriteUint48(d uint64) {
	*e = append(*e,
		byte(d),
		byte(d>>8),
		byte(d>>16),
		byte(d>>24),
		byte(d>>32),
		byte(d>>40),
	)
}

// WriteUint56 Writes a 56 bit uint as a uint64 to the byte slice.
func (e *MemLittleEndianWriter) WriteUint56(d uint64) {
	*e = append(*e,
		byte(d),
		byte(d>>8),
		byte(d>>16),
		byte(d>>24),
		byte(d>>32),
		byte(d>>40),
		byte(d>>48),
	)
}

// WriteUint64 Writes a 64 bit uint as a uint64 to the byte slice.
func (e *MemLittleEndianWriter) WriteUint64(d uint64) {
	*e = append(*e,
		byte(d),
		byte(d>>8),
		byte(d>>16),
		byte(d>>24),
		byte(d>>32),
		byte(d>>40),
		byte(d>>48),
		byte(d>>56),
	)
}

// WriteFloat32 Writes a 32 bit float as a float32 to the byte slice.
func (e *MemLittleEndianWriter) WriteFloat32(d float32) {
	c := math.Float32bits(d)
	*e = append(*e,
		byte(c),
		byte(c>>8),
		byte(c>>16),
		byte(c>>24),
	)
}

// WriteFloat64 Writes a 64 bit float as a float64 to the byte slice.
func (e *MemLittleEndianWriter) WriteFloat64(d float64) {
	c := math.Float64bits(d)
	*e = append(*e,
		byte(c),
		byte(c>>8),
		byte(c>>16),
		byte(c>>24),
		byte(c>>32),
		byte(c>>40),
		byte(c>>48),
		byte(c>>56),
	)
}

// WriteBytes Writes a []byte.
func (e *MemLittleEndianWriter) WriteBytes(d []byte) {
	*e = append(*e, d...)
}

// MemLittleEndianWriter implements io.Writer.
func (e *MemLittleEndianWriter) Write(p []byte) (int, error) {
	*e = append(*e, p...)

	return len(p), nil
}

// WriteBytesX Writes the length of the Bytes, using WriteUintX and then Writes the bytes.
func (e *MemLittleEndianWriter) WriteBytesX(p []byte) {
	e.WriteUintX(uint64(len(p)))
	e.WriteBytes(p)
}

// WriteBytes8 Writes the length of the Bytes, using WriteUint8 and then Writes the bytes.
func (e *MemLittleEndianWriter) WriteBytes8(p []byte) {
	e.WriteUint8(uint8(len(p)))
	e.WriteBytes(p)
}

// WriteBytes16 Writes the length of the Bytes, using WriteUint16 and then Writes the bytes.
func (e *MemLittleEndianWriter) WriteBytes16(p []byte) {
	e.WriteUint16(uint16(len(p)))
	e.WriteBytes(p)
}

// WriteBytes24 Writes the length of the Bytes, using WriteUint24 and then Writes the bytes.
func (e *MemLittleEndianWriter) WriteBytes24(p []byte) {
	e.WriteUint24(uint32(len(p)))
	e.WriteBytes(p)
}

// WriteBytes32 Writes the length of the Bytes, using WriteUint32 and then Writes the bytes.
func (e *MemLittleEndianWriter) WriteBytes32(p []byte) {
	e.WriteUint32(uint32(len(p)))
	e.WriteBytes(p)
}

// WriteBytes40 Writes the length of the Bytes, using WriteUint40 and then Writes the bytes.
func (e *MemLittleEndianWriter) WriteBytes40(p []byte) {
	e.WriteUint40(uint64(len(p)))
	e.WriteBytes(p)
}

// WriteBytes48 Writes the length of the Bytes, using WriteUint48 and then Writes the bytes.
func (e *MemLittleEndianWriter) WriteBytes48(p []byte) {
	e.WriteUint48(uint64(len(p)))
	e.WriteBytes(p)
}

// WriteBytes56 Writes the length of the Bytes, using WriteUint56 and then Writes the bytes.
func (e *MemLittleEndianWriter) WriteBytes56(p []byte) {
	e.WriteUint56(uint64(len(p)))
	e.WriteBytes(p)
}

// WriteBytes64 Writes the length of the Bytes, using WriteUint64 and then Writes the bytes.
func (e *MemLittleEndianWriter) WriteBytes64(p []byte) {
	e.WriteUint64(uint64(len(p)))
	e.WriteBytes(p)
}

// WriteString Writes a string.
func (e *MemLittleEndianWriter) WriteString(d string) (int, error) {
	*e = append(*e, d...)

	return len(d), nil
}

// WriteStringX Writes the length of the String, using WriteUintX and then Writes the bytes.
func (e *MemLittleEndianWriter) WriteStringX(p string) {
	e.WriteUintX(uint64(len(p)))
	e.WriteString(p)
}

// WriteString8 Writes the length of the String, using WriteUint8 and then Writes the bytes.
func (e *MemLittleEndianWriter) WriteString8(p string) {
	e.WriteUint8(uint8(len(p)))
	e.WriteString(p)
}

// WriteString16 Writes the length of the String, using WriteUint16 and then Writes the bytes.
func (e *MemLittleEndianWriter) WriteString16(p string) {
	e.WriteUint16(uint16(len(p)))
	e.WriteString(p)
}

// WriteString24 Writes the length of the String, using WriteUint24 and then Writes the bytes.
func (e *MemLittleEndianWriter) WriteString24(p string) {
	e.WriteUint24(uint32(len(p)))
	e.WriteString(p)
}

// WriteString32 Writes the length of the String, using WriteUint32 and then Writes the bytes.
func (e *MemLittleEndianWriter) WriteString32(p string) {
	e.WriteUint32(uint32(len(p)))
	e.WriteString(p)
}

// WriteString40 Writes the length of the String, using WriteUint40 and then Writes the bytes.
func (e *MemLittleEndianWriter) WriteString40(p string) {
	e.WriteUint40(uint64(len(p)))
	e.WriteString(p)
}

// WriteString48 Writes the length of the String, using WriteUint48 and then Writes the bytes.
func (e *MemLittleEndianWriter) WriteString48(p string) {
	e.WriteUint48(uint64(len(p)))
	e.WriteString(p)
}

// WriteString56 Writes the length of the String, using WriteUint56 and then Writes the bytes.
func (e *MemLittleEndianWriter) WriteString56(p string) {
	e.WriteUint56(uint64(len(p)))
	e.WriteString(p)
}

// WriteString64 Writes the length of the String, using WriteUint64 and then Writes the bytes.
func (e *MemLittleEndianWriter) WriteString64(p string) {
	e.WriteUint64(uint64(len(p)))
	e.WriteString(p)
}

// WriteString0 Writes the bytes of the string ending with a 0 byte.
func (e *MemLittleEndianWriter) WriteString0(p string) {
	e.WriteString(p)
	e.WriteUint8(0)
}
