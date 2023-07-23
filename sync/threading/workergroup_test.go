package threading

import (
	"fmt"
	"runtime"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkerGroup(t *testing.T) {
	m := make(map[string]struct{})
	var lock sync.Mutex
	var wg sync.WaitGroup
	wg.Add(runtime.NumCPU())
	group := NewWorkerGroup(func() {
		lock.Lock()
		m[fmt.Sprint(RoutineId())] = struct{}{}
		lock.Unlock()
		wg.Done()
	}, runtime.NumCPU())
	go group.Start()
	wg.Wait()
	assert.Equal(t, runtime.NumCPU(), len(m))
}