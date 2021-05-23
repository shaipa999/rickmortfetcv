FROM golang:alpine as builder



WORKDIR /app 

COPY ./files/main.go main.go

RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=off GOARCH=amd64 go build -o get-rickmortyqueryfetch -ldflags="-w -s" . 

FROM scratch

WORKDIR /app

COPY --from=builder /app/get-rickmortyqueryfetch /usr/bin/

ENTRYPOINT ["get-rickmortyqueryfetch"]

