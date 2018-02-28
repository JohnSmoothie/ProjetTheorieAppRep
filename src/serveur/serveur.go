package main

import (
  "fmt"
  "net"
  "bufio"
  "strings"
  "sync"
)

func main() {
  fromCollector := make(chan string)
  listener,err:=net.Listen("tcp",":10000")
  if err!=nil{
    fmt.Println("Erreur de Listen")
  }

  for {
    connexion, err := listener.Accept()
    if err!=nil{
      fmt.Println("Erreur de Connexion")
    }

    go collecteur(connexion, fromCollector)
  }
}

func collecteur(connexion net.Conn, fromCollector chan string) {
  
}

func repartiteur() {

}

func travailleur() {

}
