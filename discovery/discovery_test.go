// Package discovery
package discovery

import (
	"context"
	"testing"
	"time"

	"half-nothing.cn/service-core/interfaces/global"
	"half-nothing.cn/service-core/testutils"
)

func TestNewServiceDiscovery(t *testing.T) {
	lg := testutils.NewFakeLogger(t)
	discover := NewServiceDiscovery(lg, "test", 6850, global.NewVersion("1.0.0"))
	if discover == nil {
		t.Fatal("ServiceDiscovery should not be nil")
		return
	}
	if discover.serviceName != "test" {
		t.Fatal("ServiceName should be test")
		return
	}
	if discover.port != 6850 {
		t.Fatal("Port should be 6850")
		return
	}
	if discover.version == nil {
		t.Fatal("Version should not be nil")
		return
	}
	if err := discover.Start(); err != nil {
		t.Fatal(err)
		return
	}
	time.Sleep(time.Second)
	if err := discover.Stop(context.Background()); err != nil {
		t.Fatal(err)
		return
	}
}
