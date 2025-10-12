package node

import (
	"testing"

	"github.com/Genekkion/gogogadgets/pkg/test"
)

// Tests getting value from a node.
func TestNewNodeGetValue(t *testing.T) {
	v := 42
	n := NewNode(v)

	test.AssertEqual(t, "Unexpected value", v, n.GetValue())
}

// Tests setting value on a node.
func TestSetValue(t *testing.T) {
	n := NewNode(1)
	v := 2
	n.SetValue(v)
	test.AssertEqual(t, "Unexpected value", v, n.GetValue())
}
