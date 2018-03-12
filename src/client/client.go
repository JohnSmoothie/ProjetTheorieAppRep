package main

import (
  "bufio"
  "net"
  "log"
  "fmt"
  "strconv"
)


func main() {
  //var arg string = os.Args[1]
  var arg string = "5"
  fmt.Println(arg)
  a, _ := strconv.Atoi(arg)

/*
  for i:=0;i<a;i++ {
    fmt.Println(i)
  }
  */
  conn, err := net.Dial("tcp","127.0.0.1:10000")
  if err != nil {
    fmt.Println("La connexion a echoué")
    log.Fatal(err)
  }
  writer := bufio.NewWriter(conn)
  _, er := writer.WriteString(strconv.Itoa(a))
  //fmt.Println("moi"+moi)
  writer.Flush()
  if er != nil {
    fmt.Println("L'identification a echoué ")
    log.Fatal(err)
  }
}
