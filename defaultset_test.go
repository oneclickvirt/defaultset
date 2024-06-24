package defaultset

import (
	"testing"
)

func TestRed(t *testing.T) {
	expected := "\033[31m\033[01mtest\033[0m"
	result := Red("test")
	if result != expected {
		t.Errorf("Red(\"test\") = %v; want %v", result, expected)
	}
}

func TestGreen(t *testing.T) {
	expected := "\033[32m\033[01mtest\033[0m"
	result := Green("test")
	if result != expected {
		t.Errorf("Green(\"test\") = %v; want %v", result, expected)
	}
}

func TestDarkGreen(t *testing.T) {
	expected := "\033[32m\033[02mtest\033[0m"
	result := DarkGreen("test")
	if result != expected {
		t.Errorf("DarkGreen(\"test\") = %v; want %v", result, expected)
	}
}

func TestYellow(t *testing.T) {
	expected := "\033[33m\033[01mtest\033[0m"
	result := Yellow("test")
	if result != expected {
		t.Errorf("Yellow(\"test\") = %v; want %v", result, expected)
	}
}

func TestBlue(t *testing.T) {
	expected := "\033[36m\033[01mtest\033[0m"
	result := Blue("test")
	if result != expected {
		t.Errorf("Blue(\"test\") = %v; want %v", result, expected)
	}
}

func TestPurple(t *testing.T) {
	expected := "\033[35m\033[01mtest\033[0m"
	result := Purple("test")
	if result != expected {
		t.Errorf("Purple(\"test\") = %v; want %v", result, expected)
	}
}

func TestCyan(t *testing.T) {
	expected := "\033[36m\033[01mtest\033[0m"
	result := Cyan("test")
	if result != expected {
		t.Errorf("Cyan(\"test\") = %v; want %v", result, expected)
	}
}

func TestWhite(t *testing.T) {
	expected := "\033[37m\033[01mtest\033[0m"
	result := White("test")
	if result != expected {
		t.Errorf("White(\"test\") = %v; want %v", result, expected)
	}
}
