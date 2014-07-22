package main

import (
//  "fmt"

    "../"
)

func main() {
    g := gost.NewGost()

    g.AddApplication(&gost.WriterApplication{})

    g.Run()
}
