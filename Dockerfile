from debian:9.7-slim

# This dockerfile can manage de host's docker process if launched with this command
# docker run -v /var/run/docker.sock:/container/path/docker.sock -p 8000:8000


ENV GOPATH /app
ENV WORKDIR  $GOPATH/gopdater
RUN mkdir -p $WORKDIR

RUN apt update -y
RUN apt install \
  golang        \
  git           \
  gcc           \
  curl           \
  -y

RUN curl -fsSL https://get.docker.com -o get-docker.sh
RUN sh get-docker.sh
RUN apt update -y && apt upgrade -y
RUN apt install docker-compose -y

RUN go get github.com/gorilla/mux

WORKDIR $WORKDIR
COPY . $WORKDIR

RUN go build

EXPOSE 8000

USER root

CMD ./gopdater
