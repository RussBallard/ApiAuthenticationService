FROM golang:1.13-alpine AS builder

WORKDIR /src

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch

WORKDIR /src

COPY --from=builder /src/main /src

ENTRYPOINT ["./main"]