package byteio

import (
	"bytes"
	"testing"
)

func TestLittleEndianVarUint(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0, 9))
	for n, test := range [...]struct {
		Input  uint64
		Output []byte
	}{
		{0, []byte{0}},
		{1, []byte{1}},
		{127, []byte{127}},
		{128, []byte{128, 0}},
		{129, []byte{129, 0}},
		{163, []byte{163, 0}},
		{255, []byte{255, 0}},
		{256, []byte{128, 1}},
		{257, []byte{129, 1}},
		{300, []byte{172, 1}},
		{512, []byte{128, 3}},
		{1024, []byte{128, 7}},
		{2048, []byte{128, 15}},
		{4096, []byte{128, 31}},
		{8192, []byte{128, 63}},
		{16384, []byte{128, 127}},
		{16511, []byte{255, 127}},
		{16512, []byte{128, 128, 0}},
		{32768, []byte{128, 255, 0}},
		{65536, []byte{128, 255, 2}},
		{131072, []byte{128, 255, 6}},
		{262144, []byte{128, 255, 14}},
		{524288, []byte{128, 255, 30}},
		{1048576, []byte{128, 255, 62}},
		{2097152, []byte{128, 255, 126}},
		{134217728, []byte{128, 255, 254, 62}},
		{268435456, []byte{128, 255, 254, 126}},
		{17179869184, []byte{128, 255, 254, 254, 62}},
		{18446744073709551615, []byte{255, 254, 254, 254, 254, 254, 254, 254, 254}},
	} {
		buf.Reset()
		(&LittleEndianWriter{Writer: buf}).WriteUintX(test.Input)
		if !bytes.Equal(buf.Bytes(), test.Output) {
			t.Errorf("test %d.1: expecting %v, got %v", n+1, test.Output, buf.Bytes())
			buf.Reset()
			buf.Write(test.Output)
		}
		num, _, _ := (&LittleEndianReader{Reader: buf}).ReadUintX()
		if num != test.Input {
			t.Errorf("test %d.2: expecting %d, got %d", n+1, test.Input, num)
		}
	}
}

func TestStickyLittleEndianVarUint(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0, 9))
	for n, test := range [...]struct {
		Input  uint64
		Output []byte
	}{
		{0, []byte{0}},
		{1, []byte{1}},
		{127, []byte{127}},
		{128, []byte{128, 0}},
		{129, []byte{129, 0}},
		{163, []byte{163, 0}},
		{255, []byte{255, 0}},
		{256, []byte{128, 1}},
		{257, []byte{129, 1}},
		{300, []byte{172, 1}},
		{512, []byte{128, 3}},
		{1024, []byte{128, 7}},
		{2048, []byte{128, 15}},
		{4096, []byte{128, 31}},
		{8192, []byte{128, 63}},
		{16384, []byte{128, 127}},
		{16511, []byte{255, 127}},
		{16512, []byte{128, 128, 0}},
		{32768, []byte{128, 255, 0}},
		{65536, []byte{128, 255, 2}},
		{131072, []byte{128, 255, 6}},
		{262144, []byte{128, 255, 14}},
		{524288, []byte{128, 255, 30}},
		{1048576, []byte{128, 255, 62}},
		{2097152, []byte{128, 255, 126}},
		{134217728, []byte{128, 255, 254, 62}},
		{268435456, []byte{128, 255, 254, 126}},
		{17179869184, []byte{128, 255, 254, 254, 62}},
		{18446744073709551615, []byte{255, 254, 254, 254, 254, 254, 254, 254, 254}},
	} {
		buf.Reset()
		(&StickyLittleEndianWriter{Writer: buf}).WriteUintX(test.Input)
		if !bytes.Equal(buf.Bytes(), test.Output) {
			t.Errorf("test %d.1: expecting %v, got %v", n+1, test.Output, buf)
			buf.Reset()
			buf.Write(test.Output)
		}
		num := (&StickyLittleEndianReader{Reader: buf}).ReadUintX()
		if num != test.Input {
			t.Errorf("test %d.2: expecting %d, got %d", n+1, test.Input, num)
		}
	}
}

func FuzzLittleEndianVarInt(f *testing.F) {
	for _, seed := range []uint64{0, 127, 128, 255, 256, 1024, 16512, 134217728, 18446744073709551615} {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, in uint64) {
		var buf bytes.Buffer
		(&LittleEndianWriter{Writer: &buf}).WriteUintX(in)
		num, _, _ := (&LittleEndianReader{Reader: &buf}).ReadUintX()

		if num != in {
			t.Errorf("failed on %d", in)
		}
	})
}

func TestBigEndianVarUint(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0, 9))
	for n, test := range [...]struct {
		Input  uint64
		Output []byte
	}{
		{0, []byte{0}},
		{1, []byte{1}},
		{127, []byte{127}},
		{128, []byte{128, 0}},
		{129, []byte{128, 1}},
		{130, []byte{128, 2}},
		{255, []byte{128, 127}},
		{256, []byte{129, 0}},
		{257, []byte{129, 1}},
		{512, []byte{131, 0}},
		{1024, []byte{135, 0}},
		{2048, []byte{143, 0}},
		{4096, []byte{159, 0}},
		{8192, []byte{191, 0}},
		{16384, []byte{255, 0}},
		{16385, []byte{255, 1}},
		{16511, []byte{255, 127}},
		{16512, []byte{128, 128, 0}},
		{32768, []byte{128, 255, 0}},
		{65536, []byte{130, 255, 0}},
		{131072, []byte{134, 255, 0}},
		{262144, []byte{142, 255, 0}},
		{2113663, []byte{255, 255, 127}},
		{2113664, []byte{128, 128, 128, 0}},
		{18446744073709551615, []byte{254, 254, 254, 254, 254, 254, 254, 254, 255}},
		{9295997013522923646, []byte{255, 255, 255, 255, 255, 255, 255, 255, 126}},
		{9295997013522923647, []byte{255, 255, 255, 255, 255, 255, 255, 255, 127}},
		{9295997013522923648, []byte{191, 191, 191, 191, 191, 191, 191, 191, 128}},
		{9295997013522923649, []byte{191, 191, 191, 191, 191, 191, 191, 191, 129}},
	} {
		buf.Reset()
		(&BigEndianWriter{Writer: buf}).WriteUintX(test.Input)
		if !bytes.Equal(buf.Bytes(), test.Output) {
			t.Errorf("test %d.1: expecting %v, got %v", n+1, test.Output, buf)
			buf.Reset()
			buf.Write(test.Output)
		}
		num, _, _ := (&BigEndianReader{Reader: buf}).ReadUintX()
		if num != test.Input {
			t.Errorf("test %d.2: expecting %d, got %d", n+1, test.Input, num)
		}
	}
}

func FuzzBigEndianVarInt(f *testing.F) {
	for _, seed := range []uint64{0, 127, 128, 255, 256, 1024, 16512, 134217728, 18446744073709551615} {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, in uint64) {
		var buf bytes.Buffer
		(&BigEndianWriter{Writer: &buf}).WriteUintX(in)
		num, _, _ := (&BigEndianReader{Reader: &buf}).ReadUintX()

		if num != in {
			t.Errorf("failed on %d", in)
		}
	})
}
