package kata

import "testing"

func TestCase1(t *testing.T) {
	var a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	var a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}

	if !Comp(a1, a2) {
		t.Errorf("Comp(a1, a2) should be true")
	}
}

func TestCase2(t *testing.T) {
	var a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	var a2 = []int{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}

	if Comp(a1, a2) {
		t.Errorf("Comp(a1, a2) should be false")
	}
}

func TestCase3(t *testing.T) {
	var a1 []int = nil
	var a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}

	if Comp(a1, a2) {
		t.Errorf("Comp(a1, a2) should be false")
	}
}
