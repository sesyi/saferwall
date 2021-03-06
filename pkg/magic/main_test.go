// Copyright 2020 Saferwall. All rights reserved.
// Use of this source code is governed by Apache v2 license
// license that can be found in the LICENSE file.

package magic

import (
	"testing"
)

var magictests = []struct {
	in  string
	out string
}{
	{"../../test/multiav/clean/putty.exe", 
	"PE32 executable (GUI) Intel 80386, for MS Windows"},
}

func TestMagicScan(t *testing.T) {
	for _, tt := range magictests {
		t.Run(tt.in, func(t *testing.T) {
			filePath := tt.in
			got, err := Scan(filePath)
			if err != nil {
				t.Errorf("TestMagicScan(%s) got %v, want %v", tt.in, err, tt.in)
			}
			if got != tt.out {
				t.Errorf("TestMagicScan(%s) got %v, want %v", tt.in, got, tt.out)
			}
		})
	}
}
