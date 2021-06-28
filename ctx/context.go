package ctx

import (
	"context"
	"time"
)

// Context provides a wrapped context.Context with fluent API
type Context interface {
	context.Context

	GoContext() context.Context
	WithCancel() (Context, context.CancelFunc)
	WithDeadline(deadlineAt time.Time) (Context, context.CancelFunc)
	WithTimeout(timeout time.Duration) (Context, context.CancelFunc)
	WithValue(k, v interface{}) Context
}

func New() Context { return nc(context.TODO()) }

func NewFrom(c context.Context) Context { return nc(c) }

func nc(c context.Context) Context { return &base{c} }

type base struct {
	context.Context
}

func (c *base) GoContext() context.Context { return c.Context }

func (c *base) WithCancel() (Context, context.CancelFunc) {
	cc, cf := context.WithCancel(c)
	return nc(cc), cf
}

func (c *base) WithDeadline(deadlineAt time.Time) (Context, context.CancelFunc) {
	cc, cf := context.WithDeadline(c, deadlineAt)
	return nc(cc), cf
}

func (c *base) WithTimeout(timeout time.Duration) (Context, context.CancelFunc) {
	cc, cf := context.WithTimeout(c, timeout)
	return nc(cc), cf
}

func (c *base) WithValue(k, v interface{}) Context {
	cc := context.WithValue(c, k, v)
	return nc(cc)
}
