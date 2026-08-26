package config

import "time"

type Clock struct { Now func() time.Time }
func RealClock() Clock { return Clock{Now:func() time.Time{return time.Now().UTC()}} }
func (c Clock) Current() time.Time { if c.Now==nil { return time.Now().UTC() }; return c.Now() }
