### Builder
FROM golang:latest
MAINTAINER yu
RUN apt-get -y update

ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p /go/src/github.com/mail
WORKDIR /go/src/github.com/mail
COPY . .

# Download all the dependencies

RUN go get

#RUN go run ./main.go

# Install the package

# This container exposes port 8080 to the outside world
EXPOSE 8081

# Run the executable
CMD ["go run ./main.go"]

