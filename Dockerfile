FROM golang:latest
MAINTAINER 0.1 esseryu@gmail.com
#RUN mkdir /go
#RUN mkdir /go/src
RUN mkdir /go/src/github.com
RUN mkdir /go/src/github.com/hqbfs
#ADD . /go/src/github.com/hqbfs
ADD . /go/src/github.com/esseryu/hqbfs
WORKDIR /go/src/github.com/esseryu/hqbfs/hqbfs
RUN go build
#RUN go build -o ./hqbfs/hqbfs /go/src/github.com/esseryu/hqbfs/.
#CMD ["./hqbfs/hqbfs", "$MONGO_PORT_27017_TCP_ADDR", "$MONGO_PORT_27017_TCP_PORT"]
RUN chmod +x ./wait-for-it.sh

#ENTRYPOINT ["./hqbfs_start.sh"]
