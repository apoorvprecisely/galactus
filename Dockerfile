FROM golang:1.8

WORKDIR /go/src/app

RUN go get -u github.com/apoorvprecisely/galactus
RUN go install github.com/apoorvprecisely/galactus

CMD ["galactus"]
