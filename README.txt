Projet de Théorie des applications réparties
Par Victor BOIX et Florian DANIEL


--> Manuel d'installation
Go version requise : go1.8.3 linux/amd64
Positionner la variable GOPATH sur le dossier du Projet.

Compilation des fichiers "serveur" et "client" :
#go install serveur
#go install client


--> Manuel d'utilisation
Ouvrir un Terminal.

Se positionner dans le dossier adéquat "Projet"
On effectue la commande suivante, pour lancer le serveur :
#./bin/serveur

Ouvrir un autre Terminal.

Pour lancer un client, l'argument détermine le nombre de boucles que le travailleur va effectuer :
#./bin/client 5
Le travailleur associé à la requète du client, effectuera 5 boucles.
Attention ! Un client ne peut effectuer qu'une seule requète !
Dans notre programme, il n'y a que 3 travailleurs disponibles.

Le réponse suivante sera affichée sur le terminal "Serveur".
On m'a demandé de travailler : 5 fois , j'ai travaillé : 5 fois.
