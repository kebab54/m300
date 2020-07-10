# m300 LB3

## Zweck meiner LB3
Ich werde versuchen eine CI/CD Pipeline zu erstellen. Ziel ist es schlussendlich bei einem push auf Github einen automatisch bereitgestellten Container in GKE zu erhalten. Dies ist eine fortsetzung eines fehlgeschlagten versuchs des moduls w906

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

##
In unserem fall wird mit einem dockerfile "gebaut". 

Das Dockerfile sowie der code stammt aus einer demo der Google Cloud Dokumentation. 

Das dockerfile wird als Multi-stage build definiert. Kurzgefasst wird das Docker image in mehreren schritten erstellt. in unserem Fall 2 mal, jeweil bei einem FROM aufruf. zunächst wird das golang:1.13 image geladen. 

ein verzeichnis App wird erstellt und alle abhängigkeiten die von GO benötigt werden werden geladen. (Go MOD Download) 
Code wird Kopiert

und anschliessend ein build erstellt

anschliessend wird das "builder" image verworfen 

Die komplette applikation wird in einen neuen schlankeren Container geladen (alpine image)
Code wird kopiert und ausgeführt

Anschliessend wird die Applikation asgeführt

    ### Use the official Golang image to create a build artifact.`  
    ### This is based on Debian and sets the GOPATH to /go.`  
    ### https://hub.docker.com/_/golang`  
    FROM golang:1.13 as builder`  

    ### Create and change to the app directory.`  
    WORKDIR /app`  
    
    ### Retrieve application dependencies using go modules.`  
    ### Allows container builds to reuse downloaded dependencies.`  
    COPY go.* ./`  
    RUN go mod download`  

    ### Copy local code to the container image.`  
    COPY . ./`

    ### Build the binary.`  
    ### -mod=readonly ensures immutable go.mod and go.sum in container builds.`  
    RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o server`  

    ### Use the official Alpine image for a lean production container.`
    ### https://hub.docker.com/_/alpine`   
    ### https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds`  
    FROM alpine:3`  
    RUN apk add --no-cache ca-certificates` 

    ### Copy the binary to the production image from the builder stage.`  
    COPY --from=builder /app/server /server`  
    
    ### Run the web service on container startup.`  
    CMD ["/server"]`  

Ich habe das dockerfile anschliessend mit einem dockerfile eines anderen dummy apps ersetzt. Die funktion ist jedoch ähnlich. Jedoch ist es kein Multistage dockerfile mehr, sondern das multistage building wird ausgelagert in ein cloudbuild.yml file. 

### Probleme
Ich hatte das grösste problem dabei die Builds hinzukriegen. Momentan bin ich an folgendem problem. Bei jedem build erhalte ich den fehler manifest unknown: Failed to fetch "latest" 


## Container Registry

In meiner Google Cloud Container Registry werden alle gebauten Docker Images gespeichert welche erfolgreich in Cloud Build durchlaufen sind.

## Continuous Delivery
Ich wollte zunächst Spinnaker als CD tool verwenden. Jedoch gelang es mir nicht nach mehreren Versuchen/ stunden die appliance auf meiner cloud umgebung bereitzustellen. Die installation dauert jeweils mehr als 30min ist jedoch bei 3 versuchen fehlgeschlagen. Ich habe es über mehrere Guides versucht, wobei ich zunächst die Marketplace version verwendet habe. Und anschliessend einfach in der Cloud Console ein wget mit allen Scripts etc ausgeführt habe. 

Die konfiguration im skript scheint zu funktionieren, jedoch hat die Installation nie abgeschlossen. 

Anschliessend habe ich mich für WeaveWorks entschieden. Dies ist ein bezahltes cloudbasiertes tool für CD in google cloud. mir gelang es soweit meinen Cloud umgebung zu integrieren. 

Zur installation musste ich ebenfalls einen command in der cloud shell ausführen. 
curl -Ls https://get.weave.works |
  sh -s -- --token=zdag3658dbrzgarbk7bbta5u3upmyabj --gke

Dieses script installiert die weaveworks agents auf einem GKE Cluster und verbindet sich automatisch mit den diensten von WeaveWorks. 

##
