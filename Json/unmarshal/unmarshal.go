package main

import (
  "encoding/json"
  "fmt"
)
type player struct{
  FirstName string
  LastName string
  UserName string
  Age int
  PremiumMember bool
}

func main() {

  j := `[{"FirstName":"Kenneth","LastName":"Spry","UserName":"Shadowhawk","Age":55,"PremiumMember":false},{"FirstName":"Micheal","LastName":"Robby","UserName":"Elite H@xor","Age":23,"PremiumMember":true}]`

  fmt.Println("JSON: ", j)

  xpl := []player{}

  err := json.Unmarshal([]byte(j), &xpl)
  if err != nil {
    fmt.Println(err)
  }

  fmt.Printf("Go: %+v \n", xpl)

}
