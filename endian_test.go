package byteio

import (
	"bytes"
	"testing"
)

type readWrite struct {
	readU  func(StickyEndianReader) uint64
	readI  func(StickyEndianReader) int64
	writeU func(StickyEndianWriter, uint64)
	writeI func(StickyEndianWriter, int64)
}

var (
	testUints = [][]uint64{
		{0x00, 0x01, 0x0f, 0x10, 0x1f, 0x20, 0x3f, 0x41, 0x7f, 0x80, 0xff},
		{0x0100, 0x7fff, 0x8000, 0xffff},
		{0x010000, 0x7fffff, 0x800000, 0xffffff},
		{0x01000000, 0x7fffffff, 0x80000000, 0xffffffff},
		{0x0100000000, 0x7fffffffff, 0x8000000000, 0xffffffffff},
		{0x010000000000, 0x7fffffffffff, 0x800000000000, 0xffffffffffff},
		{0x01000000000000, 0x7fffffffffffff, 0x80000000000000, 0xffffffffffffff},
	}
	testInts = [][]int64{
		{-0x80, -0x40, -0x01, 0x00, 0x01, 0x7f, 0x10, 0x1f, 0x20, 0x3f, 0x40, 0x7f},
		{-0x8000, -0x00ff, -0x0081, -0x0080, 0x0100, 0x7fff},
		{-0x800000, -0x00ffff, -0x008001, -0x008000, 0x010000, 0x7fffff},
		{-0x80000000, -0x00ffffff, -0x00800001, -0x00800000, 0x01000000, 0x7fffffff},
		{-0x8000000000, -0x00ffffffff, -0x0080000001, -0x0080000000, 0x0100000000, 0x7fffffffff},
		{-0x800000000000, -0x00ffffffffff, -0x008000000001, -0x008000000000, 0x010000000000, 0x7fffffffffff},
		{-0x80000000000000, -0x00ffffffffffff, -0x00800000000001, -0x00800000000000, 0x01000000000000, 0x7fffffffffffff},
	}
	testBytes = []byte{1, 2, 3, 4, 5, 6, 7, 8}
)

func testReadWrite(t *testing.T, numBytes int, rw readWrite, little, big uint64) {
	t.Helper()

	var buf bytes.Buffer

	for name, test := range map[string]struct {
		StickyEndianReader
		StickyEndianWriter
	}{
		"Little": {
			&StickyReader{Reader: &LittleEndianReader{Reader: &buf}},
			&StickyWriter{Writer: &LittleEndianWriter{Writer: &buf}},
		},
		"Big": {
			&StickyReader{Reader: &BigEndianReader{Reader: &buf}},
			&StickyWriter{Writer: &BigEndianWriter{Writer: &buf}},
		},
		"StickyLittle": {
			&StickyLittleEndianReader{Reader: &buf},
			&StickyLittleEndianWriter{Writer: &buf},
		},
		"StickyBig": {
			&StickyBigEndianReader{Reader: &buf},
			&StickyBigEndianWriter{Writer: &buf},
		},
	} {
		var expectedReadWrite int64

		for n := range numBytes {
			uints := testUints[n]
			ints := testInts[n]

			buf.Reset()

			for i, d := range uints {
				rw.writeU(test.StickyEndianWriter, d)

				if buf.Len() != numBytes {
					t.Errorf("%sUint%d (%d.%d): expected to write %d byte, wrote %d", name, 8*numBytes, n+1, i+1, numBytes, buf.Len())
				} else if got := rw.readU(test.StickyEndianReader); got != d {
					t.Errorf("%sUint%d (%d.%d): wanted %d, got %d", name, 8*numBytes, n+1, i+1, d, got)
				} else if buf.Len() != 0 {
					t.Errorf("%sUint%d (%d.%d): expected to have read all bytes, %d remain", name, 8*numBytes, n+1, i+1, buf.Len())
				}
			}

			expectedReadWrite += int64(len(uints) * numBytes)

			if written := test.StickyEndianWriter.GetCount(); written != expectedReadWrite {
				t.Errorf("%s%d (%d.1): expected to write %d bytes, read %d", name, 8*numBytes, n+1, expectedReadWrite, written)
			} else if read := test.StickyEndianReader.GetCount(); read != expectedReadWrite {
				t.Errorf("%s%d (%d.1): expected to read %d bytes, read %d", name, 8*numBytes, n+1, expectedReadWrite, read)
			}

			buf.Reset()

			for i, d := range ints {
				rw.writeI(test.StickyEndianWriter, d)

				if buf.Len() != numBytes {
					t.Errorf("%sInt%d (%d.%d): expected to write %d byte, wrote %d", name, 8*numBytes, n+1, i+1, numBytes, buf.Len())
				} else if got := rw.readI(test.StickyEndianReader); got != d {
					t.Errorf("%sInt%d (%d.%d): wanted %d, got %d", name, 8*numBytes, n+1, i+1, d, got)
				} else if buf.Len() != 0 {
					t.Errorf("%sInt%d (%d.%d): expected to have read all bytes, %d remain", name, 8*numBytes, n+1, i+1, buf.Len())
				}
			}

			expectedReadWrite += int64(len(ints) * numBytes)

			if written := test.StickyEndianWriter.GetCount(); written != expectedReadWrite {
				t.Errorf("%s%d (%d.2): expected to write %d bytes, read %d", name, 8*numBytes, n+1, expectedReadWrite, written)
			} else if read := test.StickyEndianReader.GetCount(); read != expectedReadWrite {
				t.Errorf("%s%d (%d.2): expected to read %d bytes, read %d", name, 8*numBytes, n+1, expectedReadWrite, read)
			}
		}
	}

	buf.Write(testBytes[:numBytes])

	if read := rw.readU(&StickyLittleEndianReader{Reader: &buf}); read != little {
		t.Errorf("LittleEndian%d: expected to read value %d, got %d", 8*numBytes, little, read)
	}

	buf.Write(testBytes[:numBytes])

	if read := rw.readU(&StickyBigEndianReader{Reader: &buf}); read != big {
		t.Errorf("BigEndian%d: expected to read value %d, got %d", 8*numBytes, big, read)
	}
}

func Test8(t *testing.T) {
	testReadWrite(t, 1, readWrite{
		readU: func(s StickyEndianReader) uint64 {
			return uint64(s.ReadUint8())
		},
		readI: func(s StickyEndianReader) int64 {
			return int64(s.ReadInt8())
		},
		writeU: func(s StickyEndianWriter, n uint64) {
			s.WriteUint8(uint8(n))
		},
		writeI: func(s StickyEndianWriter, n int64) {
			s.WriteInt8(int8(n))
		},
	}, 0x01, 0x01)
}

func Test16(t *testing.T) {
	testReadWrite(t, 2, readWrite{
		readU: func(s StickyEndianReader) uint64 {
			return uint64(s.ReadUint16())
		},
		readI: func(s StickyEndianReader) int64 {
			return int64(s.ReadInt16())
		},
		writeU: func(s StickyEndianWriter, n uint64) {
			s.WriteUint16(uint16(n))
		},
		writeI: func(s StickyEndianWriter, n int64) {
			s.WriteInt16(int16(n))
		},
	}, 0x0201, 0x0102)
}

func Test24(t *testing.T) {
	testReadWrite(t, 3, readWrite{
		readU: func(s StickyEndianReader) uint64 {
			return uint64(s.ReadUint24())
		},
		readI: func(s StickyEndianReader) int64 {
			return int64(s.ReadInt24())
		},
		writeU: func(s StickyEndianWriter, n uint64) {
			s.WriteUint24(uint32(n))
		},
		writeI: func(s StickyEndianWriter, n int64) {
			s.WriteInt24(int32(n))
		},
	}, 0x030201, 0x010203)
}

func Test32(t *testing.T) {
	testReadWrite(t, 4, readWrite{
		readU: func(s StickyEndianReader) uint64 {
			return uint64(s.ReadUint32())
		},
		readI: func(s StickyEndianReader) int64 {
			return int64(s.ReadInt32())
		},
		writeU: func(s StickyEndianWriter, n uint64) {
			s.WriteUint32(uint32(n))
		},
		writeI: func(s StickyEndianWriter, n int64) {
			s.WriteInt32(int32(n))
		},
	}, 0x04030201, 0x01020304)
}

func Test40(t *testing.T) {
	testReadWrite(t, 5, readWrite{
		readU: func(s StickyEndianReader) uint64 {
			return s.ReadUint40()
		},
		readI: func(s StickyEndianReader) int64 {
			return s.ReadInt40()
		},
		writeU: func(s StickyEndianWriter, n uint64) {
			s.WriteUint40(n)
		},
		writeI: func(s StickyEndianWriter, n int64) {
			s.WriteInt40(n)
		},
	}, 0x0504030201, 0x0102030405)
}

func Test48(t *testing.T) {
	testReadWrite(t, 6, readWrite{
		readU: func(s StickyEndianReader) uint64 {
			return s.ReadUint48()
		},
		readI: func(s StickyEndianReader) int64 {
			return s.ReadInt48()
		},
		writeU: func(s StickyEndianWriter, n uint64) {
			s.WriteUint48(n)
		},
		writeI: func(s StickyEndianWriter, n int64) {
			s.WriteInt48(n)
		},
	}, 0x060504030201, 0x010203040506)
}

func Test56(t *testing.T) {
	testReadWrite(t, 7, readWrite{
		readU: func(s StickyEndianReader) uint64 {
			return s.ReadUint56()
		},
		readI: func(s StickyEndianReader) int64 {
			return s.ReadInt56()
		},
		writeU: func(s StickyEndianWriter, n uint64) {
			s.WriteUint56(n)
		},
		writeI: func(s StickyEndianWriter, n int64) {
			s.WriteInt56(n)
		},
	}, 0x07060504030201, 0x01020304050607)
}
