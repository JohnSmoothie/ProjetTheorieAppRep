package main

import (
  "net"
  "log"
  "fmt"
  "bufio"
)

func main() {
  listener, err := net.Listen("tcp",":10000")
  if err != nil {
    fmt.Print("La demande de connexion a echoué ")
    log.Fatal(err)
  }

  var tab []string

  for{
    //Acceptation de la demande de connection
    conn, er := listener.Accept()
    if er != nil{
      fmt.Print("L'acceptation de la demande de connexion a echoué ")
      log.Fatal(err)
    }
    reader := bufio.NewReader(conn)
    reponse, _ := reader.ReadString('\n')
    tab = append(tab,reponse)
    fmt.Println(tab)
    fmt.Println(reponse)
  }
}
