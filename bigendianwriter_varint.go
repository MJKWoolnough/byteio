package byteio

// WriteUintX writes the unsigned integer using a variable number of bytes
func (e *BigEndianWriter) WriteUintX(d uint64) (int, error) {
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
	return e.Write(buf[pos:])
}

// WriteIntX writes the integer using a variable number of bytes
func (e *BigEndianWriter) WriteIntX(d int64) (int, error) {
	return e.WriteUintX(zigzag(d))
}
