package kata

import "testing"

func TestCase1(t *testing.T) {
	highAndLow := HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4")
	if highAndLow != "42 -9" {
		t.Errorf("got %s, want %s", highAndLow, "42 -9")
	}
}

func TestCase2(t *testing.T) {
	highAndLow := HighAndLow("1 2 3")
	if highAndLow != "3 1" {
		t.Errorf("got %s, want %s", highAndLow, "3 1")
	}
}

func TestCase3(t *testing.T) {
	highAndLow := HighAndLow("1 2 3 4 5")
	if highAndLow != "5 1" {
		t.Errorf("got %s, want %s", highAndLow, "5 1")
	}
}

func TestCase4(t *testing.T) {
	highAndLow := HighAndLow("1 2 -3 4 5")
	if highAndLow != "5 -3" {
		t.Errorf("got %s, want %s", highAndLow, "5 -3")
	}
}

func TestCase5(t *testing.T) {
	highAndLow := HighAndLow("1 9 3 4 -5")
	if highAndLow != "9 -5" {
		t.Errorf("got %s, want %s", highAndLow, "9 -5")
	}
}
