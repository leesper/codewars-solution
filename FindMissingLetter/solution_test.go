package kata

import "testing"

func TestShouldContainRuneInSet(t *testing.T) {
	var s RuneSet = newRuneSet()
	s.Add('a')
	if s.Len() != 1 {
		t.Errorf("got %d, want %d", s.Len(), 1)
	}

	if _, ok := s['a']; !ok {
		t.Errorf("RuneSet s should contain %c", 'a')
	}
}

func TestShouldContainInitialsInRuneSet(t *testing.T) {
	var s RuneSet = newRuneSetWithInitials([]rune{'a', 'b', 'c'})
	if s.Len() != 3 {
		t.Errorf("got %d, want %d", s.Len(), 3)
	}

	if _, ok := s['a']; !ok {
		t.Errorf("RuneSet s should contain %c", 'a')
	}

	if _, ok := s['b']; !ok {
		t.Errorf("RuneSet s should contain %c", 'b')
	}

	if _, ok := s['c']; !ok {
		t.Errorf("RuneSet s should contain %c", 'c')
	}
}

func TestShouldContainDifferenceInRuneSet(t *testing.T) {
	var s RuneSet = newRuneSetWithInitials([]rune{'a', 'b', 'c', 'd'})
	var v RuneSet = newRuneSetWithInitials([]rune{'b', 'd'})
	diff := s.Diff(v)
	if len(diff) != 2 {
		t.Errorf("got %d, want %d", len(diff), 2)
	}

	if _, ok := s['a']; !ok {
		t.Errorf("RuneSet s should contain %c", 'a')
	}

	if _, ok := s['c']; !ok {
		t.Errorf("RuneSet s should contain %c", 'c')
	}
}

func TestCase1(t *testing.T) {
	missing := FindMissingLetter([]rune{'a', 'b', 'c', 'd', 'f'})
	if missing != 'e' {
		t.Errorf("got %c, want %c", missing, 'e')
	}
}
func TestCase2(t *testing.T) {
	missing := FindMissingLetter([]rune{'O', 'Q', 'R', 'S'})
	if missing != 'P' {
		t.Errorf("got %c, want %c", missing, 'P')
	}
}
