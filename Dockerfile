# build everything we need

FROM golang:latest AS builder

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./... 
RUN go install -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cmd/createuser ./cmd/createuser
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cmd/randomkey ./cmd/randomkey

# start from a fresh alpine and copy our files in

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /go/src/app .

CMD ["./app"]
