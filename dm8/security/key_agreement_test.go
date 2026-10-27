package security

import (
	"bytes"
	"sync"
	"testing"
)

func TestConcurrentKeyAgreement(t *testing.T) {
	start := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for j := 0; j < 40; j++ {
				alice, err := NewClientKeyPair()
				if err != nil {
					t.Error(err)
					return
				}
				bob, err := NewClientKeyPair()
				if err != nil {
					t.Error(err)
					return
				}
				aliceSecret := ComputeSessionKey(alice, bob.GetY().Bytes())
				bobSecret := ComputeSessionKey(bob, alice.GetY().Bytes())
				if len(aliceSecret) == 0 || !bytes.Equal(aliceSecret, bobSecret) {
					t.Error("peers derived different session keys")
					return
				}
			}
		}()
	}
	close(start)
	wg.Wait()
}
