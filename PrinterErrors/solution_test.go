package kata

import "testing"

func TestCase1(t *testing.T) {
	prints := PrinterError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz")
	if prints != "3/56" {
		t.Errorf("got %s, want %s", prints, "3/56")
	}
}

func TestCase2(t *testing.T) {
	prints := PrinterError("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz")
	if prints != "6/60" {
		t.Errorf("got %s, want %s", prints, "6/60")
	}
}

func TestCase3(t *testing.T) {
	prints := PrinterError("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyzuuuuu")
	if prints != "11/65" {
		t.Errorf("got %s, want %s", prints, "11/65")
	}
}

func TestCase4(t *testing.T) {
	prints := PrinterError("aaabbbbhaijjjm")
	if prints != "0/14" {
		t.Errorf("got %s, want %s", prints, "0/14")
	}
}

func TestCase5(t *testing.T) {
	prints := PrinterError("aaaxbbbbyyhwawiwjjjwwm")
	if prints != "8/22" {
		t.Errorf("got %s, want %s", prints, "8/22")
	}
}
