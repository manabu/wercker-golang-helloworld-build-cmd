package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrint(t *testing.T) {
	buf := &bytes.Buffer{}
	testPrint(buf)
	outputString := buf.String()
	correctString := "Hello world!!"
	if strings.Index(outputString, correctString) == -1 {
		t.Errorf("output string should be %s but %s", correctString, outputString)
		t.FailNow()
	}
}
