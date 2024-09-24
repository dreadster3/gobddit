package pkg_test

import (
	"testing"

	. "github.com/dreadster3/gobddit/pkg"
	"github.com/stretchr/testify/assert"
)

func TestInteger(t *testing.T) {
	Describe(t, "Given some integer with a starting value", func(c *BDDContext) {
		var x int

		c.BeforeEach(func(t *testing.T) {
			t.Log("Before Each")
			x = 1
		})

		c.AfterEach(func(t *testing.T) {
			t.Log("After Each")
		})

		c.When("Integer is incremented", func(c *BDDContext) {
			x++

			c.It("should be greater by 1", func(t *testing.T) {
				assert.Equal(t, 2, x)
			})

			c.It("should be greater than initial value", func(t *testing.T) {
				assert.Greater(t, x, 1)
			})
		})

		c.When("Integer is decremented", func(c *BDDContext) {
			x--

			c.It("should be lower by 1", func(t *testing.T) {
				assert.Equal(t, 0, x)
			})
			c.It("should be lower than initial value", func(t *testing.T) {
				assert.Less(t, x, 1)
			})
		})
	})
}
