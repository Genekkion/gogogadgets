package node

import "testing"

type stringerType struct{ Name string }

func (s stringerType) String() string { return "Stringer(" + s.Name + ")" }

func TestNewNode_GetValue_Int(t *testing.T) {
	n := NewNode(42)
	if got := n.GetValue(); got != 42 {
		t.Fatalf("GetValue() = %v, want %v", got, 42)
	}
}

func TestNewNode_GetValue_String(t *testing.T) {
	n := NewNode("hello")
	if got := n.GetValue(); got != "hello" {
		t.Fatalf("GetValue() = %q, want %q", got, "hello")
	}
}

func TestSetValue(t *testing.T) {
	n := NewNode(1)
	n.SetValue(2)
	if got := n.GetValue(); got != 2 {
		t.Fatalf("after SetValue, GetValue() = %v, want %v", got, 2)
	}
}

func TestString_BasicTypes(t *testing.T) {
	// int
	ni := NewNode(123)
	if got, want := ni.String(), "123"; got != want {
		t.Fatalf("String() = %q, want %q (int)", got, want)
	}

	// string
	ns := NewNode("world")
	if got, want := ns.String(), "world"; got != want {
		t.Fatalf("String() = %q, want %q (string)", got, want)
	}
}

func TestString_WithStringer(t *testing.T) {
	n := NewNode(stringerType{Name: "Ann"})
	if got, want := n.String(), "Stringer(Ann)"; got != want {
		t.Fatalf("String() = %q, want %q (stringer)", got, want)
	}
}

func TestZeroValueNode(t *testing.T) {
	var n Node[int] // zero value node
	if got, want := n.GetValue(), 0; got != want {
		t.Fatalf("zero Node[int] GetValue() = %v, want %v", got, want)
	}
	if got, want := n.String(), "0"; got != want {
		t.Fatalf("zero Node[int] String() = %q, want %q", got, want)
	}
}
