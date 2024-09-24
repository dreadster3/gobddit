package pkg

import "testing"

type BDDContext struct {
	t *testing.T

	beforeEach func(t *testing.T)
	afterEach  func(t *testing.T)
}

func NewBDDContext(t *testing.T) *BDDContext {
	return &BDDContext{t: t, beforeEach: func(t *testing.T) {}, afterEach: func(t *testing.T) {}}
}

func (c *BDDContext) Child() *BDDContext {
	ctx := NewBDDContext(c.t)

	ctx.beforeEach = c.beforeEach
	ctx.afterEach = c.afterEach

	return ctx
}

func (c *BDDContext) Log(args ...interface{}) {
	c.t.Log(args...)
}

func (c *BDDContext) BeforeEach(f func(*testing.T)) {
	if c.beforeEach == nil {
		return
	}

	c.beforeEach = f
}

func (c *BDDContext) AfterEach(f func(*testing.T)) {
	if c.afterEach == nil {
		return
	}


	c.afterEach = f
}

func (c *BDDContext) When(description string, f func(c *BDDContext)) {
	c.t.Run(description, func(t *testing.T) {
		c.beforeEach(t)
		f(NewBDDContext(t))
		c.afterEach(t)
	})
}

func (c *BDDContext) It(description string, f func(t *testing.T)) {
	c.t.Run(description, func(t *testing.T) {
		if c.beforeEach != nil {
			c.beforeEach(t)
		}
		f(t)
		if c.afterEach != nil {
			c.afterEach(t)
		}
	})
}
