FROM golang:1.22.1-alpine

RUN go version
ENV GOPATH=/

COPY ./go.* ./
RUN go mod download

COPY ./ ./

RUN go build -o main ./cmd/app/main.go

CMD ["./main"]