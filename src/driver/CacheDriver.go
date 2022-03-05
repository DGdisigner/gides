package driver

import (
	"sync"
	"time"
)

type CacheDriver struct {
	DataMap map[string]*CacheValue
	Len     int
	look    sync.RWMutex
}

type CacheValue struct {
	Val string
	Ttl int64
}

func (c *CacheDriver) Find(key string) (str string) {
	c.look.RLock()
	defer c.look.RUnlock()
	if c.DataMap == nil {
		return
	}
	target, ok := c.DataMap[key]
	if ok && target.Ttl > time.Now().Unix() {
		str = target.Val
	}
	return
}

func (c *CacheDriver) Set(key string, value string, ttl int64) {
	c.look.Lock()
	defer c.look.Unlock()
	target, ok := c.DataMap[key]
	if ok {
		target.Val = value
		target.Ttl = time.Now().Unix() + ttl
	} else {
		add := CacheValue{
			Val: value,
			Ttl: time.Now().Unix() + ttl,
		}
		if c.DataMap == nil {
			c.DataMap = make(map[string]*CacheValue)
		}
		c.DataMap[key] = &add
		c.Len += 1
	}
}

func (c *CacheDriver) Clean() {
	c.look.Lock()
	defer c.look.Unlock()
	newDataMap := make(map[string]*CacheValue)
	l := 0
	for k, v := range c.DataMap {
		if v.Ttl > time.Now().Unix() {
			newDataMap[k] = v
			l += 1
		}
	}
	c.DataMap = newDataMap
	c.Len = l
	return
}
