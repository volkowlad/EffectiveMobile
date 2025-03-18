FROM golang:1.24
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o fin-app ./cmd/main.go

CMD ["./fin-app"]