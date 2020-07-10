# m300 LB3

## Zweck meiner LB3
Ich werde versuchen eine CI/CD Pipeline zu erstellen. Ziel ist es schlussendlich bei einem push auf Github einen automatisch bereitgestellten Container in GKE zu erhalten. 

## Docker
Docker ist eine software die es ermöglich container zu erstellen und verwalten. Dies wird auch Containervirtualierung genannt. Im gegenzug zu Servervirtualisierung wird bei Containervirtualisierung nicht die Hardwareschicht abstrahiert, sondern die Betriebssystemschicht. 

docker run – Führt einen befehl in einem neuen Container aus.
docker start – Startet Container
docker stop – Stoppt Container 
docker build – Baut ein Docker image aus einem Dockerfile
docker pull – Image aus einem Registry laden 
docker push – Image oder Repository in ein Registry pushen z.B. Docker Hub
docker export – Inhalt eines Containers als TAR Archiv Exportieren
docker exec – führt Befehl in einem laufdendem Container aus. 
docker search – Docker hub nach einem Image durchsuchen
docker attach – Koppelt lokales terminal mit Terminal eines containers (vergleichbar mit ssh session bei vagrant)
docker commit – Erstellt neues image sofern änderungen getätigt wurden. 

### Docker Beispielimage
git clone https://github.com/dockersamples/node-bulletin-board
cd node-bulletin-board/bulletin-board-app
clone des Docker testimage

docker build --tag bulletinboard:1.0 .
build des images anhand dockerfile 

docker run --publish 8000:8080 --detach --name bb bulletinboard:1.0
docker image ausführen und localhost port 8000 an den docker container port 8080 binden. --detach: container in hintergrund ausführen --name erstellt einen namen um den container einfacher zu finden (anstatt Container ID)

#### Test
Zum testen wird localhost:8080 im browser geöffnet

## CI und CD
Continuous Integration und Continuous Deployment

Ci und CD sind wichtige bestandteile einer Dev Ops Pipeline
CI automatische überprüfung von Code und automatische Builds von Applikationen (Google Cloud Builder)

In meiner Demo umgebung wurden folgende komponente verwendet:

Code repository: Google Cloud Platform Source Repositories (GSR)
Continuous integration: Google Cloud Builder (GCB)
Container registry: Google Container Registry (GCR)
Continuous deployment: Weave Cloud Deploy
Orchestrator: Google Kubernetes Engine (GKE)

## Repository
Der komplette code des Demoapps
wird auf github geladen. 

## Cloud Build 
wird ein Push auf mein Github Repository ausgeführt, wird automatisch ein build ausgeführt. momentan wird die applikation als dockerfile in einen container gebaut. (Docker Build) Dazu muss in Cloud Build mein repository integriert werden und ein Trigger erstellt werden. 
### Trigger
Der trigger beinhaltet bei welchen ereignissen ein build ausgeführt wird. in diesem fall bei einem Push. Ausserdem kann angegeben werden ob ein Cloudbuild.yml file verwendet wird oder nur ein dockerfile.

mit einem yml file ist es möglich andere builder zu verwenden. es ist ausserdem möglich 3rd party builder aus einem github repository zu importieren falls der code standarmässig keinen builder hat. Es ist ebenfalls möglich selbst einen Builder zu erstellen. 

In unserem fall wird mit einem dockerfile "gebaut". 

Das Dockerfile sowie der code stammt aus einer demo der Google Cloud Dokumentation. 

