package main

import (
    "fmt"
    "strings"
)

func encode(data []byte) {
    length := len(data)
    chunks := [][]byte
    
    for i := 0; i < length; i+=10 {
        append (chunks, data[i:i+10])
        fmt.Println(chunks)
    }

}

func main() {
    data := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
    encode(data)
}
