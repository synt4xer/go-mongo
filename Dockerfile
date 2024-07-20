FROM golang:latest AS builder

# WORKDIR /app/go-mongo
WORKDIR /go/src/github.com/synt4xer/go-mongo

COPY go.mod Gopkg.lock ./
RUN go mod download

COPY . .
RUN go build -o cmd/go-mongo cmd/go-mongo.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /go/src/github.com/synt4xer/go-mongo/cmd/go-mongo .

# EXAMPLE
ENV PORT=8080
ENV MONGO_URI=mongodb://localhost:27017
ENV MONGO_DATABASE=myDatabase

EXPOSE 8080

CMD ["cmd/go-mongo"]