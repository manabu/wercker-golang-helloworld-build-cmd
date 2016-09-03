package main

import (
    "fmt"
    "io"
    "os"
)

func testPrint(w io.Writer){
	fmt.Println("Hello world!!")
}

func main() {
	testPrint(os.Stdout)
}
