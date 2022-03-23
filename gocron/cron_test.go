package gocron

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGoCron(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(6)
	c := NewCron()
	c.AddFunc("*/5 * * * * * *", func() {
		fmt.Println(5, time.Now())
		wg.Done()
	})
	c.AddFunc("*/2 * * * * * *", func() {
		fmt.Println(2, time.Now())
		wg.Done()
	})
	c.Start()

	wg.Wait()
}
