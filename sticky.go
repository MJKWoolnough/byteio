package byteio

type StickyReader struct {
	Reader EndianReader
	Err    error
	Count  int64
}

func (s *StickyReader) Read(b []byte) {
	if s.Err != nil {
		return
	}
	n, err := s.Reader.Read(b)
	s.Err = err
	s.Count += int64(n)
}

func (s *StickyReader) ReadUint8() uint8 {
	if s.Err != nil {
		return 0
	}
	i, n, err := s.Reader.ReadUint8()
	s.Err = err
	s.Count += int64(n)
	return i
}

func (s *StickyReader) ReadInt8() int8 {
	if s.Err != nil {
		return 0
	}
	i, n, err := s.Reader.ReadInt8()
	s.Err = err
	s.Count += int64(n)
	return i
}

func (s *StickyReader) ReadUint16() uint16 {
	if s.Err != nil {
		return 0
	}
	i, n, err := s.Reader.ReadUint16()
	s.Err = err
	s.Count += int64(n)
	return i
}

func (s *StickyReader) ReadInt16() int16 {
	if s.Err != nil {
		return 0
	}
	i, n, err := s.Reader.ReadInt16()
	s.Err = err
	s.Count += int64(n)
	return i
}

func (s *StickyReader) ReadUint32() uint32 {
	if s.Err != nil {
		return 0
	}
	i, n, err := s.Reader.ReadUint32()
	s.Err = err
	s.Count += int64(n)
	return i
}

func (s *StickyReader) ReadInt32() int32 {
	if s.Err != nil {
		return 0
	}
	i, n, err := s.Reader.ReadInt32()
	s.Err = err
	s.Count += int64(n)
	return i
}

func (s *StickyReader) ReadUint64() uint64 {
	if s.Err != nil {
		return 0
	}
	i, n, err := s.Reader.ReadUint64()
	s.Err = err
	s.Count += int64(n)
	return i
}

func (s *StickyReader) ReadInt64() int64 {
	if s.Err != nil {
		return 0
	}
	i, n, err := s.Reader.ReadInt64()
	s.Err = err
	s.Count += int64(n)
	return i
}

func (s *StickyReader) ReadFloat32() float32 {
	if s.Err != nil {
		return 0
	}
	i, n, err := s.Reader.ReadFloat32()
	s.Err = err
	s.Count += int64(n)
	return i
}

func (s *StickyReader) ReadFloat64() float64 {
	if s.Err != nil {
		return 0
	}
	i, n, err := s.Reader.ReadFloat64()
	s.Err = err
	s.Count += int64(n)
	return i
}

type StickyWriter struct {
	Writer EndianWriter
	Err    error
	Count  int64
}

func (s *StickyWriter) Write(p []byte) {
	if s.Err != nil {
		return
	}
	n, err := s.Writer.Write(p)
	s.Err = err
	s.Count += int64(n)
}

func (s *StickyWriter) WriteUint8(i uint8) {
	if s.Err != nil {
		return
	}
	n, err := s.Writer.WriteUint8(i)
	s.Err = err
	s.Count += int64(n)
}

func (s *StickyWriter) WriteInt8(i int8) {
	if s.Err != nil {
		return
	}
	n, err := s.Writer.WriteInt8(i)
	s.Err = err
	s.Count += int64(n)
}

func (s *StickyWriter) WriteUint16(i uint16) {
	if s.Err != nil {
		return
	}
	n, err := s.Writer.WriteUint16(i)
	s.Err = err
	s.Count += int64(n)
}

func (s *StickyWriter) WriteInt16(i int16) {
	if s.Err != nil {
		return
	}
	n, err := s.Writer.WriteInt16(i)
	s.Err = err
	s.Count += int64(n)
}

func (s *StickyWriter) WriteUint32(i uint32) {
	if s.Err != nil {
		return
	}
	n, err := s.Writer.WriteUint32(i)
	s.Err = err
	s.Count += int64(n)
}

func (s *StickyWriter) WriteInt32(i int32) {
	if s.Err != nil {
		return
	}
	n, err := s.Writer.WriteInt32(i)
	s.Err = err
	s.Count += int64(n)
}

func (s *StickyWriter) WriteUint64(i uint64) {
	if s.Err != nil {
		return
	}
	n, err := s.Writer.WriteUint64(i)
	s.Err = err
	s.Count += int64(n)
}

func (s *StickyWriter) WriteInt64(i int64) {
	if s.Err != nil {
		return
	}
	n, err := s.Writer.WriteInt64(i)
	s.Err = err
	s.Count += int64(n)
}

func (s *StickyWriter) WriteFloat32(i float32) {
	if s.Err != nil {
		return
	}
	n, err := s.Writer.WriteFloat32(i)
	s.Err = err
	s.Count += int64(n)
}

func (s *StickyWriter) WriteFloat64(i float64) {
	if s.Err != nil {
		return
	}
	n, err := s.Writer.WriteFloat64(i)
	s.Err = err
	s.Count += int64(n)
}
