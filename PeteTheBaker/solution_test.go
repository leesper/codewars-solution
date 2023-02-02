package kata

import "testing"

func TestCase1(t *testing.T) {
	num := Cakes(map[string]int{"flour": 500, "sugar": 200, "eggs": 1}, map[string]int{"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200})
	if num != 2 {
		t.Errorf("got %d, want %d", num, 2)
	}
}

func TestCase2(t *testing.T) {
	num := Cakes(map[string]int{"apples": 3, "flour": 300, "sugar": 150, "milk": 100, "oil": 100}, map[string]int{"sugar": 500, "flour": 2000, "milk": 2000})
	if num != 0 {
		t.Errorf("got %d, want %d", num, 0)
	}
}
