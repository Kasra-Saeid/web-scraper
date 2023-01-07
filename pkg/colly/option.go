package colly

import (
	"time"

	collyPkg "github.com/gocolly/colly"
)

type Option func(c *collyPkg.Collector)

func AddRandomAgent(agentGenerator func() string) Option {
	return func(c *collyPkg.Collector) {
		c.OnRequest(func(r *collyPkg.Request) {
			r.Headers.Set("User-Agent", agentGenerator())
		})
	}
}

func DoParallel(routineLimit int) Option {
	return func(c *collyPkg.Collector) {
		c.Async = true
		c.MaxDepth = 1
		c.Limit(&collyPkg.LimitRule{
			DomainRegexp: "",
			DomainGlob:   "*",
			Delay:        time.Millisecond * 3,
			RandomDelay:  0,
			Parallelism:  routineLimit,
		})
	}
}
