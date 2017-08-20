package byteio

import (
	"strings"
	"testing"
)

func TestReadBigEndian(t *testing.T) {
	const testData = "\x0a\x0b\x0c\x0d\x0c\x0d\x01\x02\x03\x04\x05\x06\x07\x08A"
	r := BigEndianReader{Reader: strings.NewReader(testData)}
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

func BenchmarkReadBigEndianUint8(b *testing.B) {
	r := BigEndianReader{Reader: neverending{1}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = r.ReadUint8()
	}
}

func BenchmarkReadBigEndianInt8(b *testing.B) {
	r := BigEndianReader{Reader: neverending{1}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = r.ReadInt8()
	}
}
func BenchmarkReadBigEndianUint16(b *testing.B) {
	r := BigEndianReader{Reader: neverending{1, 1}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = r.ReadUint16()
	}
}

func BenchmarkReadBigEndianInt16(b *testing.B) {
	r := BigEndianReader{Reader: neverending{1, 1}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = r.ReadInt16()
	}
}

func BenchmarkReadBigEndianUint32(b *testing.B) {
	r := BigEndianReader{Reader: neverending{1, 1, 1, 1}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = r.ReadUint32()
	}
}

func BenchmarkReadBigEndianInt32(b *testing.B) {
	r := BigEndianReader{Reader: neverending{1, 1, 1, 1}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = r.ReadInt32()
	}
}

func BenchmarkReadBigEndianUint64(b *testing.B) {
	r := BigEndianReader{Reader: neverending{1, 1, 1, 1, 1, 1, 1, 1}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = r.ReadUint64()
	}
}

func BenchmarkReadBigEndianInt64(b *testing.B) {
	r := BigEndianReader{Reader: neverending{1, 1, 1, 1, 1, 1, 1, 1}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = r.ReadInt64()
	}
}
