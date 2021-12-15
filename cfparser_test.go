package cfparser

import (
	"os"
	"testing"
)

func TestCFParser(t *testing.T) {
	path := "resource/test.txt"
	file, err := os.Open(path)
	if err != nil {
		t.Log(err)
	}

	parser := NewCFParser(file, "#", '=')
	validCount := parser.ReadAll()
	t.Logf("read %v valid config from %v", validCount, path)

	if pair := parser.Get("name"); pair.String() != "CentOS - Base" {
		t.Fatalf("expected 'CentOS - Base', but get '%v'", pair.String())
	}

	if pair := parser.Get("baseurl"); pair.String() != "https://mirrors.somesite.com" {
		t.Fatalf("expected 'https://mirrors.somesite.com', but get '%v'", pair.String())
	}

	if pair := parser.Get("enabled"); pair.Bool() != true {
		t.Fatal("expected true, but get false")
	}
}
