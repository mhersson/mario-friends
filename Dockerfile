FROM golang:1.23 AS builder

WORKDIR /src

COPY . /src

RUN go env -w CGO_ENABLED=0

RUN GOOS=linux go build -ldflags "-s -w" -o mario-friends main.go


FROM scratch

COPY --from=builder /src/mario-friends /bin/mario-friends

USER 1001

EXPOSE 8080

CMD ["/bin/mario-friends"]
