// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package utils
package utils

import "testing"

func TestVersion(t *testing.T) {
	v := NewVersion("1.2.3")
	if v.String() != "1.2.3" {
		t.Fail()
	}
	if v.CheckVersion(NewVersion("1.2.3")) != AllMatch {
		t.Fail()
	}
	if v.CheckVersion(NewVersion("1.2.4")) != PatchUnmatch {
		t.Fail()
	}
	if v.CheckVersion(NewVersion("1.3.3")) != MinorUnmatch {
		t.Fail()
	}
	if v.CheckVersion(NewVersion("2.2.3")) != MajorUnmatch {
		t.Fail()
	}
}
