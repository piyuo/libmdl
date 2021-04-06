package global

import (
	"context"
	"os"
	"testing"

	"github.com/piyuo/libsrv/src/google/gaccount"
	"github.com/piyuo/libsrv/src/log"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	gaccount.UseTestCredential(true)
	log.TestModeAlwaySuccess()
}

func shutdown() {
	gaccount.UseTestCredential(false)
	log.TestModeBackNormal()
}

func BenchmarkClean(b *testing.B) {
	ctx := context.Background()
	g, _ := GlobalClient(ctx)
	g.Truncate(ctx, "Account")
	g.Truncate(ctx, "Domain")
	g.Truncate(ctx, "User")
	g.Truncate(ctx, "Count")
	g.Truncate(ctx, "Code")
	g.Truncate(ctx, "Serial")
}
