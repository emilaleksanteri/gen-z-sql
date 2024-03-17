package main

import (
	"testing"
)

func TestTokenParse(t *testing.T) {
	for key, val := range tokens {
		parser := parser{input: key}
		parsed := parser.parse()
		if parsed != val {
			t.Errorf("Expected %s, got %s", val, parsed)
		}
	}
}

func TestBasicSelect(t *testing.T) {
	parser := parser{input: "yoink * skibity table"}
	parsed := parser.parse()
	if parsed != "select * from table" {
		t.Errorf("Expected select * from table, got %s", parsed)
	}
}
