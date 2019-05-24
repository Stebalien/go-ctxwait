package ctxwait

import (
	"context"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestOnDone(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				ctx, cancel := context.WithCancel(context.Background())
				ch := make(chan struct{})
				OnDone(ctx, func() {
					close(ch)
				})

				select {
				case <-time.After(time.Duration(rand.Intn(int(time.Second))) + time.Millisecond):
				case <-ch:
					t.Error("should not have fired")
					cancel()
					return
				}

				cancel()

				select {
				case <-time.After(time.Millisecond):
					t.Error("should have fired")
				case <-ch:
				}
			}
		}()
	}
	wg.Wait()
}
