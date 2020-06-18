package card

import "testing"
import "github.com/stretchr/testify/assert"

func TestFoo(t *testing.T) {
	asserts := assert.New(t)

	asserts.Equal(1, 1, "Yes")
}
