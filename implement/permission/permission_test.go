// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

// Package permission
package permission

import "testing"

func TestPermission(t *testing.T) {
	permission := Permission(1)
	if !permission.HasPermission(AdminEntry) {
		t.Error("AdminEntry not found")
	}
	permission.Grant(ActivityShowList)
	if permission.HasPermission(ActivityDelete) {
		t.Error("ActivityDelete found")
	}
	if !permission.HasPermission(ActivityShowList) {
		t.Error("ActivityShowList not found")
	}
	permission.Revoke(ActivityShowList)
	if permission.HasPermission(ActivityShowList) {
		t.Error("ActivityShowList found")
	}
	permission.Merge(ActivityShowList | ActivityEdit)
	if !permission.HasPermission(ActivityShowList | ActivityEdit) {
		t.Error("ActivityShowList | ActivityEdit not found")
	}
}
