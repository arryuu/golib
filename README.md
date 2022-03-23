## 简介

go工具库

## 安装

```shell
go get -u github.com/arryuu/golib
```

## 使用

### gogo

goroutine - 资源复用节省内存使用量, 允许并发时限制goroutine数量

```go
package main

import (
	"sync"
	"github.com/arryuu/golib/gogo"
)

var wg sync.WaitGroup
g := gogo.NewGo(10) // goroutine
wg.Add(1)
_ = g.Task(func () {
	// func

	wg.Done()
})

wg.Wait()
```

### golog

日志

```go
package main

import "github.com/arryuu/golib/golog"

var (
	Log *golog.LoggerSt
)

Log = golog.New(false) // true, 记录到日志文件上
Log.Info("日志记录")
```

### gocron

定时器

```go
package main

import "github.com/arryuu/golib/gocron"

c := gocron.NewCron()
c.AddFunc("*/5 * * * * * *", func () {
	fmt.Println(5, time.Now())
})
c.AddFunc("*/2 * * * * * *", func () {
	fmt.Println(2, time.Now())
})
c.Start()
```
