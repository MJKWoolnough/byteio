package byteio

// WriteUintX writes the unsigned integer using a variable number of bytes
func (e *StickyBigEndianWriter) WriteUintX(d uint64) {
	var (
		buf [9]byte
		pos = 8
	)
	if d > 9295997013522923647 {
		buf[8] = byte(d)
		d >>= 8
	} else {
		buf[8] = byte(d & 0x7f)
		d >>= 7
	}
	for ; d > 0; d >>= 7 {
		pos--
		d--
		buf[pos] = byte(d&0x7f) | 0x80
	}
	e.Write(buf[pos:])
}

// WriteIntX writes the integer using a variable number of bytes
func (e *StickyBigEndianWriter) WriteIntX(d int64) {
	e.WriteUintX(zigzag(d))
}
