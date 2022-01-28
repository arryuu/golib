package gogo

import (
	"sync"
)

type (
	Fn func()

	GoCls struct {
		cap         uint
		run         uint
		lock        sync.RWMutex
		goPool      sync.Pool
		goPoolCache *goPool
	}

	goPool struct {
		task chan Fn
	}
)

func NewGo(n int) *GoCls {
	if n < 1 {
		n = 1
	}

	cls := &GoCls{
		cap: uint(n),
		run: 0,
	}
	cls.goPool.New = func() interface{} {
		return &goPool{task: make(chan Fn)}
	}
	return cls
}

func (cls *GoCls) Task(fn Fn) error {
	var fp *goPool
	fp = cls.taskRun()
	fp.task <- fn

	return nil
}

func (cls *GoCls) taskRun() (gp *goPool) {
	if cls.goPoolCache == nil {
		cls.goPoolCache = cls.goPool.Get().(*goPool)
	}
	gp = cls.goPoolCache

	cls.lock.Lock()
	if cls.run == cls.cap {
		cls.lock.Unlock()
	} else {
		cls.run++
		cls.lock.Unlock()

		gp.doing(cls)
	}

	return gp
}

func (gp *goPool) doing(cls *GoCls) {
	go func() {
		defer func() {
			cls.lock.Lock()
			cls.run--
			cls.lock.Unlock()
		}()
		for fn := range gp.task {
			if fn == nil {
				return
			}
			fn()
		}
		return
	}()
}
