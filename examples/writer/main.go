package main

import (
//  "fmt"

    "../../"
)

func main() {
    g := gost.NewGost()
    g.AddApplication(gost.NewWriterApplication())
    g.Run()
}
