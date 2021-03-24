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
	client, _ := RegionalClient(ctx)
	client.Truncate(ctx, "Store", 100)
	client.Truncate(ctx, "Location", 100)
}
