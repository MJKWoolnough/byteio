package byteio

// WriteUintX writes the unsigned integer using a variable number of bytes.
func (e *MemBigEndian) WriteUintX(d uint64) {
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
func (e *MemBigEndian) WriteIntX(d int64) {
	e.WriteUintX(zigzag(d))
}

// ReadUintX reads an unsigned integer that was encoded using a variable number
// of bytes.
func (e *MemBigEndian) ReadUintX() uint64 {
	c := e.ReadUint8()
	val := uint64(c) & 0x7f

	for n := 1; c&0x80 != 0 && n < 9; n++ {
		c = e.ReadUint8()
		val++

		if n == 8 && c&0x80 > 0 {
			val = (val << 8) | uint64(c)
		} else {
			val = (val << 7) | uint64(c&0x7f)
		}
	}

	return val
}

// ReadIntX reads an integer that was encoded using a variable number of bytes.
func (e *MemBigEndian) ReadIntX() int64 {
	return unzigzag(e.ReadUintX())
}
