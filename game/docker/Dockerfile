FROM golang:1.22.2 as builder

WORKDIR /app

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY cmd cmd
COPY internal internal
COPY docs docs

COPY Makefile .

RUN make build

FROM debian:stable

WORKDIR /app
USER nobody

COPY --from=builder /app/server /app/server

EXPOSE 8000

ENTRYPOINT ["/app/server"]

