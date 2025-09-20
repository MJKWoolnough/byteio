package byteio_test

import (
	"bytes"
	"fmt"

	"vimagination.zapto.org/byteio"
)

func Example() {
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
		"Read uint32: 0x%X\n"+
			"Read string: %q\n"+
			"Read bool: %v\n",
		r.ReadUint32(),
		r.ReadString8(),
		r.ReadBool(),
	)

	fmt.Printf("\nWrote bytes: %d\nRead bytes: %d\n", w.Count, r.Count)

	// Output:
	// Read uint32: 0xDEADBEEF
	// Read string: "hello"
	// Read bool: true
	//
	// Wrote bytes: 11
	// Read bytes: 11
}
