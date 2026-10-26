package counter

import (
	"runtime"
	"sync"
	"testing"
)

func TestCounterSequential(t *testing.T) {
	var c Counter
	c.Add(2)
	c.Add(3)
	if got := c.Value(); got != 5 {
		t.Fatalf("Value() = %d, want 5", got)
	}
}

func TestCounterConcurrent(t *testing.T) {
	const workers, increments = 32, 2000
	var c Counter
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < increments; j++ {
				c.Add(1)
				runtime.Gosched()
			}
		}()
	}
	wg.Wait()
	if got, want := c.Value(), workers*increments; got != want {
		t.Fatalf("Value() = %d, want %d", got, want)
	}
}

func TestCounterReadsDuringWrites(t *testing.T) {
	var c Counter
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			c.Add(1)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			_ = c.Value()
		}
	}()
	wg.Wait()
}
