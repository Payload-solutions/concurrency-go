package main

import (
  "fmt"
  "time"
  "sync"
)

func main(){
  
  var wg sync.WaitGroup
  for x:= 0; x< 10; x++{
    wg.Add(1)
    go doSomething(x, &wg)
  }
  wg.Wait()
}

func doSomething(x int, wg *sync.WaitGroup){
  defer wg.Done()
  fmt.Printf("Started %v\n", x)
  time.Sleep(2 * time.Second)
  fmt.Println("Finished")
}
