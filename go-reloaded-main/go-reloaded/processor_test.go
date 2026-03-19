package main

import "testing"

func TestSomething(t *testing.T) {
	input := "..."
	expected := "..."

	result := ProcessText(input)

	if result != expected {
		t.Errorf("Expected: %q, got: %q", expected, result)
	}
}

func TestHex(t *testing.T) {
	input := "1E (hex)"
	expected := "30"

	result := ProcessText(input)

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestBin(t *testing.T) {
	input := "10 (bin)"
	expected := "2"

	result := ProcessText(input)

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestUpper(t *testing.T) {
	input := "go (up)"
	expected := "GO"

	result := ProcessText(input)

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestUpperN(t *testing.T) {
	input := "this is cool (up, 2)"
	expected := "this IS COOL"

	result := ProcessText(input)

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestCapitalize(t *testing.T) {
	input := "hello world (cap)"
	expected := "hello World"

	result := ProcessText(input)

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestCapitalizeN(t *testing.T) {
	input := "this is amazing (cap, 3)"
	expected := "This Is Amazing"

	result := ProcessText(input)

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestPunctuation(t *testing.T) {
	input := "hello , world !"
	expected := "hello, world!"

	result := ProcessText(input)

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestArticles(t *testing.T) {
	input := "a apple"
	expected := "an apple"

	result := ProcessText(input)

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestQuotes(t *testing.T) {
	input := "this is ' amazing '"
	expected := "this is 'amazing'"

	result := ProcessText(input)

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestFull(t *testing.T) {
	input := "1E (hex) is cool , go (up)"
	expected := "30 is cool, GO"

	result := ProcessText(input)

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}