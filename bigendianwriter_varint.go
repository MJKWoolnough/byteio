package byteio

const maxBENineByte = 0x80<<56 | 0x80<<49 | 0x80<<42 | 0x80<<35 | 0x80<<28 | 0x80<<21 | 0x80<<14 | 0x80<<7 | 0x7f

// WriteUintX writes the unsigned integer using a variable number of bytes
func (e *BigEndianWriter) WriteUintX(d uint64) (int, error) {
	pos := 8
	if d > maxBENineByte {
		e.buffer[8] = byte(d)
		d >>= 8
	} else {
		e.buffer[8] = byte(d & 0x7f)
		d >>= 7
	}
	for ; d > 0; d >>= 7 {
		pos--
		d--
		e.buffer[pos] = byte(d&0x7f) | 0x80
	}
	return e.Write(e.buffer[pos:])
}

// WriteIntX writes the integer using a variable number of bytes
func (e *BigEndianWriter) WriteIntX(d int64) (int, error) {
	return e.WriteUintX(zigzag(d))
}
