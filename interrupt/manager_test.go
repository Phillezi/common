package interrupt

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestManager_ShutdownHookIsCalled(t *testing.T) {
	resetInstance() // Ensure clean state

	mgr := GetInstance()
	var called int32

	mgr.AddShutdownHook(func() {
		atomic.StoreInt32(&called, 1)
	})

	mgr.Shutdown()

	if atomic.LoadInt32(&called) != 1 {
		t.Fatal("expected shutdown hook to be called")
	}
}

func TestManager_ShutdownIsIdempotent(t *testing.T) {
	resetInstance()

	mgr := GetInstance()
	var count int32

	mgr.AddShutdownHook(func() {
		atomic.AddInt32(&count, 1)
	})
	mgr.Shutdown()
	mgr.Shutdown()
	mgr.Shutdown()

	if atomic.LoadInt32(&count) != 1 {
		t.Fatalf("expected hook to be called once, got %d", count)
	}
}

func TestManager_ContextCancellation(t *testing.T) {
	resetInstance()

	mgr := GetInstance()
	ctx := mgr.Context()

	done := make(chan struct{})
	go func() {
		<-ctx.Done()
		close(done)
	}()

	mgr.Shutdown()

	select {
	case <-done:
		// success
	case <-time.After(1 * time.Second):
		t.Fatal("expected context to be canceled")
	}
}

func TestManager_Wait_Timeout(t *testing.T) {
	resetInstance()

	mgr := GetInstance()

	start := time.Now()
	mgr.Wait(100 * time.Millisecond)
	elapsed := time.Since(start)

	if elapsed < 90*time.Millisecond {
		t.Fatal("expected Wait to block at least 90ms")
	}
}

func TestManager_Wait_AfterShutdown(t *testing.T) {
	resetInstance()

	mgr := GetInstance()
	go func() {
		time.Sleep(50 * time.Millisecond)
		mgr.Shutdown()
	}()

	start := time.Now()
	mgr.Wait(500 * time.Millisecond)
	elapsed := time.Since(start)

	if elapsed > 300*time.Millisecond {
		t.Fatal("expected Wait to return shortly after shutdown")
	}
}
