#!/bin/bash

apt-get update
apt-get upgrade -y

# install vim.
apt-get install -y --no-install-recommends vim

# install git and curl
sudo apt install -y  git curl golang
curl https://get.docker.com/ | bash -

service start docker
docker network create proxy

git clone https://github.com/notarock/gopdater
cd gopdater
mkdir sources
git clone https://github.com/notarock/portfolio-ng sources/portfolio.ng
cd sources/portfolio.ng
git reset --hard f1464db63eeb20244d5538502c92653026ccad86
bash rebuild.sh

cd -

echo "#!/bin/bash
bash ../sources/rebuild.sh" > scripts/portfolio

chmod +x scripts/portfolio

export GOPATH=$HOME/govagrant/

exportvagrant/
PATH=$PATH:$GOROOT/bin:$GOPATH/bin

go get github.com/gorilla/mux
go build

