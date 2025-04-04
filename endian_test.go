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
		{0, 1, 15, 16, 31, 32, 63, 65, 127, 128, 255},
		{256, 32767, 32768, 65535},
	}
	testInts = [][]int64{
		{-128, -64, -1, 0, 1, 15, 16, 31, 32, 63, 64, 127},
		{-32768, -255, 256, 32767},
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

			for i, n := range uints {
				rw.writeU(test.StickyEndianWriter, n)

				if buf.Len() != numBytes {
					t.Errorf("%sUint (%d.%d): expected to write %d byte, wrote %d", name, n+1, i+1, numBytes, buf.Len())
				} else if got := rw.readU(test.StickyEndianReader); got != n {
					t.Errorf("%sUint (%d.%d): wanted %d, got %d", name, n+1, i+1, n, got)
				} else if buf.Len() != 0 {
					t.Errorf("%sUint (%d.%d): expected to have read all bytes, %d remain", name, n+1, i+1, buf.Len())
				}
			}

			expectedReadWrite += int64(len(uints) * numBytes)

			if written := test.StickyEndianWriter.GetCount(); written != expectedReadWrite {
				t.Errorf("%s (%d.1): expected to write %d bytes, read %d", name, n+1, expectedReadWrite, written)
			} else if read := test.StickyEndianReader.GetCount(); read != expectedReadWrite {
				t.Errorf("%s (%d.1): expected to read %d bytes, read %d", name, n+1, expectedReadWrite, read)
			}

			buf.Reset()

			for i, n := range ints {
				rw.writeI(test.StickyEndianWriter, n)

				if buf.Len() != numBytes {
					t.Errorf("%sInt (%d.%d): expected to write %d byte, wrote %d", name, n+1, i+1, numBytes, buf.Len())
				} else if got := rw.readI(test.StickyEndianReader); got != n {
					t.Errorf("%sInt (%d.%d): wanted %d, got %d", name, n+1, i+1, n, got)
				} else if buf.Len() != 0 {
					t.Errorf("%sInt (%d.%d): expected to have read all bytes, %d remain", name, n+1, i+1, buf.Len())
				}
			}

			expectedReadWrite += int64(len(ints) * numBytes)

			if written := test.StickyEndianWriter.GetCount(); written != expectedReadWrite {
				t.Errorf("%s (%d.2): expected to write %d bytes, read %d", name, n+1, expectedReadWrite, written)
			} else if read := test.StickyEndianReader.GetCount(); read != expectedReadWrite {
				t.Errorf("%s (%d.2): expected to read %d bytes, read %d", name, n+1, expectedReadWrite, read)
			}
		}
	}

	buf.Write(testBytes[:numBytes])

	if read := rw.readU(&StickyLittleEndianReader{Reader: &buf}); read != little {
		t.Errorf("LittleEndian: expected to read value %d, got %d", little, read)
	}

	buf.Write(testBytes[:numBytes])

	if read := rw.readU(&StickyBigEndianReader{Reader: &buf}); read != big {
		t.Errorf("BigEndian: expected to read value %d, got %d", big, read)
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
	}, 1, 1)
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
	}, 513, 258)
}
