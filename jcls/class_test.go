package jcls_test

import (
	"testing"

	"os"

	"github.com/LiterMC/wasm-jdk/jcls"
)

func TestParseClass(t *testing.T) {
	fd, err := os.Open("testdata/Test.class")
	if err != nil {
		t.Fatalf("Cannot open file: %v", err)
	}
	defer fd.Close()
	class, err := jcls.ParseClass(fd)
	if err != nil {
		t.Fatalf("Cannot ParseClass: %v", err)
	}
	t.Logf("class: %v", class)
}
