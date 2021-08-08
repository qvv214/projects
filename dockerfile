FROM golang:1.16.6

WORKDIR /project

COPY main.go .

COPY /proto ./proto

COPY go.mod .

COPY go.sum .

RUN go mod download

RUN go build

# EXPOSE 8080

CMD [ "./project"]
