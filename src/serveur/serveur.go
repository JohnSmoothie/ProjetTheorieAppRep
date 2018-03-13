package main

import (
  "net"
  "log"
  "fmt"
  "bufio"
  "strconv"
  "sync"
)

var wg sync.WaitGroup
var wgTravail sync.WaitGroup

func main() {
  //Le serveur écoute sur le port 10000 avec le protocole tcp
  listener, errCo := net.Listen("tcp",":10000")
  if errCo != nil {
    fmt.Println("La demande de connexion a echoué ")
    log.Fatal(errCo)
  }

  fmt.Println("Serveur ouvert")

  //création du channel fromCollector qui est un channel de string
  fromCollector := make(chan string, 1)
  //création du channel connChan qui est un channel de net.Conn
  connChan := make(chan net.Conn, 1)

  //lancement de l'accepteur de connecteur
  wg.Add(1)
  go AcceptConn(connChan, listener)
  //lancement du colecteur
  wg.Add(1)
  go collecteur(connChan, fromCollector)
  //lancement du répartiteur
  wg.Add(1)
  go repartiteur(fromCollector)

  wg.Wait()
}

func AcceptConn(connChan chan net.Conn, listener net.Listener) {
  defer wg.Done()
  for{
    //Acceptation de la demande de connexion
    conn, errAccept := listener.Accept()
    if errAccept != nil{
      fmt.Println("L'acceptation de la demande de connexion a echoué ")
      log.Fatal(errAccept)
    }
    connChan <- conn
    fmt.Println("Connexion d'un client")
  }
}

/*fonction collecteur
* récupère la valeur envoyé par le client et la place dans le channel fromCollector
*/
func collecteur(connChan chan net.Conn,fromCollector chan string){
  defer wg.Done()
  for {
    conn := <- connChan
    //lecture de la valeur envoyé par le client
    reader := bufio.NewReader(conn)
    reponse, _ := reader.ReadString('\n')

    //la valeur du client est mise dans le channel fromCollecteur
    fromCollector<-reponse
  }
}

/*fonction repartiteur
* répartie le travaille demandé par le client sur 3 travailleurs
*/
func repartiteur(fromCollector chan string) {
  defer wg.Done()
  //création du channel
  avaibleWorker := make(chan chan string, 3)
  //création et insertion de 3 channel workerChan dans avaibleWorker
  workerChan1 := make(chan string,1)
  workerChan2 := make(chan string,1)
  workerChan3 := make(chan string,1)
  avaibleWorker <- workerChan1
  avaibleWorker <- workerChan2
  avaibleWorker <- workerChan3
  for {
    //consomation de la valeur dans fromCollector
    msg := <- fromCollector
    //consomation d'un travailleur disponible de avaibleWorker
    workerChan := <- avaibleWorker

    //occupation du workerChan
    workerChan<-msg
    //lancement d'un travailleur
    wgTravail.Add(1)
    go travailleur(workerChan, avaibleWorker)
    wgTravail.Wait()
  }

}

func travailleur(workerChan chan string, avaibleWorker chan chan string) {
  defer wgTravail.Done()
  nbBoucle := <- workerChan
  //convertion en Entier du nombre de boucle
  nbBoucleInt,errInt := strconv.Atoi(nbBoucle)
  cmp :=0
  for i := 1; i<=nbBoucleInt; i++{
    cmp ++;
  }
  fmt.Println("On m'a demandé de travailler : ", nbBoucle, "fois , j'ai travaillé : ", cmp, " fois")
  //réinsertion du workerChan dans avaibleWorker
  avaibleWorker <- workerChan
}
