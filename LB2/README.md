# m300 LB2
## Linux
Typischerweise open source betriebssysteme basierend auf dem Linux Kernel. Es gibt dabei verschiedene Distributionen und Linux Derivate. 
Da Linux Open source ist, hat jeder die möglichkeit den Code nach freiem willen zu ändern. 
## Virtualisierung 
Abstrahierung der Hardware eines Servers. Bei der Virtualisierung wird die Physische Hardware eines Servers von einem VMM/Hypervisor verwaltet und wird auf mehrere Virtuelle Maschinen. Der Hypervisor erlaubt es Hardwareressourcen einer VM zuzuteilen. 
Die hauptsächlichen vorteile der virtualisierung ist die effizientere Nutzung von Hardwareressourcen. So wird Strom, Platz und somit auch Geld gespart. 

Die konfiguration einer VM ist primär unabhängig von der Hardware. Dies ermöglicht sogenanntes Over Provisioning. dabei kann z.B. mehr speicher einer VM zugeteilt werden als tatsächlich zur verfügung steht. Overprovisioning wird oftmals betrieben, da typischerweise sowieso nur die Tatsächli von der VM benötigten ressourcen verbraucht werden. 

## Git 
Git ist ein versionsverwaltungssystem. Es kann on premise oder auch in der cloud betireben werden. Beliebt ist z.B. Github

Wichtige Commands:

eigenen Namen und email hinterlegen

git config --global user.name "Sam Smith"
git config --global user.email sam@example.com

git init: Lokale Repo erstellen
git clone: Kopie einer repository in die lokal repository
git add: files zum ändern markieren
git commit: änderungen anwenden (lokal)
git push: änderungen and das "remote" Repository senden

## Markdown

## Vagrant
Vagrant ist eine Software welche verwendet wird um VMs zu provisionieren. 
es ermögtlich leichtes teilen von umgebungen, da nicht komplette images verteilt werden, sondern nur vagrantfiles. Die os images werden anschliessend aus der vagrant cloud heruntergeladen oder ein lokal vorhandenes image wird verwendet. 

Vagrant Init: erstellt ein Vagrantfile
Vagrant box: verwaltung von vagrant boxen, z.B. vagrant box list
Vagrant Up: box anhand vagrantfile erstellen und ausführen
Vagrant SSH: ssh session zur vagrantbox
Vagrant Destroy: löscht die Vagrant umgebung, jedoch nicht die box

# Des erstellten vagrant files
Das vagrantfile erstellt eine Ubuntu VM und installiert automatisch jenkins. nach der installation wird ausserdem automatisch das standard passwort von jenkins ausgegeben. welches benötigt wird um jenkins auszuführen. 

	Vagrant.configure("2") do |config|
	  config.vm.box = "ubuntu/bionic64"
	  config.vm.network "forwarded_port", guest: 8080, host: 8080
	  config.vm.provider "virtualbox" do |vb|
		vb.gui = false
		vb.cpus = 2
		vb.memory = "4096"
	  end
	config.vm.provision "shell" do |shell|
    shell.path = "setupscript.sh"
	end
    end

Das Vagreantfile erstellt eine Ubuntu Bionic 64 VM.
Port 8080 der VM wird auf port 8080 vom host belegt. somit ist dann jenkins mit localhost:8080 aufrufbar. 
anschliessend wird das script setupscript.sh ausgeführt

    #Hinzufügen von Jenkins Repository, 

    wget -q -O - https://pkg.jenkins.io/debian-stable/jenkins.io.key | sudo apt-key add -
    sudo sh -c 'echo deb https://pkg.jenkins.io/debian-stable binary/ > \
        /etc/apt/sources.list.d/jenkins.list'
    #apt update und upgrade
    sudo apt-get update && sudo apt-get upgrade
    #Java installieren
    echo default jre und jdk wird installiert
    sudo apt-get -y install default-jre default-jdk > /dev/null 2>&1
    #git installieren
    echo git und git ftp wird installiert 
    sudo apt-get -y install git git-ftp > /dev/null 2>&1
    #jenkins installieren
    echo Jenkins wird installiert und gestartet
    sudo apt-get -y install jenkins
    sudo service jenkins start
    #wartezeit für service start
    sleep 1m

    #jenkins initialpasswort ausgeben
    echo Jenkins Initialpasswort
    sudo cat /var/lib/jenkins/secrets/initialAdminPassword


## Test der VM
Um die VM zu testen muss sie zunächst gestartet werden mit vagrant up.

