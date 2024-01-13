// Copyright 2021 Dan Kortschak.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qr

import (
	"hash/crc32"
	"testing"
)

func TestMasks(t *testing.T) {
	tests := []struct {
		content string
		crc     uint32
	}{
		{"1", 0x047527ac},   // mask 0.
		{"db", 0x91042aa9},  // mask 1.
		{"0", 0xb61726b6},   // mask 2.
		{"7b9", 0x9a6a83db}, // mask 3.
		{"3", 0xed14d379},   // mask 4.
		{"16", 0xe1ce75fa},  // mask 5.
		{"35", 0x019a2b3a},  // mask 6.
		{"c", 0x1ec9992d},   // mask 7.
	}
	for _, test := range tests {
		qr, err := Encode(test.content, L)
		if err != nil {
			t.Fatal(err)
		}
		crc := crc32.NewIEEE()
		if _, err := crc.Write(qr.Bitmap); err != nil {
			t.Error(err)
		}
		if got := crc.Sum32(); got != test.crc {
			t.Errorf("checksum for %q is %x, want %x", test.content, got, test.crc)
		}
	}
}
