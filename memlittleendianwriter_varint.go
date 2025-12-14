package byteio

// WriteUintX writes the unsigned integer using a variable number of bytes.
func (e *MemLittleEndianWriter) WriteUintX(d uint64) {
	var pos int

	var buffer [9]byte

	for ; d > 127 && pos < 8; pos++ {
		buffer[pos] = byte(d&0x7f) | 0x80
		d >>= 7
		d--
	}

	buffer[pos] = byte(d)
	e.WriteBytes(buffer[:pos+1])
}

// WriteIntX writes the integer using a variable number of bytes.
func (e *MemLittleEndianWriter) WriteIntX(d int64) {
	e.WriteUintX(zigzag(d))
}
