package dm8

import (
	"sync"
	"testing"
)

func TestConcurrentDiagnosticConfiguration(t *testing.T) {
	start := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for j := 0; j < 100; j++ {
				props := NewProperties()
				props.Set(LogLevelKey, "error")
				connector := new(DmConnector).init()
				if err := connector.setAttributes(props); err != nil {
					t.Error(err)
					return
				}
				// These are called by active queries and background statistics workers.
				if !ConnLogger.IsErrorEnabled() || ConnLogger.IsDebugEnabled() {
					t.Error("configured log level was not preserved")
					return
				}
				_ = ConnLogger.IsInfoEnabled()
				_ = ConnLogger.IsWarnEnabled()
				_ = ConnLogger.IsSqlEnabled()
				_ = newStatFlusher()
			}
		}()
	}
	close(start)
	wg.Wait()
	// Restore the default configuration for other package tests.
	if err := new(DmConnector).init().setAttributes(NewProperties()); err != nil {
		t.Fatal(err)
	}
}
