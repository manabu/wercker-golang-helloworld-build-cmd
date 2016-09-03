package main

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
