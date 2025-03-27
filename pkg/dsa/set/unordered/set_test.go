package set

import "testing"

func TestLen(t *testing.T) {
	s := New[int]()
	for i := range 10 {
		s.Add(i)
		l := s.Len()
		if l != i+1 {
			t.Fatalf("Expected: %v, Got: %v", i+1, l)
		}
	}

	l := s.Len()
	s.Add(1)
	res := s.Len()
	if res != l {
		t.Fatalf("Expected: %v, Got: %v", l, res)
	}
	s.Add(2)
    res = s.Len()
	if res != l {
		t.Fatalf("Expected: %v, Got: %v", l, res)
	}
}
