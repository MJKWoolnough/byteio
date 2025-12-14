package byteio

// WriteUintX writes the unsigned integer using a variable number of bytes.
func (e *MemBigEndianWriter) WriteUintX(d uint64) {
	pos := 8
	var buffer [9]byte

	if d > 9295997013522923647 {
		buffer[8] = byte(d)
		d >>= 8
	} else {
		buffer[8] = byte(d & 0x7f)
		d >>= 7
	}

	for ; d > 0; d >>= 7 {
		pos--
		d--
		buffer[pos] = byte(d&0x7f) | 0x80
	}

	e.WriteBytes(buffer[pos:])
}

// WriteIntX writes the integer using a variable number of bytes.
func (e *MemBigEndianWriter) WriteIntX(d int64) {
	e.WriteUintX(zigzag(d))
}
