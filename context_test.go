package z3

import (
	"testing"
)

func TestContext(t *testing.T) {
	config := NewConfig()
	defer config.Close()

	ctx := NewContext(config)
	ctx.Close()
}
