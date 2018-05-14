package byteio

// WriteUintX writes the unsigned integer using a variable number of bytes
func (e *StickyLittleEndianWriter) WriteUintX(d uint64) {
	var (
		buf [9]byte
		pos int
	)
	for ; d > 127 && pos < 8; pos++ {
		buf[pos] = byte(d&0x7f) | 0x80
		d >>= 7
		d--
	}
	buf[pos] = byte(d)
	e.Write(buf[:pos+1])
}

// WriteIntX writes the integer using a variable number of bytes
func (e *StickyLittleEndianWriter) WriteIntX(d int64) {
	e.WriteUintX(zigzag(d))
}
