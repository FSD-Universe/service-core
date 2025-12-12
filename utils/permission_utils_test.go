// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package utils
package utils

import "testing"

func TestPermission(t *testing.T) {
	permission := Permission(1)
	if !permission.HasPermission(PermissionAdminEntry) {
		t.Error("PermissionAdminEntry not found")
	}
	permission.Grant(PermissionActivityShowList)
	if permission.HasPermission(PermissionActivityDelete) {
		t.Error("PermissionActivityDelete found")
	}
	if !permission.HasPermission(PermissionActivityShowList) {
		t.Error("PermissionActivityShowList not found")
	}
	permission.Revoke(PermissionActivityShowList)
	if permission.HasPermission(PermissionActivityShowList) {
		t.Error("PermissionActivityShowList found")
	}
	permission.Merge(PermissionActivityShowList | PermissionActivityEdit)
	if !permission.HasPermission(PermissionActivityShowList | PermissionActivityEdit) {
		t.Error("PermissionActivityShowList | PermissionActivityEdit not found")
	}
}
