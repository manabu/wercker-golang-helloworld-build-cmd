package main

import (
	"fmt"
	"io"
	"os"
)

const version = "0.1.3-dev"

func testPrint(w io.Writer) {
	fmt.Fprint(w, "Hello world!!\n")
}

func main() {
	testPrint(os.Stdout)
}
