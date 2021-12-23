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

	CFP := NewCFParser(file, "#", '=')
	numOfValidLines := CFP.ReadAll()
	t.Logf("read %v valid lines", numOfValidLines)

	if pair := CFP.Get("name"); pair.String() != "CentOS - Base" {
		t.Fatalf("expected 'CentOS - Base', but get '%v'", pair.String())
	}

	if pair := CFP.Get("baseurl"); pair.String() != "https://mirrors.somesite.com" {
		t.Fatalf("expected 'https://mirrors.somesite.com', but get '%v'", pair.String())
	}

	if pair := CFP.Get("enabled"); pair.Bool() != true {
		t.Fatal("expected true, but get false")
	}
}
