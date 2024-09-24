package pkg

import "testing"

func Describe(t *testing.T, description string, f func(c *BDDContext)) {
	t.Run(description, func(t *testing.T) {
		f(NewBDDContext(t))
	})
}
