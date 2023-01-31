package kata

import "testing"

func TestCase1(t *testing.T) {
	numOfBits := CountBits(0)
	if numOfBits != 0 {
		t.Errorf("got %d, want %d", numOfBits, 0)
	}
}

func TestCase2(t *testing.T) {
	numOfBits := CountBits(4)
	if numOfBits != 1 {
		t.Errorf("got %d, want %d", numOfBits, 1)
	}
}
func TestCase3(t *testing.T) {
	numOfBits := CountBits(7)
	if numOfBits != 3 {
		t.Errorf("got %d, want %d", numOfBits, 3)
	}
}
func TestCase4(t *testing.T) {
	numOfBits := CountBits(9)
	if numOfBits != 2 {
		t.Errorf("got %d, want %d", numOfBits, 2)
	}
}
func TestCase5(t *testing.T) {
	numOfBits := CountBits(10)
	if numOfBits != 2 {
		t.Errorf("got %d, want %d", numOfBits, 2)
	}
}
func TestCase6(t *testing.T) {
	numOfBits := CountBits(1234)
	if numOfBits != 5 {
		t.Errorf("got %d, want %d", numOfBits, 5)
	}
}
