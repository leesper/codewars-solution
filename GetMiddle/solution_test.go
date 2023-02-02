package kata

import "testing"

func TestCase1(t *testing.T) {
	middle := GetMiddle("test")
	if middle != "es" {
		t.Errorf("got %s, want %s", middle, "es")
	}
}

func TestCase2(t *testing.T) {
	middle := GetMiddle("testing")
	if middle != "t" {
		t.Errorf("got %s, want %s", middle, "t")
	}
}

func TestCase3(t *testing.T) {
	middle := GetMiddle("middle")
	if middle != "dd" {
		t.Errorf("got %s, want %s", middle, "dd")
	}
}

func TestCase4(t *testing.T) {
	middle := GetMiddle("A")
	if middle != "A" {
		t.Errorf("got %s, want %s", middle, "A")
	}
}
