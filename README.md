# byteio

[![CI](https://github.com/MJKWoolnough/byteio/actions/workflows/go-checks.yml/badge.svg)](https://github.com/MJKWoolnough/byteio/actions)
[![Go Reference](https://pkg.go.dev/badge/vimagination.zapto.org/byteio.svg)](https://pkg.go.dev/vimagination.zapto.org/byteio)
[![Go Report Card](https://goreportcard.com/badge/vimagination.zapto.org/byteio)](https://goreportcard.com/report/vimagination.zapto.org/byteio)

--
    import "vimagination.zapto.org/byteio"

Package byteio helps with writing binary data in both big and little endian.

## Highlights

 - Read and Write:
   - Integers of arbitrary widths (8->64 bits).
   - Variable-length encoded integers.
   - Float32s and float64s.
   - Booleans.
   - Length prefixed string and byte slices.
 - Read and Write float32 and float64 types.
 - Read and Write in both **big-** and **little-endian** formats.
 - Read and Write Variable Integer formats to save on space.
 - Sticky Reads and Writers that store errors and read/write counts.
 - Read and Write directly from and to byte slices with the MemLittleEndian and MemBigEndian types.

## Usage

```go
package main

import (
	"fmt"
	"bytes"

	"vimagination.zapto.org/byteio"
)

func main() {
	var buf bytes.Buffer

	w := byteio.StickyBigEndianWriter{Writer: &buf}

	// Write a uint32
	w.WriteUint32(0xDEADBEEF)

	// Write a String
	w.WriteString8("hello")

	// Write a Bool
	w.WriteBool(true)

	r := byteio.StickyBigEndianReader{Reader: &buf}

	// Read the values back
	fmt.Printf(
		"Read uint32: 0x%X\n" +
			"Read string: %q\n" +
			"Read bool: %v\n",
		r.ReadUint32(),
		r.ReadString8(),
		r.ReadBool(),
	)

	fmt.Printf("\nWrote bytes: %d\nRead bytes: %d\n", w.Count, r.Count)
}
```

## Documentation

Full API docs can be found at:

https://pkg.go.dev/vimagination.zapto.org/byteio
