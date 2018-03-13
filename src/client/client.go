package main

import (
  "bufio"
  "net"
  "log"
  "fmt"
  "os"
  "strconv"
)


func main() {
  //récupération du nombre d'arguments
  var argLen int = len(os.Args[1:])

  //on ne communique avec le serveur que si on a 1 seul argument
  if argLen == 1 {
    //récupération du premier argument, qui est la valeur que l'on veut envoyer au serveur
    var argString string = os.Args[1]
    //test si la valeur de l'agument est bien un entier
    _,errInt := strconv.Atoi(argString)
    if errInt != nil {
      fmt.Println("L'argument n'est pas un entier")
      os.Exit(0)
    }

    //connexion au serveur
    conn, err := net.Dial("tcp","127.0.0.1:10000")
    if err != nil {
      fmt.Println("La connexion a echoué")
      log.Fatal(err)
    }

    writer := bufio.NewWriter(conn)
    _, errWrite := writer.WriteString(argString)
    if errWrite != nil {
      fmt.Println("Erreur d'écriture sur le serveur")
      log.Fatal(errWrite)
    }
    //envoy de la valeur au serveur
    writer.Flush()
    fmt.Println("Vous avez demandé au serveur de travailler "+argString+ " fois")

  } else {
    fmt.Println("Veuillez rentrer exactement un argument")
    os.Exit(0)
  }
}
