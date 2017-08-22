package byteio

import (
	"strings"
	"testing"
)

func TestReadLittleEndian(t *testing.T) {
	const testData = "\x0d\x0c\x0b\x0a\x0d\x0c\x08\x07\x06\x05\x04\x03\x02\x01A"
	r := LittleEndianReader{Reader: strings.NewReader(testData)}
	if i, n, err := r.ReadUint32(); err != nil {
		t.Errorf("test 1: unexpected error: %s", err)
	} else if n != 4 {
		t.Errorf("test 1: expecting to read 4 bytes, read %d", n)
	} else if i != 168496141 {
		t.Errorf("test 1: expecting 168496141, got %d", i)
	}
	if i, n, err := r.ReadUint16(); err != nil {
		t.Errorf("test 2: unexpected error: %s", err)
	} else if n != 2 {
		t.Errorf("test 2: expecting to read 2 bytes, read %d", n)
	} else if i != 3085 {
		t.Errorf("test 2: expecting 3085, got %d", i)
	}
	if i, n, err := r.ReadUint64(); err != nil {
		t.Errorf("test 3: unexpected error: %s", err)
	} else if n != 8 {
		t.Errorf("test 3: expecting to read 8 bytes, read %d", n)
	} else if i != 72623859790382856 {
		t.Errorf("test 3: expecting 72623859790382856 got %d", i)
	}
	if i, n, err := r.ReadUint8(); err != nil {
		t.Errorf("test 4: unexpected error: %s", err)
	} else if n != 1 {
		t.Errorf("test 4: expecting to read 1 byte, read %d", n)
	} else if i != 65 {
		t.Errorf("test 4: expecting 65, got %d", i)
	}
}

type neverending []byte

func (n neverending) Read(b []byte) (int, error) {
	return copy(b, n), nil
}

func BenchmarkReadLittleEndianUint8(b *testing.B) {
	r := LittleEndianReader{Reader: neverending{1}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = r.ReadUint8()
	}
}

func BenchmarkReadLittleEndianInt8(b *testing.B) {
	r := LittleEndianReader{Reader: neverending{1}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = r.ReadInt8()
	}
}
func BenchmarkReadLittleEndianUint16(b *testing.B) {
	r := LittleEndianReader{Reader: neverending{1, 1}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = r.ReadUint16()
	}
}

func BenchmarkReadLittleEndianInt16(b *testing.B) {
	r := LittleEndianReader{Reader: neverending{1, 1}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = r.ReadInt16()
	}
}

func BenchmarkReadLittleEndianUint32(b *testing.B) {
	r := LittleEndianReader{Reader: neverending{1, 1, 1, 1}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = r.ReadUint32()
	}
}

func BenchmarkReadLittleEndianInt32(b *testing.B) {
	r := LittleEndianReader{Reader: neverending{1, 1, 1, 1}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = r.ReadInt32()
	}
}

func BenchmarkReadLittleEndianUint64(b *testing.B) {
	r := LittleEndianReader{Reader: neverending{1, 1, 1, 1, 1, 1, 1, 1}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = r.ReadUint64()
	}
}

func BenchmarkReadLittleEndianInt64(b *testing.B) {
	r := LittleEndianReader{Reader: neverending{1, 1, 1, 1, 1, 1, 1, 1}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = r.ReadInt64()
	}
}
