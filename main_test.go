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

func TestBasicWhere(t *testing.T) {
	parser := parser{input: "yoink * skibity table on god column = 5"}
	parsed := parser.parse()
	if parsed != "select * from table where column = 5" {
		t.Errorf("Expected select * from table where column = 5, got %s", parsed)
	}
}

func TestBasicAnd(t *testing.T) {
	parser := parser{input: "yoink * skibity table on god column = 5 goon column2 = 6"}
	parsed := parser.parse()
	if parsed != "select * from table where column = 5 and column2 = 6" {
		t.Errorf("Expected select * from table where column = 5 and column2 = 6, got %s", parsed)
	}
}

func TestBasicOr(t *testing.T) {
	parser := parser{input: "yoink * skibity table on god column = 5 edge column2 = 6"}
	parsed := parser.parse()
	if parsed != "select * from table where column = 5 or column2 = 6" {
		t.Errorf("Expected select * from table where column = 5 or column2 = 6, got %s", parsed)
	}
}

func TestBasicNot(t *testing.T) {
	parser := parser{input: "yoink * skibity table on god bruh column = 5"}
	parsed := parser.parse()
	if parsed != "select * from table where not column = 5" {
		t.Errorf("Expected select * from table where not column = 5, got %s", parsed)
	}
}

func TestBasicAsc(t *testing.T) {
	parser := parser{input: "yoink * skibity table short king"}
	parsed := parser.parse()
	if parsed != "select * from table asc" {
		t.Errorf("Expected select * from table asc, got %s", parsed)
	}
}

func TestBasicDesc(t *testing.T) {
	parser := parser{input: "yoink * skibity table tall king"}
	parsed := parser.parse()
	if parsed != "select * from table desc" {
		t.Errorf("Expected select * from table desc, got %s", parsed)
	}
}

func TestBasicDelete(t *testing.T) {
	parser := parser{input: "yeet skibity table on god column = 5"}
	parsed := parser.parse()
	if parsed != "delete from table where column = 5" {
		t.Errorf("Expected delete from table where column = 5, got %s", parsed)
	}
}

func TestBasicInsert(t *testing.T) {
	parser := parser{input: "slide into table (column1, column2) values (5, 6)"}
	parsed := parser.parse()
	if parsed != "insert into table (column1, column2) values (5, 6)" {
		t.Errorf("Expected insert into table (column1, column2) values (5, 6), got %s", parsed)
	}
}
