package dameng_test

import (
	"context"
	"database/sql"
	"net/url"
	"os"
	"sync"
	"testing"
	"time"

	dameng "github.com/godoes/gorm-dameng"
)

// Exercise connection-pool churn while existing connections execute queries.
func TestDMSharedConnectorQueryRace(t *testing.T) {
	for _, tc := range []struct {
		name    string
		options map[string]string
	}{
		{name: "default"},
		{name: "zlib", options: map[string]string{"compress": "1", "compressId": "0"}},
		{name: "snappy", options: map[string]string{"compress": "1", "compressId": "1"}},
		{name: "bdta", options: map[string]string{"isBdtaRS": "true"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			dsn, err := url.Parse(dmProbeDSN(t))
			if err != nil {
				t.Fatal(err)
			}
			q := dsn.Query()
			for key, value := range tc.options {
				q.Set(key, value)
			}
			dsn.RawQuery = q.Encode()
			testSharedConnectorQueries(t, dsn.String())
		})
	}
}

func testSharedConnectorQueries(t *testing.T, dsn string) {
	db, err := sql.Open(dameng.DriverName, dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()
	for {
		if err = db.PingContext(ctx); err == nil {
			break
		}
		select {
		case <-ctx.Done():
			t.Fatal("DM readiness failed")
		case <-time.After(time.Second):
		}
	}
	db.SetMaxOpenConns(8)
	db.SetMaxIdleConns(0)
	start := make(chan struct{})
	var wg sync.WaitGroup
	for n := 0; n < 8; n++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for i := 0; i < 40; i++ {
				var got int
				if err := db.QueryRowContext(ctx, "SELECT 1").Scan(&got); err != nil {
					t.Error(err)
					return
				}
				if got != 1 {
					t.Errorf("SELECT 1 returned %d", got)
					return
				}
			}
		}()
	}
	close(start)
	wg.Wait()
}

func TestDMConcurrentConnectorInitializationRace(t *testing.T) {
	dsn := dmProbeDSN(t)
	start := make(chan struct{})
	var wg sync.WaitGroup
	for n := 0; n < 8; n++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for i := 0; i < 10; i++ {
				db, err := sql.Open(dameng.DriverName, dsn)
				if err != nil {
					t.Error(err)
					return
				}
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				err = db.PingContext(ctx)
				cancel()
				_ = db.Close()
				if err != nil {
					t.Error(err)
					return
				}
			}
		}()
	}
	close(start)
	wg.Wait()
}

func dmProbeDSN(t *testing.T) string {
	t.Helper()
	dsn := os.Getenv("DM_PROBE_DSN")
	if dsn == "" {
		t.Skip("DM_PROBE_DSN is required for real-DM integration")
	}
	return dsn
}
