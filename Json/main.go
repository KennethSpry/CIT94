
package main

import (
  "fmt"
  "encoding/json"
)

type person struct {
  First string
  Last string
  TopSecret bool
  Items []string
}

func main() {
  p1 := person{
    First: "Kenneth-",
    Last: "Spry",
    TopSecret: true,
    Items: []string{
      "Software",
      "Tools",
      "USB",
      "patience",
    },
  }
  bs, err := json.Marshal(p1)
  if err != nil{
    fmt.Println(err)
  }

  fmt.Println(p1)
  fmt.Println("-----------------------------------------------------------")
  fmt.Println(string(bs))

}
