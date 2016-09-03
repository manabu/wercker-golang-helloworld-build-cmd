package main

import (
    "bytes"
    "testing"
)

func TestPrint(t *testing.T) {
    buf := &bytes.Buffer{}
    testPrint(buf)
    outputString := buf.String()
    correctString := "Hello world!!\n"
    if correctString != outputString {
        t.Errorf("output string should be %s but %s", correctString, outputString)
        t.FailNow()
    }
}
