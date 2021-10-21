package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {

    content, err := ioutil.ReadFile("text_text.txt")

     if err != nil {
          log.Fatal(err)
     }

    fmt.Println(string(content))
}
