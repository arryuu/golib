package gogo

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

const (
	_   = 1 << (10 * iota)
	MiB // 1048576
)

const (
	Size1 = 100
	Size2 = 1000

	Count1 = 10000
)

var curMem uint64

func demoFunc() {
	time.Sleep(time.Duration(10) * time.Millisecond)
}

func TestGogo(t *testing.T) {
	var wg sync.WaitGroup
	g := NewGo(Size1)
	for i := 0; i < Count1; i++ {
		wg.Add(1)
		_ = g.Task(func() {
			demoFunc()
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
			wg.Done()
		})
	}

	wg.Wait()
	t.Log("执行完毕")
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	t.Logf("memory usage:%d MB", curMem)
}
