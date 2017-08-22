package byteio

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestWriteLittleEndian(t *testing.T) {
	const testData = "\x0d\x0c\x0b\x0a\x0d\x0c\x08\x07\x06\x05\x04\x03\x02\x01A"
	var b bytes.Buffer
	w := LittleEndianWriter{Writer: &b}

	if n, err := w.WriteUint32(168496141); err != nil {
		t.Errorf("test 1: unexpected error: %s", err)
	} else if n != 4 {
		t.Errorf("test 1: expecting to write 4 bytes, read %d", n)
	}
	if n, err := w.WriteUint16(3085); err != nil {
		t.Errorf("test 2: unexpected error: %s", err)
	} else if n != 2 {
		t.Errorf("test 2: expecting to write 2 bytes, read %d", n)
	}
	if n, err := w.WriteUint64(72623859790382856); err != nil {
		t.Errorf("test 3: unexpected error: %s", err)
	} else if n != 8 {
		t.Errorf("test 3: expecting to write 8 bytes, read %d", n)
	}
	if n, err := w.WriteUint8(65); err != nil {
		t.Errorf("test 4: unexpected error: %s", err)
	} else if n != 1 {
		t.Errorf("test 4: expecting to write 1 byte, read %d", n)
	}
	if str := b.String(); str != testData {
		t.Errorf("expecting %q, got %q", testData, str)
	}
}

func BenchmarkWriteLittleUint8(b *testing.B) {
	w := LittleEndianWriter{Writer: ioutil.Discard}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.WriteUint8(1)
	}
}

func BenchmarkWriteLittleInt8(b *testing.B) {
	w := LittleEndianWriter{Writer: ioutil.Discard}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.WriteInt8(1)
	}
}

func BenchmarkWriteLittleUint16(b *testing.B) {
	w := LittleEndianWriter{Writer: ioutil.Discard}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.WriteUint16(1)
	}
}

func BenchmarkWriteLittleInt16(b *testing.B) {
	w := LittleEndianWriter{Writer: ioutil.Discard}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.WriteInt16(1)
	}
}

func BenchmarkWriteLittleUint32(b *testing.B) {
	w := LittleEndianWriter{Writer: ioutil.Discard}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.WriteUint32(1)
	}
}

func BenchmarkWriteLittleInt32(b *testing.B) {
	w := LittleEndianWriter{Writer: ioutil.Discard}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.WriteInt32(1)
	}
}

func BenchmarkWriteLittleUint64(b *testing.B) {
	w := LittleEndianWriter{Writer: ioutil.Discard}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.WriteUint64(1)
	}
}

func BenchmarkWriteLittleInt64(b *testing.B) {
	w := LittleEndianWriter{Writer: ioutil.Discard}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.WriteInt64(1)
	}
}
