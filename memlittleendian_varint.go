package byteio

// WriteUintX writes the unsigned integer using a variable number of bytes.
func (e *MemLittleEndian) WriteUintX(d uint64) {
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
func (e *MemLittleEndian) WriteIntX(d int64) {
	e.WriteUintX(zigzag(d))
}

// ReadUintX reads an unsigned integer that was encoded using a variable number
// of bytes.
func (e *MemLittleEndian) ReadUintX() uint64 {
	var (
		n   int
		val uint64
	)

	for n < 9 {
		c := e.ReadUint8()
		val += uint64(c&0xff) << uint(n*7)
		n++

		if c&0x80 == 0 {
			break
		}
	}

	return val
}

// ReadIntX reads an integer that was encoded using a variable number of bytes.
func (e *MemLittleEndian) ReadIntX() int64 {
	return unzigzag(e.ReadUintX())
}
