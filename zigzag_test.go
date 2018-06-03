package byteio

import "testing"

func TestZigZag(t *testing.T) {
	for n, test := range [...]struct {
		In  int64
		Out uint64
	}{
		{0, 0},
		{-1, 1},
		{1, 2},
		{-2, 3},
		{2, 4},
		{-9223372036854775807, 18446744073709551613},
		{9223372036854775807, 18446744073709551614},
	} {
		if i := zigzag(test.In); i != test.Out {
			t.Errorf("test %d.1: from %d expecting %d, got %d", n+1, test.In, test.Out, i)
		}
		if j := unzigzag(test.Out); j != test.In {
			t.Errorf("test %d.2: from %d expecting %d, got %d", n+1, test.Out, test.In, j)
		}
	}
}
