package trie

import (
	"testing"

	"bytes"
)

func TestCompactEncode(t *testing.T) {
	tests := []struct {
		hex, compact []byte
		isLeaf       bool
	}{
		// empty keys, with and without terminator.
		{hex: []byte{}, compact: []byte{0x00}, isLeaf: false},
		{hex: []byte{16}, compact: []byte{0x20}, isLeaf: true},
		// odd length, no terminator
		{hex: []byte{1, 2, 3, 4, 5}, compact: []byte{0x11, 0x23, 0x45}, isLeaf: false},
		// even length, no terminator
		{hex: []byte{0, 1, 2, 3, 4, 5}, compact: []byte{0x00, 0x01, 0x23, 0x45}, isLeaf: false},
		// odd length, terminator
		{hex: []byte{15, 1, 12, 11, 8, 16 /*term*/}, compact: []byte{0x3f, 0x1c, 0xb8}, isLeaf: true},
		// even length, terminator
		{hex: []byte{0, 15, 1, 12, 11, 8, 16 /*term*/}, compact: []byte{0x20, 0x0f, 0x1c, 0xb8}, isLeaf: true},
	}

	for _, test := range tests {
		if c := compactEncode(test.hex); !bytes.Equal(c, test.compact) {
			t.Errorf("hexToCompact(% 0x) -> % 0x, want % 0x", test.hex, c, test.compact)
		}
		if c := compactDecode(test.compact); !bytes.Equal(c, test.hex) {
			t.Errorf("compactToHex(% 0x) -> % 0x, want % 0x", test.compact, c, test.hex)
		}
		if l := isLeaf(test.compact); l != test.isLeaf {
			t.Errorf("isLeaf(% 0x) -> %v, want %v", test.compact, l, test.isLeaf)
		}
	}
}
