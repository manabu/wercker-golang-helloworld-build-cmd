package main

const version = "0.1.0"

import (
    "fmt"
    "io"
    "os"
)

func testPrint(w io.Writer){
	fmt.Fprint(w,"Hello world!!\n")
}

func main() {
	testPrint(os.Stdout)
}
