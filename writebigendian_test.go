package byteio

import (
	"bytes"
	"testing"
)

func TestWriteBigEndian(t *testing.T) {
	const testData = "\x0a\x0b\x0c\x0d\x0c\x0d\x01\x02\x03\x04\x05\x06\x07\x08A"
	var b bytes.Buffer
	w := BigEndianWriter{Writer: &b}

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