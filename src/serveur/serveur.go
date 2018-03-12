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

  for{
    //Acceptation de la demande de connection
    conn, er := listener.Accept()
    if er != nil{
      fmt.Print("L'acceptation de la demande de connexion a echoué ")
      log.Fatal(err)
    }

    go collecteur(conn)

  }
}

func collecteur(connexion net.Conn) {
  fromCollector := make(chan int)
  reader := bufio.NewReader(conn)
  reponse, _ := reader.ReadString('\n')
  fmt.Println(tab)
  fmt.Println(reponse)
}

func repartiteur() {

}

func travailleur(workerChan chan int) {

}
