package main

import (
  "fmt"
  "time"
)

func main() {
  timeout := make(chan bool, 1)
  go func(){
    time.Sleep(time.Millisecond)
    timeout <- false
  }()

  go func(){
    var e int
    ok:=true
    for{
      select{
      case e,ok=<-ch11:
        if !ok{
          fmt.Println("END.")
          break
        }else{
          fmt.Printf("%d\n", e)
        }
      case ok = <-timeout:
        fmt.Println("Timeout.")
        break
      }
      if ok!{
        break
      }
    }
  }()
}
