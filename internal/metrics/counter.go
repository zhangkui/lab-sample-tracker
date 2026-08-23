package metrics

import "sync/atomic"

type Counter struct{ value int64 }

func (c *Counter) Inc()         { atomic.AddInt64(&c.value, 1) }
func (c *Counter) Value() int64 { return atomic.LoadInt64(&c.value) }
