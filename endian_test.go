package byteio

import (
	"iter"
	"math"
	"testing"
)

type readWrite[T any] struct {
	read  func(StickyEndianReader) T
	write func(StickyEndianWriter, T)
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
		{0x0100000000000000, 0x7fffffffffffffff, 0x8000000000000000, 0xffffffffffffffff},
	}
	testInts = [][]int64{
		{-0x80, -0x40, -0x01, 0x00, 0x01, 0x7f, 0x10, 0x1f, 0x20, 0x3f, 0x40, 0x7f},
		{-0x8000, -0x00ff, -0x0081, -0x0080, 0x0100, 0x7fff},
		{-0x800000, -0x00ffff, -0x008001, -0x008000, 0x010000, 0x7fffff},
		{-0x80000000, -0x00ffffff, -0x00800001, -0x00800000, 0x01000000, 0x7fffffff},
		{-0x8000000000, -0x00ffffffff, -0x0080000001, -0x0080000000, 0x0100000000, 0x7fffffffff},
		{-0x800000000000, -0x00ffffffffff, -0x008000000001, -0x008000000000, 0x010000000000, 0x7fffffffffff},
		{-0x80000000000000, -0x00ffffffffffff, -0x00800000000001, -0x00800000000000, 0x01000000000000, 0x7fffffffffffff},
		{-0x8000000000000000, -0x00ffffffffffffff, -0x0080000000000001, -0x0080000000000000, 0x0100000000000000, 0x7fffffffffffffff},
	}
	testFloats = [][]float64{
		{0, 1, float64(math.Float32frombits(0x01)), float64(math.Float32frombits(0xff)), float64(math.Float32frombits(0xffffff)), float64(math.Float32frombits(0x80000000))},
		{math.Float64frombits(0xffffffffffffff), math.Float64frombits(0x8000000000000000), math.Float64frombits(0x0101010101010101)},
	}
	testBytes = []byte{1, 2, 3, 4, 5, 6, 7, 8}
)

type testStickyReadWrite struct {
	StickyEndianReader
	StickyEndianWriter
}

func testEndians(buf *MemLittleEndian) iter.Seq2[string, testStickyReadWrite] {
	return func(yield func(string, testStickyReadWrite) bool) {
		_ = yield("Little", testStickyReadWrite{
			&StickyReader{Reader: &LittleEndianReader{Reader: buf}},
			&StickyWriter{Writer: &LittleEndianWriter{Writer: buf}},
		}) && yield("Big", testStickyReadWrite{
			&StickyReader{Reader: &BigEndianReader{Reader: buf}},
			&StickyWriter{Writer: &BigEndianWriter{Writer: buf}},
		}) && yield("StickyLittle", testStickyReadWrite{
			&StickyLittleEndianReader{Reader: buf},
			&StickyLittleEndianWriter{Writer: buf},
		}) && yield("StickyBig", testStickyReadWrite{
			&StickyBigEndianReader{Reader: buf},
			&StickyBigEndianWriter{Writer: buf},
		}) && yield("MemBig", testStickyReadWrite{
			(*MemBigEndian)(buf),
			(*MemBigEndian)(buf),
		}) && yield("MemLittle", testStickyReadWrite{
			buf,
			buf,
		})
	}
}

func isMem(t testStickyReadWrite) bool {
	if _, ok := t.StickyEndianReader.(*MemLittleEndian); ok {
		return true
	}

	_, ok := t.StickyEndianReader.(*MemBigEndian)

	return ok
}

func testReadWrite[T comparable](t *testing.T, tname string, numBytes int, rw readWrite[T], tests [][]T, expectedLittle, expectedBig T) {
	t.Helper()

	var buf MemLittleEndian

	for name, test := range testEndians(&buf) {
		var expectedReadWrite int64

		for n, tests := range tests {
			buf = buf[:0]

			for i, d := range tests {
				rw.write(test.StickyEndianWriter, d)

				if len(buf) != numBytes {
					t.Errorf("%s%s%d (%d.%d): expected to write %d byte, wrote %d", name, tname, 8*numBytes, n+1, i+1, numBytes, len(buf))
				} else if got := rw.read(test.StickyEndianReader); got != d {
					t.Errorf("%s%s%d (%d.%d): wanted %v, got %v", name, tname, 8*numBytes, n+1, i+1, d, got)
				} else if len(buf) != 0 {
					t.Errorf("%s%s%d (%d.%d): expected to have read all bytes, %d remain", name, tname, 8*numBytes, n+1, i+1, len(buf))
				}
			}

			expectedReadWrite += int64(len(tests) * numBytes)

			if written := test.StickyEndianWriter.GetCount(); written != expectedReadWrite && !isMem(test) {
				t.Errorf("%s%s%d (%d.1): expected to write %d bytes, read %d", name, tname, 8*numBytes, n+1, expectedReadWrite, written)
			} else if read := test.StickyEndianReader.GetCount(); read != expectedReadWrite && !isMem(test) {
				t.Errorf("%s%s%d (%d.1): expected to read %d bytes, read %d", name, tname, 8*numBytes, n+1, expectedReadWrite, read)
			}
		}
	}

	buf.Write(testBytes[:numBytes])

	if read := rw.read(&StickyLittleEndianReader{Reader: &buf}); read != expectedLittle {
		t.Errorf("LittleEndian%s%d: expected to read value %v, got %v", tname, 8*numBytes, expectedLittle, read)
	}

	buf.Write(testBytes[:numBytes])

	if read := rw.read(&StickyBigEndianReader{Reader: &buf}); read != expectedBig {
		t.Errorf("BigEndian%s%d: expected to read value %v, got %v", tname, 8*numBytes, expectedBig, read)
	}
}

func Test8(t *testing.T) {
	testReadWrite(t, "Uint", 1, readWrite[uint64]{
		read: func(s StickyEndianReader) uint64 {
			return uint64(s.ReadUint8())
		},
		write: func(s StickyEndianWriter, n uint64) {
			s.WriteUint8(uint8(n))
		},
	}, testUints[:1], 0x01, 0x01)

	testReadWrite(t, "Int", 1, readWrite[int64]{
		read: func(s StickyEndianReader) int64 {
			return int64(s.ReadInt8())
		},
		write: func(s StickyEndianWriter, n int64) {
			s.WriteInt8(int8(n))
		},
	}, testInts[:1], 0x01, 0x01)
}

func Test16(t *testing.T) {
	testReadWrite(t, "Uint", 2, readWrite[uint64]{
		read: func(s StickyEndianReader) uint64 {
			return uint64(s.ReadUint16())
		},
		write: func(s StickyEndianWriter, n uint64) {
			s.WriteUint16(uint16(n))
		},
	}, testUints[:2], 0x0201, 0x0102)
	testReadWrite(t, "Int", 2, readWrite[int64]{
		read: func(s StickyEndianReader) int64 {
			return int64(s.ReadInt16())
		},
		write: func(s StickyEndianWriter, n int64) {
			s.WriteInt16(int16(n))
		},
	}, testInts[:2], 0x0201, 0x0102)
}

func Test24(t *testing.T) {
	testReadWrite(t, "Uint", 3, readWrite[uint64]{
		read: func(s StickyEndianReader) uint64 {
			return uint64(s.ReadUint24())
		},
		write: func(s StickyEndianWriter, n uint64) {
			s.WriteUint24(uint32(n))
		},
	}, testUints[:3], 0x030201, 0x010203)
	testReadWrite(t, "Ints", 3, readWrite[int64]{
		read: func(s StickyEndianReader) int64 {
			return int64(s.ReadInt24())
		},
		write: func(s StickyEndianWriter, n int64) {
			s.WriteInt24(int32(n))
		},
	}, testInts[:3], 0x030201, 0x010203)
}

func Test32(t *testing.T) {
	testReadWrite(t, "Uint", 4, readWrite[uint64]{
		read: func(s StickyEndianReader) uint64 {
			return uint64(s.ReadUint32())
		},
		write: func(s StickyEndianWriter, n uint64) {
			s.WriteUint32(uint32(n))
		},
	}, testUints[:4], 0x04030201, 0x01020304)
	testReadWrite(t, "Int", 4, readWrite[int64]{
		read: func(s StickyEndianReader) int64 {
			return int64(s.ReadInt32())
		},
		write: func(s StickyEndianWriter, n int64) {
			s.WriteInt32(int32(n))
		},
	}, testInts[:4], 0x04030201, 0x01020304)
}

func Test40(t *testing.T) {
	testReadWrite(t, "Uint", 5, readWrite[uint64]{
		read: func(s StickyEndianReader) uint64 {
			return s.ReadUint40()
		},
		write: func(s StickyEndianWriter, n uint64) {
			s.WriteUint40(n)
		},
	}, testUints[:5], 0x0504030201, 0x0102030405)
	testReadWrite(t, "Int", 5, readWrite[int64]{
		read: func(s StickyEndianReader) int64 {
			return s.ReadInt40()
		},
		write: func(s StickyEndianWriter, n int64) {
			s.WriteInt40(n)
		},
	}, testInts[:5], 0x0504030201, 0x0102030405)
}

func Test48(t *testing.T) {
	testReadWrite(t, "Uint", 6, readWrite[uint64]{
		read: func(s StickyEndianReader) uint64 {
			return s.ReadUint48()
		},
		write: func(s StickyEndianWriter, n uint64) {
			s.WriteUint48(n)
		},
	}, testUints[:6], 0x060504030201, 0x010203040506)
	testReadWrite(t, "Int", 6, readWrite[int64]{
		read: func(s StickyEndianReader) int64 {
			return s.ReadInt48()
		},
		write: func(s StickyEndianWriter, n int64) {
			s.WriteInt48(n)
		},
	}, testInts[:6], 0x060504030201, 0x010203040506)
}

func Test56(t *testing.T) {
	testReadWrite(t, "Uint", 7, readWrite[uint64]{
		read: func(s StickyEndianReader) uint64 {
			return s.ReadUint56()
		},
		write: func(s StickyEndianWriter, n uint64) {
			s.WriteUint56(n)
		},
	}, testUints[:7], 0x07060504030201, 0x01020304050607)
	testReadWrite(t, "Int", 7, readWrite[int64]{
		read: func(s StickyEndianReader) int64 {
			return s.ReadInt56()
		},
		write: func(s StickyEndianWriter, n int64) {
			s.WriteInt56(n)
		},
	}, testInts[:7], 0x07060504030201, 0x01020304050607)
}

func Test64(t *testing.T) {
	testReadWrite(t, "Uint", 8, readWrite[uint64]{
		read: func(s StickyEndianReader) uint64 {
			return s.ReadUint64()
		},
		write: func(s StickyEndianWriter, n uint64) {
			s.WriteUint64(n)
		},
	}, testUints, 0x0807060504030201, 0x0102030405060708)
	testReadWrite(t, "Int", 8, readWrite[int64]{
		read: func(s StickyEndianReader) int64 {
			return s.ReadInt64()
		},
		write: func(s StickyEndianWriter, n int64) {
			s.WriteInt64(n)
		},
	}, testInts, 0x0807060504030201, 0x0102030405060708)
}

func TestFloat32(t *testing.T) {
	testReadWrite(t, "Float", 4, readWrite[float64]{
		read: func(s StickyEndianReader) float64 {
			return float64(s.ReadFloat32())
		},
		write: func(s StickyEndianWriter, n float64) {
			s.WriteFloat32(float32(n))
		},
	}, testFloats[:1], float64(math.Float32frombits(0x04030201)), float64(math.Float32frombits(0x01020304)))
}

func TestFloat64(t *testing.T) {
	testReadWrite(t, "Float", 8, readWrite[float64]{
		read: func(s StickyEndianReader) float64 {
			return s.ReadFloat64()
		},
		write: func(s StickyEndianWriter, n float64) {
			s.WriteFloat64(n)
		},
	}, testFloats, math.Float64frombits(0x0807060504030201), math.Float64frombits(0x0102030405060708))
}

func TestBool(t *testing.T) {
	testReadWrite(t, "Bool", 1, readWrite[bool]{
		read: func(s StickyEndianReader) bool {
			return s.ReadBool()
		},
		write: func(s StickyEndianWriter, n bool) {
			s.WriteBool(n)
		},
	}, [][]bool{{true, false}}, true, true)
}
