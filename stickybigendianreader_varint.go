package byteio

// ReadUintX reads an unsigned integer that was encoded using a variable number
// of bytes.
func (e *StickyBigEndianReader) ReadUintX() uint64 {
	c, _ := e.ReadByte()
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
func (e *StickyBigEndianReader) ReadIntX() int64 {
	return unzigzag(e.ReadUintX())
}
