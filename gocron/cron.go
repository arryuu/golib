package gocron

import (
	"github.com/arryuu/golib/gosnowflake"
	"sync"
	"time"
)

type (
	Fn func()

	Cron struct {
		snowflake *gosnowflake.Node
		arr       sync.Map
	}

	CronSub struct {
		expr     *Expression
		nextTime time.Time
		fn       Fn
	}
)

func NewCron() *Cron {
	node, _ := gosnowflake.NewNode(1)
	return &Cron{snowflake: node}
}

func (c *Cron) AddFunc(cronLine string, fn Fn) {
	expr := MustParse(cronLine)
	newCron := &CronSub{
		expr:     expr,
		nextTime: expr.Next(time.Now()),
		fn:       fn,
	}
	c.arr.Store(c.snowflake.Generate(), newCron)
}

func (c *Cron) Start() {
	go func() {
		for {
			timeNow := time.Now()
			c.arr.Range(func(key, value interface{}) bool {
				task, ok := value.(*CronSub)
				if ok {
					if task.nextTime.Before(timeNow) || task.nextTime.Equal(timeNow) {
						go task.fn()
						task.nextTime = task.expr.Next(timeNow)
					}
				}
				return true
			})
			<-time.NewTicker(100 * time.Millisecond).C
		}
	}()
}
