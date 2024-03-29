package regional

import (
	"context"
	"os"
	"testing"

	"github.com/piyuo/libsrv/gaccount"
	"github.com/piyuo/libsrv/log"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	gaccount.ForceTestCredential(true)
	log.ForceStopLog(true)
}

func shutdown() {
	gaccount.ForceTestCredential(false)
	log.ForceStopLog(false)
}
func BenchmarkClean(b *testing.B) {
	ctx := context.Background()
	r, _ := Client(ctx)
	r.Truncate(ctx, "Store")
	r.Truncate(ctx, "Location")
	r.Truncate(ctx, "Verify")
	r.Truncate(ctx, "Count")
	r.Truncate(ctx, "Code")
	r.Truncate(ctx, "Serial")
}
