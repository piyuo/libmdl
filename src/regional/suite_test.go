package regional

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
	r, _ := RegionalClient(ctx)
	r.Truncate(ctx, "Store", 100)
	r.Truncate(ctx, "Location", 100)
	r.Truncate(ctx, "Verify", 100)
	r.Truncate(ctx, "Count", 100)
	r.Truncate(ctx, "Code", 100)
	r.Truncate(ctx, "Serial", 100)
}
