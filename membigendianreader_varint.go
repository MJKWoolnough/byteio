package byteio

// ReadUintX reads an unsigned integer that was encoded using a variable number
// of bytes.
func (e *MemBigEndianReader) ReadUintX() uint64 {
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
func (e *MemBigEndianReader) ReadIntX() int64 {
	return unzigzag(e.ReadUintX())
}
