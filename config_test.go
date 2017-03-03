package z3

import (
	"testing"
)

func TestConfig(t *testing.T) {
	c := NewConfig()
	c.SetParamValue("proof", "true")
	c.Close()
}
