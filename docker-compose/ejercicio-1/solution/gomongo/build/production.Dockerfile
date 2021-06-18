FROM golang:1.16 as builder

LABEL maintainer="carlostrejo2308"

WORKDIR /usr/local/go/app

COPY ./gomongo/cmd/main.go ./cmd/main.go
COPY ./gomongo/pkg ./pkg
COPY ./gomongo/resources/people.xml ./resources/people.xml

RUN go mod init github.com/carlostrejo2308/gomongo && go get github.com/CarlosTrejo2308/peopleApiResource@v1.0.0 && go mod tidy && go mod vendor && go mod download

RUN go build -o ./out/gomongo ./cmd/.


#Productive image
FROM golang:1.16

COPY --from=builder  ["/usr/local/go/app/resources/people.xml", "/usr/local/go/app/resources/people.xml"]
COPY --from=builder  ["/usr/local/go/app/out/gomongo", "/usr/local/go/app/out/"]

WORKDIR /usr/local/go/app

CMD ["./out/gomongo"]